
export function useCountryFullName(code: string | undefined | null): string | undefined {
      const upperCode = code.toUpperCase();

      switch (upperCode) {
            case "US":
                  return "United States";
            case "FR":
                  return "France";
            case "GB":
                  return "United Kingdom";
            case "AU":
                  return "Australia";
            case "DE":
                  return "Germany";
            case "EU":
                  return "European Union"; // Note: EU is a political/economic union, not a country.
            case "IE":
                  return "Ireland";
            case "IT":
                  return "Italy";
            default:
                  return undefined;
      }
}
