//||------------------------------------------------------------------------------------------------||
//|| Documentation Home Page
//|| src/pages/docs/Home.tsx
//||------------------------------------------------------------------------------------------------||

import React from "react";
import { Link } from "react-router-dom";
import { Lock, Eye, Cpu } from "lucide-react";
import NavDocs from "../../components/nav/NavDocs";
import FooterMain from "../../components/footer/FooterMain";

export default function Home() {
      return (
            <main className="min-h-screen bg-base-100 text-base-content">
                  {/* Navbar */}
                  <NavDocs />

                  {/* Hero Section */}
                  <section className="relative flex items-center justify-center overflow-hidden py-5">
                        <img
                              className="absolute top-0 left-0 w-full h-full object-cover"
                              src="/img/hero/docs.webp"
                              alt="Hero background"
                        />

                        <div className="relative z-10 text-white text-center px-4 mt-[70px] max-w-4xl mx-auto">
                              <img
                                    src="/img/logo-white.png"
                                    alt="ComplyAge Logo"
                                    className="mx-auto mb-6 w-28 h-28"
                              />
                              <h1 className="text-5xl font-extrabold mb-6">
                                    ComplyAge Documentation
                              </h1>
                              <p className="text-lg text-white/80">
                                    Explore our APIs and integrations to build privacy-first,
                                    age-verified, and compliant applications.
                              </p>
                        </div>
                  </section>

                  {/* Main Docs Sections */}
                  <section className="relative z-10 py-4 px-6 bg-base-200 text-center">
                        <div className="max-w-7xl mx-auto">
                              <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
                                    {/* OAuth Section */}
                                    <div className="flex flex-col items-center text-center bg-gray-900/70 p-10 rounded-2xl border border-gray-700 hover:bg-gray-900 transition-all">
                                          <Lock className="w-20 h-20 text-success mb-6" />
                                          <h3 className="text-2xl font-bold mb-4">OAuth</h3>
                                          <p className="text-white/80 mb-6">
                                                Integrate secure authentication and authorization flows
                                                for users and apps using the ComplyAge OAuth server.
                                          </p>
                                          <Link
                                                to="/oauth"
                                                className="btn btn-primary text-lg px-8 py-3"
                                          >
                                                View OAuth Docs
                                          </Link>
                                    </div>

                                    {/* Age Gate Section */}
                                    <div className="flex flex-col items-center text-center bg-gray-900/70 p-10 rounded-2xl border border-gray-700 hover:bg-gray-900 transition-all">
                                          <Eye className="w-20 h-20 text-success mb-6" />
                                          <h3 className="text-2xl font-bold mb-4">Age Gate</h3>
                                          <p className="text-white/80 mb-6">
                                                Implement seamless, privacy-respecting age verification
                                                on your website or app in minutes.
                                          </p>
                                          <Link
                                                to="/gate"
                                                className="btn btn-primary text-lg px-8 py-3"
                                          >
                                                View Age Gate Docs
                                          </Link>
                                    </div>

                                    {/* Agent Section */}
                                    <div className="flex flex-col items-center text-center bg-gray-900/70 p-10 rounded-2xl border border-gray-700 hover:bg-gray-900 transition-all">
                                          <Cpu className="w-20 h-20 text-success mb-6" />
                                          <h3 className="text-2xl font-bold mb-4">Agent</h3>
                                          <p className="text-white/80 mb-6">
                                                Learn how to use our intelligent verification agents for
                                                image, ID, and biometric analysis.
                                          </p>
                                          <Link
                                                to="/agent"
                                                className="btn btn-primary text-lg px-8 py-3"
                                          >
                                                View Agent Docs
                                          </Link>
                                    </div>
                              </div>
                        </div>
                  </section>

                  {/* Footer */}
                  <FooterMain />
            </main>
      );
}
