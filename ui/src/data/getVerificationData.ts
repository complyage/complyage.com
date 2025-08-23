//||------------------------------------------------------------------------------------------------||
//|| Verification Data
//|| src/data/getVerificationData.ts
//||------------------------------------------------------------------------------------------------||
   
      //||------------------------------------------------------------------------------------------------||      
      //|| Import
      //||------------------------------------------------------------------------------------------------||      

      import { VerificationStatuses, VerificationTypes }                            from "../interfaces/models/model.verify";
      import { Mail, Phone, IdCard, MapPin, CreditCard, User, Image as ImageIcon }  from "lucide-react";
      import { LucideIcon }                                                          from "lucide-react";

      //||------------------------------------------------------------------------------------------------||      
      //|| All Types
      //||------------------------------------------------------------------------------------------------||      

      export const VerificationTypeLabels: Record<VerificationTypes, string > = {
            MAIL: "Email",
            PHNE: "Phone",
            UAGE: "Age",
            ADDR: "Address",
            CRCD: "Credit Card",
            UNAM: "Username",
            UNKN: "Unknown"
      };

      //||------------------------------------------------------------------------------------------------||      
      //|| All Statuses
      //||------------------------------------------------------------------------------------------------||      

      export const VerificationStatusLabels: Record<VerificationStatuses, string > = {
            MISS: "N/A",
            PEND: "Pending",
            PEVF: "Pending Verification",
            APPR: "Pending Approval",
            VERF: "Verified",
            RJCT: "Rejected",
            ESCL: "Escalated",
            EXPD: "Expired",
            CNCL: "Cancelled"
      };

      //||------------------------------------------------------------------------------------------------||      
      //|| Get Verification Icon
      //||------------------------------------------------------------------------------------------------||      

      export const VerificationTypeIcons: Record<VerificationTypes, React.ComponentType<any>> = {
            MAIL: Mail,
            PHNE: Phone,
            UAGE: IdCard,
            ADDR: MapPin,
            CRCD: CreditCard,
            UNAM: User,
            UNKN: ImageIcon
      };

      //||------------------------------------------------------------------------------------------------||      
      //|| Get Verification Type
      //||------------------------------------------------------------------------------------------------||      

      export function getVerificationType(entity: VerificationTypes): string {
            return VerificationTypeLabels[entity];
      }

      //||------------------------------------------------------------------------------------------------||      
      //|| Get ALL Verification Type
      //||------------------------------------------------------------------------------------------------||      

      export function getAllVerificationTypes(): VerificationTypes[] {
            return Object.keys(VerificationTypeLabels) as VerificationTypes[];
      }

      //||------------------------------------------------------------------------------------------------||      
      //|| Get Verification Status
      //||------------------------------------------------------------------------------------------------||      

      export function getVerificationStatus(entity: VerificationStatuses): string {
            return VerificationStatusLabels[entity];
      }

      //||------------------------------------------------------------------------------------------------||      
      //|| Get Verification Icon by Type
      //||------------------------------------------------------------------------------------------------||      

      export function getVerificationIcon(type: VerificationTypes): LucideIcon | null {
            return VerificationTypeIcons[type] || null;
      }

