//||------------------------------------------------------------------------------------------------||
//|| Verification Status
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import { ConstantsEntry }                              from "../../../interfaces/aria/constants";
      import { VerificationStatuses. VerificationTypes }     from "../../models/model.verify";

      //||------------------------------------------------------------------------------------------------||
      //|| Step
      //||------------------------------------------------------------------------------------------------||

      export interface VerificationIDStatusStep {
            type                      : ConstantsEntry;
            details                   : string;
            timestamp                 : string;
      }      

      //||------------------------------------------------------------------------------------------------||
      //|| Process
      //||------------------------------------------------------------------------------------------------||

      export interface VerificationIDStatusProcess {
            type                          : VerificationTypes;
            status                        : VerificationStatuses;
            step                          : number;
            steps                         : VerificationIDStatusStep[];
            verificationUUID              : string;
      }

