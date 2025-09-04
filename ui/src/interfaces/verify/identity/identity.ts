
//||------------------------------------------------------------------------------------------------||
//|| Interfaces : Identity
//|| Identity Verification Data
//||------------------------------------------------------------------------------------------------||

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
      display           : string;
      verification      : string;
}

//||------------------------------------------------------------------------------------------------||
//|| Full Identity
//||------------------------------------------------------------------------------------------------||

export interface Identity {
      email?            : IdentityRecord;
      age?              : IdentityRecord;
      phone?            : IdentityRecord;
      address?          : IdentityRecord;
      creditCard?       : IdentityRecord;
      idCard?           : IdentityRecord;
      face?             : IdentityRecord;
      usernames?        : Record<string, IdentityUsername>;
      approved?         : string[];
      verified          : boolean;
      verifiedAge?      : number;
}
