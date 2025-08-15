//||------------------------------------------------------------------------------------------------||
//|| Face
//|| Component for Face/Selfie
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useState, useCallback, useRef, useEffect}    from "react";
import { Target }                                           from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Hooks
//||------------------------------------------------------------------------------------------------||

import { useCamera }                                        from "../../hooks/useCamera";
import { useLuma }                                          from "../../hooks/useLuma";
import { useFacePercent }                                   from "../../hooks/useFacePercent";
import { useFaceDirection }                                 from "../../hooks/useFaceDirection";
import { useCapture }                                       from "../../hooks/useCapture";

//||------------------------------------------------------------------------------------------------||
//|| Utils
//||------------------------------------------------------------------------------------------------||

import { statusLuma, iconLuma, boolLuma }                   from "../../utils/face/statusLuma"; 
import { statusFace, iconFace, boolFace }                   from "../../utils/face/statusFace";
import { statusFaceDir, iconFaceDir, 
         statusFaceCentered, boolFaceCentered, 
         iconFaceCentered }                                 from "../../utils/face/statusFaceDir";
import { statusStep }                                       from "../../utils/face/statusStep";
import { makeTarget, randomTarget }                         from "../../utils/face/makeTarget";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { FaceDirection }                                    from "../../interfaces/verify/id/types";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import ProgressSteps                                        from "../base/ProgressSteps";

//||------------------------------------------------------------------------------------------------||
//|| Const
//||------------------------------------------------------------------------------------------------||

