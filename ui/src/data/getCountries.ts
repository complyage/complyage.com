//||------------------------------------------------------------------------------------------------||
//|| Countries Data
//|| src/data/getCountries.ts
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Country Type
      //||------------------------------------------------------------------------------------------------||

      import type { Country } from "../interfaces/base/geo"; // adjust the import as needed

      //||------------------------------------------------------------------------------------------------||
      //|| Country List
      //||------------------------------------------------------------------------------------------------||

      export const CountryList: Country[] = [
      { code: "US", name: "United States",        flag: "🇺🇸", dialCode: "1" },
      { code: "GB", name: "United Kingdom",       flag: "🇬🇧", dialCode: "44" },
      { code: "CA", name: "Canada",               flag: "🇨🇦", dialCode: "1" },
      { code: "AU", name: "Australia",            flag: "🇦🇺", dialCode: "61" },
      { code: "DE", name: "Germany",              flag: "🇩🇪", dialCode: "49" },
      { code: "FR", name: "France",               flag: "🇫🇷", dialCode: "33" },
      { code: "IT", name: "Italy",                flag: "🇮🇹", dialCode: "39" },
      { code: "ES", name: "Spain",                flag: "🇪🇸", dialCode: "34" },
      { code: "BR", name: "Brazil",               flag: "🇧🇷", dialCode: "55" },
      { code: "IN", name: "India",                flag: "🇮🇳", dialCode: "91" },
      { code: "MX", name: "Mexico",               flag: "🇲🇽", dialCode: "52" },
      { code: "JP", name: "Japan",                flag: "🇯🇵", dialCode: "81" },
      { code: "KR", name: "South Korea",          flag: "🇰🇷", dialCode: "82" },
      { code: "RU", name: "Russia",               flag: "🇷🇺", dialCode: "7" },
      { code: "TR", name: "Turkey",               flag: "🇹🇷", dialCode: "90" },
      { code: "ZA", name: "South Africa",         flag: "🇿🇦", dialCode: "27" },
      { code: "SE", name: "Sweden",               flag: "🇸🇪", dialCode: "46" },
      { code: "NL", name: "Netherlands",          flag: "🇳🇱", dialCode: "31" },
      { code: "CH", name: "Switzerland",          flag: "🇨🇭", dialCode: "41" },
      { code: "BE", name: "Belgium",              flag: "🇧🇪", dialCode: "32" },
      { code: "NO", name: "Norway",               flag: "🇳🇴", dialCode: "47" },
      { code: "SG", name: "Singapore",            flag: "🇸🇬", dialCode: "65" },
      { code: "AE", name: "United Arab Emirates", flag: "🇦🇪", dialCode: "971" },
      { code: "AR", name: "Argentina",            flag: "🇦🇷", dialCode: "54" },
      { code: "SA", name: "Saudi Arabia",         flag: "🇸🇦", dialCode: "966" },
      { code: "ID", name: "Indonesia",            flag: "🇮🇩", dialCode: "62" },
      { code: "NG", name: "Nigeria",              flag: "🇳🇬", dialCode: "234" },
      { code: "PL", name: "Poland",               flag: "🇵🇱", dialCode: "48" },
      { code: "TH", name: "Thailand",             flag: "🇹🇭", dialCode: "66" },
      { code: "HK", name: "Hong Kong SAR",        flag: "🇭🇰", dialCode: "852" }
      ];

      //||------------------------------------------------------------------------------------------------||
      //|| Get Country Label
      //||------------------------------------------------------------------------------------------------||

      export function getCountryLabel(code: string): string {
            const country = CountryList.find(c => c.code === code);
            return country ? country.name : code;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Get All Countries
      //||------------------------------------------------------------------------------------------------||

      export function getAllCountries(): Country[] {
            return CountryList;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Get Country Flag
      //||------------------------------------------------------------------------------------------------||

      export function getCountryFlag(code: string): string {
            const country = CountryList.find(c => c.code === code);
            return country && country.flag ? country.flag : "";
      }
