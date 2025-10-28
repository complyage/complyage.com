//||------------------------------------------------------------------------------------------------||
//|| CCVerification.Step3
//|| src/components/payments/CCVerification.Step3.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      import React, { useEffect, useState, useCallback }   from "react";
      import InlineAlert                                   from "../../../components/base/InlineAlert";
      import { CheckCircle, AlertCircle }                  from "lucide-react";

      //||------------------------------------------------------------------------------------------------||
      //|| Hooks
      //||------------------------------------------------------------------------------------------------||

      import { useOverlayNavigate }                        from "../../../hooks/useOverlay";

      //||------------------------------------------------------------------------------------------------||
      //|| Props
      //||------------------------------------------------------------------------------------------------||

      import { CCStepProps }                               from "./CCVerification";

      //||------------------------------------------------------------------------------------------------||
      //|| Data
      //||------------------------------------------------------------------------------------------------||

      import { formatStripeAmount }                        from "../../../data/getCurrencies";

      //||------------------------------------------------------------------------------------------------||
      //|| Component
      //||------------------------------------------------------------------------------------------------||

      export default function CCVerificationStep3({ process, updateProcess }: CCStepProps) {

            //||------------------------------------------------------------------------------------------------||
            //|| Navigate
            //||------------------------------------------------------------------------------------------------||

            const navigate                     = useOverlayNavigate();  

            //||------------------------------------------------------------------------------------------------||
            //|| State
            //||------------------------------------------------------------------------------------------------||

            const [ error, setError ]                       = useState<string | null>(null);

            //||------------------------------------------------------------------------------------------------||
            //|| Generate Descriptor
            //||------------------------------------------------------------------------------------------------||

            const descriptor = import.meta.env.VITE_MERCHANT_DESCRIPTOR + " CODE-441530" || "complyage.com CODE-441530";

            //||------------------------------------------------------------------------------------------------||
            //|| Busy
            //||------------------------------------------------------------------------------------------------||

            return (
			<div className="max-w-2xl mx-auto pt-2 px-4">
				<div className="flex flex-col items-center mb-8">
					<CheckCircle className="w-14 h-14 text-green-400 mb-3" />
					<div className="text-2xl font-bold text-green-300 mb-1">Card charged successfully!</div>
					<div className="text-lg text-gray-300 text-center max-w-lg">
						Please check your credit card statement or online banking.
						<br />
						<span className="text-yellow-300 font-semibold">Look for the 6-digit code in your bank’s pending transactions!</span>
					</div>
				</div>
				{/* Fake billing statement */}
				<div className="bg-[#181818] border border-white/20 rounded-xl shadow-lg px-6 py-6 mb-8">
					<div className="text-xl font-bold text-gray-100 mb-4">Sample Bank Statement</div>
					<table className="w-full text-base text-gray-300">
						<thead>
							<tr>
								<th className="text-left font-semibold pb-2">Date</th>
								<th className="text-left font-semibold pb-2">Description</th>
								<th className="text-right font-semibold pb-2">Amount</th>
							</tr>
						</thead>
						<tbody>
							<tr>
								<td className="py-1">08/11</td>
								<td className="py-1">NETFLIX.COM</td>
								<td className="py-1 text-right">$13.99</td>
							</tr>
							<tr>
								<td className="py-1">08/12</td>
								<td className="py-1">AMAZON MKTPLACE PMTS</td>
								<td className="py-1 text-right">$44.23</td>
							</tr>
                                          <tr className="bg-yellow-400/20 ring-2 ring-yellow-400 rounded-md  shadow-lg scale-105 transition-transform duration-200 z-10 relative px-4 text-[1.4rem] leading-[2.8rem]">
                                                <td className="py-3 font-extrabold text-lg text-yellow-300 pl-2 align-middle">08/13</td>
                                                <td className="py-3 font-extrabold text-xl text-yellow-200 tracking-wide align-middle">{descriptor}</td>
                                                <td className="py-3 text-right font-extrabold text-xl text-yellow-200 pr-2 align-middle"> { formatStripeAmount(process.baseAmount, process.currency) || "0.99"}</td>
                                          </tr>
							<tr>
								<td className="py-1">08/13</td>
								<td className="py-1">UBER TRIP</td>
								<td className="py-1 text-right">$15.80</td>
							</tr>
						</tbody>
					</table>
					<div className="pt-4 text-xs text-gray-400 italic">
						<span>
							Your charge will appear just like this (amount and code may vary). If you don't see the charge, try refreshing your
							bank's page or check pending transactions.
						</span>
					</div>
				</div>
				<div className="flex flex-col items-center">
					<button
						onClick={() => { navigate(`/verification/check?identifier=${process.verificationUUID}`)} }
						className="w-full md:w-2/3 text-lg font-bold py-4 px-8 bg-gradient-to-tr from-yellow-400 to-orange-500 text-black rounded-2xl shadow-lg hover:from-yellow-300 hover:to-orange-400 transition-all">
						Enter My Code
					</button>
					<div className="pt-3 text-xs text-gray-400 max-w-md text-center">
						This will take you to the final step to complete your age verification.
					</div>
				</div>
			</div>
		);
      }
