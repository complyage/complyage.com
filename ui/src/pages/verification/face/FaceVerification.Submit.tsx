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
      //|| Hooks
      //||------------------------------------------------------------------------------------------------||

      import { useOverlayNavigate }                        from "../../../hooks/useOverlay";

      //||------------------------------------------------------------------------------------------------||
      //|| Interfaces
      //||------------------------------------------------------------------------------------------------||

      import { StepProps }                                       from "./FaceVerification";
      import { DOB }                                             from "../../../interfaces/base/user";

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
            //|| Navigate
            //||------------------------------------------------------------------------------------------------||

            const navigate                     = useOverlayNavigate();  

            //||------------------------------------------------------------------------------------------------||
            //|| Handle Submit
            //||------------------------------------------------------------------------------------------------||

            const handleSubmit = async () => {
                  setError(null);
                  setLoading(true);
                  if (!process.verificationUUID) return setError("Verification UUID is missing.");
                  if (!process.selfie) return setError("All media must be uploaded before submission.");
                  if (process.selfie.exists === false) return setError("All media must be valid before submission.");
                  if (!process.dob) return setError("Date of Birth is required.");
                  if (process.dob.year === 0 || process.dob.month === 0 || process.dob.day === 0) return setError("Date of Birth is required.");
                  try {
                        const res = await fetch("/v1/api/verify/face/success", {
                              method                  : "POST",
                              credentials             : "include",
                              headers                 : { "Content-Type": "application/json", },
                              body                    : JSON.stringify({ "identifier" : process.verificationUUID, "DOB": process.dob }),
                        });
                        const data = await res.json();
                        console.log("Submission response:", data);
                        if (!res.ok && data.message) {
                              throw new Error(data.message || "Unknown error occurred.");
                        }
                        navigate(`/verification/status?identifier=${process.verificationUUID}`);
                  } catch (err: any) {
                        setError(err.message || "Submission failed.");
                  } finally {
                        setLoading(false);
                  }
            };

            //||------------------------------------------------------------------------------------------------||
            //|| Format
            //||------------------------------------------------------------------------------------------------||

            const formatDOB = (dob: DOB) : string => {
                  const year    = dob.year;
                  const month   = (dob.month).toString().padStart(2, '0');
                  const day     = dob.day.toString().padStart(2, '0');
                  return `${year}-${month}-${day}`;
            }


            //||------------------------------------------------------------------------------------------------||
            //|| JSX
            //||------------------------------------------------------------------------------------------------||

            return (
                  <div className="w-full rounded-lg p-4 bg-transparent">
                        {error && (<InlineAlert message={error} isError />) }
                        <div className="flex flex-row w-full gap-3 mb-2">
                              <div className="flex-1 flex flex-col items-center">
                                    <b className="font-bold pb-2">
                                          {process.dob ? (
                                                <span className="text-lg">Date of Birth : {`${formatDOB(process.dob)}`}</span>
                                          ) : ("—")}                                          
                                    </b>
                                    {process.selfie ? (
                                          <div className="text-sm text-gray-400">
                                                <MediaThumb onEdit={() => updateProcess({ ...process, step: 3 })} media={process.selfie} />
                                          </div>
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