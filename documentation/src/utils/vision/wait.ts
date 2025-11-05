//||------------------------------------------------------------------------------------------------||
//|| waitForVideoReady
//|| Utility to ensure <video> has loaded enough metadata before use
//||------------------------------------------------------------------------------------------------||

      export async function waitForVideoReady(video: HTMLVideoElement): Promise<void> {
            if (video.readyState >= 2) {
                  return; // Already ready
            }
            return new Promise((resolve) => {
                  const handler = () => {
                        video.removeEventListener("loadeddata", handler);
                        resolve();
                  };
                  video.addEventListener("loadeddata", handler);
            });
      }
