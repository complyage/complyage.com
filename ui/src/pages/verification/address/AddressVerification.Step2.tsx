//||------------------------------------------------------------------------------------------------||
//|| AddressVerification.Step2
//|| src/components/security/AddressVerification.Step2.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React from "react";
      import { CheckCircle, CreditCard, Lock }                    from "lucide-react";
      import {getEnv}                                             from "../../../data/getEnv";

      //||------------------------------------------------------------------------------------------------||
      //|| Step Props
      //||------------------------------------------------------------------------------------------------||

      import { StepProps }                                        from "./AddressVerification";
      import { VerificationAddress }                              from "../../../interfaces/verify/address/process";

      //||------------------------------------------------------------------------------------------------||
      //|| Currency
      //||------------------------------------------------------------------------------------------------||

      import { formatStripeAmount, toStripeAmount }               from "../../../data/getCurrencies";

      //||------------------------------------------------------------------------------------------------||
      //|| Components
      //||------------------------------------------------------------------------------------------------||

      import InlineAlert                                          from "../../../components/base/InlineAlert";
      import Donation                                             from "../../../components/dynamic/Donate";
      import CurrencyDropdown, { CurrencyCode }                   from "../../../components/dynamic/Currency";

      //||------------------------------------------------------------------------------------------------||
      //|| Default
      //||------------------------------------------------------------------------------------------------||

      export default function AddressVerificationStep2({
            process,
            updateProcess,
      }: StepProps) {

            //||------------------------------------------------------------------------------------------------||
            //|| Confirm
            //||------------------------------------------------------------------------------------------------||

            const [confirmed, setConfirmed] = React.useState(false);

            //||------------------------------------------------------------------------------------------------||
            //|| Call API to generate the Verification Charge
            //||------------------------------------------------------------------------------------------------||

            const onSubmit = async() : Promise<string | null> => {
                  //||------------------------------------------------------------------------------------------------||
                  //|| Generate Amount
                  //||------------------------------------------------------------------------------------------------||
                  let amount = 0;
                  try { 
                        amount += Number(getEnv("VITE_VERIFICATION_ADDRESS_AMOUNT"));
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
                        "baseAmount"            : process.baseAmount, 
                        "donationAmount"        : process.donationAmount || 0,
                        "totalAmount"           : amount,
                        "currency"              : process.currency,
                        "address"               : process.verifyAddress || {},

                  });
                  console.log("Calling API with data:", processJSON);
                  //||------------------------------------------------------------------------------------------------||
                  //|| Create the Payload
                  //||------------------------------------------------------------------------------------------------||
                  try {
                        console.log("Calling /v1/api/verify/address/init", processJSON);
                        const response = await fetch("/v1/api/verify/address/init", {
                              method: "POST",
                              headers: { "Content-Type": "application/json" },
                              body: processJSON
                        });
                        //||------------------------------------------------------------------------------------------------||
                        //|| Error
                        //||------------------------------------------------------------------------------------------------||
                        if (!response.ok) {
                              const errorText = await response.text();
                              return errorText || "API call failed with an unknown error.";
                        }
                        //||------------------------------------------------------------------------------------------------||
                        //|| Got it. Update the Process
                        //||------------------------------------------------------------------------------------------------||
                        const data = await response.json();
                        console.log("API response data:", data);
                        const update: Partial<VerificationAddress> = {
                              clientSecret      : data.data.clientSecret,
                              step              : 3,
                              chargeAmount      : data.data.amount || 99,
                              currency          : process.currency,
                              verificationUUID  : data.data.identifier,
                        };
                        updateProcess({ ...process, ...update });
                        return null
                  } catch (error) {
                        console.error("API call failed:", error);
                        return "API call failed. Please try again later.";
                  }
            }    

            //||------------------------------------------------------------------------------------------------||
            //|| Vars
            //||------------------------------------------------------------------------------------------------||

            const addr = process.verifyAddress || {};

            //||------------------------------------------------------------------------------------------------||
            //|| Helpers
            //||------------------------------------------------------------------------------------------------||

            const setDonation = (amount: number) => {
                  if (amount < 0) amount = 0;
                  updateProcess({ ...process, donationAmount: amount });
            };

            const setCurrency = (currency: CurrencyCode) => {
                  updateProcess({ ...process, currency: currency });
            };            

            //||------------------------------------------------------------------------------------------------||
            //|| Render
            //||------------------------------------------------------------------------------------------------||

            return (
                  <div className="space-y-5 p-5">
                        <div>
                              <div className="flex items-center justify-between border-b border-gray-400 py-2 mb-4">
                                    <span className="font-bold">Please verify your address is correct:</span>
                              </div>
                        </div>

                        <div className="flex flex-col bg-black/20 rounded-lg p-4">
                              <div>
                                    <div className="text-2xl font-bold">
                                          <div>{addr.line1}</div>
                                          {addr.line2 ? <div>{addr.line2}</div> : null}
                                          <div>
                                                {addr.city}, {addr.state} {addr.postal}
                                          </div>
                                          <div>{addr.country}</div>
                                    </div>
                              </div>                              
                              {!confirmed ? (
                                    <div className="flex items-center gap-2 mt-3 justify-end">
                                          <button onClick={() => setConfirmed(!confirmed)} className="mt-3 btn btn-secondary ">
                                          <Lock className="w-4 h-4" />
                                          I confirm this address is correct
                                    </button>
                              </div>) : (
                                    <div className="flex items-center justify-end ">
                                          <span className="flex gap-3 items-center bg-black/30 p-2 rounded-lg text-orange-400">
                                          <CheckCircle className="w-4 h-4" />
                                          Confirmed
                                          </span>
                                    </div>
                              )}
                        </div>

				<div>
					<div className="flex items-center justify-between border-b border-gray-400 py-2 mb-4">
						<span className="font-bold">Summary</span>
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
                                    <div className="text-sm font-medium">{ formatStripeAmount(process.baseAmount, process.currency)  } { process.currency }</div>
                              </div>
                              <div className="flex items-center justify-between py-1">
                                    <div className="text-sm">Donation</div>
                                    <div className="text-sm font-medium">{ formatStripeAmount(process.donationAmount, process.currency) } { process.currency }</div>
                              </div>
                              <div className="h-px border my-3 border-gray-500" />
                                    <div className="flex items-center justify-between  py-3">
                                          <div className="text-lg font-bold">Total</div>
                                          <div className="text-lg font-bold">{ formatStripeAmount(process.chargeAmount, process.currency) } { process.currency }</div>
                                    </div>
					</div>
				</div>              

				<div>
					<div className="flex items-center justify-between border-b border-gray-400 py-2 mb-4">
						<div className="font-bold">We run solely on donations, please help keep this private service running!</div>
					</div>
					<Donation currency={ process.currency } donation={process.donationAmount} setDonation={setDonation} />
				</div>

                        <hr className="mt-5 border-gray-500" />
                        
                        {!confirmed && (<InlineAlert message="Please confirm your address." isError={true} />) }

                        <div className="flex items-center justify-between">
                              <button
                                    onClick={() => updateProcess({ ...process, step: 1 })}
                                    className="px-4 py-2 rounded-2xl font-bold bg-black/30 opacity-70 hover:opacity-100"
                              >
                                    Edit Address
                              </button>
                              <button
                                    disabled={!confirmed || !process.chargeAmount}
                                    onClick={() => onSubmit()}
                                    className="inline-flex items-center gap-2 px-5 py-2.5 rounded-2xl shadow-sm bg-orange-400 font-bold text-white border-0 opacity-80 hover:opacity-100"
                              >
                                    Continue
                              </button>
                        </div>
                  </div>
            );
      }
