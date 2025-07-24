export interface ZoneRequirementType {
      IDUpload: string;
      GovID: string;
      TxnData: string;
      DigitalID: string;
      Biometric: string;
      ThirdPartyAV: string;
      DeviceSignal: string;
      FacialEst: string;
      OpenBanking: string;
      HandAnalysis: string;
      AnonAV: string;
      ZKP: string;
      CreditCard: string;
   }
   
   export const ZoneRequirements: ZoneRequirementType = {
      IDUpload:     "ID_UPLOAD",
      GovID:        "GOV_ID",
      TxnData:      "TXN_DATA",
      DigitalID:    "DIGITAL_ID",
      Biometric:    "BIOMETRIC",
      ThirdPartyAV: "3P_AV",
      DeviceSignal: "DEVICE_SIGNAL",
      FacialEst:    "FACIAL_EST",
      OpenBanking:  "OPEN_BANKING",
      HandAnalysis: "HAND_ANALYSIS",
      AnonAV:       "ANON_AV",
      ZKP:          "ZKP",
      CreditCard:   "CREDIT_CARD",
   };
   
   export function ZoneRequirementPlain(requirement: string): string {
      switch (requirement) {
         case ZoneRequirements.IDUpload:
            return "Upload of government-issued ID";
         case ZoneRequirements.GovID:
            return "Government-issued ID check";
         case ZoneRequirements.TxnData:
            return "Transactional data check";
         case ZoneRequirements.DigitalID:
            return "National digital ID system (SPID, eID, EUDI Wallet)";
         case ZoneRequirements.Biometric:
            return "Biometric check (facial recognition, facial age estimation)";
         case ZoneRequirements.ThirdPartyAV:
            return "Third-party age verification service";
         case ZoneRequirements.DeviceSignal:
            return "Device-based age signal/API";
         case ZoneRequirements.FacialEst:
            return "Facial age estimation";
         case ZoneRequirements.OpenBanking:
            return "Open banking verification";
         case ZoneRequirements.HandAnalysis:
            return "Hand movement analysis";
         case ZoneRequirements.AnonAV:
            return "Double anonymity system (France)";
         case ZoneRequirements.ZKP:
            return "Zero-knowledge proof system";
         case ZoneRequirements.CreditCard:
            return "Credit/debit card (only if paired with ID)";
         default:
            return "Unknown requirement";
      }
   }
   