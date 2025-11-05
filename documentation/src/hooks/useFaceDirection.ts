/*||------------------------------------------------------------------------------------------------||
//|| Face Direction Hook (self-contained)
//|| Estimates gaze/pose as left|center|right and up|center|down
//|| Returns a combined label like: "left-up", "center-center", "right-down"
//|| hooks/useFaceDirection.ts
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

type UseFaceDirectionOpts = {
      size?               : number;   // normalized canvas size (default 500)
      enabled?            : boolean;  // start/stop RAF loop
      ema?                : number;   // smoothing 0..1 for dx/dy (omit for raw)
      flipHorizontal?     : boolean;  // forwarded to estimateFaces
      mirroredPreview?    : boolean;  // set true if your on-screen <video> is mirrored (e.g., CSS -scale-x-100)
      hThreshold?         : number;   // horizontal deadzone (normalized, default 0.08)
      vThreshold?         : number;   // vertical deadzone (normalized, default 0.08)
      smoothWindow?       : number;   // N frames to vote over (default 7)
      smoothConsensus?    : number;   // K votes required (default 5)
      holdMs?             : number;   // must persist before commit (default 120ms)
      cooldownMs?         : number;   // min time between commits (default 600ms)
      hysteresis?         : number;   // add/subtract around thresholds (default 0.02)
      minBoxFrac?         : number;   // ignore updates if face is tiny (default 0.28)
};

type Box = { x: number; y: number; width: number; height: number };
type FaceDir      = "left" | "center" | "right";
type FaceVert     = "bottom" | "center" | "top";

import { FaceDirection }                  from "../interfaces/verify/id/types";

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
//|| Hook: useFaceDirection
//|| - Draws video → normalized square canvas (size x size)
//|| - Detects landmarks with MediaPipe FaceMesh (tfjs)
//|| - Finds a nose anchor (closest Z if available, else bbox center)
//|| - Computes normalized offsets (dx, dy) relative to face bbox center
//|| - Applies thresholds + optional EMA smoothing
//|| - Emits horiz/vert labels and combined "horiz-vert" string
//||------------------------------------------------------------------------------------------------||*/
export function useFaceDirection(
      videoRef: React.RefObject<HTMLVideoElement>,
      opts: UseFaceDirectionOpts = {}
) {
      /*||------------------------------------------------------------------------------------------------||
      //|| Options
      //||------------------------------------------------------------------------------------------------||*/
      const {
            size            = 500,
            enabled         = true,
            ema,
            flipHorizontal  = false,
            mirroredPreview = false,
            hThreshold      = 0.08,
            vThreshold      = 0.08,
            smoothWindow    = 7,
            smoothConsensus = 5,
            holdMs          = 120,
            cooldownMs      = 600,
            hysteresis      = 0.02,
            minBoxFrac      = 0.28,
      } = opts;

      /*||------------------------------------------------------------------------------------------------||
      //|| State / Debug
      //||------------------------------------------------------------------------------------------------||*/
      const [horiz,     setHoriz]               = useState<FaceDir>("center");
      const [vert,      setVert ]               = useState<FaceVert>("center");
      const [faceDir,   setFaceDir ]            = useState<FaceDirection>("center-center");
      const [faceDirReady, setFaceDirReady]     = useState(false);
      const [faceDirError, setFaceDirError]     = useState<string | null>(null);

      /*||------------------------------------------------------------------------------------------------||
      //|| State / Debug
      //||------------------------------------------------------------------------------------------------||*/
      const [facesCount, setFacesCount] = useState(0);
      const [vw, setVw]                 = useState(0);
      const [vh, setVh]                 = useState(0);
      const [lastBox, setLastBox]       = useState<Box | null>(null);
      const [lastNose, setLastNose]     = useState<{x:number;y:number;z?:number}|null>(null);
      const [lastDxDy, setLastDxDy]     = useState<{dx:number;dy:number}>({dx:0,dy:0});

      /*||------------------------------------------------------------------------------------------------||
      //|| Refs
      //||------------------------------------------------------------------------------------------------||*/
      const detectorRef                 = useRef<Detector | null>(null);
      const rafRef                      = useRef<number | null>(null);
      const smoothDxRef                 = useRef<number | null>(null);
      const smoothDyRef                 = useRef<number | null>(null);

      // Label smoothing / gating
      const labelBufRef                 = useRef<FaceDirection[]>([]);
      const committedRef                = useRef<{h:FaceDir; v:FaceVert; label:FaceDirection; at:number}>({
            h: "center", v: "center", label: "center-center", at: performance.now()
      });
      const lastAcceptAtRef             = useRef<number>(0);
      const rawLabelRef                 = useRef<FaceDirection>("center-center");

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
      // Build bbox from a set of 2D or 3D points
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

      // Pick a "nose" anchor (prefer smallest Z; fallback to bbox center)
      const pickNose = (pts: any[], box: Box): {x:number;y:number;z?:number} => {
            let best = null as {x:number;y:number;z?:number} | null;
            let bestZ = Infinity;
            for (const p of pts) {
                  const x = typeof p?.x === "number" ? p.x : (Array.isArray(p) ? p[0] : undefined);
                  const y = typeof p?.y === "number" ? p.y : (Array.isArray(p) ? p[1] : undefined);
                  const z = typeof p?.z === "number" ? p.z : (Array.isArray(p) ? p[2] : undefined);
                  if (typeof x !== "number" || typeof y !== "number") continue;
                  if (typeof z === "number") {
                        if (z < bestZ) { bestZ = z; best = { x, y, z }; }
                  }
            }
            if (best) return best;
            return { x: box.x + box.width/2, y: box.y + box.height/2 };
      };

      // Baseline quantizer (kept for reference)
      const quantizeDir = (dx: number, dy: number): {h: FaceDir; v: FaceVert} => {
            let h: FaceDir = "center";
            let v: FaceVert = "center";
            const adjDx = mirroredPreview ? -dx : dx;
            if (adjDx >  hThreshold) h = "right";
            else if (adjDx < -hThreshold) h = "left";
            if (dy >  vThreshold)         v = "top";
            else if (dy < -vThreshold)    v = "bottom";
            return { h, v };
      };

      // Hysteresis helpers
      function hysteresis1D(prev: "left"|"center"|"right", x: number, th: number, m: number): FaceDir {
            const enter = th + m;
            const exit  = th - m;
            switch (prev) {
                  case "center":
                        if (x >  enter)  return "right";
                        if (x < -enter)  return "left";
                        return "center";
                  case "right":
                        if (x <  exit)   return "center";
                        if (x < -enter)  return "left";
                        return "right";
                  case "left":
                        if (x > -exit)   return "center";
                        if (x >  enter)  return "right";
                        return "left";
            }
      }
      function quantizeWithHys(
            dx: number, dy: number,
            prevH: FaceDir, prevV: FaceVert,
            hTh: number, vTh: number, m: number,
            mirrored: boolean
      ): { h: FaceDir; v: FaceVert } {
            const adjDx = mirrored ? -dx : dx;
            const h = hysteresis1D(prevH, adjDx, hTh, m);
            const vDir = hysteresis1D(
                  prevV === "top" ? "right" : prevV === "bottom" ? "left" : "center",
                  dy, vTh, m
            );
            let v: FaceVert = "center";
            if (vDir === "right") v = "top";
            else if (vDir === "left") v = "bottom";
            else v = "center";
            return { h, v };
      }

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

            // require real source dimensions
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
                  if (srcAspect > 1) {
                        const newW = srcH;
                        sx = (srcW - newW) / 2; sw = newW;
                  } else if (srcAspect < 1) {
                        const newH = srcW;
                        sy = (srcH - newH) / 2; sh = newH;
                  }
                  ctx.drawImage(video, sx, sy, sw, sh, 0, 0, size, size);

                  // detect on normalized canvas
                  const faces: any[] = await detector.estimateFaces(offscreen, { flipHorizontal });
                  setFacesCount(Array.isArray(faces) ? faces.length : 0);

                  if (!faces || faces.length === 0) {
                        setLastBox(null);
                        setLastNose(null);
                        // decay toward center if using EMA
                        if (typeof ema === "number") {
                              if (smoothDxRef.current != null) smoothDxRef.current = smoothDxRef.current * (1 - ema);
                              if (smoothDyRef.current != null) smoothDyRef.current = smoothDyRef.current * (1 - ema);
                        }
                        rafRef.current = requestAnimationFrame(tick);
                        return;
                  }

                  const f = faces[0] as any;

                  // landmarks (keypoints or scaledMesh)
                  const pts = Array.isArray(f.keypoints) && f.keypoints.length
                        ? f.keypoints
                        : (Array.isArray(f.scaledMesh) ? f.scaledMesh : null);

                  // bbox
                  let box: Box | null = null;
                  if (f?.box && f.box.width > 0 && f.box.height > 0) {
                        box = {
                              x: f.box.xMin ?? f.box.x ?? 0,
                              y: f.box.yMin ?? f.box.y ?? 0,
                              width: f.box.width,
                              height: f.box.height,
                        };
                  } else if (pts) {
                        box = bboxFromPoints(pts);
                  }

                  if (!box) {
                        setLastBox(null);
                        setLastNose(null);
                        rafRef.current = requestAnimationFrame(tick);
                        return;
                  }

                  // nose anchor
                  const nose = pts ? pickNose(pts, box) : { x: box.x + box.width/2, y: box.y + box.height/2 };
                  setLastBox(box);
                  setLastNose(nose);

                  // offsets normalized within the face box (center = 0,0)
                  let dx = (nose.x - (box.x + box.width / 2)) / box.width;
                  let dy = (nose.y - (box.y + box.height / 2)) / box.height;

                  // optional EMA smoothing on dx/dy
                  if (typeof ema === "number") {
                        if (smoothDxRef.current == null) smoothDxRef.current = dx;
                        if (smoothDyRef.current == null) smoothDyRef.current = dy;
                        smoothDxRef.current = smoothDxRef.current + ema * (dx - smoothDxRef.current);
                        smoothDyRef.current = smoothDyRef.current + ema * (dy - smoothDyRef.current);
                        dx = smoothDxRef.current;
                        dy = smoothDyRef.current;
                  }

                  setLastDxDy({ dx, dy });

                  // Gate tiny faces: skip updates if face is very small on screen
                  const faceFrac = Math.min(1, Math.max(0, box.width / size));
                  const prevH = committedRef.current.h;
                  const prevV = committedRef.current.v;

                  if (faceFrac < minBoxFrac) {
                        const { h: hTmp, v: vTmp } = quantizeWithHys(dx, dy, prevH, prevV, hThreshold, vThreshold, hysteresis, mirroredPreview);
                        rawLabelRef.current = `${hTmp}-${vTmp}` as FaceDirection;
                        // keep committed state in React state
                        if (horiz !== committedRef.current.h) setHoriz(committedRef.current.h);
                        if (vert  !== committedRef.current.v) setVert (committedRef.current.v);
                        if (faceDir !== committedRef.current.label) setFaceDir(committedRef.current.label);
                        rafRef.current = requestAnimationFrame(tick);
                        return;
                  }

                  // Candidate label with hysteresis around committed state
                  const { h: hCand, v: vCand } = quantizeWithHys(dx, dy, prevH, prevV, hThreshold, vThreshold, hysteresis, mirroredPreview);
                  const cand = `${hCand}-${vCand}` as FaceDirection;
                  rawLabelRef.current = cand;

                  // Consensus over buffer
                  const buf = labelBufRef.current;
                  buf.push(cand);
                  if (buf.length > smoothWindow) buf.shift();

                  let mode: FaceDirection = committedRef.current.label;
                  let best = -1;
                  const counts = new Map<FaceDirection, number>();
                  for (const l of buf) counts.set(l, (counts.get(l) || 0) + 1);
                  for (const [k, c] of counts) { if (c > best) { best = c; mode = k as FaceDirection; } }

                  // Dwell + cooldown
                  const now    = performance.now();
                  const enough = best >= smoothConsensus;
                  const dwell  = (now - committedRef.current.at) >= holdMs;
                  const cool   = (now - lastAcceptAtRef.current) >= cooldownMs;

                  if (mode !== committedRef.current.label && enough && dwell && cool) {
                        const [mh, mv] = mode.split("-") as [FaceDir, FaceVert];
                        committedRef.current = { h: mh, v: mv, label: mode, at: now };
                        lastAcceptAtRef.current = now;
                        if (horiz !== mh) setHoriz(mh);
                        if (vert  !== mv) setVert (mv);
                        if (faceDir !== mode) setFaceDir(mode);
                  } else {
                        if (horiz !== committedRef.current.h) setHoriz(committedRef.current.h);
                        if (vert  !== committedRef.current.v) setVert (committedRef.current.v);
                        if (faceDir !== committedRef.current.label) setFaceDir(committedRef.current.label);
                  }

            } catch (e: any) {
                  // keep looping even on a bad frame
            }

            rafRef.current = requestAnimationFrame(tick);
      }, [videoRef, enabled, ema, offscreen, size, flipHorizontal, mirroredPreview, hThreshold, vThreshold, smoothWindow, smoothConsensus, holdMs, cooldownMs, hysteresis, minBoxFrac, horiz, vert, faceDir]);

      /*||------------------------------------------------------------------------------------------------||
      //|| Init / Cleanup
      //||------------------------------------------------------------------------------------------------||*/
      useEffect(() => {
            let mounted = true;
            (async () => {
                  try {
                        setFaceDirError(null);
                        detectorRef.current = await initDetector();
                        if (!mounted) return;
                        setFaceDirReady(true);
                        if (enabled) rafRef.current = requestAnimationFrame(tick);
                  } catch (e: any) {
                        setFaceDirError(e?.message || "Failed to initialize face detector");
                        setFaceDirReady(false);
                  }
            })();
            return () => {
                  mounted = false;
                  if (rafRef.current) cancelAnimationFrame(rafRef.current);
                  // detector is a shared singleton; don't dispose here
            };
      }, [tick, enabled]);

      /*||------------------------------------------------------------------------------------------------||
      //|| Enable/Disable (start/stop RAF)
      //||------------------------------------------------------------------------------------------------||*/
      useEffect(() => {
            if (!faceDirReady) return;
            if (enabled) {
                  if (!rafRef.current) rafRef.current = requestAnimationFrame(tick);
            } else if (rafRef.current) {
                  cancelAnimationFrame(rafRef.current);
                  rafRef.current = null;
            }
      }, [enabled, faceDirReady, tick]);

      /*||------------------------------------------------------------------------------------------------||
      //|| Return
      //||------------------------------------------------------------------------------------------------||*/
      return {
            faceDir,
            faceDirReady,
            faceDirError,
            debug: {
                  facesCount,
                  vw, vh,
                  lastBox,
                  lastNose,
                  lastDxDy,
                  thresholds: { hThreshold, vThreshold, hysteresis },
                  options: { mirroredPreview, flipHorizontal },
                  rawDir: rawLabelRef.current,
                  buffer: [...labelBufRef.current],
            }
      };
}
