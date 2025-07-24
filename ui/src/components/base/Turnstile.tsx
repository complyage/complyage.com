import React, { useEffect, useRef } from "react";

declare global {
   interface Window {
      turnstile?: any;
   }
}

interface TurnstileProps {
   siteKey: string;
   onSuccess: (token: string) => void;
   options?: Record<string, any>; // optional extra props
}

const Turnstile: React.FC<TurnstileProps> = ({ siteKey, onSuccess, options = {} }) => {
   const ref = useRef<HTMLDivElement>(null);

   useEffect(() => {
      // Only render if window.turnstile exists
      if (!window.turnstile || !ref.current) return;

      const widgetId = window.turnstile.render(ref.current, {
         sitekey: siteKey,
         callback: onSuccess,
         ...options,
      });

      return () => {
         if (window.turnstile && widgetId) {
            window.turnstile.remove(widgetId);
         }
      };
   }, [siteKey, onSuccess, options]);

   return <div ref={ref} className="my-4"></div>;
};

export default Turnstile;
