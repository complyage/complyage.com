//||------------------------------------------------------------------------------------------------||
//|| Verification : ID
//|| Handles the ID Verification Process
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Verification Media
      //||------------------------------------------------------------------------------------------------||

      export interface VerificationMedia {
            type            : "image" | "video";
            section         : "front" | "back" | "selfie";
            blob            : Blob;
            mime            : string;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Verification Media
      //||------------------------------------------------------------------------------------------------||

      export interface VerificationProcessID {
            identifier      : string;
            status          : "pending" | "approved" | "rejected" | "expired";
            level           : "basic" | "advanced" | "human";
            idType          : "passport" | "drivers" | "national" | "other";
            error           : null | string;
            front           : VerificationMedia | null;
            back            : VerificationMedia | null;
            selfie          : VerificationMedia | null;
            steps           : {                  
                  parsedTextFront   : boolean;
                  parsedTextBack    : boolean;
                  dataParsed        : boolean;
                  hasDOB            : boolean;
                  hasName           : boolean;
                  hasAddress        : boolean;
                  faceMatch         : boolean;
                  verified          : boolean;
            }
      }
