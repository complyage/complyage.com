//||------------------------------------------------------------------------------------------------||
//|| SiteTestingSection
//|| Developer tools for testing Age Gate functionality, webhooks, and OAuth flows
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useState }          from "react";
import { Globe, RefreshCcw }        from "lucide-react";
import { ModelSite }                from "../../../interfaces/models/model.sites";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface SiteTestingSectionProps {
      data          : ModelSite;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function SiteTestingSection({ data }: SiteTestingSectionProps) {

      //||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||

      const [webhookStatus, setWebhookStatus]       = useState<string>("Idle");
      const [webhookResponse, setWebhookResponse]   = useState<string>("");

      //||------------------------------------------------------------------------------------------------||
      //|| Handlers
      //||------------------------------------------------------------------------------------------------||

      const handlePreviewGate = () => {
            window.open(`${import.meta.env.VITE_COMPLYAGE_OAUTH_URL}/v1/gate/preview/${data.clientId}`, "_blank", "noopener,noreferrer");
      };

      const handleTestWebhook = async () => {
            try {
                  setWebhookStatus("Sending...");
                  const res  = await fetch(`/v1/test/webhook/${data.clientId}`, { method: "POST", credentials: "include" });
                  const json = await res.json();
                  setWebhookStatus(res.ok ? "Success" : "Failed");
                  setWebhookResponse(JSON.stringify(json, null, 2));
            } catch (err) {
                  setWebhookStatus("Error");
                  setWebhookResponse(String(err));
                  console.error("Webhook test failed:", err);
            }
      };

      const oauthTestURL = `${import.meta.env.VITE_COMPLYAGE_OAUTH_URL}/test?client_id=${data.clientId}&access_key=${data.private}`;

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <div className="w-full bg-base-100 shadow-lg rounded-lg p-8 mb-5">
                  <h2 className="text-2xl font-bold mb-6">Testing & Developer Tools</h2>
                  <p className="text-sm text-gray-400 mb-6 leading-6">
                        Use these tools to preview and test your Age Gate integration, webhook delivery, 
                        and OAuth access flow before going live. These actions are only available for test-mode sites.
                  </p>

                  {/* Preview Age Gate */}
                  <div className="border-t border-gray-700 pt-4 mt-2">
                        <h3 className="text-lg font-semibold text-white mb-1 flex items-center gap-2">
                              <Globe size={18} /> Preview Age Gate
                        </h3>
                        <p className="text-sm text-gray-400 mb-3">
                              Opens a live preview of your Age Gate as it appears to end users.
                        </p>
                        <button
                              onClick={handlePreviewGate}
                              className="bg-blue-600 hover:bg-blue-700 text-white px-5 py-2 rounded-md transition"
                        >
                              Launch Preview
                        </button>
                  </div>

                  {/* Test Webhook */}
                  <div className="border-t border-gray-700 pt-6 mt-6">
                        <h3 className="text-lg font-semibold text-white mb-1 flex items-center gap-2">
                              <RefreshCcw size={18} /> Test Age Gate Webhook
                        </h3>
                        <p className="text-sm text-gray-400 mb-3">
                              Sends a test verification payload to your configured webhook endpoint.
                        </p>

                        <div className="flex items-center gap-3 mb-3">
                              <button
                                    onClick={handleTestWebhook}
                                    className="bg-emerald-600 hover:bg-emerald-700 text-white px-5 py-2 rounded-md transition"
                              >
                                    Send Test
                              </button>
                              <span
                                    className={`text-sm font-semibold px-3 py-1 rounded ${
                                          webhookStatus === "Success"
                                                ? "bg-green-600 text-white"
                                                : webhookStatus === "Failed" || webhookStatus === "Error"
                                                ? "bg-red-600 text-white"
                                                : webhookStatus === "Sending..."
                                                ? "bg-yellow-500 text-black"
                                                : "bg-gray-700 text-gray-300"
                                    }`}
                              >
                                    {webhookStatus}
                              </span>
                        </div>

                        <textarea
                              className="textarea textarea-bordered w-full text-xs font-mono bg-black/30 text-gray-200"
                              rows={8}
                              readOnly
                              value={webhookResponse || "No response yet."}
                        ></textarea>
                  </div>

                  {/* OAuth Test URL */}
                  <div className="border-t border-gray-700 pt-6 mt-6">
                        <h3 className="text-lg font-semibold text-white mb-1">OAuth Test URL</h3>
                        <p className="text-sm text-gray-400 mb-3">
                              Use this link to initiate a test OAuth authorization request. Copy it or open directly in a new tab.
                        </p>
                        <div className="flex flex-col sm:flex-row gap-3">
                              <input
                                    type="text"
                                    readOnly
                                    className="input input-bordered w-full bg-black/30 text-gray-300 text-sm font-mono"
                                    value={oauthTestURL}
                              />
                              <button
                                    onClick={() => navigator.clipboard.writeText(oauthTestURL)}
                                    className="bg-gray-700 hover:bg-gray-600 text-white px-5 py-2 rounded-md transition"
                              >
                                    Copy
                              </button>
                        </div>
                  </div>
            </div>
      );
}
