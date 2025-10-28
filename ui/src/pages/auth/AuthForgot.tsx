//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useState }         from "react";
import { useNavigate, useLocation } from "react-router-dom";


//||------------------------------------------------------------------------------------------------||
//|| API
//||------------------------------------------------------------------------------------------------||

import apiURL                                                 from "../../utils/apiURL";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import Turnstile                   from "../../components/base/Turnstile";
import SpinnerCircle               from "../../components/base/SpinnerCircle";
import LinkQuery                   from "../../components/dynamic/LinkQuery";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function AuthForgot() {

      const navigate                            = useNavigate();
      const [captchaToken, setCaptchaToken]     = useState("asdasd");
      const [email, setEmail]                   = useState("");
      const [statusMessage, setStatusMessage]   = useState("");
      const [captcha, setCaptcha]               = useState<string | null>(null);      

      // Extract oauth query param
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
                  const res = await fetch(apiURL("/auth/forgot"), {
                        method  : "POST",
                        credentials : "include",
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
            <div className="flex-1 flex items-center justify-center p-5">
                  <div className="w-full p-8">
                        <h2 className="text-3xl font-bold mb-6 text-center  pb-4">
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
                              <button
                                    type="submit"
                                    className="btn btn-secondary"
                                    disabled={!captcha}
                              >
                                    {captcha ? "Reset My Password" : "Verifying Captcha"}
                              </button>
                              <Turnstile siteKey={ import.meta.env.VITE_TURNSTILE_PUBLIC } onVerify={ (token) => setCaptcha(token) } />
                        </form>


                        <div className="flex justify-between mt-6 text-sm">
                              <LinkQuery to="/login" className="btn btn-tertiary">
                                    Return to Login
                              </LinkQuery>
                              <LinkQuery to="/signup" className="btn btn-tertiary">
                                    Create Account
                              </LinkQuery>
                        </div>
                  </div>
            </div>
      );
}
