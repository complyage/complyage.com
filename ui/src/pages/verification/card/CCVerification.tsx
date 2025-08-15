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

      //||------------------------------------------------------------------------------------------------||
      //|| Pages
      //||------------------------------------------------------------------------------------------------||

      import CCVerificationStep1                            from "./CCVerification.Step1";
      import StripeWrapper                                  from "./StripeWrapper";
      import CCVerificationStep3                            from "./CCVerification.Step3";

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

      export default function CCVerification() {

            //||------------------------------------------------------------------------------------------------||
            //|| Verification
            //||------------------------------------------------------------------------------------------------||

            const [searchParams] = useSearchParams();
            const verificationId = searchParams.get("verification") || "";
      
            //||------------------------------------------------------------------------------------------------||
            //|| Navigate
            //||------------------------------------------------------------------------------------------------||

            const navigate                     = useNavigate();

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
                              const res = await fetch(`/api/v1/verify/card/lookup?verification=${verificationId}`);
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
                              <ProgressSteps steps={ steps } currentStep={ process.step } className="mb-6" />
                              { process.step === 1 && 
                                    <CCVerificationStep1 
                                          process={ process } 
                                          updateProcess={updateProcess}
                                    /> 
                              }    

                              { process.step === 2 &&                               
                                    <StripeWrapper 
                                          process={process}
                                          updateProcess={updateProcess}
                                    />
                              }    
                                    
                              {process.step === 3 && (
                                    <CCVerificationStep3
                                          process={process}
                                          updateProcess={updateProcess}
                                    />
                              )}                                                                                
                        </div>
                  </MembersLayout>                  
            );
      }
