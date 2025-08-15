//||------------------------------------------------------------------------------------------------||
//|| Verification : Card
//|| Handles the Card Verification Process
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Currency
      //||------------------------------------------------------------------------------------------------||

      import { CurrencyAmount }                 from "../../currency/currency.amount";
      import { CurrencyCode }                   from "../../../components/dynamic/Currency";

      //||------------------------------------------------------------------------------------------------||
      //|| Verification Media
      //||------------------------------------------------------------------------------------------------||

      export type VerificationAddress = {
            baseAmount            : CurrencyAmount;
            chargeAmount          : CurrencyAmount;
            donation              : number;
            currency              : CurrencyCode;
            cardNumber?           : string;
            expMonth?             : string;
            expYear?              : string;
            cvc?                  : string;            
            billingZip?           : string;
            //Response Based Ignore
            lastFour?             : string;
            cardType?             : string;
            transactionId?        : string;
            verificationUUID?    : string;
      }
      