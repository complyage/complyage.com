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
      //|| Render Widget (load script if needed, render once)
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            // if we already have a captcha token, skip entirely
            if (captchaToken) return;

            // if widget already rendered, do nothing
            if (widgetId.current) return;

            // helper to render widget once script is available
            const renderWidget = () => {
                  if (!window.turnstile || !ref.current || widgetId.current) return;
                  widgetId.current = window.turnstile.render(ref.current, {
                        sitekey : siteKey,
                        callback: (token: string) => {
                              setCaptchaToken(token);
                              setTimeout(() => {
                                    setHideBanner(true);
                              }, 2000);
                              if (onVerify) onVerify(token);
                        },
                        ...options,
                  });
            };

            // if script already present, render immediately
            if (window.turnstile) {
                  renderWidget();
            } else {
                  // otherwise load the script once
                  const script    = document.createElement("script");
                  script.src      = "https://challenges.cloudflare.com/turnstile/v0/api.js";
                  script.async    = true;
                  script.defer    = true;
                  script.onload   = renderWidget;
                  document.head.appendChild(script);
            }

            // cleanup on unmount
            return () => {
                  if (window.turnstile && widgetId.current) {
                        window.turnstile.remove(widgetId.current);
                        widgetId.current = null;
                  }
            };
      }, [siteKey, onVerify, options, captchaToken]);

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
