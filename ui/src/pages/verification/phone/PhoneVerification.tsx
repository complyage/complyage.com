//||------------------------------------------------------------------------------------------------||
//|| PhoneVerification
//|| src/components/security/PhoneVerification.tsx
//||------------------------------------------------------------------------------------------------||

      import React, {useMemo, useRef, useState, useEffect}        from "react";
      import {Phone, ShieldCheck, Key, RefreshCw, Lock}           from "lucide-react";
      import {useNavigate}                                        from "react-router-dom";

      //||------------------------------------------------------------------------------------------------||
      //|| Components
      //||------------------------------------------------------------------------------------------------||

      import MembersLayout                                        from "../../../layouts/MembersLayout";
      import ProgressSteps, { ProgessStep }                       from "../../../components/base/ProgressSteps";

      //||------------------------------------------------------------------------------------------------||
      //|| Interfaces
      //||------------------------------------------------------------------------------------------------||

      import { VerificationPhone }                               from "../../../interfaces/verify/phone/process";

      //||------------------------------------------------------------------------------------------------||
      //|| Pages
      //||------------------------------------------------------------------------------------------------||

      import PhoneVerificationStep1                               from "./PhoneVerification.Step1";
      import PhoneVerificationStep2                               from "./PhoneVerification.Step2";

      //||------------------------------------------------------------------------------------------------||
      //|| Types
      //||------------------------------------------------------------------------------------------------||

      type Country = {
            code: string; // ISO code
            name: string; // display name
            dial: string; // dialing prefix
            flag: string; // emoji flag
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Country list (minimal — extend as needed)
      //||------------------------------------------------------------------------------------------------||

      const countries: Country[] = [
            {code: "US", name: "United States", dial: "+1", flag: "🇺🇸"},
            {code: "CA", name: "Canada", dial: "+1", flag: "🇨🇦"},
            {code: "GB", name: "United Kingdom", dial: "+44", flag: "🇬🇧"},
            {code: "AU", name: "Australia", dial: "+61", flag: "🇦🇺"},
            {code: "DE", name: "Germany", dial: "+49", flag: "🇩🇪"},
      ];


      //||------------------------------------------------------------------------------------------------||
      //|| Component
      //||------------------------------------------------------------------------------------------------||

      export default function PhoneVerification() {

            //||------------------------------------------------------------------------------------------------||
            //|| Navigate
            //||------------------------------------------------------------------------------------------------||

            const navigate                     = useNavigate();

            //||------------------------------------------------------------------------------------------------||
            //|| Process
            //||------------------------------------------------------------------------------------------------||

            const [process, setProcess] = useState<VerificationPhone>({
                  countryCode       : "",
                  phoneNumber       : "",
                  step              : 1,
            });

            //||------------------------------------------------------------------------------------------------||
            //|| Set Process
            //||------------------------------------------------------------------------------------------------||

            const updateProcess = (update: Partial<VerificationPhone>) => {
                  setProcess((prev) => ({
                        ...prev,
                        ...update
                  }));
            }


            //||------------------------------------------------------------------------------------------------||
            //|| Stepper
            //||------------------------------------------------------------------------------------------------||

            const steps: ProgessStep[] = [
                  { label: "Enter Phone Number", description: "Enter your phone number" },
                  { label: "Enter Verification Code ", description: "Enter the code we sent your." },
            ]
                  
            //||------------------------------------------------------------------------------------------------||
            //|| Send Code
            //||------------------------------------------------------------------------------------------------||

            return (
                  <MembersLayout title="Phone Verification">
                        <div className="w-full mx-auto max-w-2xl">
                              {/* Stepper */}
                              <ProgressSteps steps={ steps } currentStep={ process.step } className="mb-6" />

                              {/* Step 1 */}
                              {process.step === 1 && (
                                    <PhoneVerificationStep1 process={ process } updateProcess={ updateProcess } />
                              )}

                              {/* Step 2 */}
                              {process.step === 2 && (
                                    <PhoneVerificationStep2 process={ process } updateProcess={ updateProcess } />
                              )}
                        </div>
                  </MembersLayout>
            );
      }
