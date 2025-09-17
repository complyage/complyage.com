//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useState }         from "react";
import { useNavigate, useLocation } from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import NavMain                     from "../../components/nav/NavMain";
import FooterMain                  from "../../components/footer/FooterMain";
import Turnstile                   from "../../components/base/Turnstile";
import SpinnerCircle               from "../../components/base/SpinnerCircle";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Forgot() {

      //||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||

      const navigate                            = useNavigate();
      const [captchaToken, setCaptchaToken]     = useState("asdasd");
      const [email, setEmail]                   = useState("");
      const [statusMessage, setStatusMessage]   = useState("");
      const [captcha, setCaptcha]               = useState<string | null>(null);      

      //||------------------------------------------------------------------------------------------------||
      //|| Extract oauth query param from current URL
      //||------------------------------------------------------------------------------------------------||

      const location      = useLocation();
      const params        = new URLSearchParams(location.search);
      const oauthParam    = params.get("oauth");

      //||------------------------------------------------------------------------------------------------||
      //|| Submit Handler
      //||------------------------------------------------------------------------------------------------||

      const handleSubmit = async (e: React.FormEvent) => {
            e.preventDefault();

            const payload = {
                  area         : "FORGOT",
                  captchaToken : captchaToken,
                  email        : email,
            };

            try {
                  const res = await fetch("/auth/forgot", {
                        method  : "POST",
                        headers : { "Content-Type": "application/x-www-form-urlencoded" },
                        body    : new URLSearchParams(payload).toString(),
                  });

                  const json = await res.json();

                  if (json.success) {
                        const { token } = json.data;
                        const redirectURL = `/members/settings/?token=${encodeURIComponent(token)}` + (oauthParam ? `&oauth=${oauthParam}` : "");
                        navigate(redirectURL);
                  } else {
                        setStatusMessage(json.error || "Error: Unable to process request");
                  }

            } catch (err) {
                  console.error(err);
                  setStatusMessage("Something went wrong. Please try again.");
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <main className="min-h-screen flex flex-col">
                  <NavMain />

                  {/* Background wrapper */}
                  <div className="relative flex-1">
                        {/* Background Image */}
                        <img 
                              src="/img/hero/forgot.webp"
                              alt="Background"
                              className="absolute inset-0 w-full h-full object-cover" 
                        />

                        {/* 50/50 Split */}
                        <div className="relative z-10 flex flex-col md:flex-row min-h-[calc(100vh-60px)] mt-[60px]">

                              {/* Left Side */}
                              <div className="flex-1 flex max-w-xl items-center justify-center bg-black/60 text-primary-content p-12">
                                    <div className="max-w-md">
                                          <h1 className="text-3xl font-bold mb-6">Privacy. Freedom. Compliance.</h1>
                                          <ul className="list-disc list-inside space-y-4 text-md">
                                                <li>Stay age-compliant worldwide with one simple integration.</li>
                                                <li>Protect your users’ privacy — we never track or sell data.</li>
                                                <li>Transparent, open-source code for ultimate trust.</li>
                                                <li>Verify once, stay verified across your favorite sites.</li>
                                          </ul>
                                    </div>
                              </div>

                              {/* Right Side */}
                              <div className="flex-1 flex items-center justify-center p-12">
                                    <div className="w-full max-w-lg bg-black/80 border-gray-700 rounded-lg border p-8">
                                          <h2 className="text-3xl font-bold mb-6 text-center border-b border-base-content/20 pb-4">
                                                Forgot Your Password?
                                          </h2>

                                          {statusMessage && (
                                                <div role="alert" className="alert alert-error mb-5">
                                                      <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 shrink-0 stroke-current text-white" fill="none" viewBox="0 0 24 24">
                                                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
                                                      </svg>
                                                      <span className="text-white">{statusMessage}</span>
                                                </div>
                                          )}

                                          {/* Forgot Form */}
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
                                                      {captchaToken ? (
                                                            <button type="submit" className="btn btn-secondary w-full" disabled={!captchaToken}>
                                                                  Reset your password
                                                            </button>
                                                      ) : (
                                                            <span className="flex items-center gap-x-2 text-base-content/60">
                                                                  <SpinnerCircle /><span className="text-sm">Loading CAPTCHA...</span>
                                                            </span>
                                                      )}

                                                      <Turnstile siteKey={ import.meta.env.VITE_TURNSTILE_PUBLIC } onVerify={ (token) => setCaptcha(token) } />
                                                </div>
                                          </form>

                                          {/* Links */}
                                          <div className="flex justify-between mt-6 text-sm">
                                                <button onClick={() => navigate("/login")} className="btn btn-tertiary">
                                                      Back to Log In
                                                </button>
                                                <button onClick={() => navigate("/signup")} className="btn btn-tertiary">
                                                      Create Account
                                                </button>
                                          </div>
                                    </div>
                              </div>
                        </div>
                  </div>

                  <FooterMain />
            </main>
      );
}
