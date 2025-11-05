/*||------------------------------------------------------------------------------------------------||
//|| Gate Overview Section
//|| src/sections/gate/GateOverview.tsx
//||------------------------------------------------------------------------------------------------||*/

import React from "react";
import PageHeader from "../../components/base/PageHeader";

export default function GateOverview() {
      return (
            <div className="prose prose-invert max-w-none space-y-10 leading-relaxed text-sm">
                  
                  <PageHeader title="1. Age Gate - Overview" />
                  {/* Header */}
                  <div className="mb-8">
                        <p className=" mt-3">
                              The <strong>ComplyAge Gate</strong> is a modular verification system that ensures
                              users meet age or consent requirements before accessing restricted content.
                              It can be deployed automatically with one line of code, or manually with full backend control.
                        </p>
                  </div>

                  {/* Section: Setup */}
                  <section className="space-y-4">
                        <h2 className="text-xl text-gray-400 font-bold mb-3">Account & Site Setup</h2>
                        <p>
                              Before you can enable the Age Gate, create a <strong>ComplyAge account</strong> and
                              register your website.
                        </p>
                        <ol className="list-decimal list-inside space-y-2 bg-black/20 p-4 text-sm pl-4">
                              <li>
                                    Sign up at{" "}
                                    <a
                                          href="https://complyage.com"
                                          target="_blank"
                                          rel="noreferrer"
                                          className="hover:underline"
                                    >
                                          complyage.com
                                    </a>.
                              </li>
                              <li>Go to <strong>Sites</strong> and click <em>Add New Site</em>.</li>
                              <li>Provide your website’s domain and identity details.</li>
                              <li>In the <strong>Age Gate</strong> section, choose how traffic should be enforced:</li>
                        </ol>

                        <h2 className="text-xl text-gray-400 font-bold mb-3">Choosing Enforcement Mode</h2>

                        <div className="bg-black/20 px-4 pb-4 rounded-lg">
                              <div className="border border-base-300 rounded-xl p-4 w-fit">                                    
                                    <select className="select select-bordered w-full sm:max-w-xs">
                                          <option value="ALLZ" selected>Force All Traffic</option>
                                          <option value="REGU">Enforce Regulated Zones Only</option>
                                          <option value="CSTM">Custom</option>
                                    </select>
                              </div>                        

                              <p className=" mt-2">
                                    <ul className="ml-5 list list-disc indent-2">
                                          <li className="text-gray-300 py-0.5"><strong className="text-white">Force All Traffic</strong> requires every visitor to verify age.</li>
                                          <li className="text-gray-300 py-0.5"><strong className="text-white"> Regulated Zones Only</strong> targets jurisdictions with legal mandates</li>
                                          <li className="text-gray-300 py-0.5"><strong className="text-white">Custom</strong> allows selective enforcement.</li>
                                    </ul>
                              </p>
                        </div>
                  </section>


                  {/* Section: Summary */}
                  <section className="space-y-3 border-t border-base-300 pt-8 text-sm">
                        <h2 className="text-2xl font-bold mb-3">TLDR</h2>
                        <ul className="list-disc list-inside space-y-2 ">
                              <li><strong>Automated:</strong> Add one script tag — instant enforcement.</li>
                              <li><strong>Manual:</strong> API-based flow with complete developer control.</li>
                              <li>Both require a registered site and selected enforcement mode.</li>
                              <li>Webhooks optionally return verification data for integration with your stack.</li>
                        </ul>
                  </section>
            </div>
      );
}
