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
      //|| Get Verification Title (Page)
      //||------------------------------------------------------------------------------------------------||

      export const VerificationPageTitles: Record<VerificationTypes, string> = {
            MAIL: "Email Verification",
            PHNE: "Phone Verification",
            ADDR: "Address Verification",
            CRCD: "Credit Card Verification",
            UNAM: "Username Verification",
            IDEN: "Identification Verification",
            FACE: "Liveness / Face Verification",
            UNKN: "Verification"
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
      //|| Get Description by Type
      //||------------------------------------------------------------------------------------------------||      
      
      export const VerificationTypeDescriptions: Record<VerificationTypes, string> = {
            MAIL: "Confirm your email address so we can send account updates, receipts, and password resets.",
            PHNE: "Add and verify a mobile number for security alerts, account recovery, and optional 2-factor codes.",
            ADDR: "Validate your billing or residential address to enable region-specific features and compliance checks.",
            CRCD: "Place a valid card on file to confirm identity for payments and prevent fraudulent activity.",
            UNAM: "Choose a unique public username so others can find and identify your account.",
            IDEN: "Verify your legal name and age using a government-issued ID for higher-trust actions.",
            FACE: "Complete a quick liveness/face check to confirm you’re a real person and estimate age where required.",
            UNKN: "Unrecognized verification method.",
      };

      //||------------------------------------------------------------------------------------------------||      
      //|| Tutle
      //||------------------------------------------------------------------------------------------------||      

      export const VerificationStatusTitle: Record<VerificationStatuses, string > = {
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

      export const VerificationTypeIcons: Record<VerificationTypes, LucideIcon> = {
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

      export function getVerificationIcon(type: VerificationTypes): LucideIcon {
            return VerificationTypeIcons[type] || User;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Get Verification Description
      //||------------------------------------------------------------------------------------------------||

      export function getVerificationDescription(type: VerificationTypes): string {
            return VerificationTypeDescriptions[type] ?? "Unrecognized verification method.";
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Get Verification Page Title
      //||------------------------------------------------------------------------------------------------||

      export function getVerificationPageTitle(type: VerificationTypes): string {
           return VerificationPageTitles[type] ?? "Verification";
      
      }
