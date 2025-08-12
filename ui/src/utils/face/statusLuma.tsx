/*||------------------------------------------------------------------------------------------------||
//|| Return Back the Status of the Luma
//|| statusLuma.tsx
//||------------------------------------------------------------------------------------------------||*/

      /*||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||*/

      import React, { JSX }                           from "react";

      /*||------------------------------------------------------------------------------------------------||
      //|| Icons
      //||------------------------------------------------------------------------------------------------||*/

      import {Check, Lightbulb, LightbulbOff}         from "lucide-react";

      /*||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||*/

      const MAX_LUMA = 60;
      const MIN_LUMA = 10;

      /*||------------------------------------------------------------------------------------------------||
      //|| Status
      //||------------------------------------------------------------------------------------------------||*/

      export function statusLuma(luma: number): JSX.Element {
            const textClass = "flex flex-col px-2 py-1 text-md font-bold border-gray-500 border-t mt-2 w-[80%] text-center";
            if (luma >= MAX_LUMA) return <span className={`${textClass} text-yellow-400`}>Too Bright</span>;
            if (luma <= MIN_LUMA) return <span className={`${textClass} text-red-400`}>Too Dark</span>;
            return <span className={`${textClass} text-green-400`}>Good</span>;
      }

      /*||------------------------------------------------------------------------------------------------||
      //|| Icon
      //||------------------------------------------------------------------------------------------------||*/

      export function iconLuma(luma: number): JSX.Element {
            const iconClass = "w-10 h-10 rounded-full flex flex-col";
            if (luma >= MAX_LUMA) return <Lightbulb className={`${iconClass} text-yellow-500`} />;
            if (luma <= MIN_LUMA) return <LightbulbOff className={`${iconClass} text-red-500`} />;
            return <Check className={`${iconClass} text-green-400`} />;                           
      }

      /*||------------------------------------------------------------------------------------------------||
      //|| Safe Luma
      //||------------------------------------------------------------------------------------------------||*/

      export function boolLuma(luma : number) : boolean {
            return (luma >= MIN_LUMA && luma <= MAX_LUMA);
      }
