//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useState, useEffect } from "react";

//||------------------------------------------------------------------------------------------------||
//|| HealthBanner Component
//||------------------------------------------------------------------------------------------------||

export default function HealthBanner() {
      const [online, setOnline] = useState(true);

      useEffect(() => {
            let mounted = true;

            async function checkHealth() {
                  try {
                        const res = await fetch(import.meta.env.VITE_COMPLYAGE_API_URL + "/health");
                        if (!mounted) return;
                        const json = await res.json();
                        setOnline(res.ok);
                  } catch(e) {
                        if (!mounted) return;
                        setOnline(false);
                  }
            }

            // initial check and poll every 10s
            checkHealth();
            const interval = setInterval(checkHealth, 10000);

            return () => {
                  mounted = false;
                  clearInterval(interval);
            };
      }, []);

      if (online) return null;

      return (
            <div className="fixed bottom-0 left-0 right-0 bg-red-600 text-white text-center p-2">
                  <p>⚠️ Connection lost. Attempting to reconnect…</p>
            </div>
      );
}
