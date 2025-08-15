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
            //|| State
            //||------------------------------------------------------------------------------------------------||

            const [ error, setError ]                       = useState<string | null>(null);
            const [ busy, setBusy ]                         = useState<boolean>(true);

            //||------------------------------------------------------------------------------------------------||
            //|| Refresh Handler
            //||------------------------------------------------------------------------------------------------||

            const handleLoad = useCallback(async () => {
                  setBusy(true);
                  setError(null);
                  //||------------------------------------------------------------------------------------------------||
                  //|| Payload
                  //||------------------------------------------------------------------------------------------------||
                  const payload : {
                        amount        : number;
                        currency      : string;
                        uuid          : string;
                        billingZip?   : string;
                        clientSecret  : string;
                        transactionId?: string;
                  } = {
                        amount        : process.chargeAmount.amount,
                        billingZip    : process.billingZip || "",
                        currency      : process.chargeAmount.currency,
                        uuid          : process.verificationUUID || "",
                        clientSecret  : process.clientSecret || "",
                        transactionId : process.transactionId || "",
                  };
                  //||------------------------------------------------------------------------------------------------||
                  //|| Pull
                  //||------------------------------------------------------------------------------------------------||
                  try {
                        const res = await fetch("/v1/api/verify/card/success", {
                              method      : "POST",
                              headers     : { "Content-Type": "application/json" },
                              body        : JSON.stringify(payload),
                        });
                        if (!res.ok) {
                              let msg     = "Unknown error";
                              try { msg = (await res.json()).message || msg; } catch {}
                              setError(msg);
                              setBusy(false);
                              return;
                        }
                        setBusy(false);
                  } catch (err: any) {
                        setError("Network error, please try again.");
                        setBusy(false);
                  }
            }, [process]);

            //||------------------------------------------------------------------------------------------------||
            //|| Auto-Call on Mount/Update
            //||------------------------------------------------------------------------------------------------||

            useEffect(() => { handleLoad(); }, [handleLoad]);

            //||------------------------------------------------------------------------------------------------||
            //|| Handlers
            //||------------------------------------------------------------------------------------------------||

            const onEnterCode = () => {
                  window.location.href = `/verification/card/verify?uuid=${process.verificationUUID}`;
            };

            //||------------------------------------------------------------------------------------------------||
            //|| Generate Descriptor
            //||------------------------------------------------------------------------------------------------||

            const descriptor = import.meta.env.VITE_MERCHANT_DESCRIPTOR + " CODE-441530" || "complyage.com CODE-441530";

            //||------------------------------------------------------------------------------------------------||
            //|| Busy
            //||------------------------------------------------------------------------------------------------||

            if (busy) {
                  return (
                        <div className="flex flex-col items-center gap-4 p-12">
                              <CheckCircle className="w-10 h-10 text-yellow-300 animate-spin" />
                              <div className="text-gray-200 font-medium">Finalizing your payment...</div>
                        </div>
                  );
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Error
            //||------------------------------------------------------------------------------------------------||

            if (error) {
                  return (
                        <div className="max-w-xl mx-auto p-10 flex flex-col items-center gap-6">
                              <AlertCircle className="w-12 h-12 text-red-500" />
                              <InlineAlert message={error} isError={true} />
                              <button 
                                    onClick={handleLoad} 
                                    className="px-6 py-3 bg-red-500 text-white rounded-lg hover:bg-red-600 transition-colors">
                                    Retry
                              </button>
                              <textarea defaultValue={JSON.stringify(process, null, 2)} className="h-96 w-full" readOnly />

                        </div>
                  );
            }

            return (
			<div className="max-w-2xl mx-auto py-8 px-4">
                        <textarea defaultValue={JSON.stringify(process, null, 2)} className="h-96 w-full" readOnly />
				<div className="flex flex-col items-center mb-8">
					<CheckCircle className="w-14 h-14 text-green-400 mb-3" />
					<div className="text-2xl font-bold text-green-300 mb-1">Card charged successfully!</div>
					<div className="text-lg text-gray-300 text-center max-w-lg">
						Please check your credit card statement or online banking.
						<br />
						<span className="text-yellow-300 font-semibold">Look for the 6-digit code in your bank’s pending transactions!</span>
					</div>
				</div>
				<div className="bg-yellow-100/10 border border-yellow-400/40 rounded-lg px-5 py-4 mb-6 shadow">
					<p className="text-yellow-300 text-base font-medium mb-1">
						<span className="font-bold">Good news:</span> We’ll <span className="font-bold">refund</span> this charge as soon as you
						complete verification.
					</p>
					<p className="text-sm text-yellow-100/80">You’ll only see a temporary charge on your statement.</p>
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
                                                <td className="py-3 text-right font-extrabold text-xl text-yellow-200 pr-2 align-middle"> { formatStripeAmount(process.chargeAmount.amount, process.chargeAmount.currency) || "0.99"}</td>
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
						onClick={onEnterCode}
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
