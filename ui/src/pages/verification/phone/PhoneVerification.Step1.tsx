//||------------------------------------------------------------------------------------------------||
//|| PhoneVerification.Step1
//|| src/components/payments/PhoneVerification.Step1.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      import React                                          from "react";
      import { Phone, Lock }                                from "lucide-react";
      
      //||------------------------------------------------------------------------------------------------||
      //|| Interfacss
      //||------------------------------------------------------------------------------------------------||

      import { VerificationPhone }                          from "../../../interfaces/verify/phone/process";

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      interface PhoneStep1Props {
            process           : VerificationPhone;
            setProcess        : (process: Partial<VerificationPhone>) => void;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      export default function PhoneVerificationStep1({ process, setProcess }: PhoneStep1Props) {
            return (
                  <div className="space-y-4 p-5">

                        <div className="flex items-center gap-2 text-gray-300">
                              <Phone className="w-5 h-5" />
                              <span>We’ll text a 6-digit verification code to your phone.</span>
                        </div>

                        <div className="flex gap-2">
                              <select
                                    className="select select-bordered w-28 h-16 text-2xl">
                              </select>
                              <input
                                    placeholder="Phone number"
                                    inputMode="tel"
                                    onChange={(e) => setLocalPhone(e.target.value)}
                                    className="input input-bordered input-primary flex-1 h-16 text-2xl"
                              />
                        </div>

                        <div className="flex items-center justify-end pt-2">
                              <button
                                    onClick={ () => { setStep(2) } }
                                    type="button"
                                    className={`inline-flex items-center gap-2 px-5 py-2.5 rounded-2xl shadow-sm bg-orange-400 cursor-pointer font-bold text-white border-0 opacity-80 hover:opacity-100"}`}>
                                    <Lock className="w-4 h-4" /> Continue to verification
                              </button>
                        </div>
                  </div>	
            );
      }
