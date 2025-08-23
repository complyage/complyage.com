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
      //|| Data
      //||------------------------------------------------------------------------------------------------||

      import { getAllCountries }                             from "../../../data/getCountries";
      
      //||------------------------------------------------------------------------------------------------||
      //|| Components
      //||------------------------------------------------------------------------------------------------||

      import PhoneNumber                                     from "../../../components/base/PhoneNumber";

      //||------------------------------------------------------------------------------------------------||
      //|| Interfacss
      //||------------------------------------------------------------------------------------------------||

      import { Country }                                    from "../../../interfaces/base/geo";
      import { VerificationPhone }                          from "../../../interfaces/verify/phone/process";

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      interface PhoneStep1Props {
            process              : VerificationPhone;
            updateProcess        : (process: Partial<VerificationPhone>) => void;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      export default function PhoneVerificationStep1({ process, updateProcess }: PhoneStep1Props) {

            const countries = getAllCountries() as Country[];

            return (
                  <div className="space-y-4 p-5">

                        <div className="flex gap-2 mx-auto items-center justify-center text-center text-gray-200">
                              <PhoneNumber
                                    value={{ countryCode: process.countryCode || "1", phoneNumber: process.phoneNumber || "" }}
                                    onChange={(val) => updateProcess({ countryCode: val.countryCode, phoneNumber: val.phoneNumber })}
                              />                           
                        </div>

                        <div className="flex items-center justify-center pt-5 border-t border-gray-700 mt-5 w-full">
                              <button
                                    onClick={ () => updateProcess({ step: 2 }) }
                                    type="button"
                                    className={`inline-flex items-center gap-2 px-5 py-3 rounded-2xl text-2xl h-auto shadow-sm bg-orange-400 cursor-pointer font-bold text-white border-0 opacity-80 hover:opacity-100"}`}>
                                    <Phone className="w-4 h-4" /> Continue to verification
                              </button>
                        </div>
                  </div>	
            );
      }
