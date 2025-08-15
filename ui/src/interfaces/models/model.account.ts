//||------------------------------------------------------------------------------------------------||
//|| Interfaces : Account
//|| Model Account
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||      
      //|| Model Account
      //||------------------------------------------------------------------------------------------------||

      export interface ModelAccount {
            id_account           : number;
            account_type?        : string | null;
            account_salt?        : string | null;
            account_username?    : string | null;
            account_email?       : string | null;
            account_password?    : string | null;
            account_security?    : number | null;
            account_public?      : string | null;
            account_private?     : string | null;
            account_private_hash?: string | null;
            account_status?      : string | null;
            account_level?       : number | null;
            account_advanced?    : number | null;
            account_identity?    : string | null;
      }

      //||------------------------------------------------------------------------------------------------||      
      //|| Types
      //||------------------------------------------------------------------------------------------------||

      export type AccountStatusTypes = "NONE" | "PNEW" | "PEND" | "VERF" | "ACTV" | "BNND" | "RMVD";

