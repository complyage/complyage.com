/*||------------------------------------------------------------------------------------------------||
//|| Initiates the camera and provides a video element reference
//|| useCamrea
//||------------------------------------------------------------------------------------------------||*/

      /*||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||*/
      
      import { useEffect, useRef, useState, useCallback }                     from "react";
      import { waitForVideoReady }                                            from "../utils/vision/wait";

      /*||------------------------------------------------------------------------------------------------||
      //|| Hook
      //||------------------------------------------------------------------------------------------------||*/

      export function useCamera( videoRef: React.RefObject<HTMLVideoElement> | null ) {

            /*||------------------------------------------------------------------------------------------------||
            //|| Ref/State
            //||------------------------------------------------------------------------------------------------||*/

            const [cameraError,   setCameraError]     = useState<string | null>(null);
            const [cameraLoading, setCameraLoading]   = useState(false);
            const [cameraReady,   setCameraReady]     = useState(false);

            /*||------------------------------------------------------------------------------------------------||
            //|| Start Camera
            //||------------------------------------------------------------------------------------------------||*/

            const startCamera = useCallback(async () => {
                  if (!videoRef.current) return;
                  try {
                        setCameraLoading(true);
                        setCameraError(null);
                        setCameraReady(false);

                        const stream = await navigator.mediaDevices.getUserMedia({ video: true });
                        videoRef.current.srcObject = stream;
                        await waitForVideoReady(videoRef.current);

                        setCameraReady(true); // ✅ fully initiated
                  } catch (e: any) {
                        setCameraError(e?.message || "Failed to initialize camera");
                  } finally {
                        setCameraLoading(false);
                  }
            }, []);

            /*||------------------------------------------------------------------------------------------------||
            //|| Handle Loading
            //||------------------------------------------------------------------------------------------------||*/

            useEffect(() => {
                  startCamera();

                  return () => {
                        const v = videoRef.current;
                        const src = v?.srcObject as MediaStream | null;
                        src?.getTracks().forEach((t) => t.stop());
                  };
            }, [startCamera]);

            /*||------------------------------------------------------------------------------------------------||
            //|| Return Values
            //||------------------------------------------------------------------------------------------------||*/

            return { cameraError, cameraLoading, cameraReady, startCamera };
      }
