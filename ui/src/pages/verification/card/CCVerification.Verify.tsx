//||------------------------------------------------------------------------------------------------||
//|| CCVerification (Container)
//|| src/pages/verification/CCVerification.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React, {useRef, useState, useEffect}           from "react";
      import {useNavigate, useSearchParams}                 from "react-router-dom";
      import {convertCurrency}                              from "../../../utils/convertCurrency";

      //||------------------------------------------------------------------------------------------------||
      //|| Components
      //||------------------------------------------------------------------------------------------------||

      import MembersLayout                                  from "../../../layouts/MembersLayout";
      import ProgressSteps, { ProgessStep }                 from "../../../components/base/ProgressSteps";
      import InlineAlert                                    from "../../../components/base/InlineAlert";

      //||------------------------------------------------------------------------------------------------||
      //|| Interfaces
      //||------------------------------------------------------------------------------------------------||

      import { VerificationCard }                           from "../../../interfaces/verify/card/process";

      //||------------------------------------------------------------------------------------------------||
      //|| Step Props
      //||------------------------------------------------------------------------------------------------||

      export interface CCStepProps {
            process           : VerificationCard;
            updateProcess     : (process: VerificationCard) => void;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Page
      //||------------------------------------------------------------------------------------------------||

      export default function CCVerificationVerify() {

            //||------------------------------------------------------------------------------------------------||
            //|| Verification
            //||------------------------------------------------------------------------------------------------||

            const [searchParams] = useSearchParams();
            const verificationId = searchParams.get("identifier") || "";
      
            //||------------------------------------------------------------------------------------------------||
            //|| Navigate
            //||------------------------------------------------------------------------------------------------||

            const navigate                     = useNavigate();

            //||------------------------------------------------------------------------------------------------||
            //|| useState
            //||------------------------------------------------------------------------------------------------||

            const [code, setCode] = useState("");
            const [busy, setBusy] = useState(false);
            const [error, setError] = useState<string | null>(null);

            //||------------------------------------------------------------------------------------------------||
            //|| Process
            //||------------------------------------------------------------------------------------------------||

            const [process, setProcess] = useState<VerificationCard>({
                  step : 1,
                  donation : 0,
                  baseAmount : { 
                        currency : "USD",
                        amount   : import.meta.env.VITE_VERIFICATION_CARD_AMOUNT || 0.75
                  },
                  currency : "USD",
                  chargeAmount : { 
                        currency : "USD",
                        amount   : import.meta.env.VITE_VERIFICATION_CARD_AMOUNT || 0.75
                  }
            }); 

            //||------------------------------------------------------------------------------------------------||
            //|| Handle Returning
            //||------------------------------------------------------------------------------------------------||

            useEffect(() => {
                  if (!verificationId) return;
                  (async () => {
                        try {
                              const res = await fetch(`/api/v1/verify/card/lookup?uuid=${verificationId}`);
                              if (res.ok) {
                                    const data = await res.json();
                                    if (data?.status === "pending") {
                                          updateProcess({...process, step : 2});
                                    }
                              }
                        } catch {}
                  })();
            }, [verificationId]);

            //||------------------------------------------------------------------------------------------------||
            //|| Stepper
            //||------------------------------------------------------------------------------------------------||

            const steps: ProgessStep[] = [
                  { label: "How it works", description: "Provide your physical address details." },
                  { label: "Enter Credit Card", description: "Review and confirm your address." },
                  { label: "Processed", description: "Transaction Complete!" },
                  { label: "Confirm Statement Code", description: "Enter the statement code" },
            ]
            
            //||------------------------------------------------------------------------------------------------||
            //|| Set Process
            //||------------------------------------------------------------------------------------------------||

            const updateProcess = (update: Partial<VerificationCard>) => {
                  setProcess((prev) => ({
                        ...prev,
                        ...update
                  }));
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Props
            //||------------------------------------------------------------------------------------------------||

            const submitCode = async () => {
                  setBusy(true);
                  setError(null);
                  const payload = {
                        "uuid" : verificationId,
                        "code" : code                    
                  }
                  try {
                        const res = await fetch(`/v1/api/verify/card/check`, {
                              method: "POST",
                              headers: { "Content-Type": "application/json" },
                              body: JSON.stringify(payload),
                        });
                        const data = await res.json();
                        if (!res.ok) {
                              setError(data?.message || "Verification failed.");
                              return;
                        }
                        alert("Verification successful!");
                  } catch (err) {
                        setError("Network error. Please try again.");
                  } finally {
                        setBusy(false);
                  }
            };

            /*||------------------------------------------------------------------------------------------------||
            //|| Effect : Charge Amount (currency/base change)
            //||------------------------------------------------------------------------------------------------||*/

            useEffect(() => {
                  const API       = import.meta.env.VITE_COMPLYAGE_API_URL as string;
                  const currency  = (process as any).currency || process.chargeAmount?.currency || process.baseAmount.currency;
                  const baseAmt   = Number(process.baseAmount.amount) + process.donation;

                  if (!API || !currency || Number.isNaN(baseAmt)) return;

                  const ac = new AbortController();

                  (async () => {
                        const converted = await convertCurrency(baseAmt, currency);                        
                        if (converted === null) return;
                        setProcess(prev => ({
                              ...prev,
                              chargeAmount: {
                                    currency : String(currency),
                                    amount   : converted,
                                    country  : prev.chargeAmount.country 
                              },
                        }));
                  })();

                  return () => ac.abort();
            }, [process.donation, process.chargeAmount?.currency, process.baseAmount.amount, (process as any).currency]);

            //||------------------------------------------------------------------------------------------------||
            //|| Step
            //||------------------------------------------------------------------------------------------------||

            return (
                  <MembersLayout title="Credit Card Verification">
                        <div className="w-full max-w-2xl mx-auto">                              
                              <ProgressSteps steps={ steps } currentStep={ 4 } className="mb-6" />
                              <div className="space-y-6 p-6">
                                    <p className="text-gray-100 text-center font-bold">
                                          Please check your credit card statement and enter the 6-digit code to complete verification.
                                    </p>
                                    {error && <InlineAlert message={error} isError />}
                                    <div className="px-10">
                                          <input
                                                type="text"
                                                inputMode="numeric"
                                                maxLength={6}
                                                placeholder="_ _ _ _ _ _"
                                                className="input input-bordered w-full text-2xl h-16 text-center tracking-widest"
                                                value={code}
                                                onChange={(e) => setCode(e.target.value)}
                                          />
                                          <span className="text-gray-500 w-full flex p-0 m-0 justify-center">{ verificationId }</span>
                                    </div>
                                    <div className="w-full max-w-lg mx-auto text-center">
                                          <button
                                                onClick={submitCode}
                                                disabled={busy || code.length !== 6}
                                                className="btn btn-primary bg-orange-500 max-w-lg mx-auto text-xl px-10 py-7"
                                          >
                                                {busy ? "Verifying..." : "Submit Code"}
                                          </button>
                                    </div>
                              </div>         
                              <textarea defaultValue={ JSON.stringify(process, null, 2) } className="w-full bg-black/40 p-5 h-96 p-2 text-xs rounded-md" readOnly />                                                                                                 
                        </div>
                  </MembersLayout>                  
            );
      }

