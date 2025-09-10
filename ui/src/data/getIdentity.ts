//||------------------------------------------------------------------------------------------------||
//|| Identity Data
//|| src/data/getIdentity.ts
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Country Type
      //||------------------------------------------------------------------------------------------------||

      import type { Identity, IdentityRecord }   from "../interfaces/verify/identity/identity"; 

      //||------------------------------------------------------------------------------------------------||
      //|| Is verified
      //||------------------------------------------------------------------------------------------------||

      export function isVerified(code: string, identity : Identity): boolean {
            const record = getIdentityVerification(code, identity);
            if (!record || record === null) return false;
            return (record.verification && record.verification !== "") ? true : false;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Get The Verification
      //||------------------------------------------------------------------------------------------------||

      export function getIdentityVerification(code: string, identity : Identity): IdentityRecord | null {
            switch(code) { 
                  case "MAIL": return identity.MAIL ? identity.MAIL : null;
                  case "PHNE": return identity.PHNE ? identity.PHNE : null;
                  case "ADDR": return identity.ADDR ? identity.ADDR : null;
                  case "CRCD": return identity.CRCD ? identity.CRCD : null;
                  case "IDEN": return identity.IDEN ? identity.IDEN : null;
                  case "FACE": return identity.FACE ? identity.FACE : null;
                  default:     
                  return null;
            }
      }
