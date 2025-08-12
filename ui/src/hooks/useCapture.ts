/*||------------------------------------------------------------------------------------------------||
//|| usePhotoCapture
//|| Capture a photo from a <video>, preview it, and keep a Blob for upload
//|| hooks/usePhotoCapture.ts
//||------------------------------------------------------------------------------------------------||*/

import { useCallback, useEffect, useMemo, useRef, useState } from "react";

/*||------------------------------------------------------------------------------------------------||
//|| Types
//||------------------------------------------------------------------------------------------------||*/

type CaptureOpts = {
      quality?      : number;   // 0..1 (jpeg/webp)
      maxSize?      : number;   // longest side in px
      unmirror?     : boolean;  // flip back if your preview video is mirrored (-scale-x-100)
      format?       : "image/jpeg" | "image/webp" | "image/png";
};

type UploadFn = (blob: Blob) => Promise<any>;

/*||------------------------------------------------------------------------------------------------||
//|| Hook
//||------------------------------------------------------------------------------------------------||*/

export function useCapture(
      videoRef: React.RefObject<HTMLVideoElement>,
      opts: CaptureOpts = {}
) {

      /*||------------------------------------------------------------------------------------------------||
      //|| Options
      //||------------------------------------------------------------------------------------------------||*/

      const {
            quality  = 0.92,
            maxSize  = 1080,
            unmirror = true,
            format   = "image/jpeg",
      } = opts;

      /*||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||*/

      const [capturing, setCapturing]     = useState(false);
      const [countdown, setCountdown]     = useState<number | null>(null);
      const [error, setError]             = useState<string | null>(null);

      const [blob, setBlob]               = useState<Blob | null>(null);
      const [previewUrl, setPreviewUrl]   = useState<string | null>(null);

      /*||------------------------------------------------------------------------------------------------||
      //|| Refs
      //||------------------------------------------------------------------------------------------------||*/

      const timerRef                      = useRef<number | null>(null);

      /*||------------------------------------------------------------------------------------------------||
      //|| Offscreen Canvas
      //||------------------------------------------------------------------------------------------------||*/

      const canvas = useMemo(() => {
            if (typeof document === "undefined") return null;
            return document.createElement("canvas");
      }, []);

      /*||------------------------------------------------------------------------------------------------||
      //|| Helpers
      //||------------------------------------------------------------------------------------------------||*/

      const revokeURL = useCallback((url?: string | null) => {
            try { if (url) URL.revokeObjectURL(url); } catch {}
      }, []);

      const drawToCanvas = useCallback((video: HTMLVideoElement) => {
            if (!canvas) throw new Error("Canvas unavailable");

            // source dims
            const sw = video.videoWidth;
            const sh = video.videoHeight;
            if (!sw || !sh) throw new Error("Video not ready");

            // scale to fit within maxSize (preserve aspect)
            const scale = Math.min(1, maxSize / Math.max(sw, sh));
            const dw = Math.round(sw * scale);
            const dh = Math.round(sh * scale);

            canvas.width  = dw;
            canvas.height = dh;

            const ctx = canvas.getContext("2d", { willReadFrequently: true })!;
            ctx.save();
            if (unmirror) {
                  // flip horizontally to undo preview mirroring
                  ctx.translate(dw, 0);
                  ctx.scale(-1, 1);
            }
            ctx.drawImage(video, 0, 0, dw, dh);
            ctx.restore();
      }, [canvas, maxSize, unmirror]);

      const canvasToBlob = useCallback(async (): Promise<Blob> => {
            if (!canvas) throw new Error("Canvas unavailable");
            return new Promise<Blob>((resolve, reject) => {
                  if (format === "image/png") {
                        canvas.toBlob((b) => (b ? resolve(b) : reject(new Error("toBlob failed"))), format);
                  } else {
                        canvas.toBlob((b) => (b ? resolve(b) : reject(new Error("toBlob failed"))), format, quality);
                  }
            });
      }, [canvas, format, quality]);

      /*||------------------------------------------------------------------------------------------------||
      //|| Capture (immediate)
      //||------------------------------------------------------------------------------------------------||*/

      const captureNow = useCallback(async () => {
            const video = videoRef.current;
            if (!video) return;

            try {
                  setCapturing(true);
                  setError(null);

                  // draw current frame
                  drawToCanvas(video);

                  // encode → blob
                  const b = await canvasToBlob();

                  // preview url
                  revokeURL(previewUrl);
                  const url = URL.createObjectURL(b);

                  setBlob(b);
                  setPreviewUrl(url);
            } catch (e: any) {
                  setError(e?.message || "Failed to capture");
            } finally {
                  setCapturing(false);
            }
      }, [videoRef, drawToCanvas, canvasToBlob, revokeURL, previewUrl]);

      /*||------------------------------------------------------------------------------------------------||
      //|| Countdown → Capture
      //||------------------------------------------------------------------------------------------------||*/

      const startCountdown = useCallback((sec = 3) => {
            if (timerRef.current) window.clearInterval(timerRef.current);
            setBlob(null);
            revokeURL(previewUrl);
            setPreviewUrl(null);
            setError(null);
            setCountdown(sec);

            timerRef.current = window.setInterval(async () => {
                  setCountdown((c) => {
                        if (c === null) return c;
                        if (c <= 1) {
                              if (timerRef.current) {
                                    window.clearInterval(timerRef.current);
                                    timerRef.current = null;
                              }
                              // fire capture
                              captureNow();
                              return null;
                        }
                        return c - 1;
                  });
            }, 1000);
      }, [captureNow, previewUrl, revokeURL]);

      /*||------------------------------------------------------------------------------------------------||
      //|| Retake / Clear
      //||------------------------------------------------------------------------------------------------||*/

      const retake = useCallback(() => {
            if (timerRef.current) {
                  window.clearInterval(timerRef.current);
                  timerRef.current = null;
            }
            setCountdown(null);
            setBlob(null);
            revokeURL(previewUrl);
            setPreviewUrl(null);
            setError(null);
      }, [previewUrl, revokeURL]);

      /*||------------------------------------------------------------------------------------------------||
      //|| Upload (pass your own uploader)
      //||------------------------------------------------------------------------------------------------||*/

      const upload = useCallback(async (uploader: UploadFn) => {
            if (!blob) throw new Error("No photo to upload");
            return uploader(blob);
      }, [blob]);

      /*||------------------------------------------------------------------------------------------------||
      //|| Cleanup on unmount
      //||------------------------------------------------------------------------------------------------||*/

      useEffect(() => {
            return () => {
                  if (timerRef.current) window.clearInterval(timerRef.current);
                  revokeURL(previewUrl);
            };
      }, [previewUrl, revokeURL]);

      /*||------------------------------------------------------------------------------------------------||
      //|| Return
      //||------------------------------------------------------------------------------------------------||*/

      return {
            // state
            capturing,
            countdown,
            error,

            // result
            blob,
            previewUrl,

            // actions
            captureNow,
            startCountdown,
            retake,
            upload,

            // convenience
            hasPreview: !!previewUrl,
      };
}
