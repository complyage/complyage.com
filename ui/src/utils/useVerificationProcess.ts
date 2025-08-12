//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import { useEffect, useState }            from "react";
import { useSearchParams }                from "react-router-dom";
import type { VerificationProcessID }     from "../interfaces/verification.process.id";

//||------------------------------------------------------------------------------------------------||
//|| Hook
//||------------------------------------------------------------------------------------------------||

export function useVerificationProcess(pollInterval: number = 5000) {

      //||------------------------------------------------------------------------------------------------||      
      //|| Var
      //||------------------------------------------------------------------------------------------------||

      const [searchParams]                         = useSearchParams();
      const [process, setProcess]                  = useState<VerificationProcessID | null>(null);
      const [error, setError]                      = useState<string | null>(null);
      const [loading, setLoading]                  = useState(true);

      //||------------------------------------------------------------------------------------------------||      
      //|| Effect
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {

            const identifier = searchParams.get("identifier");
            if (!identifier) {
                  setError("Missing identifier in URL");
                  setLoading(false);
                  return;
            }

            const run = async () => {
                  try {
                        const res  = await fetch(`/v1/api/verify/id/process?identifier=${identifier}`);
                        const json = await res.json();
                        console.log("Verification Process Response:", json);
                        if (!res.ok || !json?.success) {
                              throw new Error(json?.message || "Failed to fetch verification process");
                        }
                        setProcess(json.data as VerificationProcessID);
                        setError(null);
                  } catch (err: any) {
                        setError(err.message || "Unknown error");
                  } finally {
                        setLoading(false);
                  }
            };

            run();
            const interval = setInterval(run, pollInterval);
            return () => clearInterval(interval);

      }, [searchParams]);

      //||------------------------------------------------------------------------------------------------||      
      //|| Return
      //||------------------------------------------------------------------------------------------------||

      return { process, error, loading };
}
