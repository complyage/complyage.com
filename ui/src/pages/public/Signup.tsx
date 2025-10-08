//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React                           from "react";
import NavMain                         from "../../components/nav/NavMain";
import FooterMain                      from "../../components/footer/FooterMain";
import AuthSignup                      from "../auth/AuthSignup";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Signup() {
      return (
            <main className="min-h-screen flex flex-col">
                  <NavMain />

                  <div className="relative flex-1">
                        <img
                              src="/img/hero/signup.webp"
                              alt="Background"
                              className="absolute inset-0 w-full h-full object-cover"
                        />

                        <div className="relative z-10 flex flex-col md:flex-row min-h-[calc(100vh-60px)] mt-[60px]">
                              {/* Left Side */}
                              <div className="flex-1 flex max-w-xl items-center justify-center bg-black/60 text-primary-content p-12">
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
                              <AuthSignup />
                        </div>
                  </div>

                  <FooterMain />
            </main>
      );
}
