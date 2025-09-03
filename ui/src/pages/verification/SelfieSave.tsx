//||------------------------------------------------------------------------------------------------||
//|| Selfie
//|| component/dynamic/vision/Selfie.tsx
//||------------------------------------------------------------------------------------------------||

import React, {useEffect, useRef, useState, useCallback}                from "react";
import {initDetector, type Detector}                                    from "../../utils/vision/detector";
import {computeDirection, type DirVector}                               from "../../utils/vision/computeDirection";
import {getAverageLuma}                                                 from "../../utils/vision/brightness";
import {getFaceAreaPct}                                                 from "../../utils/vision/faceMetrics";
import {captureFrameURL, revokeBlobURL, uploadBlob}                     from "../../utils/vision/capture";
import {X as XIcon, Check, Lightbulb, LightbulbOff, UserCheck, UserX }  from "lucide-react";
import {
      pickRandomSteps, getTargetPoint, humanizeDir, 
      ALLOWED_STEPS, type Dir, type StepId
}                                                                       from "../../utils/vision/steps";
import {waitForVideoReady}                                              from "../../utils/vision/wait";

//||------------------------------------------------------------------------------------------------||
//|| Const 
//||------------------------------------------------------------------------------------------------||

const CANVAS_SIZE       = 500;
const LUMA_MIN          = 10;
const LUMA_MAX          = 30;
const MAX_FACE_PCT      = 30; // % required to ENTER gaze mode
const MIN_FACE_PCT      = 20; // % required to ENTER gaze mode
const MIN_FACE_PCT_STEP = 10; // % required to COUNT dwell while IN gaze mode
const MIRRORED_PREVIEW  = true;
const DWELL_MS          = 700;

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Selfie() {

      //||------------------------------------------------------------------------------------------------||
      //|| useState
      //|| Mode // 0 = get in frame, 1 = direction steps, 2 = ready to snap (countdown), 3 = uploading 4 done      
      //||------------------------------------------------------------------------------------------------||

      const [error, setError]             = useState<string | null>(null);
	const [loading, setLoading]         = useState(true);
	const videoRef                      = useRef<HTMLVideoElement | null>(null);
	const canvasRef                     = useRef<HTMLCanvasElement | null>(null);
	const detectorRef                   = useRef<Detector | null>(null);
	const rafRef                        = useRef<number | null>(null);
	const [debug, setDebug]             = useState(false);
	const [gaze, setGaze]               = useState<Dir>("unknown-unknown");
	const [luma, setLuma]               = useState(0);
	const [facePct, setFacePct]         = useState(0);
	const [mode, setMode]               = useState<0 | 1 | 2 | 3>(0);
	const [steps, setSteps]             = useState<StepId[]>(() => pickRandomSteps(3));
	const [passed, setPassed]           = useState<Record<Dir, boolean>>({});
	const [stepIndex, setStepIndex]     = useState(0);
	const [countdown, setCountdown]     = useState<number | null>(null);
      const [capturedBlob, setCapturedBlob]     = useState<Blob | null>(null);
      const [previewUrl, setPreviewUrl]         = useState<string | null>(null);
      
      //||------------------------------------------------------------------------------------------------||
      //|| useRef
      //||------------------------------------------------------------------------------------------------||

	const holdMsRef                     = useRef(0);
	const lastTsRef                     = useRef<number | null>(null);

      //||------------------------------------------------------------------------------------------------||
      //|| Draw Outline
      //||------------------------------------------------------------------------------------------------||

	const draw = useCallback((faces: any[], ctx: CanvasRenderingContext2D) => {
		ctx.clearRect(0, 0, CANVAS_SIZE, CANVAS_SIZE);
		const f = faces?.[0];
		if (f?.box) {
			const {xMin, yMin, width, height} = f.box;
			ctx.strokeStyle = "rgba(0,255,0,0.8)";
			ctx.lineWidth = 2;
			ctx.strokeRect(xMin, yMin, width, height);
		}
	}, []);

      //||------------------------------------------------------------------------------------------------||
      //|| Draw the Tick
      //||------------------------------------------------------------------------------------------------||

      const tick = useCallback(async () => {
		const video       = videoRef.current!;
		const canvas      = canvasRef.current!;
		const ctx         = canvas.getContext("2d")!;
		const detector    = detectorRef.current;
		if (!detector || !video || video.readyState < 2) {
			rafRef.current = requestAnimationFrame(tick);
			return;
		}
		video.width       = CANVAS_SIZE;
		video.height      = CANVAS_SIZE;
		canvas.width      = CANVAS_SIZE;
		canvas.height     = CANVAS_SIZE;

		// faces
		let faces: any[] = [];
		try {
			faces = await detector.estimateFaces(video);
		} catch {
			// keep going
		}

            //||------------------------------------------------------------------------------------------------||
            //|| Brightness
            //||------------------------------------------------------------------------------------------------||

            const lv = await getAverageLuma(video, 96);
		setLuma(lv);

            //||------------------------------------------------------------------------------------------------||
            //|| Face Detection
            //||------------------------------------------------------------------------------------------------||

		if (faces.length) {
			const f = faces[0];
			setFacePct(getFaceAreaPct(f, CANVAS_SIZE));
			setGaze( computeDirection(f, { mirrored: MIRRORED_PREVIEW}) as Dir);
		} else {
			setFacePct(0);
			setGaze("unknown-unknown");
		}
		draw(faces, ctx);
		rafRef.current = requestAnimationFrame(tick);
	}, [draw]);

      //||------------------------------------------------------------------------------------------------||
      //|| When to pause
      //||------------------------------------------------------------------------------------------------||

	useEffect(() => {
		if (mode !== 1) return;

		const now = performance.now();
		if (lastTsRef.current == null) lastTsRef.current = now;
		const dt = now - lastTsRef.current;
		lastTsRef.current = now;

		// lenient "in frame" while stepping
		const inFrame = facePct >= MIN_FACE_PCT_STEP && luma >= LUMA_MIN;
		if (!inFrame) {
			// pause progress; do not reset/decay
			return;
		}

		// Are we done? (check against the CHOSEN steps)
		if (steps.every((s) => passed[s])) {
			setMode(2);
			return;
		}

		// ensure current step points at the first incomplete
		let idx = stepIndex;
		if (passed[steps[idx]]) {
			const firstIncomplete = steps.findIndex((s) => !passed[s]);
			idx = firstIncomplete === -1 ? steps.length : firstIncomplete;
			if (idx !== stepIndex) setStepIndex(idx);
		}
		if (idx >= steps.length) {
			setMode(2);
			return;
		}

		const target = steps[idx];
		const match = gaze === target;

		if (match) {
			holdMsRef.current += dt;
			if (holdMsRef.current >= DWELL_MS) {
				setPassed((p) => ({...p, [target]: true}));
				holdMsRef.current = 0;
				const nextIdx = idx + 1;
				if (nextIdx >= steps.length) setMode(2);
				else setStepIndex(nextIdx);
			}
		}
		// else: no reset, no decay — just wait
	}, [mode, facePct, luma, steps, stepIndex, passed, gaze]);

      //||------------------------------------------------------------------------------------------------||
      //|| Countdown and Upload
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
		if (mode !== 2) return;

		setCountdown(3); // start fresh
		const id = setInterval(() => {
			setCountdown((c) => {
				if (c === null) return c; // already done
				if (c <= 1) {
					clearInterval(id);
					(async () => {
						try {
							const video = videoRef.current!;
                                          const { blob, url } = await captureFrameURL(videoRef.current!, { quality: 0.92, maxSize: 1080, unmirror: true });
                                          setCapturedBlob(blob);
                                          setPreviewUrl(url);							
                                          setMode(3);
						} catch (e: any) {
							console.error(e);
							setError(e?.message || "Upload failed");
							setMode(0); // restart
						} finally {
							setCountdown(null);
						}
					})();
					return null;
				}
				return c - 1;
			});
		}, 1000);

		return () => clearInterval(id);
	}, [mode]);

      //||------------------------------------------------------------------------------------------------||
      //|| Draw Set Everything up
      //||------------------------------------------------------------------------------------------------||

	useEffect(() => {
		(async () => {
			try {
				setLoading(true);
				setError(null);

				const v = videoRef.current!;
				const stream = await navigator.mediaDevices.getUserMedia({video: true});
				v.srcObject = stream;
				await waitForVideoReady(v);

				detectorRef.current = await initDetector();

				setLoading(false);
				rafRef.current = requestAnimationFrame(tick);
			} catch (e: any) {
				console.error(e);
				setError(e?.message || "Failed to initialize");
				setLoading(false);
			}
		})();

		return () => {
			if (rafRef.current) cancelAnimationFrame(rafRef.current);
			const v = videoRef.current;
			const src = v?.srcObject as MediaStream | null;
			src?.getTracks().forEach((t) => t.stop());
		};
	}, [tick]);

      //||------------------------------------------------------------------------------------------------||
      //|| When to transition mode
      //||------------------------------------------------------------------------------------------------||

	useEffect(() => {
		if (mode === 0) {
			const inFrame = facePct >= MIN_FACE_PCT && luma >= LUMA_MIN;
			if (inFrame) {
				// reset flow
				setPassed({"left-center": false, "right-center": false, "center-top": false});
				setSteps(pickRandomSteps(3));
				setStepIndex(0);
				holdMsRef.current = 0;
				lastTsRef.current = null;
				setMode(1);
			}
		}
	}, [mode, facePct, luma]);

      //||------------------------------------------------------------------------------------------------||
      //|| UI Helpers
      //||------------------------------------------------------------------------------------------------||

	const currentStep       = steps[stepIndex];
	const marker            = currentStep ? getTargetPoint(currentStep) : null;


      const brightIndicator = ( luma : number ) => {
            return luma < LUMA_MIN
            ? <span className="inline-flex items-center gap-1 text-red-500"><LightbulbOff />Too Dark</span>
            : luma > LUMA_MAX
                  ? <span className="inline-flex items-center gap-1 text-red-500"><Lightbulb />Too Bright</span>
                  : <span className="inline-flex items-center gap-1 text-green-500"><Check />Brightness</span>;      
      }

      const faceIndicator = (facePerc: number) => {
            return facePerc < MIN_FACE_PCT
                  ? <span className="inline-flex items-center gap-1 text-red-500"><UserX />Move closer to the camera</span>
                  : <span className="inline-flex items-center gap-1 text-green-500"><UserCheck />Face size is good</span>;
      }; 

      //||------------------------------------------------------------------------------------------------||
      //|| Countdown + Manual Capture (Take → countdown → capture) & Upload
      //||------------------------------------------------------------------------------------------------||

      const countdownRef = useRef<number | null>(null);

      const takePhoto = useCallback(() => {
            // start a fresh countdown
            if (countdownRef.current) clearInterval(countdownRef.current);
            setCapturedBlob(null);
            setPreviewUrl(null);
            setCountdown(3);

            countdownRef.current = window.setInterval(async () => {
                  setCountdown((c) => {
                        if (c === null) return c;
                        if (c <= 1) {
                              if (countdownRef.current) {
                                    clearInterval(countdownRef.current);
                                    countdownRef.current = null;
                              }
                              (async () => {
                                    try {
                                          const { blob, url } = await captureFrameURL(videoRef.current!, { quality: 0.92, maxSize: 1080, unmirror: true });
                                          setCapturedBlob(blob);
                                          setPreviewUrl(url);
                                    } catch (e: any) {
                                          console.error(e);
                                          setError(e?.message || "Failed to capture");
                                    } finally {
                                          setCountdown(null);
                                    }
                              })();
                              return null;
                        }
                        return c - 1;
                  });
            }, 1000);
      }, []);

      const retake = useCallback(() => {
            if (countdownRef.current) {
                  clearInterval(countdownRef.current);
                  countdownRef.current = null;
            }
            revokeBlobURL(previewUrl);
            setCapturedBlob(null);
            setPreviewUrl(null);
            setCountdown(null);
            setMode(2);   // stay on capture step
      }, [previewUrl]);

      const nextUpload = useCallback(async () => {
            if (!capturedBlob) return;
            try {
                  setMode(3); // uploading
                  await uploadBlob("/api/selfie", capturedBlob, { fields: { flow: "face-verification" } });
                  // success: navigate or mark success here
            } catch (e: any) {
                  console.error(e);
                  setError(e?.message || "Upload failed");
                  setMode(0); // restart
            }
      }, [capturedBlob]);

      // cleanup countdown on unmount
      useEffect(() => {
            return () => {
                  if (countdownRef.current) clearInterval(countdownRef.current);
            };
      }, []);




      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

	return (
		<div className="text-white">
			<div className="relative mx-auto mt-6 flex h-[650px] max-w-4xl items-center justify-center text-center">
				{error && <div className="absolute left-1/2 top-2 z-50 -translate-x-1/2 rounded bg-red-600/80 px-3 py-1 text-sm">{error}</div>}                        

				{loading && !error && (
					<div className="absolute left-1/2 top-2 z-50 -translate-x-1/2 rounded bg-black/50 px-3 py-1 text-xs font-bold">Loading Please wait...</div>
				)}

                        {!loading && !error && (
                              <div className="absolute left-1/2 top-2 z-50 -translate-x-1/2 rounded bg-black/50 border-2 w-[50%] h-[50px] px-3 py-1 text-sm">
                                    <div className="flex flex-row h-full">
                                          <div className="flex flex-1 items-center justify-center text-xs">{brightIndicator(luma) }</div>
                                          <div className="flex flex-1 items-center justify-center text-xs">{faceIndicator(facePct) }</div>
                                          <div className="flex flex-1 items-center justify-center text-xs">asdasd</div>
                                    </div>
                              </div>
                        )}

				<div className="relative mx-auto h-[500px] w-[500px] overflow-hidden rounded-full border-8 border-white/20">
					<video ref={videoRef} playsInline muted autoPlay className="absolute inset-0 z-0 h-full w-full -scale-x-100 object-cover" />
					<canvas ref={canvasRef} className="absolute inset-0 z-10 h-full w-full -scale-x-100" />

					{/* oval guide + markers */}
					<div className="pointer-events-none absolute inset-0 z-20 flex items-center justify-center">
						<svg viewBox="0 0 500 500" className="h-full w-full">
							<ellipse cx="250" cy="250" rx="180" ry="230" fill="none" stroke="white" strokeWidth="8" />
							<ellipse cx="250" cy="250" rx="180" ry="230" fill="none" stroke="rgba(0,0,0,0.25)" strokeWidth="1" />

							{/* target marker for current step */}
							{mode === 1 && marker && <circle cx={marker.x} cy={marker.y} r="10" fill="rgba(255,255,0,0.95)" />}

							{/* check marks for passed */}
							{mode <= 2 && Object.entries(passed).map(([k, v]) => {
								if (!v) return null;
								const p = getTargetPoint(k as StepId);
								return (
									<text key={k} x={p.x} y={p.y - 14} textAnchor="middle" fontSize="18" fill="#00ff7f"><Check /></text>
								);
							})}
						</svg>
					</div>

                              {/* capture overlay (manual + countdown) */}
                              {mode === 2 && (
                                    <div className="absolute inset-0 z-30 flex items-center justify-center bg-black/30">
                                          {countdown !== null && capturedBlob === null ? (
                                                <div className="pointer-events-none text-7xl font-bold">
                                                      {countdown}
                                                </div>
                                          ) : capturedBlob ? (
                                                <div className="flex gap-3">
                                                      <button onClick={retake} className="px-4 py-2 rounded bg-neutral-800 hover:bg-neutral-700">
                                                            Retake
                                                      </button>
                                                      <button onClick={nextUpload} className="px-4 py-2 rounded bg-blue-600 hover:bg-blue-500 text-white">
                                                            Next
                                                      </button>
                                                </div>
                                          ) : (
                                                <button onClick={takePhoto} className="px-5 py-3 rounded bg-green-600 hover:bg-green-500 text-white">
                                                      Take Photo
                                                </button>
                                          )}
                                    </div>
                              )}                          

					{/* countdown overlay */}
					{mode === 2 && countdown !== null && (
						<div className="pointer-events-none absolute inset-0 z-30 flex items-center justify-center bg-black/30 text-7xl font-bold">
							{countdown}
						</div>
					)}
				</div>
			</div>
                  
                  <div className="max-w-xl mx-auto text-center bg-black/50 p-3 rounded-lg">
                        <div className="font-semibold text-center py-2">
						{mode === 0 && "Step 1: Get in frame (bright & close enough)"}
						{mode === 1 && `Step 2: Look where shown - ${humanizeDir(currentStep || "center-center")}`}
						{mode === 2 && "Step 3: Hold still for a photo…"}
						{mode === 3 && "Uploading selfie…"}
				</div>                        
                  </div>

			{/* bottom walkthrough bar */}
			<div className="mx-auto mt-4 w-full max-w-2xl rounded-md bg-neutral-900 px-4 py-3 text-sm">
				<div className="mb-2 flex items-center justify-between">
					<label className="flex items-center gap-2 text-xs opacity-80">
						<input type="checkbox" checked={debug} onChange={(e) => setDebug(e.target.checked)} className="h-4 w-4" />
						Debug
					</label>
				</div>

				{/* progress chips */}
				<div className="mb-2 flex flex-wrap gap-2">
					{steps.map((s, i) => {
						const isCurrent = i === stepIndex && mode === 1 && !passed[s];
						const done = !!passed[s];
						return (
							<div
								key={s}
								className={[
									"rounded px-2 py-1 text-xs",
									done ? "bg-green-700/70" : isCurrent ? "bg-yellow-600/70" : "bg-neutral-800",
								].join(" ")}>
								{humanizeDir(s)}
							</div>
						);
					})}
				</div>


				{/* Counting gate hint while in gaze mode */}
				{mode === 1 && (facePct < MIN_FACE_PCT_STEP || luma < LUMA_MIN) && (
					<div className="mb-2 text-xs text-yellow-300/90">
						Not counting yet:&nbsp;
						{facePct < MIN_FACE_PCT_STEP ? `get closer (face ${facePct.toFixed(0)}% < ${MIN_FACE_PCT_STEP}%)` : ""}
						{facePct < MIN_FACE_PCT_STEP && luma < LUMA_MIN ? " and " : ""}
						{luma < LUMA_MIN ? `raise light (luma ${luma.toFixed(0)} < ${LUMA_MIN})` : ""}
					</div>
				)}

				{debug && (
					<div className="grid grid-cols-2 gap-2 text-xs opacity-90">
						<div>
							Direction: <b>{gaze}</b>
						</div>
						<div>
							Face area: <b>{facePct.toFixed(0)}%</b>
						</div>
						<div>
							Luma: <b>{luma.toFixed(0)}</b> (min {LUMA_MIN})
						</div>
						<div>
							Mode: <b>{mode}</b>
						</div>
					</div>
				)}
			</div>
		</div>
	);
}
