//||------------------------------------------------------------------------------------------------||
//|| Google Address Utilities
//|| src/util/googleAddress.ts
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Imports & Types
      //||------------------------------------------------------------------------------------------------||

      import type { Address } from "../../interfaces/base/geo";

      //||------------------------------------------------------------------------------------------------||
      //|| API Key (from Vite env)
      //||------------------------------------------------------------------------------------------------||

      const GOOGLE_API_KEY = import.meta.env.VITE_GOOGLE_API_PUBLIC as string;

      //||------------------------------------------------------------------------------------------------||
      //|| Format Address for Google Maps
      //||------------------------------------------------------------------------------------------------||

      export function formatAddressForGoogle(addr: Address): string {
            // "123 Main St, Apt 2, Springfield, IL 12345, US"
            return [
                  addr.line1,
                  addr.line2,
                  [addr.city, addr.state, addr.postal].filter(Boolean).join(" "),
                  addr.country
            ]
                  .filter(Boolean)
                  .join(", ");
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Verify Address with Google Maps Geocoding API
      //||------------------------------------------------------------------------------------------------||

      export async function verifyAddressWithGoogle(addr: Address): Promise<{
            ok: boolean;
            status: string;
            message: string;
            results?: any[];
      }> {
            if (!GOOGLE_API_KEY) {
                  return {
                        ok: false,
                        status: "NO_API_KEY",
                        message: "Google API key is not set in VITE_GOOGLE_API_PUBLIC."
                  };
            }
            const address = formatAddressForGoogle(addr);
            const url = `https://maps.googleapis.com/maps/api/geocode/json?address=${encodeURIComponent(address)}&key=${GOOGLE_API_KEY}`;

            try {
                  const resp = await fetch(url);
                  const data = await resp.json();
                  console.log("Google Maps API Response:", data);
                  if (data.status === "OK" && Array.isArray(data.results) && data.results.length > 0) {
                        return {
                              ok: true,
                              status: "OK",
                              message: "Address verified with Google Maps.",
                              results: data.results
                        };
                  } else {
                        return {
                              ok: false,
                              status: data.status,
                              message: `Google Maps could not verify this address.`
                        };
                  }
            } catch (err: any) {
                  return {
                        ok: false,
                        status: "ERROR",
                        message: "Error contacting Google Maps API."
                  };
            }
      }
