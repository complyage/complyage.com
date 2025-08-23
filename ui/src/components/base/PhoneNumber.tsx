/*||------------------------------------------------------------------------------------------------||
//|| Phone Number 
//|| 
//||------------------------------------------------------------------------------------------------||*/

/*||------------------------------------------------------------------------------------------------||
//|| Import 
//||------------------------------------------------------------------------------------------------||*/

import React                                          from "react";
import { CountryList }                                from "../../data/getCountries";
import type { Country }                               from "../../interfaces/base/geo";

/*||------------------------------------------------------------------------------------------------||
//|| Utils
//||------------------------------------------------------------------------------------------------||*/

import { formatSimplePhone }                          from "../../utils/phoneUtils";

/*||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||*/

type Props = {
      value    : { countryCode: string; phoneNumber: string };
      onChange : (val: { countryCode: string; phoneNumber: string }) => void;
};


/*||------------------------------------------------------------------------------------------------||
//|| Phone Input Component
//||------------------------------------------------------------------------------------------------||*/

export default function PhoneInput({ value, onChange }: Props) {
      const selectedCountry =
            CountryList.find(
                  c => c.dialCode === value.countryCode || c.code === value.countryCode
            ) || CountryList[0];

      const handleCountry = (e: React.ChangeEvent<HTMLSelectElement>) => {
            onChange({ countryCode: e.target.value, phoneNumber: "" });
      };

      const handleNumber = (e: React.ChangeEvent<HTMLInputElement>) => {
            // Always strip non-digits for storage
            const digits = e.target.value.replace(/\D/g, "");
            onChange({ countryCode: value.countryCode, phoneNumber: digits });
      };

      const prettyNumber = formatSimplePhone(value.phoneNumber, value.countryCode);

      return (
            <div className="flex gap-2 items-center">
                  <select
                        className="select select-bordered w-32 h-16 text-2xl rounded-lg"
                        value={selectedCountry.dialCode}
                        onChange={handleCountry}
                  >
                        {CountryList.map((country) => (
                              <option key={country.code} value={country.dialCode}>
                                    {country.flag} +{country.dialCode}
                              </option>
                        ))}
                  </select>
                  <input
                        type="tel"
                        inputMode="tel"
                        autoComplete="tel"
                        className="input input-bordered flex-1 h-16 text-2xl tracking-wider rounded-lg text-center"
                        placeholder="Phone number"
                        value={prettyNumber}
                        onChange={handleNumber}
                        maxLength={20}
                  />
            </div>
      );
}