//||------------------------------------------------------------------------------------------------||
//|| FaceVerification (Container)
//|| src/pages/verification/face/IDVerification.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React, {useState, useEffect}                   from "react";
      import {useNavigate, useSearchParams}                 from "react-router-dom";
      import { Smile }                                      from "lucide-react";

      //||------------------------------------------------------------------------------------------------||
      //|| Interfaces
      //||------------------------------------------------------------------------------------------------||

      import { VerificationFace, VerificationMedia }          from "../../../interfaces/verify/face/process";

      //||------------------------------------------------------------------------------------------------||
      //|| Steps
      //||------------------------------------------------------------------------------------------------||

      import FaceVerificationInitial                         from "./FaceVerification.Initial";
      import FaceVerificationStep1                           from "./FaceVerification.Step1";
      import FaceVerificationStep2                           from "./FaceVerification.Step2";
      import FaceVerificationSubmit                          from "./FaceVerification.Submit";

      //||------------------------------------------------------------------------------------------------||
      //|| Components
      //||------------------------------------------------------------------------------------------------||

      import InlineAlert                                    from "../../../components/base/InlineAlert";
      import MembersLayout                                  from "../../../layouts/MembersLayout";
      import ProgressSteps, { ProgressStep }                from "../../../components/base/ProgressSteps";

      //||------------------------------------------------------------------------------------------------||
      //|| Step Props
      //||------------------------------------------------------------------------------------------------||

      export interface StepProps {
            which?            : "front" | "back" | "selfie";
            process           : VerificationFace;
            updateProcess     : (process: VerificationFace) => void;
            onUpload          : (file: File | Blob | null, which: "front" | "back" | "selfie") => Promise<void>;
            getUpload         : (which: "front" | "back" | "selfie") => Promise<VerificationMedia | null>;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Page
      //||------------------------------------------------------------------------------------------------||

      export default function FaceVerification({overlay}: { overlay?: boolean }) {

            //||------------------------------------------------------------------------------------------------||
            //|| Verification
            //||------------------------------------------------------------------------------------------------||

            const [searchParams]                = useSearchParams();
            const verificationUUID              = searchParams.get("identifier") || "";
            const [mainError, setMainError]     = useState<string | null>(null);

            //||------------------------------------------------------------------------------------------------||
            //|| Navigate
            //||------------------------------------------------------------------------------------------------||

            const navigate                     = useNavigate();  

            //||------------------------------------------------------------------------------------------------||
            //|| Process
            //||------------------------------------------------------------------------------------------------||

            const [process, setProcess] = useState<VerificationFace>({
                  step                : 0,
                  status              : "PEND",
                  verificationUUID    : verificationUUID,
            });            

            //||------------------------------------------------------------------------------------------------||
            //|| onUpload
            //||------------------------------------------------------------------------------------------------||

            const onUpload = async (file: File | Blob | null) => {
                  try {
                        if (!process.verificationUUID) throw new Error("Missing verification process identifier");

                        //||------------------------------------------------------------------------------------------------||
                        //|| Form Data / Request
                        //||------------------------------------------------------------------------------------------------||
                        const formData = new FormData();
                        formData.append("media", file || "");
                        const res = await fetch(`/v1/api/verify/face/media/upload?identifier=${process.verificationUUID}`, {
                              method            : "POST",
                              body              : formData,
                              credentials       : "include"
                        });
                        //||------------------------------------------------------------------------------------------------||
                        //|| Response
                        //||------------------------------------------------------------------------------------------------||
                        const json = await res.json();
                        if (!res.ok || !json.success) {
                              throw new Error(json.message || "Upload failed");
                        }
                        //||------------------------------------------------------------------------------------------------||
                        //|| Response
                        //||------------------------------------------------------------------------------------------------||
                        alert('Upload successful!');
                        console.log(`[UPLOAD] selfie`, json);
                        const media = json.data as VerificationMedia || null;
                        updateProcess({ ["selfie"]: media });
                  } catch (err: any) {
                        console.error("Upload error:", err.message || err);
                        alert(`Upload failed: ${err.message || err}`);
                  }
            };       

            //||------------------------------------------------------------------------------------------------||
            //|| Continue Enabled
            //||------------------------------------------------------------------------------------------------||

            const isContinueEnabled = () => {
                  if (process.step === 2) return true;
                  if (process.step === 3 && (!process.selfie || process.selfie.exists === false)) return false;
                  return true;
            }

            //||------------------------------------------------------------------------------------------------||
            //|| getUpload
            //||------------------------------------------------------------------------------------------------||

            const getUpload = async () : Promise<VerificationMedia | null> => {
                  try {
                        if (!process.verificationUUID) throw new Error("Missing verification process identifier");
                        //||------------------------------------------------------------------------------------------------||
                        //|| Request
                        //||------------------------------------------------------------------------------------------------||
                        const res = await fetch(`/v1/api/verify/face/media/fetch?identifier=${process.verificationUUID}`, {
                              method            : "GET",
                              credentials       : "include"
                        });
                        //||------------------------------------------------------------------------------------------------||
                        //|| Response
                        //||------------------------------------------------------------------------------------------------||
                        const json = await res.json();
                        console.log(`[FETCH=UPLOAD] selfie`, json);
                        //||------------------------------------------------------------------------------------------------||
                        //|| Request
                        //||------------------------------------------------------------------------------------------------||
                        if (!res.ok || !json.success) {
                              updateProcess({ ["selfie"]: null });
                              throw new Error(json.message || "Upload failed");
                        }
                        console.log("Fetched upload:", json.data);
                        return json.data as VerificationMedia;
                  } catch (err: any) {
                        console.error("Upload error:", err.message || err);
                        return null;
                  }
            };                 

            //||------------------------------------------------------------------------------------------------||
            //|| Continue
            //||------------------------------------------------------------------------------------------------||

            const handleContinue = () => {
                  updateProcess({
                        step: (process.step || 0) + 1
                  });
            };

            //||------------------------------------------------------------------------------------------------||
            //|| Set Process
            //||------------------------------------------------------------------------------------------------||

            const updateProcess = (update: Partial<VerificationFace>) => {
                  setProcess((prev) => ({
                        ...prev,
                        ...update
                  }));
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Handle Returning
            //||------------------------------------------------------------------------------------------------||

            useEffect(() => {
                  if (!process.verificationUUID) return;
                  (async () => {
                        try {
                              const res = await fetch(`/v1/api/verify/status?identifier=${process.verificationUUID}`);
                              if (!res.ok) {
                                    throw new Error("Failed to load verification process");
                              }
                              if (res.ok) {
                                    const data = await res.json();                                    
                                    console.log("/FACE/LOAD/", data);
                                    const newProcess = process;
                                    if (data.data.type !== "FACE") setMainError("Invalid verification type. Whatcha trying to do?");
                                    newProcess.status = data.data.status;
                                    switch(data.data.status) {
                                          case "PEND": break;
                                          case "INPG": break;
                                          case "PEVF":
                                                setMainError("This verification is pending approval. Please wait..");
                                                break;
                                          case "APPR":
                                                setMainError("This verification is already approved. Please view your dashboard for more information.");
                                                break;
                                          case "REJT":
                                                setMainError("This verification has been rejected. Please view your dashboard for more information.");
                                                break;
                                          default:
                                                setMainError("Invalid Status.");
                                                newProcess.step = 1;
                                    }
                                    updateProcess({...process, ...newProcess });
                              }
                        } catch {}
                  })();
            }, [process.verificationUUID]);

            //||------------------------------------------------------------------------------------------------||
            //|| Stepper
            //||------------------------------------------------------------------------------------------------||

            const steps: ProgressStep[] = [
                  { label: "How it works",            description: "How to get your Age verified" },
                  { label: "Enter DOB",               description: "Enter your birthday." },
                  { label: "Take a Selfie",           description: "Take a Selfie" },
                  { label: "Submit for Approval",     description: "Submit for Approval" },
            ];    

            //||------------------------------------------------------------------------------------------------||
            //|| Main
            //||------------------------------------------------------------------------------------------------||

            if (mainError !== null ) {
                  return (
                        <MembersLayout title="Face erification" icon={ Smile } overlay={overlay}>
                              <div className="w-full max-w-2xl mx-auto text-center p-6">
                                    <InlineAlert message={ mainError } isError={false} />
                                    {process.verificationUUID !== "" && (<button onClick={() => navigate(`/verification/status?identifier=${ process.verificationUUID}`)} className="mt-4 btn bg-blue-600 text-2xl h-auto py-2">View Status</button>)}
                              </div>
                              <div className="w-full max-w-2xl mx-auto text-center p-6 border-t border-gray-500">
                                    <button onClick={() => navigate("/members/verification")} className="btn btn-primary text-2xl bg-black/40">Go Back</button>                              
                              </div>
                        </MembersLayout>
                  );
            }      

            //||------------------------------------------------------------------------------------------------||
            //|| Step
            //||------------------------------------------------------------------------------------------------||

            return (
                  <MembersLayout overlay={overlay}>
                        <div className="w-full max-w-5xl mx-auto">                              
                              <ProgressSteps steps={ steps } currentStep={ process.step } className="mb-6" verificationType="FACE" />
                              { process.step < 1 && <FaceVerificationInitial updateProcess={ updateProcess } process={process} />}
                              { process.step === 2 && <FaceVerificationStep1 which="front" updateProcess={ updateProcess } process={process} onUpload={ onUpload } getUpload={ getUpload } />}
                              { process.step === 3 && <FaceVerificationStep2 which="back" updateProcess={ updateProcess } process={process}  onUpload={ onUpload } getUpload={ getUpload } />}
                              { process.step === 4 && <FaceVerificationSubmit updateProcess={ updateProcess } process={process} onUpload={ onUpload } getUpload={ getUpload } />}
                              {process.step > 0 && process.step < 4 && (                              
                                    <div className="flex flex-row items-center justify-between w-full border-t border-gray-500 mt-3 mx-auto pt-4">
                                          <button className="btn btn-primary text-2xl py-3 px-10 h-auto" onClick={() => updateProcess({ step: process.step - 1 })} >Back</button>
                                          <button disabled={ !isContinueEnabled() } className="btn btn-secondary text-2xl py-3 px-10 h-auto" onClick={() => handleContinue()}>Continue</button>
                                    </div>
                              )}
                              <div className="text-gray-500 text-sm block text-center pt-5">
                                    {process.verificationUUID} {process.step}
                              </div>

                        </div>

                  </MembersLayout>                  
            );
      }
