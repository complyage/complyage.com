/*||------------------------------------------------------------------------------------------------||
//|| Initiates the camera and provides a video element reference
//|| useCamrea
//||------------------------------------------------------------------------------------------------||*/

      /*||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||*/
      
      import {useCallback, useEffect, useRef, useState}                       from "react";

      /*||------------------------------------------------------------------------------------------------||
      //|| Props
      //||------------------------------------------------------------------------------------------------||*/

      type UseLumaOpts = {      
            sampleSize?       : number;   // canvas sample size for brightness.ts
            enabled?          : boolean;  // start/stop the loop
            ema?              : number;   // 0..1 smoothing factor (e.g. 0.2). Omit for raw luma
      };

      /*||------------------------------------------------------------------------------------------------||
      //|| Hook
      //||------------------------------------------------------------------------------------------------||*/

      export function useLuma( videoRef : React.RefObject<HTMLVideoElement> | null, opts : UseLumaOpts = {} ) {

            /*||------------------------------------------------------------------------------------------------||
            //|| Options
            //||------------------------------------------------------------------------------------------------||*/

            const {sampleSize = 96, enabled = true, ema} = opts;

            /*||------------------------------------------------------------------------------------------------||
            //|| Var
            //||------------------------------------------------------------------------------------------------||*/

            const [luma, setLuma]         = useState(0);
            const rafRef                  = useRef<number | null>(null);
            const smoothedRef             = useRef<number | null>(null);

            /*||------------------------------------------------------------------------------------------------||
            //|| Get Average
            //||------------------------------------------------------------------------------------------------||*/

            const getAverageLuma = (video: HTMLVideoElement, sampleSize = 96): number  => {
                  if (!(video.readyState >= 2) || video.videoWidth === 0 || video.videoHeight === 0) return 0;
                  const c     = document.createElement("canvas");
                  c.width     = sampleSize;
                  c.height    = sampleSize;
                  const ctx   = c.getContext("2d", {willReadFrequently: true})!;
                  ctx.drawImage(video, 0, 0, sampleSize, sampleSize);
                  const {data} = ctx.getImageData(0, 0, sampleSize, sampleSize);
                  let sum = 0;
                  for (let i = 0; i < data.length; i += 4) {
                        const r = data[i],
                              g = data[i + 1],
                              b = data[i + 2];
                        sum += 0.2126 * r + 0.7152 * g + 0.0722 * b;
                  }
                  const pixels = sampleSize * sampleSize;
                  return sum / pixels;
            }


            /*||------------------------------------------------------------------------------------------------||
            //|| Ref
            //||------------------------------------------------------------------------------------------------||*/
            
            if (videoRef == null) return { "luma" : -1, getAverageLuma };
            
            /*||------------------------------------------------------------------------------------------------||
            //|| Tick
            //||------------------------------------------------------------------------------------------------||*/

            const tick = useCallback(async () => {
                  const v = videoRef.current;
                  if (!v || v.readyState < 2) {
                        rafRef.current = requestAnimationFrame(tick);
                        return;
                  }

                  const raw = await getAverageLuma(v, sampleSize);

                  if (typeof ema === "number") {
                        // exponential moving average smoothing
                        if (smoothedRef.current == null) smoothedRef.current = raw;
                        smoothedRef.current = smoothedRef.current + ema * (raw - smoothedRef.current);
                        setLuma(smoothedRef.current);
                  } else {
                        setLuma(raw);
                  }

                  rafRef.current = requestAnimationFrame(tick);
            }, [videoRef, sampleSize, ema]);

            /*||------------------------------------------------------------------------------------------------||
            //|| Get Average
            //||------------------------------------------------------------------------------------------------||*/

            useEffect(() => {
                  if (!enabled) return;
                  rafRef.current = requestAnimationFrame(tick);
                  return () => {
                        if (rafRef.current) cancelAnimationFrame(rafRef.current);
                  };
            }, [enabled, tick]);

            /*||------------------------------------------------------------------------------------------------||
            //|| Return
            //||------------------------------------------------------------------------------------------------||*/

            return { luma, getAverageLuma };
      }
