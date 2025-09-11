//||------------------------------------------------------------------------------------------------||
//|| Badge Statuses
//|| /src/components/base/BadgeStatuses.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React                                           from "react";
      import { 
            ArrowUpRight, 
            XCircle, 
            Timer, 
            CheckCircle,
            ClockArrowUp,
            Loader
      }                                                     from "lucide-react";
      import { LucideIcon }                                  from "lucide-react";
      import { VerificationTypes, VerificationStatuses }     from "../../interfaces/models/model.verify";
      import { getVerificationStatus }                       from "../../data/getVerificationData";

      //||------------------------------------------------------------------------------------------------||
      //|| Types
      //||------------------------------------------------------------------------------------------------||

      export interface BadgeDefaults {
            type         : VerificationTypes;
            status?      : VerificationStatuses;
            handleInit?  : () => void;
            handleVerify?: (uuid: string) => void;
            checkStatus? : (uuid: string) => void;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Base
      //||------------------------------------------------------------------------------------------------||

      export const BaseBadge: React.FC<{
            text     : string;
            icon     : LucideIcon;
            color    : string;             // e.g. "white", "green-600", "yellow-500"
            fill?    : string;             // e.g. "bg-green-600", "bg-blue-400/60"
            border?  : string;             // e.g. "border-transparent", "border-blue-200"
            pulse?   : boolean;
            onClick? : () => void;
      }> = ({ text, icon: Icon, color, fill, border, pulse, onClick }) => {
            const clickable  = onClick ? "cursor-pointer shadow-lg" : "cursor-default";
            const pulseClass = pulse ? "animate-pulse" : "";
            return (
                  <span
                        className={`badge border-2 ${border ? border : `border-${color}`} text-${color} ${fill || ""} flex items-center gap-2 ml-auto py-4 px-5 w-40 ${clickable} ${pulseClass}`}
                        onClick={onClick}
                  >
                        <span className="block font-bold whitespace-nowrap">{text}</span>
                        <Icon size={20} className="inline-block" />
                  </span>
            );
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Badge: Pending  (uses BaseBadge + click via defaults.handleInit)
      //||------------------------------------------------------------------------------------------------||

      export const BadgePending: React.FC<{ defaults: BadgeDefaults }> = ({ defaults }) => {
            return (
                  <BaseBadge
                        text="Continue"
                        icon={Loader}
                        color="white"
                        fill="bg-white/10"
                        border="border-transparent"
                        onClick={defaults.handleInit}
                  />
            );
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Badge: Pending Verification (routes by type; uses defaults handlers)
      //||------------------------------------------------------------------------------------------------||

      export const BadgePendingVerification: React.FC<{
            uuid     : string;
            defaults : BadgeDefaults;
      }> = ({ uuid, defaults }) => {
            if (defaults.type === "IDEN") {
                  return (
                        <BaseBadge
                              text="View Status"
                              icon={CheckCircle}
                              color="white"
                              fill="bg-green-600"
                              border="border-transparent"
                              onClick={() => defaults.checkStatus && defaults.checkStatus(uuid)}
                        />
                  );
            }
            return (
                  <BaseBadge
                        text="Enter Code"
                        icon={ArrowUpRight}
                        color="white"
                        fill="bg-blue-400/60"
                        border="border-transparent"
                        onClick={() => defaults.handleVerify && defaults.handleVerify(uuid)}
                  />
            );
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Badge: In Progress  (BaseBadge)
      //||------------------------------------------------------------------------------------------------||

      export const BadgeInProgress: React.FC<{ defaults: BadgeDefaults }> = () => (
            <BaseBadge
                  text="In Progress"
                  icon={Timer}
                  color="white"
                  border="border-blue-200"
                  fill="bg-blue-400/30"
                  pulse={true}
            />
      );

      //||------------------------------------------------------------------------------------------------||
      //|| Badge: Verified  (BaseBadge)
      //||------------------------------------------------------------------------------------------------||

      export const BadgeVerified: React.FC<{ defaults?: BadgeDefaults }> = () => (
            <BaseBadge
                  text="Verified"
                  icon={CheckCircle}
                  color="white"
                  fill="bg-green-600"
                  border="border-transparent"
            />
      );

      //||------------------------------------------------------------------------------------------------||
      //|| Badge: Rejected  (BaseBadge + click to retry via defaults.handleInit)
      //||------------------------------------------------------------------------------------------------||

      export const BadgeRejected: React.FC<{ defaults: BadgeDefaults }> = ({ defaults }) => (
            <BaseBadge
                  text="Rejected"
                  icon={XCircle}
                  color="white"
                  fill="bg-red-600"
                  border="border-transparent"
                  onClick={defaults.handleInit}
            />
      );

      //||------------------------------------------------------------------------------------------------||
      //|| Badge: Escalated  (BaseBadge with dynamic text)
      //||------------------------------------------------------------------------------------------------||

      export const BadgeEscalated: React.FC<{ defaults: BadgeDefaults }> = ({ defaults }) => (
            <BaseBadge
                  text={getVerificationStatus(defaults.status as VerificationStatuses)}
                  icon={ClockArrowUp}
                  color="yellow-500"
                  fill="bg-yellow-500/10"
                  border="border-yellow-500"
            />
      );

      //||------------------------------------------------------------------------------------------------||
      //|| Badge: Expired  (BaseBadge with dynamic text)
      //||------------------------------------------------------------------------------------------------||

      export const BadgeExpired: React.FC<{ defaults: BadgeDefaults }> = ({ defaults }) => (
            <BaseBadge
                  text={getVerificationStatus(defaults.status as VerificationStatuses)}
                  icon={Timer}
                  color="gray-400"
                  fill="bg-gray-700/20"
                  border="border-transparent"
            />
      );

      //||------------------------------------------------------------------------------------------------||
      //|| Badge: Error / Default  (BaseBadge)
      //||------------------------------------------------------------------------------------------------||

      export const BadgeError: React.FC<{ defaults: BadgeDefaults }> = ({ defaults }) => (
            <BaseBadge
                  text={`Error - ${defaults.status ?? "UNKN"}`}
                  icon={XCircle}
                  color="red-500"
                  fill="bg-red-500/10"
                  border="border-red-500"
            />
      );

      //||------------------------------------------------------------------------------------------------||
      //|| Master BadgeStatus (routes to the correct badge) — passes badgeDefaults to all
      //||------------------------------------------------------------------------------------------------||

      export const BadgeStatus: React.FC<{
            uuid         : string;
            type         : VerificationTypes;
            status       : VerificationStatuses;
            handleInit?  : () => void;
            handleVerify?: (uuid: string) => void;
            checkStatus? : (uuid: string) => void;
      }> = ({ uuid, type, status, handleInit, handleVerify, checkStatus }) => {

            //||------------------------------------------------------------------------------------------------||
            //|| Defaults bundle (shared across all sub-badges)
            //||------------------------------------------------------------------------------------------------||
            const badgeDefaults: BadgeDefaults = {
                  type        : type,
                  status      : status,
                  handleInit  : handleInit,
                  handleVerify: handleVerify,
                  checkStatus : checkStatus
            };

            //||------------------------------------------------------------------------------------------------||
            //|| Route
            //||------------------------------------------------------------------------------------------------||
            switch (status) {
                  case "PEND": return <BadgePending              defaults={badgeDefaults} />;
                  case "PEVF": return <BadgePendingVerification  uuid={uuid} defaults={badgeDefaults} />;
                  case "INPR": return <BadgeInProgress           defaults={badgeDefaults} />;
                  case "VERF": return <BadgeVerified             defaults={badgeDefaults} />;
                  case "RJCT": return <BadgeRejected             defaults={badgeDefaults} />;
                  case "ESCL": return <BadgeEscalated            defaults={badgeDefaults} />;
                  case "EXPD": return <BadgeExpired              defaults={badgeDefaults} />;
                  default:      return <BadgeError               defaults={badgeDefaults} />;
            }
      };
