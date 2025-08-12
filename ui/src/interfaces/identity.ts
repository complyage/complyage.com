
import { VerificationTypeValues } from "./types.verification";

//||------------------------------------------------------------------------------------------------||
//|| Verified Profile
//||------------------------------------------------------------------------------------------------||

interface Identity { 
      email       : VerifiedEmail           | "MISS",
      age         : VerifiedAge             | "MISS",
      phone       : VerifiedPhone           | "MISS",
      address     : VerifiedAddress         | "MISS",
      creditCard  : VerifiedCreditCard      | "MISS",
      usernames   : Record<number, VerifiedUsername>
      approved    : VerificationTypeValues[];
}
      
//||------------------------------------------------------------------------------------------------||
//|| Global Media
//||------------------------------------------------------------------------------------------------||

interface VerifiedMedia {
      display   : string;
      data      : string;
      decrypted : null | { 
            hash      : string;
            size      : number;
      }
      timestamp : Date;
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Email
//||------------------------------------------------------------------------------------------------||

interface VerifiedEmail {
      display     : string;
      data        : string;
      decrypted   : null | string;
      timestamp   : Date;
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Identification
//||------------------------------------------------------------------------------------------------||

interface VerifiedIdentification {
      display     : string;
      data        : string;
      decrypted   : null | {
            name : {
                  "first"     : string;
                  "last"      : string;
                  "middle"    : string;
            },
            dob  : {
                  month       : number;
                  day         : number;
                  year        : number;
            },
            address : {
                  street1     : string;
                  street2     : string;
                  city        : string;
                  state       : string;
                  zip         : string;
                  country     : string;
            },
            idNumber    : string;
      }
      timestamp   : Date
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Age
//||------------------------------------------------------------------------------------------------||

interface VerifiedAge {
      display     : string;
      data        : string;
      decrypted   : null | {
            month       : number;
            day         : number;
            year        : number;
      };
      media       : VerifiedMedia[]
      timestamp   : Date;
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Phone
//||------------------------------------------------------------------------------------------------||

interface VerifiedPhone {
      display     : string;
      data        : string;
      decrypted   : null | {
            countryCode : string;
            number      : string;
      };
      timestamp   : Date;
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Address
//||------------------------------------------------------------------------------------------------||

interface VerifiedAddress { 
      display     : string;
      data        : string;
      decrypted   : null | {
            street1     : string;
            street2     : string;
            city        : string;
            state       : string;
            zip         : string;
            country     : string;
      };
      timestamp   : Date;
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Credit Card
//||------------------------------------------------------------------------------------------------||

interface VerifiedCreditCard {
      display     : string;
      data        : string;
      decrypted   : null | {
            last4         : string;
            type          : string;
            expires : {
                  month       : number;
                  year        : number;
            };
            transactionId : string;
      }
      timestamp : Date;
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Profile Pic
//||------------------------------------------------------------------------------------------------||

interface VerifiedUsername {
      display     : string;
      data        : string;
      decrypted   : null | { 
            reference   : VerifiedMedia
            media       : VerifiedMedia[]
            siteId      : number;
            username    : string;
      }
      timestamp   : Date;      
}