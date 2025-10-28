//||------------------------------------------------------------------------------------------------||
//|| AddressVerification.Step3
//|| src/components/security/AddressVerification.Step3.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Hooks
      //||------------------------------------------------------------------------------------------------||

      import { useOverlayNavigate }                        from "../../../hooks/useOverlay";
      import { CheckCircle }                               from "lucide-react";

      //||------------------------------------------------------------------------------------------------||
      //|| Props
      //||------------------------------------------------------------------------------------------------||

      import { StepProps }                                 from "./AddressVerification";
      
      //||------------------------------------------------------------------------------------------------||
      //|| Component
      //||------------------------------------------------------------------------------------------------||

      export default function AddressVerificationStep4({ process, updateProcess }: StepProps) {

            //||------------------------------------------------------------------------------------------------||
            //|| Navigate
            //||------------------------------------------------------------------------------------------------||

            const navigate                     = useOverlayNavigate();  

            //||------------------------------------------------------------------------------------------------||
            //|| Render
            //||------------------------------------------------------------------------------------------------||

            return (
                  <div className="space-y-6 p-6">
				<div className="flex flex-col items-center mb-8">
					<CheckCircle className="w-14 h-14 text-green-400 mb-3" />
					<div className="text-2xl font-bold text-green-300 mb-1">Postcard Sent!</div>
				</div>                        
                        <div className="bg-black/20 rounded-lg px-6 py-6 text-lg text-white shadow-md flex flex-col items-center">
                              <p className="mb-4 text-base text-gray-200 text-center">
                                    To complete your address verification, we’ll mail a <span className="font-semibold text-yellow-300">postcard</span> to the address you provided. This postcard will contain a unique <span className="font-semibold text-yellow-300">6-digit code</span>.
                              </p>
                              <p className="mb-4 text-base text-gray-200 text-center">
                                    Once you receive your postcard, simply log into your dashboard, click on <span className="font-semibold text-white">Address Verify</span>, and select the pending address record. Enter your code to confirm your address and complete the verification.
                              </p>
                              <p className="text-sm text-gray-400 text-center">
                                    <span className="font-semibold">Note:</span> Mailing times may vary depending on your location. Please allow several business days for delivery.
                              </p>                              
                        </div>
                        <div className="flex justify-center">
                              <button
                                    className="btn btn-secondary"
                                    onClick={() => {
                                          navigate("/verification/init/?type=ADDR")
                                    }}
                              >
                                    Back to Dashboard
                              </button>
                        </div>
                  </div>
		);
      }
