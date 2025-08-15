//||------------------------------------------------------------------------------------------------||
//|| Interfaces : Verification
//|| Model Verification
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Model
      //||------------------------------------------------------------------------------------------------||

      export interface ModelVerification {
            idVerification      : number;
            fidAccount          : number;
            verificationUUID?   : string | null;
            verificationYype?   : string | null;
            verificationData?   : Buffer | null;
            verificationMeta?   : string | null;
            verificationStatus? : string | null;
            createdAt           : string; // ISO datetime string
            updatedAt           : string; // ISO datetime string
      }

      //||------------------------------------------------------------------------------------------------||      
      //|| Types
      //||------------------------------------------------------------------------------------------------||

      export type VerificationTypes = "MAIL" | "PHNE" | "UAGE" | "ADDR" | "CRCD" | "UNAM";

      //||------------------------------------------------------------------------------------------------||      
      //|| Status
      //||------------------------------------------------------------------------------------------------||

      export type VerificationStatuses = "PEND" | "PEVF" | "APPR" | "VERF" | "RJCT" | "ESCL" | "EXPD" | "CNCL" | "MISS";
