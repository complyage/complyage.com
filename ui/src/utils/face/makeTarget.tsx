/*||------------------------------------------------------------------------------------------------||
//|| Make a Target
//|| makeTarget.tsx
//||------------------------------------------------------------------------------------------------||*/

      /*||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||*/

      import React, { JSX }                                 from "react";
      import { Target, ScanEye }                            from "lucide-react";
      import { FaceDirection }                              from "../../interfaces/verify/id/types";

      /*||------------------------------------------------------------------------------------------------||
      //|| Icons
      //||------------------------------------------------------------------------------------------------||*/

      import { Check, ZoomIn, ZoomOut, UserRoundX }         from "lucide-react";

      /*||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||*/

      const MAX_FACE = 35; // Too close
      const MIN_FACE = 24; // Too far

      //||------------------------------------------------------------------------------------------------||
      //|| Random Target
      //||------------------------------------------------------------------------------------------------||

      export function randomTarget(): FaceDirection {
            const directions: FaceDirection[] = [ "left-top", "center-top", "right-top", "left-center", "center-center", "right-center", "left-bottom", "center-bottom", "right-bottom" ];
            const randomIndex = Math.floor(Math.random() * directions.length);
            return directions[randomIndex];
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Make Target
      //||------------------------------------------------------------------------------------------------||

      export function makeTarget(area: FaceDirection, active : boolean): JSX.Element {
            const SpotComponent = ( active ) ? Target : ScanEye; // Use Target if active, otherwise use ScanEye
            let spotClass = "absolute z-30 w-6 h-6 p-1 bg-black border rounded-sm";
            spotClass +=  (active) ? " text-yellow-400 border-yellow-400" : " text-gray-500 border-gray-500";            
            switch (area) {
                  case "left-top":        return <SpotComponent id="left-top" className={`${spotClass} left-[15%] top-[15%]`} />;
                  case "center-top":      return <SpotComponent id="center-top" className={`${spotClass} -translate-x-1/2 left-[50%] top-[0%]`} />;
                  case "right-top":       return <SpotComponent id="right-top" className={`${spotClass} right-[15%] top-[15%]`} />;
                  case "left-center":     return <SpotComponent id="left-center" className={`${spotClass} left-[7%] -translate-y-1/2 top-[50%]`} />;
                  case "center-center":   return <SpotComponent id="center-center" className={`${spotClass} -translate-x-1/2 left-[50%] -translate-y-1/2 top-[50%]`} />;
                  case "right-center":    return <SpotComponent id="right-center" className={`${spotClass} right-[7%] -translate-y-1/2 top-[50%]`} />;
                  case "left-bottom":     return <SpotComponent id="left-bottom" className={`${spotClass} left-[15%] bottom-[15%]`} />;
                  case "center-bottom":   return <SpotComponent id="center-bottom" className={`${spotClass} -translate-x-1/2 left-[50%] bottom-0`} />;
                  case "right-bottom":    return <SpotComponent id="right-bottom" className={`${spotClass} right-[15%] bottom-[15%]`} />;
                  default:                return <></>; // empty element if unknown direction
            }
      }


