/*||------------------------------------------------------------------------------------------------||
//|| Gate Manual Section
//|| src/sections/gate/GateManual.tsx
//||------------------------------------------------------------------------------------------------||*/

import React from "react";
import PageHeader from "../../components/base/PageHeader";

export default function GateManual() {
      return (
            <div className="prose prose-invert max-w-none space-y-10 leading-relaxed text-sm">
                  <PageHeader title="Age Gate - Manual Integration" />

                  {/* Section: Manual */}
                  <section className="space-y-4 border-t border-base-300">
                        <p className="">
                              The <strong>Manual</strong> method gives you full control over when and how users verify.
                              Use the API endpoints directly in your app or backend service.
                        </p>

                        <div className="bg-base-200 rounded-xl overflow-hidden">
                              <pre className="text-sm  p-4 overflow-x-auto">
{`fetch("https://api.complyage.com/v1/gate/start", {
   method: "POST",
   headers: { "Content-Type": "application/json" },
   body: JSON.stringify({ site: "YOUR_SITE_KEY" })
});`}
                              </pre>
                        </div>

                        <p className="">
                              Perfect for single-page apps, membership systems, or custom flows.
                              You decide how and when to display the verification prompt — with or without the default overlay.
                        </p>
                  </section>

            </div>
      );
}
