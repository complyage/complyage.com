//||------------------------------------------------------------------------------------------------||
//|| Get Average Luma from a Video Element
//|| utils/vision/brightness.ts
//||------------------------------------------------------------------------------------------------||

      export async function getAverageLuma(video: HTMLVideoElement, sampleSize = 96): Promise<number> {
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
