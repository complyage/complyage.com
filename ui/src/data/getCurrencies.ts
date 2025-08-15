//||------------------------------------------------------------------------------------------------||
//|| Currencies
//|| src/utils/getCurrencies.ts
//||------------------------------------------------------------------------------------------------||
   
export function getCurrencyCodes(): string[] { return ["USD","EUR","JPY","GBP","AUD","CAD","CHF","CNY","HKD","NZD"]; }
export function getCurrencyZeros(): string[] { return ["BIF","CLP","DJF","GNF","JPY","KMF","KRW","MGA","PYG","RWF","UGX","VND","VUV","XAF","XOF","XPF"]; }

export function formatStripeAmount(amount: number | string, currency: string): string {
      const amt = typeof amount === "string" ? parseFloat(amount) : amount;
      const zeroDecimals = getCurrencyZeros();
      if (zeroDecimals.includes(currency.toUpperCase())) {
            return amt.toString();
      }
      return (amt / 100).toFixed(2);
}