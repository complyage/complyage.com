//||------------------------------------------------------------------------------------------------||
//|| Step 1
//|| ID Verification
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import { Camera as CameraIcon, Upload as UploadIcon }       from "lucide-react";
      import React, {useEffect, useRef, useState}                 from "react";
      
      //||------------------------------------------------------------------------------------------------||
      //|| Props/Interfaces
      //||------------------------------------------------------------------------------------------------||

      import { Media }                                            from "../../../interfaces/base/media";
      import { StepProps }                                        from "./IDVerification";
      import { VerificationMedia }                                from "../../../interfaces/verify/id/process";

      //||------------------------------------------------------------------------------------------------||
      //|| Components
      //||------------------------------------------------------------------------------------------------||

      import Camera                                               from "../../../components/dynamic/Camera";
      import Upload                                               from "../../../components/dynamic/Upload";
      import MediaPreview                                         from "../../../components/dynamic/MediaPreview";

      //||------------------------------------------------------------------------------------------------||
      //|| ID Verification
      //||------------------------------------------------------------------------------------------------||

      export default function IDVerificationMedia({ which, updateProcess, process, onUpload, getUpload }: StepProps) {
            //||------------------------------------------------------------------------------------------------||
            //|| Const
            //||------------------------------------------------------------------------------------------------||
            const [captureMode, setCaptureMode]  = useState<"camera" | "upload" | null>(null);
            //||------------------------------------------------------------------------------------------------||
            //|| uploadFile
            //||------------------------------------------------------------------------------------------------||
            const [media, setMedia]              = useState<VerificationMedia | null>(null);
            //||------------------------------------------------------------------------------------------------||
            //|| Load
            //||------------------------------------------------------------------------------------------------||
            useEffect(() => {
                  (async () => {
                        try {
                              const media = await getUpload(which);
                              console.log("EFFECT MEDIA : ", media);
                              if (media === null || media.exists === false) return setMedia(null);
                              setMedia(media as VerificationMedia);
                              updateProcess({ [which]: media });
                        } catch (error) {
                              console.error("Error loading front ID image:", error);
                              setMedia(null);
                        }
                  })();
            }, []);
            //||------------------------------------------------------------------------------------------------||
            //|| Handler
            //||------------------------------------------------------------------------------------------------||
            const handleUpload = async (file: File | Blob) => {
                  onUpload(file, which);
            }
            //||------------------------------------------------------------------------------------------------||
            //|| Upload
            //||------------------------------------------------------------------------------------------------||
            if (captureMode === "upload") return (
                  <div className="w-full mx-auto text-center p-4 flex flex-col items-center gap-y-5 rounded-lg">
                        <div className="flex flex-row items-center justify-center w-full mx-auto">
                              <Upload which={ which } onUpload={ handleUpload } onClose={ () => setCaptureMode(null) } /> 
                        </div>
                  </div>
            );
            //||------------------------------------------------------------------------------------------------||
            //|| Capture
            //||------------------------------------------------------------------------------------------------||
            if (captureMode === "camera") return ( <Camera onUpload={ handleUpload } onClose={ () => setCaptureMode(null) } /> );
            //||------------------------------------------------------------------------------------------------||
            //|| JSX
            //||------------------------------------------------------------------------------------------------||
            return (
                  <>
                        {media !== null ? (
                              <MediaPreview media={media} onReset={() => setMedia(null)} />
                        ) : (
                              <div className="w-full mx-auto text-center p-4 flex flex-row rounded-lg gap-5">
                                    {/* Camera Button */}
                                    <div className="flex w-1/2 justify-center">
                                          <button
                                                onClick={() => setCaptureMode("camera")}
                                                className="flex flex-col items-center justify-center p-4 w-full bg-gray-800 text-white rounded-l-lg hover:bg-gray-600 transition"
                                          >
                                                <CameraIcon className="w-8 h-8 mb-2" />
                                                <span className="text-sm">Use Camera</span>
                                          </button>
                                    </div>
                                    {/* Upload Button */}
                                    <div className="flex w-1/2 justify-center">
                                          <button
                                                onClick={() => setCaptureMode("upload")}
                                                className="flex flex-col items-center justify-center p-4 w-full bg-gray-800 text-white rounded-r-lg hover:bg-gray-600 transition"
                                          >
                                                <UploadIcon className="w-8 h-8 mb-2" />
                                                <span className="text-sm">Upload File</span>
                                          </button>
                                    </div>
                              </div>
                        )}
                  </>
            );
      }