const DEBUG = true;

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface FaceProps {
      onUpload    : (file: File | Blob) => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function Face({ onUpload } : FaceProps) {

      //||------------------------------------------------------------------------------------------------||
	//|| Steps
      //|| 0 = Load Camera
      //|| 1 = Face In Frame
      //|| 2 = Face Movement
      //|| 3 = Face Capture
      //|| 4 = Face Upload
      //|| 5 = Complete
	//||------------------------------------------------------------------------------------------------||

      const [currentStep, setCurrentStep]                   = useState(3);

      //||------------------------------------------------------------------------------------------------||
	//|| State
	//||------------------------------------------------------------------------------------------------||

      const [currentLuma, setCurrentLuma]                   = useState(0);
      const [currentFace, setCurrentFace]                   = useState(0);
      const [currentDir,  setCurrentDir]                    = useState<FaceDirection>("unknown-unknown");

      //||------------------------------------------------------------------------------------------------||
	//|| Face Live
	//||------------------------------------------------------------------------------------------------||

      const [faceMoveCount, setFaceMoveCount]               = useState<number>(0);
      const [faceMoveCheck, setFaceMoveCheck]               = useState<FaceDirection>("unknown-unknown");
      const [faceInFrame, setFaceInFrame]                   = useState<boolean>(false);
      
      //||------------------------------------------------------------------------------------------------||
	//|| useRef
	//||------------------------------------------------------------------------------------------------||

	const canvasRef                                       = useRef<HTMLCanvasElement | null>(null);
      const videoRef                                        = useRef<HTMLVideoElement | null>(null);

      //||------------------------------------------------------------------------------------------------||
	//|| Hooks
	//||------------------------------------------------------------------------------------------------||

      const { cameraError, cameraLoading, cameraReady }     = useCamera( videoRef );
      const { luma }                                        = useLuma(videoRef, { sampleSize: 96, ema: 0.2, enabled: true });
      const { facePct, faceReady, faceError, facesCount }   = useFacePercent(videoRef, { ema: 0.2, enabled: cameraReady });
      const { faceDir, faceDirReady, faceDirError }         = useFaceDirection(videoRef, { enabled: cameraReady });
      const {
            blob,
            previewUrl,
            hasPreview,
            countdown,
            error: captureError,
            captureNow,
            startCountdown,
            retake,
            upload,
      } = useCapture(videoRef, { quality: 0.92, maxSize: 1080, unmirror: true, format: "image/jpeg" });

      // Take-photo handler you can pass around (3s countdown). You were calling `takePhoto` before.
      const takePhoto = () => startCountdown(3);

      const handleNext = async () => {
            if (!blob || !onUpload) return;
            await upload(onUpload);
            setCurrentStep(4);
      };      

      //||------------------------------------------------------------------------------------------------||
	//|| Handle Camera Load
	//||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (cameraReady && currentStep === 0) setCurrentStep(1);
      }, [cameraReady, currentStep]);

      //||------------------------------------------------------------------------------------------------||
	//|| Handle Luma
	//||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            const myLuma = Math.floor(luma);
            if (currentLuma !== myLuma) setCurrentLuma( myLuma );
      }, [ luma ]);

      //||------------------------------------------------------------------------------------------------||
	//|| Handle Face Percentage
	//||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (!faceReady || faceError || facesCount === 0) return setCurrentFace(-1);
            const pct = Math.max(0, Math.min(100, Math.round(facePct)));
            if (pct !== currentFace) setCurrentFace(pct);
            if (facesCount !== 1) return setFaceInFrame(false);
            setFaceInFrame(true);
      }, [ faceReady, facesCount, faceError, facePct ]);

      //||------------------------------------------------------------------------------------------------||
	//|| Handle Face Direction
	//||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (faceDir !== currentDir) setCurrentDir(faceDir);
      }, [ faceDir ]);

      //||------------------------------------------------------------------------------------------------||
	//|| Handle the Steps
	//||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            switch(currentStep) {
                  //||------------------------------------------------------------------------------------------------||
                  //|| Step 1 
                  //||------------------------------------------------------------------------------------------------||
                  case 1 : if (!faceInFrame && currentStep === 1) return setCurrentStep(1); break;
                  //||------------------------------------------------------------------------------------------------||
                  //|| Step 2
                  //||------------------------------------------------------------------------------------------------||
                  case 2 : 
                        if (faceMoveCheck === "unknown-unknown") return setFaceMoveCheck(randomTarget());
                        if (faceMoveCount >= 3) return setCurrentStep(3); 
                        if (faceMoveCheck === currentDir) {
                              setFaceMoveCount(faceMoveCount + 1);
                              setFaceMoveCheck(randomTarget());
                              return;
                        }
                        break;
                  //||------------------------------------------------------------------------------------------------||
                  //|| Step 3
                  //||------------------------------------------------------------------------------------------------||
                  case 3 : break;
                  //||------------------------------------------------------------------------------------------------||
                  //|| Step 3
                  //||------------------------------------------------------------------------------------------------||
            }
      }, [currentStep, faceInFrame, currentFace, currentDir, currentLuma]);

      //||------------------------------------------------------------------------------------------------||
	//|| Spot 
	//||------------------------------------------------------------------------------------------------||

      const spotClass = "w-6 h-6 absolute z-30 text-yellow-400 p-1 border bg-black border-yellow-400"

      //||------------------------------------------------------------------------------------------------||
	//|| JSX
	//||------------------------------------------------------------------------------------------------||

      return (
            <div className="w-full my-5 px-2">
                  <div className="mx-auto max-w-5xl rounded-lg bg-black/30 p-4">
                        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                              {/* Left: Camera Preview */}
                              <div className="relative aspect-square overflow-hidden rounded-xl">
                                    <video
                                          ref={videoRef}
                                          playsInline
                                          muted
                                          autoPlay
                                          className="absolute inset-0 z-0 h-full w-full -scale-x-100 object-cover"
                                    />
                                    <canvas
                                          ref={canvasRef}
                                          className="absolute inset-0 z-10 h-full w-full -scale-x-100"
                                    />
                                    {/* Face Oval Guide */}
                                    <div className="pointer-events-none absolute inset-0 z-20 flex items-center justify-center">
                                          <svg viewBox="0 0 500 500" className="h-full w-full">
                                                <ellipse cx="250" cy="250" rx="180" ry="230" fill="none" stroke="white" strokeWidth="8" />
                                                <ellipse cx="250" cy="250" rx="180" ry="230" fill="none" stroke="rgba(0,0,0,0.25)" strokeWidth="1" />
                                          </svg>
                                    </div>
                                    { currentStep === 2 && ( 
                                          <div className="relative border border-yellow-400 h-full rounded-none">
                                                { makeTarget(currentDir, false) }
                                                { makeTarget(currentDir, true) }
                                          </div>
                                    )}

                              </div>

                              {/* Right: Status / Instructions */}
                              <div className="flex flex-col gap-4 rounded-xl p-2 w-full">
                                    { currentStep > 0 && (
                                          <div className="rounded-md bg-black/30 p-3 text-sm flex justify-center gap-4 w-[95%] mx-auto">
                                                <div className="flex flex-col items-center w-[32%] bg-white/20 rounded-lg p-3">
                                                      { iconLuma(currentLuma) }
                                                      { statusLuma(currentLuma) }
                                                      { DEBUG && ( <span className="text-xs text-gray-400">Luma: { currentLuma }</span> ) }
                                                </div>
                                                <div className="flex flex-col items-center w-[32%] bg-white/20 rounded-lg p-3">
                                                      { iconFace(currentFace) }
                                                      { statusFace(currentFace) }
                                                      { DEBUG && ( <span className="text-xs text-gray-400">Face: { currentFace }</span> ) }
                                                </div>
                                                { currentStep < 3 && (
                                                      <div className="flex flex-col items-center w-[32%] bg-white/20 rounded-lg p-3">
                                                            { iconFaceDir(currentDir) }
                                                            { statusFaceDir(currentDir) }
                                                            { DEBUG && ( <span className="text-xs text-gray-400">Face: { faceDir }</span> ) }
                                                      </div>
                                                )}
                                                { currentStep === 3 && (
                                                      <div className="flex flex-col items-center w-[32%] bg-white/20 rounded-lg p-3">
                                                            { iconFaceCentered(currentDir) }
                                                            { statusFaceCentered(currentDir) }
                                                            { DEBUG && ( <span className="text-xs text-gray-400">Face: { faceDir }</span> ) }
                                                      </div>
                                                )}
                                          </div>
                                    )}

                                    <div id="statusText" className="mt-auto bg-black/40 font-bold text-center py-2 text-2xl">
                                          { cameraReady ?  statusStep(currentStep, currentLuma, currentFace, currentDir, takePhoto) : "Initiating Camera" }
                                    </div>

                                    <div className="mt-auto flex items-center justify-between gap-2">
                                          <div className="flex gap-2">
                                                { currentStep === 4 && (<button className="rounded-md bg-blue-600 px-4 py-2 text-white text-sm hover:bg-blue-500 justify-end">Next</button>) }
                                          </div>
                                    </div>
                              </div>
                        </div>

                        {/* Bottom: Progress */}
                        {currentStep > 0 && ( 
                              <div className="mt-4">
                                    <ProgressSteps maxSteps={4} currentStep={currentStep} />
                              </div>
                        )}
                  </div>
            </div>
      );
}
