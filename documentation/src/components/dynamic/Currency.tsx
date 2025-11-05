//||------------------------------------------------------------------------------------------------||
//|| CurrencyDropdown
//|| src/components/payments/CurrencyDropdown.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      import React                        from "react";

      //||------------------------------------------------------------------------------------------------||
      //|| Currency Codes
      //||------------------------------------------------------------------------------------------------||

      import { getCurrencyCodes, getCurrencyZeros, getCurrencySymbol }        from "../../data/getCurrencies";
      export type CurrencyCode = (typeof CURRENCY_CODES)[number];

      //||------------------------------------------------------------------------------------------------||
      //|| Currency Codes
      //||------------------------------------------------------------------------------------------------||

      type CurrencyDropdownProps = {
            currency                : CurrencyCode | "";
            setCurrency             : (code: CurrencyCode) => void;
            includeZeroDecimalNote? : boolean; // shows "(no decimals)" for JPY, KRW, etc.
            placeholder?            : string;
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Currency Codes
      //||------------------------------------------------------------------------------------------------||

      const CURRENCY_CODES                = getCurrencyCodes();
      const ZERO_DECIMAL: Set<string>     = new Set(getCurrencyZeros());

      //||------------------------------------------------------------------------------------------------||
      //|| Component
      //||------------------------------------------------------------------------------------------------||

      export default function CurrencyDropdown({
            currency,
            setCurrency,
            includeZeroDecimalNote = true,
            placeholder = "Select currency…",
      }: CurrencyDropdownProps) {
            return (
                  <div className="">
                        <select className="select select-bordered w-full" value={currency} onChange={(e) => setCurrency(e.target.value as CurrencyCode)}>
                              <option value="">{placeholder}</option>
                              {CURRENCY_CODES.map((code) => (
                                    <option key={code} value={code}>
                                          {code} - [{getCurrencySymbol(code)}]
                                          {includeZeroDecimalNote && ZERO_DECIMAL.has(code as CurrencyCode) ? "" : ""}
                                    </option>
                              ))}
                        </select>
                  </div>
            );
      }
