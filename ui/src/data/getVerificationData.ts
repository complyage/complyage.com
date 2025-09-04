//||------------------------------------------------------------------------------------------------||
//|| Verification Data
//|| src/data/getVerificationData.ts
//||------------------------------------------------------------------------------------------------||
   
      //||------------------------------------------------------------------------------------------------||      
      //|| Import
      //||------------------------------------------------------------------------------------------------||      

      import { VerificationStatuses, VerificationTypes }                            from "../interfaces/models/model.verify";
      import { Mail, Phone, IdCard, MapPin, CreditCard, User, Image as ImageIcon, Smile }  from "lucide-react";
      import { LucideIcon }                                                          from "lucide-react";

      //||------------------------------------------------------------------------------------------------||      
      //|| All Types
      //||------------------------------------------------------------------------------------------------||      

      export const VerificationTypeLabels: Record<VerificationTypes, string > = {
            MAIL: "Email",
            PHNE: "Phone",
            ADDR: "Address",
            CRCD: "Credit Card",
            UNAM: "Username",
            IDEN: "ID Document",
            FACE: "Facial Recognition",
            UNKN: "Unknown"
      };

      //||------------------------------------------------------------------------------------------------||      
      //|| All Statuses
      //||------------------------------------------------------------------------------------------------||      

      export const VerificationStatusLabels: Record<VerificationStatuses, string > = {
            PEND: "Pending",
            PEVF: "Pending Verification",
            VERF: "Verified",
            INPR: "In Progress",
            RJCT: "Rejected",
            ESCL: "Escalated",
            EXPD: "Expired",
      };

      //||------------------------------------------------------------------------------------------------||      
      //|| Get Verification Icon
      //||------------------------------------------------------------------------------------------------||      

      export const VerificationTypeIcons: Record<VerificationTypes, React.ComponentType<any>> = {
            MAIL: Mail,
            PHNE: Phone,
            IDEN: IdCard,
            ADDR: MapPin,
            CRCD: CreditCard,
            UNAM: User,
            FACE: Smile,
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

