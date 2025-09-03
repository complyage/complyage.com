
//||------------------------------------------------------------------------------------------------||
//|| Interfaces : Identity
//|| Identity Verification Data
//||------------------------------------------------------------------------------------------------||

import { VerificationTypes }                                from "../models/model.verify";
import { CreditCard }                                       from "../base/transaction";
import { Address }                                          from "../base/geo";
import { DOB, EmailAddress, PhoneNumber, IDCard, Username}  from "../base/user";

//||------------------------------------------------------------------------------------------------||
//|| Identity (Epic Verified Profile) - TypeScript Interfaces
//||------------------------------------------------------------------------------------------------||

export interface IdentityUsername {
      idSite            : number;
      username          : string;
      verification      : string;
}

export interface IdentityRecord {
      display           : string;
      verification      : string;
}

export interface Identity {
      email?            : IdentityRecord;
      age?              : IdentityRecord;
      phone?            : IdentityRecord;
      address?          : IdentityRecord;
      creditCard?       : IdentityRecord;
      idCard?           : IdentityRecord;
      usernames?        : Record<string, IdentityUsername>;
      approved?         : string[];
      verified          : boolean;
      verifiedAge?      : number;
}
