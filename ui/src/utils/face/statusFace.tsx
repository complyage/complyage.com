/*||------------------------------------------------------------------------------------------------||
//|| Return Back the Status of the Face Percentage
//|| statusFace.tsx
//||------------------------------------------------------------------------------------------------||*/

      /*||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||*/

      import React, { JSX }                                 from "react";

      /*||------------------------------------------------------------------------------------------------||
      //|| Icons
      //||------------------------------------------------------------------------------------------------||*/

      import { Check, ZoomIn, ZoomOut, UserRoundX }         from "lucide-react";

      /*||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||*/

      const MAX_FACE = 35; // Too close
      const MIN_FACE = 24; // Too far

      /*||------------------------------------------------------------------------------------------------||
      //|| Status
      //||------------------------------------------------------------------------------------------------||*/

      export function statusFace(facePercent: number): JSX.Element {
            const textClass = "flex flex-col px-2 py-1 text-md font-bold border-gray-500 border-t mt-2 w-[80%] text-center";
            if (facePercent === -1)      return <span className={`${textClass} text-yellow-400`}>Face Not Found</span>;
            if (facePercent >= MAX_FACE) return <span className={`${textClass} text-yellow-400`}>Face Too Close</span>;
            if (facePercent <= MIN_FACE) return <span className={`${textClass} text-red-400`}>Face Too Far</span>;
            return <span className={`${textClass} text-green-400`}>Good</span>;
      }

      /*||------------------------------------------------------------------------------------------------||
      //|| Icon
      //||------------------------------------------------------------------------------------------------||*/

      export function iconFace(facePercent: number): JSX.Element {
            const iconClass = "w-10 h-10 rounded-full flex flex-col";
            if (facePercent === -1)             return <UserRoundX className={`${iconClass} text-gray-500`} />;
            if (facePercent >= MAX_FACE)        return <ZoomOut className={`${iconClass} text-yellow-500`} />;
            if (facePercent <= MIN_FACE)        return <ZoomIn className={`${iconClass} text-red-500`} />;
            return <Check className={`${iconClass} text-green-400`} />;
      }

      /*||------------------------------------------------------------------------------------------------||
      //|| Safe Face
      //||------------------------------------------------------------------------------------------------||*/

      export function boolFace(facePercent: number): boolean {
            return facePercent >= MIN_FACE && facePercent <= MAX_FACE;
      }

      /*||------------------------------------------------------------------------------------------------||
      //|| Int Face
      //||------------------------------------------------------------------------------------------------||*/

      export function intFace(facePercent: number): number {
            if (facePercent === -1) return -1; // Face not found
            if (facePercent >= MAX_FACE) return 100; // Too close
            if (facePercent <= MIN_FACE) return 0; // Too far
            return 50; // Valid face percentage
      }

