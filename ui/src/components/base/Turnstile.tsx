//||------------------------------------------------------------------------------------------------||
//|| Turnstile Component
//|| src/components/base/Turnstile.tsx
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useRef } from "react";

//||------------------------------------------------------------------------------------------------||
//|| Global Window
//||------------------------------------------------------------------------------------------------||

declare global {
      interface Window {
            turnstile?: {
                  render: (el: HTMLElement, opts: Record<string, any>) => string;
                  remove: (id: string) => void;
            };
      }
}

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface TurnstileProps {
      siteKey   : string;
      options?  : Record<string, any>;
      className?: string;
      onVerify? : (token: string) => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function Turnstile({ siteKey, options = {}, className, onVerify }: TurnstileProps) {

      const ref       = useRef<HTMLDivElement>(null);
      const widgetId  = useRef<string | null>(null);

      //||------------------------------------------------------------------------------------------------||
      //|| Load Script (only once)
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (window.turnstile) return;

            const script = document.createElement("script");
            script.src   = "https://challenges.cloudflare.com/turnstile/v0/api.js";
            script.async = true;
            script.defer = true;
            document.head.appendChild(script);

            return () => {
                  if (script.parentNode) script.parentNode.removeChild(script);
            };
      }, []);

      //||------------------------------------------------------------------------------------------------||
      //|| Render Widget (only once)
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (!window.turnstile || !ref.current || widgetId.current) return;

            widgetId.current = window.turnstile.render(ref.current, {
                  sitekey: siteKey,
                  callback: (token: string) => {
                        if (onVerify) onVerify(token);   // 🔑 call the hook
                  },
                  ...options,
            });

            return () => {
                  if (window.turnstile && widgetId.current) {
                        window.turnstile.remove(widgetId.current);
                        widgetId.current = null;
                  }
            };
      }, [siteKey, onVerify]); // keep deps minimal

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <div ref={ref} className={className ?? "my-4"}>
                  {!window.turnstile && (
                        <div className="p-4 bg-yellow-100 border border-yellow-300 text-yellow-800 rounded">
                              Loading verification...
                        </div>
                  )}
            </div>
      );
}
