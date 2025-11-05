//||------------------------------------------------------------------------------------------------||
//|| Turnstile Component
//|| src/components/base/Turnstile.tsx
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useRef, useState } from "react";

//||------------------------------------------------------------------------------------------------||
//|| Spinner
//||------------------------------------------------------------------------------------------------||

import SpinnerCircle from "./SpinnerCircle";

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

      const ref            = useRef<HTMLDivElement>(null);
      const widgetId       = useRef<string | null>(null);
      const [captchaToken, setCaptchaToken] = useState<string | null>(null);
      const [hideBanner, setHideBanner] = useState(false);

      //||------------------------------------------------------------------------------------------------||
      //|| Render Widget (script is already loaded globally)
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (widgetId.current || !ref.current) return;

            // wait for turnstile to be available if script loads slightly late
            const initTurnstile = () => {
                  if (!window.turnstile || !ref.current) return false;

                  widgetId.current = window.turnstile.render(ref.current, {
                        sitekey : siteKey,
                        callback: (token: string) => {
                              setCaptchaToken(token);
                              setHideBanner(true);
                              onVerify?.(token);
                        },
                        ...options,
                  });
                  return true;
            };

            if (!initTurnstile()) {
                  const timeout = setInterval(() => {
                        if (initTurnstile()) clearInterval(timeout);
                  }, 100);
                  return () => clearInterval(timeout);
            }
      }, [siteKey, onVerify, options]);

      //||------------------------------------------------------------------------------------------------||
      //|| Cleanup on Unmount
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            return () => {
                  if (window.turnstile && widgetId.current) {
                        window.turnstile.remove(widgetId.current);
                        widgetId.current = null;
                  }
            };
      }, []);

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <div
                  ref={ref}
                  className={`${className ?? "bg-black/80 rounded-md w-full p-3 my-4 min-h-[30px]"} flex items-center justify-center transition-opacity duration-500 ${
                        hideBanner ? "opacity-0" : "opacity-100"
                  }`}
            >
                  { captchaToken ? (
                        <span className="text-green-500">
                              You may proceed...
                        </span>
                  ) : (
                        <div className="flex items-center text-yellow-500 text-sm">
                              <SpinnerCircle className="mr-3" />
                              <span>Checking to see if you're a bot...</span>
                        </div>
                  ) }
            </div>
      );
}
