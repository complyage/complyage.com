//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect } from "react";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Overlay({ children }: { children : React.ReactNode}) {
      useEffect(() => {
            const original = document.body.style.background;
            document.body.style.background = "transparent";
            return () => { document.body.style.background = original };
      }, []);

      return (
            <main className="min-h-screen flex flex-col bg-black">
                  <div className="relative flex-1">
                        { children }
                  </div>
            </main>
      );
}