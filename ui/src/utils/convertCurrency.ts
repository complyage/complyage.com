
//||------------------------------------------------------------------------------------------------||
//|| Clean 
//|| src/utils/convertCurrency.ts
//||------------------------------------------------------------------------------------------------||

      import { CurrencyCode } from "../components/dynamic/Currency";

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      export async function convertCurrency(amount: number, toCurrency: CurrencyCode): Promise<number | null> {
            const API       = import.meta.env.VITE_COMPLYAGE_API_URL as string;
            if (!API || Number.isNaN(amount)) return null;
            try {
                  const url = `${API}/v1/api/currency?amount=${encodeURIComponent(amount)}&currency=${encodeURIComponent(toCurrency)}`;
                  const res = await fetch(url, { method: "GET" });
                  if (!res.ok) return null;
                  const json = await res.json() as {
                        success?: boolean;
                        data?: { converted?: number|string; currency?: string };
                  };
                  const d = json?.data ?? {};
                  if (!d.converted || !d.currency) return null;
                  return (d.converted && Number(d.converted) > 0) ? Number(d.converted) : null;
            } catch (err: any) {
                  return null;
            }
      }