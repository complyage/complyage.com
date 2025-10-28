//||------------------------------------------------------------------------------------------------||
//|| ID Verification Status
//|| src/pages/verification/id/IDVerification.Status.tsx
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useState, useEffect }                       from "react";
import { Clock, AlertCircle, UserPlus, CheckCircle, Loader2 }        from "lucide-react";
import { useSearchParams }                                  from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Data
//||------------------------------------------------------------------------------------------------||
      
import { getVerificationStatus, getVerificationIcon }       from "../../../data/getVerificationData";
      
//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { VerificationID }                                         from "../../../interfaces/verify/id/process";
import { VerificationStatuses, VerificationTypes }                from "../../../interfaces/models/model.verify";
import { VerificationIDStatusProcess, VerificationIDStatusStep }  from "../../../interfaces/verify/status/process";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import MembersLayout                                        from "../../../layouts/MembersLayout";
import ProgressSteps, { ProgressStep }                      from "../../../components/base/ProgressSteps";
import SpinnerCircle                                        from "../../../components/base/SpinnerCircle";
import InlineAlert                                          from "../../../components/base/InlineAlert";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface IDVerificationStatusProps {
      process        : VerificationID;
      onTryAgain?    : () => void;
      onEscalate?    : () => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Status Icon Map
//||------------------------------------------------------------------------------------------------||

const StatusIcon: Record<VerificationStatuses, React.ComponentType<any>> = {
      INPR: Loader2,       // swapped Cpu -> Loader2
      PEND: Clock,
      PEVF: Clock,
      VERF: CheckCircle,
      RJCT: AlertCircle,
      ESCL: UserPlus,
      EXPD: AlertCircle,
};

const StatusColor: Record<VerificationStatuses, string> = {
      PEND: "text-yellow-500",
      PEVF: "text-blue-500",
      VERF: "text-green-600",
      RJCT: "text-red-600",
      ESCL: "text-orange-600",
      INPR: "text-green-300",   
      EXPD: "text-gray-500",
};

//||------------------------------------------------------------------------------------------------||
//|| Steps IDEN
//||------------------------------------------------------------------------------------------------||

const stepsIDEN: ProgressStep[] = [
      { label: "Awaiting Agent",          description: "Awaiting Agent" },
      { label: "Parse ID",         description: "Parsing ID Data" },
      { label: "Verify DOB",           description: "Verifying DOB" },
      { label: "Verify Photo",         description: "Upload the back of your ID" },
      { label: "Match Photo",          description: "Matching ID to Selfie" },
      { label: "Encrypt Data",         description: "Encrypting your photos" },
];

//||------------------------------------------------------------------------------------------------||
//|| Steps FACE
//||------------------------------------------------------------------------------------------------||

const stepsFACE: ProgressStep[] = [
      { label: "Awaiting Agent",          description: "Awaiting Agent" },
      { label: "Estimate Age",            description: "Estimating Facial Age" },
      { label: "Match DOB",               description: "Matching Date of Birth" },
      { label: "Complete",                description: "Complete" },
];

//||------------------------------------------------------------------------------------------------||
//|| Major Statuses
//||------------------------------------------------------------------------------------------------||

const majorGreen = ['AGENT_L1','DOB_MATCH', 'APPROVAL', 'STATUS_REJT'];
const majorRed   = ['REVIEW', 'DOB_MISMATCH', 'REJECTED', 'EXPIRED', 'CANCELLED'];
const emphasize = (status : string) =>{ 
      return majorGreen.includes(status) ? "text-green-400 font-bold" :
             majorRed.includes(status)   ? "text-red-400 font-bold"   : 
                                          "text-gray-400"; 
}     

//||------------------------------------------------------------------------------------------------||
//|| Get Steps
//||------------------------------------------------------------------------------------------------||

function getStepsFromProcess(vType: VerificationTypes): ProgressStep[] {
      switch (vType) {
            case "IDEN": return stepsIDEN;
            case "FACE": return stepsFACE;
            default:     return [];
      }
}
            
//||------------------------------------------------------------------------------------------------||
//|| Main Component
//||------------------------------------------------------------------------------------------------||

export default function VerificationStatus({overlay}: { overlay?: boolean }) {

      //||------------------------------------------------------------------------------------------------||
      //|| Verification
      //||------------------------------------------------------------------------------------------------||

      const [searchParams] = useSearchParams();
      const verificationId = searchParams.get("identifier") || "";

      //||----------------------------------------------------------------------------------------||
      //|| State
      //||----------------------------------------------------------------------------------------||

      const [loading, setLoading]   = useState<boolean>(true);
      const [reload, setReload]     = useState<boolean>(false);
      const [error , setError]      = useState<string | null>(null);
      const [process, setProcess]   = useState<VerificationIDStatusProcess>({
            status            : "UNKN" as VerificationStatuses,
            type              : "IDEN",
            step              : 0,
            verificationUUID  : verificationId,
            steps             : [],
      });

      //||----------------------------------------------------------------------------------------||
      //|| Data
      //||----------------------------------------------------------------------------------------||

      const fetchStatus = () => {
            fetch(`/v1/api/verify/status?identifier=${verificationId}`).then(resp => {
                  if (!resp.ok) {
                        throw new Error("Failed to fetch verification status.");
                  }
                  return resp.json();
            }).then(response => {
                  console.log("Verification Status Data:", response);
                  if (response.success !== true) {
                        throw new Error(response.message || "Unknown error fetching verification status.");
                  }
                  setProcess({
                        type              : response.data.type,
                        status            : response.data.status,
                        step              : response.data.step,
                        steps             : response.data.steps,
                        verificationUUID  : response.data.verificationUUID,
                  } as VerificationIDStatusProcess);
                  setLoading(false);
                  window.setTimeout(() => setReload(false), 2000);
            }).catch(err => {
                  setError(err.message);
                  setLoading(false);
                  window.setTimeout(() => setReload(false), 2000);
            });
      }

      //||----------------------------------------------------------------------------------------||
      //|| Data
      //||----------------------------------------------------------------------------------------||

      const testReset = () => {
            fetch(`/v1/api/verify/id/reset?identifier=${verificationId}`).then(resp => {
                  if (!resp.ok) {
                        alert("Failed to reset verification status.");
                  }
                  return resp.json();
            }).then(response => {
                  console.log("Verification Status Data:", response);
            }).catch(err => {
                  alert(err.message);
            });
      }            

      //||----------------------------------------------------------------------------------------||
      //|| Poll Status Every 5 Seconds
      //||----------------------------------------------------------------------------------------||

      useEffect(() => {
            if (!verificationId) return;			
            const interval = setInterval(() => { 
                  setReload(true); 
                  fetchStatus(); 
            }, 2000);
            fetchStatus();
            return () => clearInterval(interval);
      }, [verificationId]);

      //||----------------------------------------------------------------------------------------||
      //|| Data
      //||----------------------------------------------------------------------------------------||

      const status      = process.status || "MISS" as VerificationStatuses;
      const Icon        = StatusIcon[status] || Clock;
      const color       = StatusColor[status] || "text-gray-700";
      const label       = getVerificationStatus(status || "MISS");
      
      //||----------------------------------------------------------------------------------------||
      //|| Show Final / Escalate
      //||----------------------------------------------------------------------------------------||

      const isFinal = ["VERF", "RJCT", "EXPD", "CNCL"].includes(status);

      //||----------------------------------------------------------------------------------------||
      //|| Loading
      //||----------------------------------------------------------------------------------------||

      useEffect(() => {
            if (!verificationId) setError("No verification ID provided.");
      }, [verificationId]);
      
      //||----------------------------------------------------------------------------------------||
      //|| Use Load
      //||----------------------------------------------------------------------------------------||

      useEffect(() => {
            fetchStatus();
      }, []);
                  
      //||----------------------------------------------------------------------------------------||
      //|| Loading
      //||----------------------------------------------------------------------------------------||

      if (loading) {
            return (
                  <MembersLayout overlay={overlay}>
                        <div className="w-full max-w-2xl mx-auto flex flex-col items-center justify-center p-8 bg-black/20 rounded-xl mt-10">
                              <SpinnerCircle className="w-16 h-16 text-gray-500 animate-spin" />
                              <span className="text-lg text-gray-600 mt-4">Please wait...</span>
                        </div>
                  </MembersLayout>
            );
      }

      //||----------------------------------------------------------------------------------------||
      //|| JSX
      //||----------------------------------------------------------------------------------------||

      return (
            <MembersLayout overlay={overlay}>
                  <div className="w-full max-w-full mx-auto flex flex-col items-center justify-center rounded-lg">


                        <div className="w-full mb-8">
                              <ProgressSteps steps={getStepsFromProcess(process.type)} currentStep={process.step || 1} />
                        </div>

                        <div className="bg-black/20 pt-4 w-full max-w-2xl rounded-xl relative">
                              <div className="absolute top-3 right-5 z-10">
                                    { reload && ( <SpinnerCircle className="inline-block w-4 h-4" /> )}   
                              </div>

                              <div className="flex flex-col items-center mb-5">
                                    <button className="btn btn-sm btn-secondary absolute top-3 left-5" onClick={testReset}>Reset (Test Only)</button>
                                    <Icon className={`w-24 h-24 ${color} mx-auto ${status === "INPR" ? "animate-spin" : ""}`} />
                              </div>
                              
                              <span className={`text-3xl py-5 text-center block font-bold tracking-wide items-center gap-2 bg-white/80 text-gray-800 shadow-xl border-black mx-[-20px]` }>
                                    {label}
                              </span>

                              <div className="p-4">
                              { process.steps && process.steps.length > 0 && (
                                    <table className="table table-auto w-full mt-4">
                                          <thead className="bg-gray-800">
                                                <tr className="text-sm font-bold text-gray-200">
                                                      <th className="">Step</th>
                                                      <th className=" text-center">Details</th>
                                                      <th className="  text-right">Timestamp</th>
                                                </tr>
                                          </thead>
                                          <tbody>
                                          {process.steps.map((step, index) => {
                                                return (
                                                <tr key={index}>
                                                      <td className={`text-xs font-bold ${emphasize(step.type.name)}`} colSpan={ step.details ? 1 : 2 }>
                                                            { step.type.value }
                                                      </td>
                                                      {step.details && (<td className="text-xs text-center font-bold">
                                                            <span className="bg-black/20 p-1 px-2 rounded-lg text-green-300">{step.details}</span>
                                                      </td>)}
                                                      <td className="text-xs text-right text-gray-400">
                                                            {
                                                                  (() => {
                                                                  const d = new Date(step.timestamp); // parses as UTC, displays local
                                                                  const now = new Date();
                                                                  return d.getFullYear() === now.getFullYear() &&
                                                                              d.getMonth()    === now.getMonth() &&
                                                                              d.getDate()     === now.getDate()
                                                                        ? d.toLocaleTimeString("en-US", { hour: "2-digit", minute: "2-digit", second: "2-digit", timeZone: "UTC" })
                                                                        : d.toLocaleString();
                                                                  })()
                                                            }
                                                      </td>
                                                </tr>
                                          )})}
                                          </tbody>
                                    </table>
                              )}
                              </div>
                        </div>

                        {/* Escalate Option */}
                        {status === "ESCL"  && (
                              <div className="mt-2 flex flex-col items-center">
                                    <div className="text-orange-700 mb-2 text-lg font-medium">
                                          Automatic verification couldn't complete.<br />
                                          <span className="text-base text-gray-600">You can request a manual review.</span>
                                    </div>
                                    <button
                                          className="btn btn-primary text-lg px-8 py-2"
                                    >
                                          Escalate to Human Verification
                                    </button>
                              </div>
                        )}

                        {/* Final Statuses */}
                        {isFinal && (
                              <div className="mt-6 flex flex-col items-center">
                                    <div className="mb-2 text-lg font-semibold">
                                          {status === "VERF" && (
                                                <span className="text-green-700">✅ Your ID has been verified!</span>
                                          )}
                                          {status === "RJCT" && (
                                                <span className="text-red-700">❌ Your ID verification was rejected.</span>
                                          )}
                                          {status === "EXPD" && (
                                                <span className="text-gray-500">⏰ Your verification has expired.</span>
                                          )}
                                    </div>
                                    <button
                                          className="btn btn-secondary text-lg px-10 py-2 mt-4"
                                    >
                                          {status === "VERF" ? "Verify Another ID" : "Try Again"}
                                    </button>
                              </div>
                        )}

                  </div>

                  <div className="block text-xs text-gray-500 pt-4 text-center mt-1 tracking-widest uppercase">
                        {process.verificationUUID}             
                  </div>

            </MembersLayout>
      );
};
