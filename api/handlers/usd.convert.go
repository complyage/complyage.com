package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"base/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| In-memory daily cache for USD rates
//||------------------------------------------------------------------------------------------------||

var (
	usdRatesCache = struct {
		mu    sync.RWMutex
		rates map[string]float64 // upper-case ISO code -> rate (1 USD = rate * CURRENCY)
		asOf  time.Time          // provider's date or fetch time
	}{rates: make(map[string]float64)}

	cacheTTL = 24 * time.Hour
)

//||------------------------------------------------------------------------------------------------||
//|| Provider response (Frankfurter)
//|| Example: GET https://api.frankfurter.app/latest?from=USD
//|| { "amount":1.0, "base":"USD", "date":"2025-08-12", "rates": { "EUR":0.92, ... } }
//||------------------------------------------------------------------------------------------------||

type frankfurterLatest struct {
	Amount float64            `json:"amount"`
	Base   string             `json:"base"`
	Date   string             `json:"date"`
	Rates  map[string]float64 `json:"rates"`
}

//||------------------------------------------------------------------------------------------------||
//|| Helpers
//||------------------------------------------------------------------------------------------------||

func getUSDRate(ctx context.Context, code string) (rate float64, asOf time.Time, err error) {
	code = strings.ToUpper(strings.TrimSpace(code))
	if code == "" {
		return 0, time.Time{}, errors.New("missing currency code")
	}
	if code == "USD" {
		return 1.0, time.Now(), nil
	}

	// Fast path: read cache if fresh.
	usdRatesCache.mu.RLock()
	fresh := time.Since(usdRatesCache.asOf) < cacheTTL && len(usdRatesCache.rates) > 0
	if fresh {
		rate, ok := usdRatesCache.rates[code]
		asOf = usdRatesCache.asOf
		usdRatesCache.mu.RUnlock()
		if !ok {
			return 0, asOf, fmt.Errorf("unsupported currency: %s", code)
		}
		return rate, asOf, nil
	}
	usdRatesCache.mu.RUnlock()

	// Slow path: refresh from provider.
	if err := refreshUSDRates(ctx); err != nil {
		return 0, time.Time{}, fmt.Errorf("failed to refresh rates: %w", err)
	}

	usdRatesCache.mu.RLock()
	rate, ok := usdRatesCache.rates[code]
	asOf = usdRatesCache.asOf
	usdRatesCache.mu.RUnlock()
	if !ok {
		return 0, asOf, fmt.Errorf("unsupported currency: %s", code)
	}
	return rate, asOf, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Refresh Daily
//||------------------------------------------------------------------------------------------------||

func refreshUSDRates(ctx context.Context) error {
	// Small timeout if caller didn't set one.
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 7*time.Second)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.frankfurter.app/latest?from=USD", nil)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 6 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("provider status %d", resp.StatusCode)
	}

	var data frankfurterLatest
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}
	if strings.ToUpper(data.Base) != "USD" || len(data.Rates) == 0 {
		return errors.New("invalid provider payload")
	}

	// Update cache.
	asOf := time.Now()
	if t, err := time.Parse("2006-01-02", data.Date); err == nil {
		asOf = t
	}

	usdRatesCache.mu.Lock()
	usdRatesCache.rates = data.Rates
	usdRatesCache.asOf = asOf
	usdRatesCache.mu.Unlock()

	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Handler: ConvertUSDHandler
//|| GET /v1/api/fx/convert?amount=12.34&currency=EUR
//|| Response:
//|| { success:true, amountUSD:12.34, currency:"EUR", rate:0.92, converted:11.35, asOf:"2025-08-12" }
//||------------------------------------------------------------------------------------------------||

func ConvertUSDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		responses.Error(w, http.StatusMethodNotAllowed, "use GET with query params")
		return
	}

	amountStr := strings.TrimSpace(r.URL.Query().Get("amount"))
	code := strings.ToUpper(strings.TrimSpace(r.URL.Query().Get("currency")))

	if amountStr == "" || code == "" {
		responses.Error(w, http.StatusBadRequest, "missing amount or currency")
		return
	}

	amt, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || !isFinite(amt) {
		responses.Error(w, http.StatusBadRequest, "invalid amount")
		return
	}
	if amt < 0 {
		responses.Error(w, http.StatusBadRequest, "amount must be >= 0")
		return
	}

	rate, asOf, err := getUSDRate(r.Context(), code)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	converted := amt * rate

	responses.Success(w, http.StatusOK, map[string]any{
		"success":    true,
		"amountUSD":  round2(amt),
		"currency":   code,
		"rate":       rate,
		"converted":  round2(converted),
		"asOf":       asOf.Format("2006-01-02"),
		"provider":   "frankfurter.app",
		"cacheFresh": time.Since(asOf) < cacheTTL,
	})
}

//||------------------------------------------------------------------------------------------------||
//|| Small numeric helpers
//||------------------------------------------------------------------------------------------------||

func round2(v float64) float64 { // bank/retail rounding to 2 decimals for display
	return float64(int64(v*100+0.5)) / 100
}

func isFinite(f float64) bool {
	return !math.IsNaN(f) && !math.IsInf(f, 0) // 0 => check both +Inf and -Inf
}
