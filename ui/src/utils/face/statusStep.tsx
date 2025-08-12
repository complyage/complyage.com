/*||------------------------------------------------------------------------------------------------||
//|| Return Back the Status of the Step
//|| statusStep.tsx
//||------------------------------------------------------------------------------------------------||*/

      /*||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||*/

      import React, { JSX }                           from "react";

      /*||------------------------------------------------------------------------------------------------||
      //|| Utils
      //||------------------------------------------------------------------------------------------------||*/

      import { boolFace }                             from "./statusFace";
      import { boolLuma }                             from "./statusLuma";
      import { boolFaceCentered }                     from "./statusFaceDir";

      /*||------------------------------------------------------------------------------------------------||
      //|| Interface
      //||------------------------------------------------------------------------------------------------||*/

      import { FaceDirection }                        from "../../interfaces/types.face.direction";

      /*||------------------------------------------------------------------------------------------------||
      //|| Status
      //||------------------------------------------------------------------------------------------------||*/

      export function statusStep(step: number, luma? : number, face? : number, direction? : FaceDirection, takePhoto? : () => void ): JSX.Element {
            const textClass = "flex flex-col px-2 py-1 text-xl font-bold";            
            switch(step) { 
                  case 0 : return <span className={`${textClass} text-gray-400`}>Waiting for Camera</span>; break;
                  case 1 : return <span className={`${textClass} text-yellow-400`}>Position Your Face in Center</span>; break;
                  case 2 : return <span className={`${textClass} text-green-400`}>Look at the Spot</span>; break;
                  case 3 : 
                  return (<button onClick={ takePhoto } className="rounded-md bg-green-600 text-2xl px-10 py-3 text-white hover:bg-green-500">Take a Selfie</button>);
                  return (boolLuma(luma || 0) && boolFace(face || 0) && boolFaceCentered(direction || "unknown-unknown") ) ? (<button onClick={ takePhoto } className="rounded-md bg-green-600 text-2xl px-10 py-3 text-white hover:bg-green-500">Take a Selfie</button>) : ( <span className={`${textClass} text-green-400`}>Take a Selfie</span> ); break;
                  case 4 : return <span className={`${textClass} text-green-400`}>Submit</span>; break;
                  case 5 : return <span className={`${textClass} text-green-400`}>Complete</span>; break;
            }
      }
