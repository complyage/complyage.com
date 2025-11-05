/*||------------------------------------------------------------------------------------------------||
//|| Gate Automated Section
//|| src/sections/gate/GateAutomated.tsx
//||------------------------------------------------------------------------------------------------||*/

import React from "react";
import PageHeader from "../../components/base/PageHeader";

export default function GateAutomated() {
      return (
            <div className="prose prose-invert max-w-none space-y-10 leading-relaxed text-sm">

                  {/* Section: Automated */}

                  <PageHeader title="2. Age Gate - Automated Integration" />
                  
                  <section className="space-y-4 border-t border-base-300">
                        <p className="">
                              The <strong>Automated</strong> method is quick to implement and requires no backend configuration.
                              Copy the integration script from your dashboard under <em>Websites → Age Gate</em>
                              and paste it inside your site’s <code>&lt;head&gt;</code>.
                        </p>

                        <div className="bg-base-200 rounded-xl overflow-hidden">
                              <pre className="text-sm font-mono p-4 overflow-x-auto">
{`<script src="https://gate.complyage.com//v1/complyage.js?client_id=CLID-XXX-INSERT-CLIENT-ID--"></script>`}
                              </pre>
                        </div>

                        <p className="">
                              Once added, the Age Gate will automatically appear to users according to your site’s
                              enforcement settings. Verification sessions are stored securely and reused until expiration.
                        </p>
                  </section>

            </div>
      );
}
