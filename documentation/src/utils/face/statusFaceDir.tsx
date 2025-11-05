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

      import { Check, CircleDashed, CircleDot, CircleChevronUp }   from "lucide-react";

      /*||------------------------------------------------------------------------------------------------||
      //|| Interfaces
      //||------------------------------------------------------------------------------------------------||*/

      import { FaceDirection }                              from "../../interfaces/verify/id/types";

      /*||------------------------------------------------------------------------------------------------||
      //|| Status
      //||------------------------------------------------------------------------------------------------||*/

      export function statusFaceDir(faceDir: FaceDirection): JSX.Element {
            const textClass = "flex flex-col px-2 py-1 text-md font-bold border-gray-500 border-t mt-2 w-[80%] text-center";
            if (faceDir === "unknown-unknown") return <span className={`${textClass} text-blue-400`}>N/A</span>;      
            return <span className={`${textClass} text-green-400`}>{ faceDir }</span>;
      }

      /*||------------------------------------------------------------------------------------------------||
      //|| Icon
      //||------------------------------------------------------------------------------------------------||*/

      export function iconFaceDir(faceDir: FaceDirection): JSX.Element {
            const iconClass = "w-10 h-10 rounded-full flex flex-col";      
            switch (faceDir) {
                  case "center-center":         return <CircleDot className={`${iconClass} text-blue-400`} />;
                  case "center-top":            return <CircleChevronUp className={`${iconClass} text-red-400 rotate-0`} />;
                  case "center-bottom":         return <CircleChevronUp className={`${iconClass} text-red-400 rotate-180`} />;
                  case "left-center":           return <CircleChevronUp className={`${iconClass} text-red-400 -rotate-90`} />;
                  case "left-top":              return <CircleChevronUp className={`${iconClass} text-red-400 -rotate-45`} />;
                  case "left-bottom":           return <CircleChevronUp className={`${iconClass} text-red-400 -rotate-135`} />;
                  case "right-center":          return <CircleChevronUp className={`${iconClass} text-red-400 rotate-90`} />;
                  case "right-top":             return <CircleChevronUp className={`${iconClass} text-red-400 rotate-45`} />;
                  case "right-bottom":          return <CircleChevronUp className={`${iconClass} text-red-400 rotate-135`} />;
            }
            return <CircleDashed className={`${iconClass} text-blue-400`} />;
      }

      /*||------------------------------------------------------------------------------------------------||
      //|| Bool FaceDir
      //||------------------------------------------------------------------------------------------------||*/

      export function boolFaceCentered(dir : FaceDirection) : boolean {
            return (dir === "center-center");
      }

      /*||------------------------------------------------------------------------------------------------||
      //|| Status FaceDir
      //||------------------------------------------------------------------------------------------------||*/

      export function statusFaceCentered(faceDir : FaceDirection) : JSX.Element {
            const textClass = "flex flex-col px-2 py-1 text-md font-bold border-gray-500 border-t mt-2 w-[80%] text-center";
            if (faceDir !== "center-center") return <span className={`${textClass} text-blue-400`}>Not Centered</span>;      
            return <span className={`${textClass} text-green-400`}>Centered</span>;
      }      

      /*||------------------------------------------------------------------------------------------------||
      //|| Icon
      //||------------------------------------------------------------------------------------------------||*/

      export function iconFaceCentered(faceDir: FaceDirection): JSX.Element {
            const iconClass = "w-10 h-10 rounded-full flex flex-col";      
            switch (faceDir) {
                  case "center-center":         return <Check className={`${iconClass} text-green-400`} />;
                  case "center-top":            return <CircleChevronUp className={`${iconClass} text-red-400 rotate-0`} />;
                  case "center-bottom":         return <CircleChevronUp className={`${iconClass} text-red-400 rotate-180`} />;
                  case "left-center":           return <CircleChevronUp className={`${iconClass} text-red-400 -rotate-90`} />;
                  case "left-top":              return <CircleChevronUp className={`${iconClass} text-red-400 -rotate-45`} />;
                  case "left-bottom":           return <CircleChevronUp className={`${iconClass} text-red-400 -rotate-135`} />;
                  case "right-center":          return <CircleChevronUp className={`${iconClass} text-red-400 rotate-90`} />;
                  case "right-top":             return <CircleChevronUp className={`${iconClass} text-red-400 rotate-45`} />;
                  case "right-bottom":          return <CircleChevronUp className={`${iconClass} text-red-400 rotate-135`} />;
            }
            return <CircleDashed className={`${iconClass} text-blue-400`} />;
      }      

      
