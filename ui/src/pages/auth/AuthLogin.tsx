//||------------------------------------------------------------------------------------------------||
//|| React
//||------------------------------------------------------------------------------------------------||

import React, { useState, useEffect }                from "react";

//||------------------------------------------------------------------------------------------------||
//|| Hooks
//||------------------------------------------------------------------------------------------------||

import { useOverlayNavigate }                        from "../../hooks/useOverlay";

//||------------------------------------------------------------------------------------------------||
//|| API
//||------------------------------------------------------------------------------------------------||

import apiURL                                                 from "../../utils/apiURL";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import LinkQuery                          from "../../components/dynamic/LinkQuery";
import SpinnerCircle                      from "../../components/base/SpinnerCircle";
import Turnstile                          from "../../components/base/Turnstile";

//||------------------------------------------------------------------------------------------------||
//|| JSX
//||------------------------------------------------------------------------------------------------||

export default function AuthLogin() {

      //||------------------------------------------------------------------------------------------------||
      //|| Var
      //||------------------------------------------------------------------------------------------------||

      const navigate                            = useOverlayNavigate();
      const [captchaToken, setCaptchaToken]     = useState("asd");
      const [email, setEmail]                   = useState("");
      const [password, setPassword]             = useState("");
      const [statusMessage, setStatusMessage]   = useState("");
      const [captcha, setCaptcha]               = useState<string | null>(null);      

      //||------------------------------------------------------------------------------------------------||
      //|| Hack Weird Google Autofill Styles
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            const fixAutofill = (el: HTMLInputElement) => {
                  el.style.fontSize = "1.25rem";       // Tailwind text-xl
                  el.style.lineHeight = "1.75rem";     // match leading
                  el.style.webkitTextFillColor = "#fff"; 
                  (el.style as any).webkitBoxShadow = "0 0 0px 1000px #1a1a1a inset";
            };

            const inputs = document.querySelectorAll<HTMLInputElement>("input:-webkit-autofill");
            inputs.forEach(fixAutofill);

            const observer = new MutationObserver(() => {
                  document.querySelectorAll<HTMLInputElement>("input:-webkit-autofill").forEach(fixAutofill);
            });

            observer.observe(document.body, { attributes: true, subtree: true });

            return () => observer.disconnect();
      }, []);

      //||------------------------------------------------------------------------------------------------||
      //|| Submit
      //||------------------------------------------------------------------------------------------------||

      const handleSubmit = async (e: React.FormEvent) => {
            e.preventDefault();

            const payload = {
                  captchaToken : captchaToken,
                  identifier   : email,
                  password     : password,
            };

            try {
                  const serv = import.meta.env.VITE_API_SERVER;
                  const res = await fetch(apiURL("/auth/login"), {
                        method   : "POST",
                        credentials : "include",
                        headers  : { "Content-Type": "application/x-www-form-urlencoded" },
                        body     : new URLSearchParams(payload).toString(),
                  });

                  const json = await res.json();
                  if (json.success) {
                        navigate("/members?refresh");
                  } else {
                        setStatusMessage(json.message);
                  }
            } catch (err) {
                  setStatusMessage("Something went wrong. Please try again.");
            }
      };

      //||------------------------------------------------------------------------------------------------||  
      //|| JSX  
      //||------------------------------------------------------------------------------------------------||  
      
      return (
            <div className="flex-1 flex items-center justify-center p-12">
                  <div className="w-full max-w-lg p-8">
                        <h2 className="text-3xl font-bold mb-6 text-center border-b border-base-content/20 pb-4">
                              Log In
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

                        {/* Login Form */}
                        <form className="flex flex-col gap-4" onSubmit={handleSubmit}>
                              <label className="label">Email Address</label>
                              <input
                                    type="email"
                                    placeholder="Your Email"
                                    autoComplete="email"
                                    value={email}
                                    onChange={(e) => setEmail(e.target.value)}
                                    className="input input-bordered w-full py-2 text-xl h-auto"
                                    required
                              />

                              <label className="label">Password</label>
                              <input
                                    type="password"
                                    placeholder="Your Password"
                                    value={password}
                                    onChange={(e) => setPassword(e.target.value)}
                                    className="input input-bordered w-full py-2 text-xl h-auto"
                                    required
                              />

                              <div className="w-full min-h-[70px] flex flex-col items-center justify-center mt-4">
                                     <button
                                          type="submit"
                                          className="btn btn-secondary w-full"
                                          disabled={!captcha}
                                    >
                                          {captcha ? "Login" : "Verifying Captcha"}
                                    </button>

                                    <Turnstile
                                          siteKey={import.meta.env.VITE_TURNSTILE_PUBLIC}
                                          onVerify={(token) => setCaptcha(token)}
                                    />
                              </div>
                        </form>

                        {/* Links */}
                        <div className="flex justify-between mt-6 text-sm">
                              <LinkQuery to="/forgot" className="btn btn-tertiary">
                                    Forgot Password?
                              </LinkQuery>
                              <LinkQuery to="/signup" className="btn btn-tertiary">
                                    Create Account
                              </LinkQuery>
                        </div>
                  </div>
            </div>
      );
}
