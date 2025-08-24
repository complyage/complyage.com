//||------------------------------------------------------------------------------------------------||
//|| CCVerification.Step1
//|| src/components/payments/CCVerification.Step1.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      import React                                          from "react";
      import { CreditCard, Lock }                           from "lucide-react";

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      import Donation                                       from "../../../components/dynamic/Donate";
      import CurrencyDropdown, { CurrencyCode }             from "../../../components/dynamic/Currency";

      //||------------------------------------------------------------------------------------------------||
      //|| Interfacss
      //||------------------------------------------------------------------------------------------------||

      import { VerificationCard }                           from "../../../interfaces/verify/card/process";
      
      //||------------------------------------------------------------------------------------------------||
      //|| Data
      //||------------------------------------------------------------------------------------------------||

      import { formatStripeAmount }                        from "../../../data/getCurrencies";
      
      //||------------------------------------------------------------------------------------------------||
      //|| Props
      //||------------------------------------------------------------------------------------------------||

      import { CCStepProps }                                from "./CCVerification";

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      export default function CCVerificationStep1({ process, updateProcess }: CCStepProps) {

            //||------------------------------------------------------------------------------------------------||
            //|| Call API to generate the Verification Charge
            //||------------------------------------------------------------------------------------------------||

            const onSubmit = async() : Promise<string | null> => {
                  //||------------------------------------------------------------------------------------------------||
                  //|| Generate Amount
                  //||------------------------------------------------------------------------------------------------||
                  let amount = 0;
                  try { 
                        amount += process.baseAmount || import.meta.env.VITE_VERIFICATION_CARD_AMOUNT;
                        amount += process.donationAmount || 0;
                  } catch (e) {
                        console.error("Error calculating amount:", e);
                        return "Invalid amount calculation.";
                  }
                  if (amount === 0) return "Amount cannot be zero.";
                  //||------------------------------------------------------------------------------------------------||
                  //|| Create the Payload
                  //||------------------------------------------------------------------------------------------------||
                  const processJSON = JSON.stringify({ 
                        "baseAmount"      : Number(process.baseAmount || 0), 
                        "donationAmount"  : Number(process.donationAmount || 0),
                        "totalAmount"     : Number(amount),
                        "currency"        : process.currency 
                  });
                  console.log("Calling Init with data:", "/v1/api/verify/card/init", processJSON);
                  //||------------------------------------------------------------------------------------------------||
                  //|| Create the Payload
                  //||------------------------------------------------------------------------------------------------||
                  try {
                        const resp = await fetch("/v1/api/verify/card/init", {
                              method: "POST",
                              headers: { "Content-Type": "application/json" },
                              body: processJSON
                        });
                        //||------------------------------------------------------------------------------------------------||
                        //|| Error
                        //||------------------------------------------------------------------------------------------------||
                        if (!resp.ok) {
                              const errorText = await resp.text();
                              return errorText || "API call failed with an unknown error.";
                        }
                        //||------------------------------------------------------------------------------------------------||
                        //|| Got it. Update the Process
                        //||------------------------------------------------------------------------------------------------||
                        const response = await resp.json();
                        const update: Partial<VerificationCard> = {
                              step              : 2,
                              clientSecret      : response.data.clientSecret,
                              currency          : response.data.currency,
                              chargeAmount      : response.data.amount || 99,
                              verificationUUID  : response.data.identifier
                        };
                        console.log("API Init","/v1/api/verify/card/init", update);
                        updateProcess({ ...process, ...update });
                        return null
                  } catch (error) {
                        console.error("API call failed:", error);
                        return "API call failed. Please try again later.";
                  }
            }            

            //||------------------------------------------------------------------------------------------------||
            //|| Set Donation
            //||------------------------------------------------------------------------------------------------||

            const setDonation = (amount: number) => {
                  if (amount < 0) {
                        console.warn("Donation cannot be negative. Setting to 0.");
                        amount = 0;
                  }
                  updateProcess({ ...process, donationAmount: amount });
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Set Currency
            //||------------------------------------------------------------------------------------------------||

            const setCurrency = (currency: CurrencyCode) => {
                  updateProcess({ ...process, currency: currency });
            }

            //||------------------------------------------------------------------------------------------------||
            //|| JSX
            //||------------------------------------------------------------------------------------------------||

            return (
			<div className="space-y-6 p-6">
				<div className="flex items-center gap-2 text-gray-300">
					<CreditCard className="w-5 h-5" />
					<span>Choose how your charge appears and (optionally) add a donation.</span>
				</div>

				<div>
					<div className="flex items-center justify-between border-b border-gray-400 py-2 mb-4">
						<div className="flex items-center gap-2">
							<span className="font-bold">Statement descriptor</span>
						</div>
					</div>

					<div className="text-sm 70 mt-1">
						<span className="block font-bold text-blue-500">{import.meta.env.VITE_MERCHANT_DESCRIPTOR} - [XXXX]</span>
						<span className="block text-gray-300">
							Your bank statement will show a small charge 4-digit code in place of <code>[XXXX]</code>.
						</span>
					</div>
				</div>

				<div>
					<div className="flex items-center justify-between border-b border-gray-400 py-2 mb-4">
						<div className="block items-center gap-2 w-full">
							<span className="font-bold">Summary</span>
						</div>
					</div>
				</div>

				<div className="bg-black/20 rounded">
					<div className="flex items-center justify-between py-2 bg-black/20 mb-2 px-2 shadow-lg border-b border-gray-800">
						<div className="text-sm pl-2">Currency</div>
						<div className="text-sm font-medium pr-1">
							<CurrencyDropdown currency={process.currency} setCurrency={setCurrency} />
						</div>
					</div>
					<div className="px-4">
						<div className="flex items-center justify-between py-1 pt-4">
							<div className="text-sm">Postage and Processing</div>
							<div className="text-sm font-medium">
								{formatStripeAmount(process.baseAmount, process.currency)} {process.currency}
							</div>
						</div>
						<div className="flex items-center justify-between py-1">
							<div className="text-sm">Donation</div>
							<div className="text-sm font-medium">
								{formatStripeAmount(process.donationAmount, process.currency)} {process.currency}
							</div>
						</div>
						<div className="h-px border my-3 border-gray-500" />
						<div className="flex items-center justify-between  py-3">
							<div className="text-lg font-bold">Total</div>
							<div className="text-lg font-bold">
								{formatStripeAmount(process.chargeAmount, process.currency)} {process.currency}
							</div>
						</div>
					</div>
				</div>

				<div>
					<div className="flex items-center justify-between border-b border-gray-400 py-2 mb-4">
						<div className="font-bold">We run solely on donations, please help keep this private service running!</div>
					</div>
					<Donation currency={process.currency} donation={process.donationAmount} setDonation={setDonation} />
				</div>

				<div className="flex items-center justify-between">
					<div className="text-xs opacity-75">{""}</div>
					<button
						onClick={() => {
							onSubmit();
						}}
						type="button"
						className={`inline-flex items-center gap-2 px-5 py-2.5 rounded-2xl shadow-sm bg-orange-400 cursor-pointer font-bold text-white border-0 opacity-80 hover:opacity-100"}`}>
						<Lock className="w-4 h-4" /> Continue to verification
					</button>
				</div>
			</div>
		);
      }
