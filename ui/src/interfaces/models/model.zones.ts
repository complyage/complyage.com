//||------------------------------------------------------------------------------------------------||
//|| Interfaces : Zone
//|| Model Zone
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Model
      //||------------------------------------------------------------------------------------------------||

      export interface ModelZone {
            id               : string;
            state            : string;
            country          : string;
            requirements     : string;
            law?             : string;
            lawDescription?  : string;
            penalties?       : string;
            effective?       : string;
            meta?            : string;
            latitude?        : string;
            longitude?       : string;
      }

      //||------------------------------------------------------------------------------------------------||      
      //|| Requirement Keys
      //||------------------------------------------------------------------------------------------------||

      export type ZoneRequirement = "IDEN"|"CRCD"|"FACE"|"DIGITAL_ID"|"BIOMETRIC"|"3P_AV"|"DEVICE_SIGNAL"|"FACIAL_EST"|"OPEN_BANKING"|"HAND_ANALYSIS"|"ANON_AV"|"ZKP"|"CREDIT_CARD";
