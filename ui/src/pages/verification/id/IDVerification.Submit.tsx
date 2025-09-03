//||------------------------------------------------------------------------------------------------||
//|| Step 1
//|| ID Verification
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import { Camera as CameraIcon, Upload as UploadIcon }      from "lucide-react";
      import React, {useEffect, useRef, useState}                from "react";

      //||------------------------------------------------------------------------------------------------||
      //|| Interfaces
      //||------------------------------------------------------------------------------------------------||

      import { StepProps }                                       from "./IDVerification";
      import { VerificationID, VerificationMedia }               from "../../../interfaces/verify/id/process";

      //||------------------------------------------------------------------------------------------------||
      //|| Components
      //||------------------------------------------------------------------------------------------------||

      import MediaThumb                                          from "../../../components/dynamic/MediaThumb";
      import InlineAlert                                         from "../../../components/base/InlineAlert";

      //||------------------------------------------------------------------------------------------------||
      //|| ID Verification
      //||------------------------------------------------------------------------------------------------||

      export default function IDVerificationSubmit({ process, getUpload, updateProcess } : StepProps) {

            //||------------------------------------------------------------------------------------------------||
            //|| ID Verification
            //||------------------------------------------------------------------------------------------------||
            
            const [loading, setLoading]   = useState(false);
            const [error, setError]       = useState<string | null>(null);

            //||------------------------------------------------------------------------------------------------||
            //|| Handle Submit
            //||------------------------------------------------------------------------------------------------||

            const handleSubmit = async () => {
                  setError(null);
                  setLoading(true);
                  if (!process.verificationUUID) return setError("Verification UUID is missing.");
                  if (!process.front || !process.back || !process.selfie) return setError("All media must be uploaded before submission.");
                  if (process.front.exists === false || process.back.exists === false || process.selfie.exists === false) return setError("All media must be valid before submission.");
                  try {
                        const res = await fetch("/v1/api/verify/id/success", {
                              method                  : "POST",
                              credentials             : "include",
                              headers                 : { "Content-Type": "application/json", },
                              body                    : JSON.stringify({ "identifier" : process.verificationUUID }),
                        });
                        const data = await res.json();
                        console.log("Submission response:", data);
                        if (!res.ok && data.message) {
                              throw new Error(data.message || "Unknown error occurred.");
                        }
                        window.location.href = `/verification/status?identifier=${process.verificationUUID}`;
                  } catch (err: any) {
                        setError(err.message || "Submission failed.");
                  } finally {
                        setLoading(false);
                  }
            };

            //||------------------------------------------------------------------------------------------------||
            //|| JSX
            //||------------------------------------------------------------------------------------------------||

            return (
                  <div className="w-full max-w-2xl rounded-lg p-4 bg-transparent">
                        {error && (<InlineAlert message={error} isError />) }
                        <div className="flex flex-row w-full gap-3 mb-2">
                              <div className="flex-1 flex flex-col items-center">
                                    <b className="font-bold pb-2">Front of ID</b>
                                    {process.front ? (
                                          <MediaThumb onEdit={() => updateProcess({ ...process, step: 2 })} media={process.front} />
                                    ) : (
                                          <div className="w-20 h-20 bg-gray-100 rounded border flex items-center justify-center text-gray-400">—</div>
                                    )}
                              </div>
                              <div className="flex-1 flex flex-col items-center">
                                    <b className="font-bold pb-2">Back of ID</b>
                                    {process.back ? (
                                          <MediaThumb onEdit={() => updateProcess({ ...process, step: 3 })} media={process.back} />
                                    ) : (
                                          <div className="w-20 h-20 bg-gray-100 rounded border flex items-center justify-center text-gray-400">—</div>
                                    )}
                              </div>
                              <div className="flex-1 flex flex-col items-center">
                                    <b className="font-bold pb-2">Selfie</b>
                                    {process.selfie ? (
                                          <MediaThumb onEdit={() => updateProcess({ ...process, step: 4 })} media={process.selfie} />
                                    ) : (
                                          <div className="w-20 h-20 bg-gray-100 rounded border flex items-center justify-center text-gray-400">—</div>
                                    )}
                              </div>
                        </div>

                        <div className="flex justify-center mt-10">
                              <button
                                    className="btn btn-secondary text-lg px-10 py-2"
                                    onClick={handleSubmit}
                                    disabled={loading}
                              >
                                    {loading ? "Submitting..." : "Submit for Approval"}
                              </button>
                        </div>
                  </div>
            );
      }