//||------------------------------------------------------------------------------------------------||
//|| Verification : Card
//|| Handles the Card Verification Process
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Currency
      //||------------------------------------------------------------------------------------------------||

      import { CurrencyAmount }                 from "../../base/transaction";
      import { CurrencyCode }                   from "../../../components/dynamic/Currency";

      //||------------------------------------------------------------------------------------------------||
      //|| Verification Media
      //||------------------------------------------------------------------------------------------------||

      export type VerificationCard = {
            step                  : 1 | 2 | 3;
            baseAmount            : CurrencyAmount;
            chargeAmount          : CurrencyAmount;
            donation              : number;
            currency              : CurrencyCode;
            clientSecret?         : string;
            //Response Based Ignore
            lastFour?             : string;
            cardType?             : string;
            billingZip?           : string;
            transactionId?        : string;
            verificationUUID?    : string;
      }
      