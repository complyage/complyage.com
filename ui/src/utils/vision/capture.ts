/*||------------------------------------------------------------------------------------------------||
//|| Manual Capture Flow Helpers
//|| - Manage countdown, capture, retake, and upload
//||------------------------------------------------------------------------------------------------||*/

      /*||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||*/

      export type ManualCaptureState = {
            capturedBlob            : Blob | null;
            previewUrl              : string | null;
            countdown               : number | null;
      };

      /*||------------------------------------------------------------------------------------------------||
      //|| Create Manual Capture
      //||------------------------------------------------------------------------------------------------||*/

      export function createManualCapture(
            videoRef: React.RefObject<HTMLVideoElement>,
            setState: React.Dispatch<React.SetStateAction<ManualCaptureState>>,
            setMode: (mode: number) => void,
            setError: (err: string | null) => void
      ) {
            let countdownTimer: number | null = null;

            const takePhoto = () => {
                  if (countdownTimer) clearInterval(countdownTimer);
                  setState((s) => ({ ...s, capturedBlob: null, previewUrl: null, countdown: 3 }));

                  countdownTimer = window.setInterval(async () => {
                        setState((s) => {
                              if (s.countdown === null) return s;
                              if (s.countdown <= 1) {
                                    if (countdownTimer) {
                                          clearInterval(countdownTimer);
                                          countdownTimer = null;
                                    }
                                    (async () => {
                                          try {
                                                const { blob, url } = await captureFrameURL(videoRef.current!, { quality: 0.92, maxSize: 1080, unmirror: true });
                                                setState((prev) => ({ ...prev, capturedBlob: blob, previewUrl: url, countdown: null }));
                                          } catch (e: any) {
                                                console.error(e);
                                                setError(e?.message || "Failed to capture");
                                          }
                                    })();
                                    return { ...s, countdown: null };
                              }
                              return { ...s, countdown: (s.countdown ?? 0) - 1 };
                        });
                  }, 1000);
            };

            /*||------------------------------------------------------------------------------------------------||
            //|| Retake
            //||------------------------------------------------------------------------------------------------||*/

            const retake = () => {
                  if (countdownTimer) {
                        clearInterval(countdownTimer);
                        countdownTimer = null;
                  }
                  setState((s) => {
                        revokeBlobURL(s.previewUrl);
                        return { ...s, capturedBlob: null, previewUrl: null, countdown: null };
                  });
                  setMode(2);
            };

            /*||------------------------------------------------------------------------------------------------||
            //|| Next Upload
            //||------------------------------------------------------------------------------------------------||*/

            const nextUpload = async (uploadUrl: string, fields?: Record<string, any>) => {
                  setState((s) => s); // ensure state up to date
                  const blob = (stateRef as any)?.current?.capturedBlob; // you'll need to pass stateRef if needed
                  if (!blob) return;
                  try {
                        setMode(3);
                        await uploadBlob(uploadUrl, blob, { fields });
                  } catch (e: any) {
                        console.error(e);
                        setError(e?.message || "Upload failed");
                        setMode(0);
                  }
            };

            return { takePhoto, retake, nextUpload };
      }
