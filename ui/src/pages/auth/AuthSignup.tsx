//||------------------------------------------------------------------------------------------------||
//|| React
//||------------------------------------------------------------------------------------------------||

import React, { useState }          from "react";
import { useNavigate, useLocation } from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| API
//||------------------------------------------------------------------------------------------------||

import apiURL                                                 from "../../utils/apiURL";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import LinkQuery                          from "../../components/dynamic/LinkQuery";
import Turnstile                    from "../../components/base/Turnstile";
import SpinnerCircle                from "../../components/base/SpinnerCircle";

//||------------------------------------------------------------------------------------------------||
//|| Auth Signup
//||------------------------------------------------------------------------------------------------||

export default function AuthSignup() {

      //||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||
      const navigate                            = useNavigate();
      const location                            = useLocation();
      const prefix                              = location.pathname.startsWith("/overlay") ? "/overlay" : "";
      const params                              = new URLSearchParams(location.search);
      const oauthParam                          = params.get("oauth");
      const [email, setEmail]                   = useState(() => `user${Math.floor(Math.random() * 10000)}@example.com`);
      const [captchaToken, setCaptchaToken]     = useState<string | null>(null);
      const [statusMessage, setStatusMessage]   = useState("");

      //||------------------------------------------------------------------------------------------------||
      //|| Submit
      //||------------------------------------------------------------------------------------------------||

      const handleSubmit = async (e: React.FormEvent) => {
            e.preventDefault();

            const payload = {
                  area         : "SIGNUP",
                  email        : email,
                  captchaToken : captchaToken || "",
            };

            try {
                  const res = await fetch(apiURL("/auth/signup"), {
                        method  : "POST",
                        credentials : "include",
                        headers : { "Content-Type": "application/x-www-form-urlencoded" },
                        body    : new URLSearchParams(payload).toString(),
                  });

                  const json = await res.json();
                  if (json.success) {
                        const { token } = json.data;
                        navigate(`${prefix}/verify?token=${encodeURIComponent(token)}`);
                        return;
                  }

                  setStatusMessage(`❌ Error: ${json.error || "Unknown error"}`);
            } catch (err) {
                  console.error(err);
                  setStatusMessage("❌ Error: Something went wrong. Please try again.");
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||
      return (
            <div className="flex-1 flex items-center justify-center p-12">
                  <div className="w-full max-w-lg p-8">
                        <h2 className="text-3xl font-bold mb-6 text-center border-b border-base-content/20 pb-4">
                              Sign Up
                        </h2>

                        {statusMessage && (
                              <div role="alert" className="alert alert-error mb-5">
                                    <svg xmlns="http://www.w3.org/2000/svg"
                                         className="h-6 w-6 shrink-0 stroke-current text-white"
                                         fill="none" viewBox="0 0 24 24">
                                          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2"
                                                d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
                                    </svg>
                                    <span className="text-white">{statusMessage}</span>
                              </div>
                        )}

                        <form className="flex flex-col gap-4" onSubmit={handleSubmit}>
                              <label className="label">Email Address</label>
                              <input
                                    type="email"
                                    placeholder="Your Email"
                                    autoComplete="email"
                                    value={email}
                                    onChange={(e) => setEmail(e.target.value)}
                                    className="input input-bordered w-full py-5 text-xl h-12"
                                    required
                              />

                              <div className="w-full min-h-[70px] flex flex-col items-center justify-center mt-4">
                                    <button
                                          type="submit"
                                          className="btn btn-secondary w-full"
                                          disabled={!captchaToken}
                                    >
                                          {captchaToken ? "Create an Account" : "Verifying Captcha"}
                                    </button>

                                    <Turnstile
                                          siteKey={import.meta.env.VITE_TURNSTILE_PUBLIC || ""}
                                          onVerify={(token) => setCaptchaToken(token)}
                                    />
                              </div>
                        </form>

                        <div className="flex justify-between mt-6 text-sm">
                              <LinkQuery to="/login" className="btn btn-black">
                                    Already have an account?
                              </LinkQuery>
                              <LinkQuery to="/forgot" className="btn btn-tertiary">
                                    Forgot your password?
                              </LinkQuery>
                        </div>
                  </div>
            </div>
      );
}
