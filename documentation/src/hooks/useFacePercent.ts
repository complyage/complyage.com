/*||------------------------------------------------------------------------------------------------||
//|| Face Percent Hook (self-contained)
//|| Combines detector init + useFacePercent into one module
//|| hooks/useFacePercent.ts
//||------------------------------------------------------------------------------------------------||*/

      /*||------------------------------------------------------------------------------------------------||
      //|| Imports
      //||------------------------------------------------------------------------------------------------||*/

      import { useCallback, useEffect, useMemo, useRef, useState } from "react";
      import * as tf                                               from "@tensorflow/tfjs";
      import * as faceLandmarksDetection                           from "@tensorflow-models/face-landmarks-detection";
      import "@tensorflow/tfjs-backend-webgl";

      /*||------------------------------------------------------------------------------------------------||
      //|| Types
      //||------------------------------------------------------------------------------------------------||*/

      export type Detector = Awaited<ReturnType<typeof faceLandmarksDetection.createDetector>>;

      type UseFacePercentOpts = {
            size?            : number;     // normalized canvas size (default 500)
            enabled?         : boolean;    // start/stop RAF loop
            ema?             : number;     // smoothing 0..1 (omit for raw)
            flipHorizontal?  : boolean;    // forwarded to estimateFaces
      };

      type Box = { x: number; y: number; width: number; height: number };

      /*||------------------------------------------------------------------------------------------------||
      //|| Detector (lazy singleton)
      //||------------------------------------------------------------------------------------------------||*/

      let detectorPromise: Promise<Detector> | null = null;

      async function initDetector(): Promise<Detector> {
            if (!detectorPromise) {
                  detectorPromise = (async () => {
                        await tf.setBackend("webgl");
                        await tf.ready();
                        return faceLandmarksDetection.createDetector(
                              faceLandmarksDetection.SupportedModels.MediaPipeFaceMesh,
                              { runtime: "tfjs", refineLandmarks: true, maxFaces: 1 }
                        );
                  })();
            }
            return detectorPromise;
      }

      /*||------------------------------------------------------------------------------------------------||
      //|| (Optional) Wait for <video> readiness
      //||------------------------------------------------------------------------------------------------||*/

      export async function waitForVideoReady(video: HTMLVideoElement): Promise<void> {
            await video.play().catch(() => {});
            await new Promise<void>((resolve) => {
                  const tryReady = () => {
                        if (video.readyState >= 2 && video.videoWidth > 0 && video.videoHeight > 0) resolve();
                        else requestAnimationFrame(tryReady);
                  };
                  tryReady();
            });
      }

      /*||------------------------------------------------------------------------------------------------||
      //|| Hook: useFacePercent
      //|| - Draws video onto a normalized offscreen canvas (size x size)
      //|| - Runs detector.estimateFaces(offscreen)
      //|| - Uses face.box when present, else derives bbox from landmarks
      //|| - Returns smoothed % of frame area occupied by the first face
      //||------------------------------------------------------------------------------------------------||*/

      export function useFacePercent(
            videoRef: React.RefObject<HTMLVideoElement>,
            opts: UseFacePercentOpts = {}
      ) {

            /*||------------------------------------------------------------------------------------------------||
            //|| Options
            //||------------------------------------------------------------------------------------------------||*/

            const {
                  size = 500,
                  enabled = true,
                  ema,
                  flipHorizontal = false,
            } = opts;

            /*||------------------------------------------------------------------------------------------------||
            //|| State / Debug
            //||------------------------------------------------------------------------------------------------||*/

            const [facePct, setFacePct]       = useState(0);
            const [faceReady, setFaceReady]   = useState(false);
            const [faceError, setError]           = useState<string | null>(null);

            const [facesCount, setFacesCount] = useState(0);
            const [vw, setVw]                 = useState(0);
            const [vh, setVh]                 = useState(0);
            const [lastBox, setLastBox]       = useState<Box | null>(null);

            /*||------------------------------------------------------------------------------------------------||
            //|| Refs
            //||------------------------------------------------------------------------------------------------||*/

            const detectorRef                 = useRef<Detector | null>(null);
            const rafRef                      = useRef<number | null>(null);
            const smoothRef                   = useRef<number | null>(null);

            /*||------------------------------------------------------------------------------------------------||
            //|| Offscreen canvas (reused)
            //||------------------------------------------------------------------------------------------------||*/

            const offscreen = useMemo(() => {
                  if (typeof document === "undefined") return null;
                  const c = document.createElement("canvas");
                  c.width = size;
                  c.height = size;
                  return c;
            }, [size]);

            /*||------------------------------------------------------------------------------------------------||
            //|| Helpers
            //||------------------------------------------------------------------------------------------------||*/

            const bboxFromPoints = (pts: any[]): Box | null => {
                  if (!Array.isArray(pts) || pts.length === 0) return null;
                  let minX = Infinity, minY = Infinity, maxX = -Infinity, maxY = -Infinity;
                  for (const p of pts) {
                        const x = typeof p?.x === "number" ? p.x : (Array.isArray(p) ? p[0] : undefined);
                        const y = typeof p?.y === "number" ? p.y : (Array.isArray(p) ? p[1] : undefined);
                        if (typeof x !== "number" || typeof y !== "number") continue;
                        if (x < minX) minX = x; if (x > maxX) maxX = x;
                        if (y < minY) minY = y; if (y > maxY) maxY = y;
                  }
                  if (!(isFinite(minX) && isFinite(minY) && isFinite(maxX) && isFinite(maxY))) return null;
                  const w = maxX - minX, h = maxY - minY;
                  if (w <= 0 || h <= 0) return null;
                  return { x: minX, y: minY, width: w, height: h };
            };

            /*||------------------------------------------------------------------------------------------------||
            //|| Tick (RAF)
            //||------------------------------------------------------------------------------------------------||*/

            const tick = useCallback(async () => {
                  const video = videoRef.current;
                  const detector = detectorRef.current;

                  if (!enabled) return;
                  if (!video || !detector || !offscreen) {
                        rafRef.current = requestAnimationFrame(tick);
                        return;
                  }

                  // ensure source dimensions are ready
                  if (video.readyState < 2 || video.videoWidth === 0 || video.videoHeight === 0) {
                        rafRef.current = requestAnimationFrame(tick);
                        return;
                  }

                  setVw(video.videoWidth);
                  setVh(video.videoHeight);

                  try {
                        // draw video → square offscreen (object-fit: cover)
                        const ctx = offscreen.getContext("2d", { willReadFrequently: true })!;
                        ctx.clearRect(0, 0, size, size);

                        const srcW = video.videoWidth, srcH = video.videoHeight;
                        const srcAspect = srcW / srcH;
                        let sx = 0, sy = 0, sw = srcW, sh = srcH;
                        if (srcAspect > 1) {                       // crop sides for square
                              const newW = srcH;
                              sx = (srcW - newW) / 2; sw = newW;
                        } else if (srcAspect < 1) {                // crop top/bottom for square
                              const newH = srcW;
                              sy = (srcH - newH) / 2; sh = newH;
                        }
                        ctx.drawImage(video, sx, sy, sw, sh, 0, 0, size, size);

                        // detect on normalized canvas
                        const faces: any[] = await detector.estimateFaces(offscreen, { flipHorizontal });
                        setFacesCount(Array.isArray(faces) ? faces.length : 0);

                        let pct = 0;
                        let box: Box | null = null;

                        if (faces && faces.length > 0) {
                              const f = faces[0] as any;

                              // 1) direct box if available
                              const b = f?.box;
                              if (b && b.width > 0 && b.height > 0) {
                                    box = {
                                          x: b.xMin ?? b.x ?? 0,
                                          y: b.yMin ?? b.y ?? 0,
                                          width: b.width,
                                          height: b.height,
                                    };
                              }

                              // 2) derive from landmarks (keypoints or scaledMesh)
                              if (!box) {
                                    const pts = Array.isArray(f.keypoints) && f.keypoints.length
                                          ? f.keypoints
                                          : (Array.isArray(f.scaledMesh) ? f.scaledMesh : null);
                                    if (pts) box = bboxFromPoints(pts);
                              }

                              if (box) {
                                    pct = ((box.width * box.height) / (size * size)) * 100;
                                    if (pct < 0) pct = 0;
                                    if (pct > 100) pct = 100;
                              }
                        }

                        setLastBox(box);

                        // smoothing
                        if (typeof ema === "number") {
                              if (smoothRef.current == null) smoothRef.current = pct;
                              smoothRef.current = smoothRef.current + ema * (pct - smoothRef.current);
                              setFacePct(smoothRef.current);
                        } else {
                              setFacePct(pct);
                        }
                  } catch (e: any) {
                        // keep looping; optionally surface per-frame errors if needed
                  }

                  rafRef.current = requestAnimationFrame(tick);
            }, [videoRef, enabled, ema, offscreen, size, flipHorizontal]);

            /*||------------------------------------------------------------------------------------------------||
            //|| Init / Cleanup
            //||------------------------------------------------------------------------------------------------||*/

            useEffect(() => {
                  let mounted = true;
                  (async () => {
                        try {
                              setError(null);
                              detectorRef.current = await initDetector();
                              if (!mounted) return;
                              setFaceReady(true);
                              if (enabled) rafRef.current = requestAnimationFrame(tick);
                        } catch (e: any) {
                              setError(e?.message || "Failed to initialize face detector");
                              setFaceReady(false);
                        }
                  })();
                  return () => {
                        mounted = false;
                        if (rafRef.current) cancelAnimationFrame(rafRef.current);
                        // note: detector is a shared singleton; do not dispose here
                        // if you need disposal across app lifetime, add a central disposer
                  };
            }, [tick, enabled]);

            /*||------------------------------------------------------------------------------------------------||
            //|| Enable/Disable (start/stop RAF)
            //||------------------------------------------------------------------------------------------------||*/

            useEffect(() => {
                  if (!faceReady) return;
                  if (enabled) {
                        if (!rafRef.current) rafRef.current = requestAnimationFrame(tick);
                  } else if (rafRef.current) {
                        cancelAnimationFrame(rafRef.current);
                        rafRef.current = null;
                  }
            }, [enabled, faceReady, tick]);

            /*||------------------------------------------------------------------------------------------------||
            //|| Return
            //||------------------------------------------------------------------------------------------------||*/

            return { facePct, faceReady, faceError, facesCount, faceDebug: { vw, vh, lastBox } };
      }
