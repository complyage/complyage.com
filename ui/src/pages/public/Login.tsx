//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React                           from "react";
import { useLocation }                 from "react-router-dom";
import { useState }                    from "react";
import { useNavigate }                 from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import NavMain                         from "../../components/nav/NavMain";
import FooterMain                      from "../../components/footer/FooterMain";
import Turnstile                       from "../../components/base/Turnstile";
import SpinnerCircle                   from "../../components/base/SpinnerCircle";
import LinkQuery                       from "../../components/dynamic/LinkQuery";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Login() {

      //||------------------------------------------------------------------------------------------------||
      //|| Var
      //||------------------------------------------------------------------------------------------------||

      const navigate                          = useNavigate();
      const [captchaToken, setCaptchaToken]   = useState("asd");
      const [email, setEmail]                 = useState("");
      const [password, setPassword]           = useState("");
      const [statusMessage, setStatusMessage] = useState("");

      //||------------------------------------------------------------------------------------------------||
      //|| Extract oauth query param from current URL
      //||------------------------------------------------------------------------------------------------||

      const location      = useLocation();
      const params        = new URLSearchParams(location.search);
      const oauthParam    = params.get("oauth");

      //||------------------------------------------------------------------------------------------------||
      //|| Verify
      //||------------------------------------------------------------------------------------------------||

      const handleVerify = (token: string) => {
            setCaptchaToken(token);
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Submit
      //||------------------------------------------------------------------------------------------------||

      const handleSubmit = async (e: React.FormEvent) => {
            e.preventDefault();

            const payload = {
                  captchaToken   : captchaToken,
                  identifier     : email,
                  password       : password,
            };

            try {
                  const res  = await fetch("/auth/login", {
                        method   : "POST",
                        headers  : { "Content-Type" : "application/x-www-form-urlencoded" },
                        body     : new URLSearchParams(payload).toString(),
                  });

                  const json = await res.json();
                  console.log("RESPONSE", json);

                  if (json.success) {
                        if (oauthParam) {
                              window.location.href = `${import.meta.env.VITE_COMPLYAGE_OAUTH_URL}/v1/return?oauth=${oauthParam}`;
                        } else {
                              navigate("/members");
                        }
                  } else {
                        setStatusMessage(json.message);
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
                        <img src="https://picsum.photos/1920/1080" alt="Background"
                             className="absolute inset-0 w-full h-full object-cover" />
                        <div className="absolute inset-0 bg-black/70"></div>

                        {/* 50/50 Split */}
                        <div className="relative z-10 flex flex-col md:flex-row min-h-[calc(100vh-60px)]">

                              {/* Left Side */}
                              <div className="flex-1 flex items-center justify-center bg-primary/80 text-primary-content p-12">
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
                                    <div className="w-full max-w-lg bg-black/40 p-8">
                                          <h2 className="text-3xl font-bold mb-6 text-center border-b border-base-content/20 pb-4">Log In</h2>

                                          {statusMessage && (
                                                <div role="alert" className="alert alert-error mb-5">
                                                      <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 shrink-0 stroke-current text-white" fill="none" viewBox="0 0 24 24">
                                                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
                                                      </svg>
                                                      <span className="text-white">{statusMessage}</span>
                                                </div>
                                          )}

                                          {/* Login Form */}
                                          <form className="flex flex-col gap-4" onSubmit={handleSubmit}>
                                                <label className="label">Email Address</label>
                                                <input type="email" placeholder="Your Email" autoComplete="email"
                                                       value={email} onChange={(e) => setEmail(e.target.value)}
                                                       className="input input-bordered w-full py-5 text-xl h-12" required />

                                                <label className="label">Password</label>
                                                <input type="password" placeholder="Your Password"
                                                       value={password} onChange={(e) => setPassword(e.target.value)}
                                                       className="input input-bordered w-full text-xl" required />

                                                <div className="w-full min-h-[70px] flex flex-col items-center justify-center mt-4">
                                                      {captchaToken ? (
                                                            <button type="submit" className="btn btn-secondary w-full" disabled={!captchaToken}>
                                                                  Log In
                                                            </button>
                                                      ) : (
                                                            <span className="flex items-center gap-x-2 text-base-content/60">
                                                                  <SpinnerCircle /><span className="text-sm">Loading CAPTCHA...</span>
                                                            </span>
                                                      )}

                                                      <Turnstile siteKey={import.meta.env.VITE_TURNSTILE_SITEKEY || ""} onSuccess={handleVerify} />
                                                </div>
                                          </form>

                                          {/* Links */}
                                          <div className="flex justify-between mt-6 text-sm">
                                                <LinkQuery to="/forgot" className="btn btn-tertiary">Forgot Password?</LinkQuery>
                                                <LinkQuery to="/signup" className="btn btn-tertiary">Create Account</LinkQuery>
                                          </div>
                                    </div>
                              </div>
                        </div>
                  </div>

                  <FooterMain />
            </main>
      );
}