//||------------------------------------------------------------------------------------------------||
//|| Currencies
//|| src/utils/getCurrencies.ts
//||------------------------------------------------------------------------------------------------||
   
export function getCurrencyCodes(): string[] { return ["USD","EUR","JPY","GBP","AUD","CAD","CHF","CNY","HKD","NZD"]; }
export function getCurrencyZeros(): string[] { return ["BIF","CLP","DJF","GNF","JPY","KMF","KRW","MGA","PYG","RWF","UGX","VND","VUV","XAF","XOF","XPF"]; }

//||------------------------------------------------------------------------------------------------||
//|| Denote Based on Locale
//||------------------------------------------------------------------------------------------------||

function formatCurrency(amount: number, currency: string, locale?: string): string {
      const usedLocale = locale || (typeof navigator !== "undefined" ? navigator.language : "en-US");
      return new Intl.NumberFormat(usedLocale, {
            style: 'currency',
            currency,
      }).format(amount);
}

//||------------------------------------------------------------------------------------------------||
//|| Convert to Lowest Denomination/Cents etc
//||------------------------------------------------------------------------------------------------||

export function formatStripeAmount(amount: number | string, currency: string, noFormat = false): string {
      const amt = typeof amount === "string" ? parseFloat(amount) : amount;
      const zeroDecimals = getCurrencyZeros();
      if (typeof(currency) !== "string" || !currency) {
            console.warn("No currency provided for formatStripeAmount");
            return "0.00";
      }
      if (zeroDecimals.includes(currency.toUpperCase())) {
            if (!noFormat) return formatCurrency(amt, currency);
            return amt.toString();
      }      
      if (!noFormat) return formatCurrency((amt / 100), currency);
      return (amt / 100).toString();
}

//||------------------------------------------------------------------------------------------------||
//|| Decimal Based on Currency
//||------------------------------------------------------------------------------------------------||

export function toStripeAmount(amount: number | string, currency: string): number {
      const amt = typeof amount === "string" ? parseFloat(amount) : amount;
      const zeroDecimals = getCurrencyZeros();
      if (zeroDecimals.includes(currency.toUpperCase())) {
            // No decimal, don't multiply
            return Math.round(amt);
      }
      // Multiply by 100 (for cents), and round to avoid floating point bugs
      return Math.round(amt * 100);
}

//||------------------------------------------------------------------------------------------------||
//|| Currency Symbols (Expanded)
//||------------------------------------------------------------------------------------------------||

export const CurrencySymbols: Record<string, string> = {
      USD: "$",          // US Dollar
      CAD: "$",          // Canadian Dollar
      AUD: "$",          // Australian Dollar
      NZD: "$",          // New Zealand Dollar
      HKD: "$",          // Hong Kong Dollar
      SGD: "$",          // Singapore Dollar

      EUR: "€",          // Euro
      GBP: "£",          // British Pound
      JPY: "¥",          // Japanese Yen
      CNY: "¥",          // Chinese Yuan/Renminbi

      CHF: "CHF",        // Swiss Franc

      SEK: "kr",         // Swedish Krona
      NOK: "kr",         // Norwegian Krone
      DKK: "kr",         // Danish Krone
      ISK: "kr",         // Icelandic Krona

      RUB: "₽",          // Russian Ruble
      KRW: "₩",          // South Korean Won

      INR: "₹",          // Indian Rupee
      MXN: "$",          // Mexican Peso
      BRL: "R$",         // Brazilian Real
      ZAR: "R",          // South African Rand
      TRY: "₺",          // Turkish Lira
      AED: "د.إ",         // UAE Dirham
      SAR: "ر.س",         // Saudi Riyal
      PLN: "zł",         // Polish Zloty
      CZK: "Kč",         // Czech Koruna
      HUF: "Ft",         // Hungarian Forint
      // Add more as needed
};

// Fallback: use code if symbol not found
export function getCurrencySymbol(code: string): string {
      if (!code || typeof code !== "string") return "???";
      return CurrencySymbols[code.toUpperCase()] || code;
}
