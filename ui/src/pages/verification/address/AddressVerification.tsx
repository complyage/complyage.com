//||------------------------------------------------------------------------------------------------||
//|| CCVerification (Container)
//|| src/pages/verification/CCVerification.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React, {useRef, useState, useEffect}           from "react";
      import {useNavigate, useSearchParams}                 from "react-router-dom";
      import { Home }                                       from "lucide-react";
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

      import AddressVerificationStep1                       from "./AddressVerification.Step1";
      import AddressVerificationStep2                       from "./AddressVerification.Step2";
      import StripeWrapper                                  from "./StripeWrapper";
      import AddressVerificationStep4                       from "./AddressVerification.Step4";

      //||------------------------------------------------------------------------------------------------||
      //|| Interfaces
      //||------------------------------------------------------------------------------------------------||

      import { VerificationAddress }                        from "../../../interfaces/verify/address/process";

      //||------------------------------------------------------------------------------------------------||
      //|| Step Props
      //||------------------------------------------------------------------------------------------------||

      export interface StepProps {
            process           : VerificationAddress;
            updateProcess     : (process: VerificationAddress) => void;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Page
      //||------------------------------------------------------------------------------------------------||

      export default function AddressVerification() {

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

            const [process, setProcess] = useState<VerificationAddress>({
                  verifyAddress : {
                        line1: "234 Main",
                        line2: "",
                        city: "Phoenix",
                        state: "AZ",
                        postal: "85024",
                        country: "US"
                  },
                  step              : 1,
                  currency          : "USD",
                  baseAmount        : Number(getEnv("VITE_VERIFICATION_ADDRESS_AMOUNT")) || 99,
                  donationAmount    : 0,
                  chargeAmount      : Number(getEnv("VITE_VERIFICATION_ADDRESS_AMOUNT")) || 99,
            });

            //||------------------------------------------------------------------------------------------------||
            //|| Set Process
            //||------------------------------------------------------------------------------------------------||

            const updateProcess = (update: Partial<VerificationAddress>) => {
                  setProcess((prev) => ({
                        ...prev,
                        ...update
                  }));
            }

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
                  { label: "How it works",      description: "Provide your physical address details." },
                  { label: "Verify Address",    description: "Review and confirm your address." },
                  { label: "Pay Postage",       description: "Pay for the postage to send" },
            ];    
            
            //||------------------------------------------------------------------------------------------------||
            //|| Update Currency
            //||------------------------------------------------------------------------------------------------||
                        
            useEffect(() => {
                  const API      = getEnv("VITE_COMPLYAGE_API_URL") as string;
                  const currency = process.currency;
                  const baseAmt  = Number(process.baseAmount) || 0;
                  const donation = Number(process.donationAmount) || 0;
                  const totalAmt = baseAmt + donation;

                  if (!API || !currency || Number.isNaN(totalAmt)) return;

                  const ac = new AbortController();

                  (async () => {
                        // convertCurrency(totalAmt, currency) should return the converted number, or null
                        const converted = await convertCurrency(totalAmt, currency);
                        if (converted === null) return;
                        setProcess(prev => ({
                              ...prev,
                              chargeAmount: converted,
                        }));
                  })();

                  return () => ac.abort();
            }, [process.donationAmount, process.baseAmount, process.currency]);

            //||------------------------------------------------------------------------------------------------||
            //|| Step
            //||------------------------------------------------------------------------------------------------||

            return (
                  <MembersLayout title="Address Verification" icon={ Home }>
                        <div className="w-full max-w-2xl mx-auto">                              
                              <ProgressSteps steps={ steps } currentStep={ process.step } className="mb-6" />
                              { process.step === 1 && 
                                    <AddressVerificationStep1 
                                          process={ process } 
                                          updateProcess={updateProcess}
                                    /> 
                              }    

                              { process.step === 2 &&                               
                                    <AddressVerificationStep2 
                                          process={process}
                                          updateProcess={updateProcess}
                                    />
                              }    
                                    
                              {process.step === 3 && (
                                    <StripeWrapper
                                          process={process}
                                          updateProcess={updateProcess}
                                    />
                              )}                                                                                

                              {process.step === 4 && (
                                    <AddressVerificationStep4
                                          process={process}
                                          updateProcess={updateProcess}
                                    />
                              )}                                                                                
                        </div>
                  </MembersLayout>                  
            );
      }
