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
      import {getEnv}                                       from "../../../data/getEnv";

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
            //|| UUID
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
                  step              : 1,
                  baseAmount        : import.meta.env.VITE_VERIFICATION_CARD_AMOUNT || 99,
                  donationAmount    : 0,
                  chargeAmount      : import.meta.env.VITE_VERIFICATION_CARD_AMOUNT || 99,
                  currency          : "USD"
            }); 


            //||------------------------------------------------------------------------------------------------||
            //|| Stepper
            //||------------------------------------------------------------------------------------------------||

            const steps: ProgessStep[] = [
                  { label: "How it works", description: "A Valid credit card is required" },
                  { label: "Enter Credit Card", description: "Review and confirm your address." },
                  { label: "Processed", description: "Transaction Complete!" },
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
                  const API      = getEnv("VITE_COMPLYAGE_API_URL") as string;
                  const currency = process.currency;
                  const baseAmt  = Number(process.baseAmount) || 0;
                  const donation = Number(process.donationAmount) || 0;
                  const totalAmt = baseAmt + donation;

                  if (!API || !currency || Number.isNaN(totalAmt)) return;

                  (async () => {
                        setProcess(prev => ({ ...prev, chargeAmount: totalAmt }));
                  })();

                  const ac = new AbortController();
                  return () => ac.abort();
            }, [process.donationAmount, process.baseAmount, process.currency]);

            //||------------------------------------------------------------------------------------------------||
            //|| Step
            //||------------------------------------------------------------------------------------------------||

            return (
                  <MembersLayout>
                        <div className="w-full max-w-2xl mx-auto">                              
                              <ProgressSteps steps={ steps } currentStep={ process.step } className="mb-2" />
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
                                    
                              {process.step >= 3 && (
                                    <CCVerificationStep3
                                          process={process}
                                          updateProcess={updateProcess}
                                    />
                              )}                                                                                
                        </div>
                  </MembersLayout>                  
            );
      }
