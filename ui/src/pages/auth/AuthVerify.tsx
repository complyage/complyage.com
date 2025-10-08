//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useState }                  from "react";
import { useSearchParams, useNavigate, useLocation } from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function AuthVerify() {

      const [searchParams]                    = useSearchParams();
      const navigate                          = useNavigate();
      const token                             = searchParams.get("token") || "";
      const [code, setCode]                   = useState("");
      const [statusMessage, setStatusMessage] = useState("");
      const [loading, setLoading]             = useState(false);

      const location      = useLocation();
      const prefix        = location.pathname.startsWith("/overlay") ? "/overlay" : "";
      const params        = new URLSearchParams(location.search);
      const oauthParam    = params.get("oauth");

      const handleSubmit = async (e: React.FormEvent) => {
            e.preventDefault();

            const payload = {
                  token : token,
                  code  : code.trim(),
            };

            try {
                  setLoading(true);
                  const serv = "";
                  const res = await fetch(serv+"/auth/twofactor", {
                        method  : "POST",
                        headers : { "Content-Type": "application/x-www-form-urlencoded" },
                        body    : new URLSearchParams(payload).toString(),
                  });

                  const json = await res.json();

                  if (json.success) {
                        return navigate(`${prefix}/complete`);
                  }

                  setStatusMessage(`❌ ${json.error || "Verification failed."}`);
            } catch (err) {
                  console.error(err);
                  setStatusMessage("❌ Something went wrong. Please try again.");
            } finally {
                  setLoading(false);
            }
      };

      return (
            <div className="flex-1 flex items-center justify-center p-12">
                  <div className="w-full max-w-md bg-black/80 border-gray-700 rounded-lg border p-8 text-center">
                        <h1 className="text-3xl font-bold mb-6 text-gray-300">
                              Two-Factor Verification
                        </h1>
                        <p className="mb-4 text-base-content/70">
                              Enter the 6-digit code we sent to your email.
                        </p>

                        <form onSubmit={handleSubmit} className="flex flex-col gap-4">
                              <input
                                    type="text"
                                    placeholder="Code"
                                    value={code}
                                    onChange={(e) => setCode(e.target.value)}
                                    className="border border-white w-full text-center tracking-widest text-4xl p-5"
                                    required
                              />

                              <button
                                    type="submit"
                                    disabled={code.trim().length === 0 || loading}
                                    className="btn btn-secondary h-auto w-full text-2xl py-3"
                              >
                                    {loading ? "Verifying..." : "Verify"}
                              </button>
                        </form>

                        {statusMessage && (
                              <p className="mt-4 text-sm text-red-400">{statusMessage}</p>
                        )}
                  </div>
            </div>
      );
}
