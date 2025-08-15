//||------------------------------------------------------------------------------------------------||
//|| AddressVerification.Step2
//|| src/components/security/AddressVerification.Step2.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React                        from "react";
      import { Address }                  from "../../../interfaces/verification.location";

      //||------------------------------------------------------------------------------------------------||
      //|| Props
      //||------------------------------------------------------------------------------------------------||
      
      type Props = {
		addr                    : Address;
		std                     : Address;
		notes?                  : string;
		onBack                  : () => void;
		onConfirm               : () => void;
	};

      //||------------------------------------------------------------------------------------------------||
      //|| Default
      //||------------------------------------------------------------------------------------------------||

	export default function AddressVerificationStep2({addr, std, notes, onBack, onConfirm}: Props) {
		return (
			<div className="space-y-5 p-5">
				<div className="bg-black/20 rounded-lg p-4">
					<div className="text-sm opacity-80 mb-2">Please confirm this is the correct address:</div>
						<div>
                                    <span className="text-xs opacity-70 mb-1 border-b py-2 border-gray-400">You entered</span>
                                    <div className="text-md">
                                          <div>{addr.line1}</div>
                                          {addr.line2 ? <div>{addr.line2}</div> : null}
                                          <div>
                                                {addr.city}, {addr.state} {addr.postal}
                                          </div>
                                          <div>{addr.country}</div>
                                    </div>
					</div>
				</div>

				<div className="flex items-center justify-between">
					<button onClick={onBack} className="px-4 py-2 rounded-2xl font-bold bg-black/30 opacity-70 hover:opacity-100">
						Edit Address
					</button>
					<button
						onClick={onConfirm}
						className="inline-flex items-center gap-2 px-5 py-2.5 rounded-2xl shadow-sm bg-orange-400 font-bold text-white border-0 opacity-80 hover:opacity-100">
						Continue
					</button>
				</div>
			</div>
		);
	}
