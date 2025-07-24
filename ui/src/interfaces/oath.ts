//||------------------------------------------------------------------------------------------------||
//|| OAuth Response Data
//|| Includes Site, Location, User, and Token information
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Interfaces
      //||------------------------------------------------------------------------------------------------||

      import { VerificationTypes }                    from "./types.verification";
      import { VerificationStatuses }                 from "./status.verification";

      //||------------------------------------------------------------------------------------------------||
      //|| Types
      //||------------------------------------------------------------------------------------------------||

      type VerificationTypeValue   = typeof VerificationTypes[keyof typeof VerificationTypes];
      type VerificationStatusValue = typeof VerificationStatuses[keyof typeof VerificationStatuses];

      //||------------------------------------------------------------------------------------------------||
      //|| OAuth Verification
      //||------------------------------------------------------------------------------------------------||

      interface OAuthVerification {
            type             : VerificationTypeValue;
            status           : VerificationStatusValue;
            data             : object | string;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| OAuth User
      //||------------------------------------------------------------------------------------------------||

      interface OAuthRequirements {
            type             : VerificationTypeValue;
            optional         : boolean;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| OAuth User
      //||------------------------------------------------------------------------------------------------||

      interface OAuthUser {
            id                : number;
            status            : string;
            username          : string;
            verifications     : OAuthVerification[]
      }

      //||------------------------------------------------------------------------------------------------||
      //|| OAuth Site
      //||------------------------------------------------------------------------------------------------||

      interface OAuthSite {
            name              : string;
            url               : string;
            logo              : string;
            description       : string;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| OAuth Location
      //||------------------------------------------------------------------------------------------------||

      interface OAuthZone {
            state             : string;
            country           : string;
            ip                : string;
            requirements      : string[];
            description       : string;
            law               : string;
            effectiveDate     : string;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| OAuth Site
      //||------------------------------------------------------------------------------------------------||

      interface OAuthResponse {
            site              : OAuthSite;
            user              : OAuthUser;
            zone              : OAuthZone;
            status            : string;
            requirements      : OAuthRequirements[];
      }

