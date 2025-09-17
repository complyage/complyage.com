package public

import (
	"net/http"
	"os"

	"github.com/complyage/base/db/models"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Donate API Response Types
//||------------------------------------------------------------------------------------------------||

type DonateCryptoOption struct {
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Prefix  string `json:"prefix"`
	Address string `json:"address"`
	QR      string `json:"qr"`
	Color   string `json:"color"`
}

type DonateAddress struct {
	Name     string `json:"name"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Postal   string `json:"postal"`
	Country  string `json:"country"`
}

type DonateApiResponse struct {
	Crypto  []DonateCryptoOption `json:"crypto"`
	Address DonateAddress        `json:"address"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler: Donate API (Crypto List + Address)
//||------------------------------------------------------------------------------------------------||

func DonateHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch cryptos from DB
	//||------------------------------------------------------------------------------------------------||

	var cryptos []models.Crypto
	if err := app.SQLDB["main"].DB.Order("id_crypto ASC").Find(&cryptos).Error; err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to fetch crypto options")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Compose crypto options with hardcoded symbol, QR, color
	//|| (You could extend your DB for this if you want!)
	//||------------------------------------------------------------------------------------------------||

	symbolMap := map[string]string{
		"Bitcoin":  "BTC",
		"Ethereum": "ETH",
		"USDT":     "USDT",
		// Add more if needed
	}
	colorMap := map[string]string{
		"BTC":  "bg-[#f7931a] text-white",
		"ETH":  "bg-[#627eea] text-white",
		"USDT": "bg-[#26a17b] text-white",
	}
	qrMap := map[string]string{
		"BTC":  "/img/donate/bitcoin-qr.png",
		"ETH":  "/img/donate/ethereum-qr.png",
		"USDT": "/img/donate/usdt-qr.png",
	}

	cryptoOptions := make([]DonateCryptoOption, 0, len(cryptos))
	for _, c := range cryptos {
		symbol := symbolMap[c.Name]
		qr := qrMap[symbol]
		color := colorMap[symbol]
		cryptoOptions = append(cryptoOptions, DonateCryptoOption{
			Name:    c.Name,
			Symbol:  symbol,
			Prefix:  c.Prefix,
			Address: c.Address,
			QR:      qr,
			Color:   color,
		})
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Address block from env (failover: empty string)
	//||------------------------------------------------------------------------------------------------||

	addr := DonateAddress{
		Name:     os.Getenv("VERIFICATION_ADDRESS_RETURN_NAME"),
		Address1: os.Getenv("VERIFICATION_ADDRESS_RETURN_ADDRESS1"),
		Address2: os.Getenv("VERIFICATION_ADDRESS_RETURN_ADDRESS2"),
		City:     os.Getenv("VERIFICATION_ADDRESS_RETURN_CITY"),
		State:    os.Getenv("VERIFICATION_ADDRESS_RETURN_STATE"),
		Postal:   os.Getenv("VERIFICATION_ADDRESS_RETURN_POSTAL"),
		Country:  os.Getenv("VERIFICATION_ADDRESS_RETURN_COUNTRY"),
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Final JSON Response
	//||------------------------------------------------------------------------------------------------||

	resp := DonateApiResponse{
		Crypto:  cryptoOptions,
		Address: addr,
	}
	responses.Success(w, http.StatusOK, resp)
}
