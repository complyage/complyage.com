//||------------------------------------------------------------------------------------------------||
//|| PhoneVerification.Step1
//|| src/components/payments/PhoneVerification.Step1.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      import React, { useEffect }                           from "react";
      import { CheckCircle, RefreshCw }                     from "lucide-react";
      import { onlyDigits }                                 from "../../../utils/clean";

      //||------------------------------------------------------------------------------------------------||
      //|| Interfacss
      //||------------------------------------------------------------------------------------------------||

      import { VerificationPhone }                          from "../../../interfaces/verify/phone/process";

      //||------------------------------------------------------------------------------------------------||
      //|| format
      //||------------------------------------------------------------------------------------------------||

      import { formatSimplePhone }                          from "../../../utils/phoneUtils";

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      interface PhoneStep2Props {
            process           : VerificationPhone;
            updateProcess     : (process: Partial<VerificationPhone>) => void;
      }


      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      export default function PhoneVerificationStep2({ process, updateProcess }: PhoneStep2Props) {
            //||------------------------------------------------------------------------------------------------||
            //|| React
            //||------------------------------------------------------------------------------------------------||

            const [cooldown, setCooldown] = React.useState(30);

            //||------------------------------------------------------------------------------------------------||
            //|| React
            //||------------------------------------------------------------------------------------------------||

            const handleResend = () => {
                  fetch(`/v1/api/verify/phone/resend`, {
                        method: "POST",
                        headers: {
                              "Content-Type": "application/json",
                        },
                        body: JSON.stringify(payload),
                  })
                  .then(response => response.json())
                  .then(data => {
                        console.log("Verification data sent:", data);
                        if (data.success) {
                              updateProcess({
                                    verificationUUID: data.data.uuid                                    
                              });
                              setCooldown(30); // Reset cooldown
                        } else {
                              console.error("Failed to send verification code:", data.message);
                        }
                  });
            }            
            //||------------------------------------------------------------------------------------------------||
            //|| Cooldown
            //||------------------------------------------------------------------------------------------------||

            useEffect(() => {
                  console.log(process.verificationUUID);
                  if (process.verificationUUID && process.verificationUUID !== "") return;
                  const payload = {
                        countryCode: process.countryCode,
                        phone: process.phoneNumber,
                  }
                  fetch(`/v1/api/verify/phone`, {
                        method: "POST",
                        headers: {
                              "Content-Type": "application/json",
                        },
                        body: JSON.stringify(payload),
                  })
                  .then(response => response.json())
                  .then(data => {
                        console.log("Verification data sent:", data);
                        if (data.success) {
                              updateProcess({
                                    verificationUUID: data.data.identifier
                              });
                              setCooldown(30); // Reset cooldown
                        } else {
                              console.error("Failed to send verification code:", data.message);
                        }
                  });
            }, []);

            //||------------------------------------------------------------------------------------------------||
            //|| Cooldown
            //||------------------------------------------------------------------------------------------------||
            
            useEffect(() => {
                  if (cooldown > 0) {
                        const id = setTimeout(() => setCooldown(cooldown - 1), 1000);
                        return () => clearTimeout(id);
                  }
                  return () => {};
            }, [cooldown]);
            
            //||------------------------------------------------------------------------------------------------||
            //|| React
            //||------------------------------------------------------------------------------------------------||

            return (
                  <>
			<div className="flex w-full flex-col items-center justify-center min-h-[300px] py-4 px-4 bg-black/20">
				<CheckCircle className="w-14 h-14 text-green-400 mb-4" />
				<div className="text-xl font-bold text-gray-100 mb-2">Verification Code Sent!</div>
				<div className="text-md text-gray-300 text-center mb-4">
					We’ve sent a 6-digit code to
					<br />
					<span className="font-mono text-yellow-300 text-lg">
						{formatSimplePhone(String(process.phoneNumber), String(process.countryCode))}
					</span>
					<br />
					When you receive it, click Continue below.
				</div>
				<div className="flex flex-row gap-4 w-full justify-center items-center mt-4 px-4">
					<button
						className="btn btn-primary text-2xl py-2 font-bold transition h-auto w-auto px-4"
						onClick={() => updateProcess({step: 1})}
						type="button">
						Back
					</button>
					<button
						className="btn btn-secondary text-2xl py-2 font-bold transition h-auto px-4"
						onClick={ () => { 
                                          window.location.href = `/verification/check?identifier=${process.verificationUUID}`
                                    }}
						type="button">
						Continue
					</button>
                              </div>
                        <div className="flex flex-row gap-4 w-full justify-center items-center mt-4 border-t border-gray-600 pt-4">
					<button
						className={`btn btn-outline flex items-center gap-2 px-6 py-3 font-semibold text-gray-500 transition ${cooldown > 0 ? "opacity-60 cursor-not-allowed" : "hover:bg-black/20 "}`}
						disabled={cooldown > 0}
						onClick={handleResend}
						type="button">
						<RefreshCw className="w-4 h-4" />
						{cooldown > 0 ? `Resend (${cooldown})` : "Resend"}
					</button>
				</div>
			</div>
                        <span className="opacity-10 block pt-10 text-center">UUID : { process.verificationUUID }</span>                  
                        </>
		);
      }

