//||------------------------------------------------------------------------------------------------||
//|| AddressVerification (Container)
//|| src/components/security/AddressVerification.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import { CardData, Address, Country }           from "../interfaces/verification.location";
      import { onlyDigits }                           from "./clean";

      //||------------------------------------------------------------------------------------------------||
      //|| Is Digit
      //||------------------------------------------------------------------------------------------------||

      export function isDigit(ch: string) {
             return ch >= "0" && ch <= "9"
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Credit Card Valid
      //||------------------------------------------------------------------------------------------------||

      export function isAddressValid(a: Address) {
            if ((a.line1 || "").trim().length < 4) return false;
            if ((a.city || "").trim().length < 2) return false;
            if ((a.state || "").trim().length < 2) return false;
            if ((a.postal || "").trim().length < 3) return false;
            if (!a.country) return false;
            return true;
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Credit Card Luhn
      //||------------------------------------------------------------------------------------------------||

      export function isLuhn(digits: string) {
            let sum = 0,
                  alt = false;
            for (let i = digits.length - 1; i >= 0; i--) {
                  let n = Number(digits[i]);
                  if (alt) {
                        n *= 2;
                        if (n > 9) n -= 9;
                  }
                  sum += n;
                  alt = !alt;
            }
            return sum % 10 === 0;
      };

      //||------------------------------------------------------------------------------------------------||
      //|| URL
      //||------------------------------------------------------------------------------------------------||

      export function isValidURL(url: string): boolean {
            try {
                  new URL(url);
                  return true;
            } catch {
                  return false;
            }
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Credit Card Luhn
      //||------------------------------------------------------------------------------------------------||

      export function isNumber(value: unknown, options?: {allowNaN?: boolean; allowInfinity?: boolean}): value is number {
		const {allowNaN = false, allowInfinity = false} = options ?? {};

		let n: number;

		if (typeof value === "number") {
			n = value;
		} else if (value instanceof Number) {
			n = value.valueOf();
		} else {
			return false;
		}

		if (Number.isNaN(n)) {
			return allowNaN; // NaN is of type number, but usually not desired
		}

		if (!allowInfinity && !Number.isFinite(n)) {
			// catches +Infinity and -Infinity
			return false;
		}

		return true;
	}