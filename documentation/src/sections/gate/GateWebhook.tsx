/*||------------------------------------------------------------------------------------------------||
//|| Gate WEbhook Section
//|| src/sections/gate/GateWebhook.tsx
//||------------------------------------------------------------------------------------------------||*/

import React from "react";
import PageHeader from "../../components/base/PageHeader";

export default function GateWebhook() {
      return (
            <div className="prose prose-invert max-w-none space-y-10 leading-relaxed text-sm">

                  {/* Section: Automated */}

                  <PageHeader title="4. Age Gate - Optional Webhook" />                  

                  {/* Section: Webhook */}
                  <section className="space-y-4 border-t border-base-300">
                        <p className="">
                              For advanced setups, you can configure a <strong>webhook</strong> endpoint to receive 
                              verification payloads whenever a user completes the age gate and confirms their age.
                        </p>

                        <div className="bg-base-200 rounded-xl overflow-hidden">
                              <pre className="text-sm  p-4 overflow-x-auto">
{`{
   "event": "age_verified",
   "site": "example.com",
   "verified_age": 21,
   "region": "US-CA",
   "timestamp": "2025-11-03T22:15:00Z"
}`}
                              </pre>
                        </div>

                        <p className="">
                              Webhook events are cryptographically signed using your API secret key.  
                              Use them to trigger content unlocks, update CRM entries, or log audit data.
                        </p>
                  </section>

            </div>
      );
}
