//||------------------------------------------------------------------------------------------------||
//|| AddressVerification.Step3
//|| src/components/security/AddressVerification.Step3.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React, { useState, useCallback }              from "react";
      import { useStripe, useElements }                    from "@stripe/react-stripe-js";
      
      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||
      
      import { Address }                                   from "../../../interfaces/base/geo";

      //||------------------------------------------------------------------------------------------------||
      //|| Props
      //||------------------------------------------------------------------------------------------------||

      import { StepProps }                                 from "./AddressVerification";
      
      //||------------------------------------------------------------------------------------------------||
      //|| Components
      //||------------------------------------------------------------------------------------------------||

      import StripePaymentForm                             from "../../../components/dynamic/StripePaymentForm";
      import SpinnerCircle                                 from "../../../components/base/SpinnerCircle";

      //||------------------------------------------------------------------------------------------------||
      //|| Component
      //||------------------------------------------------------------------------------------------------||

      export default function CCVerificationStep2({ process, updateProcess }: StepProps) {

            //||------------------------------------------------------------------------------------------------||
            //|| Stripe
            //||------------------------------------------------------------------------------------------------||

            const stripe                              = useStripe();
            const elements                            = useElements();

            //||------------------------------------------------------------------------------------------------||
            //|| State
            //||------------------------------------------------------------------------------------------------||

            const [processError, setProcessError]     = useState<string | null>(null);
            const [busy, setBusy]                     = useState(false);
            const [zip, setZip]                       = useState(process.billingZip || "");

            //||------------------------------------------------------------------------------------------------||
            //|| onSubmit
            //||------------------------------------------------------------------------------------------------||

            const handleSuccess = useCallback(async (transactionId : string) => {
                  setBusy(true);
                  setProcessError(null);
                  updateProcess({ ...process, 
                        transactionId: transactionId,
                        billingZip: zip,
                        step: 4,
                        lastFour: process.lastFour || "",
                        cardType: process.cardType || "",
                  });
                  //||------------------------------------------------------------------------------------------------||
                  //|| Payload
                  //||------------------------------------------------------------------------------------------------||
                  const payload : {
                        identifier    : string;
                        amount        : number;
                        currency      : string;
                        uuid          : string;
                        billingZip?   : string;
                        clientSecret  : string;
                        transactionId?: string;
                        address       : Address;
                  } = {
                        identifier    : process.verificationUUID || "",
                        amount        : process.chargeAmount,
                        billingZip    : process.billingZip || "",
                        currency      : process.currency,
                        uuid          : process.verificationUUID || "",
                        clientSecret  : process.clientSecret || "",
                        transactionId : transactionId,
                        address       : process.verifyAddress
                  };
                  //||------------------------------------------------------------------------------------------------||
                  //|| Pull
                  //||------------------------------------------------------------------------------------------------||
                  try {
                        console.log("Payment success","/v1/api/verify/card/success", payload);
                        const res = await fetch("/v1/api/verify/card/success", {
                              method      : "POST",
                              headers     : { "Content-Type": "application/json" },
                              body        : JSON.stringify(payload),
                        });
                        if (!res.ok) {
                              let msg     = "Unknown error";
                              try { msg = (await res.json()).message || msg; } catch {}
                              setProcessError(msg);
                              setBusy(false);
                              return;
                        }
                        const response = await res.json();
                        setBusy(false);
                        console.log("Payment success complete", response);
                  } catch (err: any) {
                        console.error("Error during payment success:", err);
                        setBusy(false);
                  }
            }, [process]);

            //||------------------------------------------------------------------------------------------------||
            //|| Render
            //||------------------------------------------------------------------------------------------------||

            return (
			<div className="space-y-6 p-6">
                        { busy && ( <SpinnerCircle /> ) }
                        {!busy && ( <StripePaymentForm
                              clientSecret={process.clientSecret || ""} 
                              chargeAmount={process.chargeAmount}
                              donationAmount={process.donationAmount}
                              baseAmount={process.baseAmount}
                              currency={process.currency}
                              onSuccess={(paymentIntentId) => {  handleSuccess(paymentIntentId); }}
                              onBack={() => { updateProcess({ ...process, step: 2 }); }}
                              setZip={setZip}
                              billingZip={zip}
                        />) }
                        <span className="opacity-10 block pt-10 text-center">{ process.verificationUUID }</span>
			</div>
		);
      }
