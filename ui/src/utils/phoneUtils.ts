/*||------------------------------------------------------------------------------------------------||
//|| Simple Phone Formatter
//||------------------------------------------------------------------------------------------------||*/

export function formatSimplePhone(input: string, countryCode: string) {
      const digits = input.replace(/\D/g, "");
      if (countryCode === "1" && digits.length >= 7) {
            if (digits.length >= 11) {
                  return `+${digits[0]} (${digits.slice(1, 4)}) ${digits.slice(4, 7)}-${digits.slice(7, 11)}`;
            } else if (digits.length >= 10) {
                  return `(${digits.slice(0, 3)}) ${digits.slice(3, 6)}-${digits.slice(6, 10)}`;
            } else if (digits.length > 6) {
                  return `(${digits.slice(0, 3)}) ${digits.slice(3, 6)}-${digits.slice(6)}`;
            } else if (digits.length > 3) {
                  return `(${digits.slice(0, 3)}) ${digits.slice(3)}`;
            } else {
                  return digits;
            }
      } else {
            // Group by 3s for other countries
            return digits.replace(/(\d{1,3})(?=(\d{3})+(?!\d))/g, "$1 ");
      }
}