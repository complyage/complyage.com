
//||------------------------------------------------------------------------------------------------||
//|| Interfaces : Identity
//|| Identity Verification Data
//||------------------------------------------------------------------------------------------------||

import { VerificationTypes }                                from "../models/model.verify";
import { CreditCard }                                       from "../base/transaction";
import { Address }                                          from "../base/geo";
import { DOB, EmailAddress, PhoneNumber, IDCard, Username}  from "../base/user";

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
      approved    : VerificationTypes[];
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
      decrypted   : null | EmailAddress;
      timestamp   : Date;
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Identification
//||------------------------------------------------------------------------------------------------||

interface VerifiedIdentification {
      display     : string;
      data        : string;
      decrypted   : null | IDCard;
      media       : VerifiedMedia[];
      timestamp   : Date
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Age
//||------------------------------------------------------------------------------------------------||

interface VerifiedAge {
      display     : string;
      data        : string;
      decrypted   : null | DOB;
      media       : VerifiedMedia[]
      timestamp   : Date;
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Phone
//||------------------------------------------------------------------------------------------------||

interface VerifiedPhone {
      display     : string;
      data        : string;
      decrypted   : null | PhoneNumber;
      timestamp   : Date;
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Address
//||------------------------------------------------------------------------------------------------||

interface VerifiedAddress { 
      display     : string;
      data        : string;
      decrypted   : null | Address;
      timestamp   : Date;
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Credit Card
//||------------------------------------------------------------------------------------------------||

interface VerifiedCreditCard {
      display     : string;
      data        : string;
      decrypted   : null | CreditCard;
      timestamp   : Date;
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Profile Pic
//||------------------------------------------------------------------------------------------------||

interface VerifiedUsername {
      display     : string;
      data        : string;
      decrypted   : null | Username;
      timestamp   : Date;      
}