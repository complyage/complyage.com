
//||------------------------------------------------------------------------------------------------||
//|| Interfaces : Identity
//|| Identity Verification Data
//||------------------------------------------------------------------------------------------------||

import { DOB }          from "../../base/user";

//||------------------------------------------------------------------------------------------------||
//|| Identity Username
//||------------------------------------------------------------------------------------------------||

export interface IdentityUsername {
      idSite            : number;
      username          : string;
      verification      : string;
}

//||------------------------------------------------------------------------------------------------||
//|| Record
//||------------------------------------------------------------------------------------------------||

export interface IdentityRecord {
      verified?         : boolean;
      age?              : number;
      dob?              : DOB;
      display           : string;
      verification      : string;
}

//||------------------------------------------------------------------------------------------------||
//|| Full Identity
//||------------------------------------------------------------------------------------------------||

export interface Identity {
      MAIL?       : IdentityRecord;      
      PHNE?       : IdentityRecord;
      ADDR?       : IdentityRecord;
      CRCD?       : IdentityRecord;
      IDEN?       : IdentityRecord;
      FACE?       : IdentityRecord;
      usernames?  : Record<string, IdentityUsername>;
      approved?   : string[];
}
