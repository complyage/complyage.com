//||------------------------------------------------------------------------------------------------||
//|| Zone Data
//|| src/data/getZoneData.ts
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import { ZoneRequirement } from "../interfaces/models/model.zones";

      //||------------------------------------------------------------------------------------------------||
      //|| All Requirements
      //||------------------------------------------------------------------------------------------------||

      export const ZoneRequirementLabels: Record<ZoneRequirement, string > = {
            ID_UPLOAD     : "Upload of government-issued ID",
            GOV_ID        : "Government-issued ID check",
            TXN_DATA      : "Transactional data check",
            DIGITAL_ID    : "National digital ID system (SPID, eID, EUDI Wallet)",
            BIOMETRIC     : "Biometric check (facial recognition, facial age estimation)",
            "3P_AV"       : "Third-party age verification service",
            DEVICE_SIGNAL : "Device-based age signal/API",
            FACIAL_EST    : "Facial age estimation",
            OPEN_BANKING  : "Open banking verification",
            HAND_ANALYSIS : "Hand movement analysis",
            ANON_AV       : "Double anonymity system (France)",
            ZKP           : "Zero-knowledge proof system",
            CREDIT_CARD   : "Credit/debit card (only if paired with ID)",
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Get Requirement Label
      //||------------------------------------------------------------------------------------------------||

      export function getZoneRequirement(requirement: ZoneRequirement): string {
            return ZoneRequirementLabels[requirement];
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Get All Requirements
      //||------------------------------------------------------------------------------------------------||

      export function getAllZoneRequirements(): ZoneRequirement[] {
            return Object.keys(ZoneRequirementLabels) as ZoneRequirement[];
      }
