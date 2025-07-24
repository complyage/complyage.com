//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import NavMain from "../../components/nav/NavMain";
import FooterMain from "../../components/footer/FooterMain";
import Turnstile from "../../components/base/Turnstile";
import SpinnerCircle from "../../components/base/SpinnerCircle";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Signup() {

      //||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||

      const navigate                = useNavigate();
      const [captchaToken, setCaptchaToken] = useState("asdasd");
      const [tab, setTab]          = useState<"USER" | "VNDR">("USER");
      const [email, setEmail]      = useState("");
      const [statusMessage, setStatusMessage] = useState("");

      //||------------------------------------------------------------------------------------------------||
      //|| CAPTCHA Handler
      //||------------------------------------------------------------------------------------------------||

      const handleVerify = (token: string) => {
            setCaptchaToken(token);
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Submit Handler
      //||------------------------------------------------------------------------------------------------||

      const handleSubmit = async (e: React.FormEvent) => {
            e.preventDefault();

            const payload = {
                  captchaToken: captchaToken,
                  email: email,
                  type: tab,
            };

            try {
                  const res = await fetch("/auth/signup", {
                        method  : "POST",
                        headers : { "Content-Type": "application/x-www-form-urlencoded" },
                        body    : new URLSearchParams(payload).toString(),
                  });

                  const json = await res.json();
                  console.log("RESPONSE", json);

                  if (json.success) {
                        const { token } = json.data;
                        navigate(`/verify?token=${encodeURIComponent(token)}`);
                  } else {
                        setStatusMessage(`❌ Error: ${json.error}`);
                  }

            } catch (err) {
                  console.error(err);
                  setStatusMessage("❌ Error: Something went wrong. Please try again.");
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <main className="min-h-screen flex flex-col">
                  <NavMain />

                  <div className="relative flex-1">
                        <img
                              src="https://picsum.photos/1920/1080"
                              alt="Background"
                              className="absolute inset-0 w-full h-full object-cover"
                        />
                        <div className="absolute inset-0 bg-black/70"></div>

                        <div className="relative z-10 flex flex-col md:flex-row min-h-[calc(100vh-60px)]">
                              {/* Left Side */}
                              <div className="flex-1 flex items-center justify-center bg-primary/80 text-primary-content p-12">
                                    <div className="max-w-md">
                                          <h1 className="text-3xl font-bold mb-6">
                                                Privacy. Freedom. Compliance.
                                          </h1>
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
                                    <div className="w-full max-w-lg bg-black/40 p-8">
                                          <h2 className="text-3xl font-bold mb-6 text-center border-b border-base-content/20 pb-4">
                                                Sign Up
                                          </h2>

                                          {statusMessage && (
                                                <div role="alert" className="alert alert-error">
                                                      <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 shrink-0 stroke-current text-white" fill="none" viewBox="0 0 24 24">
                                                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
                                                      </svg>
                                                      <span className="text-white">{statusMessage}</span>
                                                </div>
                                          )}

                                          <div className="flex justify-center gap-4 mb-4">
                                                <button
                                                      className={`text-md btn btn-sm ${tab === "USER" ? "btn-primary" : "btn-ghost"}`}
                                                      onClick={() => setTab("USER")}>
                                                      User
                                                </button>
                                                <button
                                                      className={`text-md btn btn-sm ${tab === "VNDR" ? "btn-primary" : "btn-ghost"}`}
                                                      onClick={() => setTab("VNDR")}>
                                                      Vendor / Website Owner
                                                </button>
                                          </div>

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
                                                            <button
                                                                  type="submit"
                                                                  className="btn btn-primary w-full"
                                                                  disabled={!captchaToken}>
                                                                  Sign Up as {tab === "USER" ? "User" : "Vendor"}
                                                            </button>
                                                      ) : (
                                                            <span className="flex items-center gap-x-2 text-base-content/60">
                                                                  <SpinnerCircle />
                                                                  <span className="text-sm">Loading CAPTCHA...</span>
                                                            </span>
                                                      )}

                                                      <Turnstile
                                                            siteKey={import.meta.env.VITE_TURNSTILE_SITEKEY || ""}
                                                            onSuccess={handleVerify}
                                                      />
                                                </div>
                                          </form>

                                          <div className="flex justify-between mt-6 text-sm">
                                                <button onClick={() => navigate("/login")} className="btn btn-black">
                                                      Already have an account?
                                                </button>
                                                <button onClick={() => setEmail("")} className="btn btn-secondary">
                                                      Reset Form
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
