//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||
import React from "react";
import { ShieldCheck, EyeOff, UserX } from "lucide-react";
import LinkQuery from "../../components/dynamic/LinkQuery";

//||------------------------------------------------------------------------------------------------||
//|| Environment
//||------------------------------------------------------------------------------------------------||
const UI_URL = import.meta.env.VITE_COMPLYAGE_UI_URL || "https://complyage.com";

//||------------------------------------------------------------------------------------------------||
//|| Default Component
//||------------------------------------------------------------------------------------------------||
export default function FooterMain() {
      return (
            <footer className="bg-black text-white py-12 px-4">
                  <div className="max-w-6xl mx-auto grid md:grid-cols-4 gap-8">
                        {/* Logo & Mission */}
                        <div>
                              <img src="/img/logo-white.png" alt="ComplyAge Logo" className="w-20 mb-4" />
                              <p className="text-gray-400">
                                    Taking back your privacy & protecting your freedom since 2024.
                              </p>
                        </div>

                        {/* Quick Links */}
                        <div>
                              <h4 className="font-bold text-lg mb-4">Quick Links</h4>
                              <ul className="space-y-2">
                                    <li>
                                          <a href={`${UI_URL}/about`} target="_blank" rel="noopener noreferrer" className="hover:underline">
                                                About Us
                                          </a>
                                    </li>
                                    <li>
                                          <a href={`${UI_URL}/contact`} target="_blank" rel="noopener noreferrer" className="hover:underline">
                                                Contact Us
                                          </a>
                                    </li>
                                    <li>
                                          <a href={`${UI_URL}/terms`} target="_blank" rel="noopener noreferrer" className="hover:underline">
                                                Terms of Service
                                          </a>
                                    </li>
                                    <li>
                                          <a href={`${UI_URL}/privacy`} target="_blank" rel="noopener noreferrer" className="hover:underline">
                                                Privacy Policy
                                          </a>
                                    </li>
                              </ul>
                        </div>

                        {/* Trust Badges */}
                        <div>
                              <h4 className="font-bold text-lg mb-4">We Stand For</h4>
                              <div className="flex flex-col gap-3">
                                    <div className="flex items-center gap-2">
                                          <ShieldCheck className="w-5 h-5 text-success" />
                                          <span>Open Source</span>
                                    </div>
                                    <div className="flex items-center gap-2">
                                          <EyeOff className="w-5 h-5 text-success" />
                                          <span>No Tracking</span>
                                    </div>
                                    <div className="flex items-center gap-2">
                                          <UserX className="w-5 h-5 text-success" />
                                          <span>Stay Anonymous</span>
                                    </div>
                              </div>
                        </div>

                        {/* Privacy Message */}
                        <div>
                              <h4 className="font-bold text-lg mb-4">We Don’t Want to Know</h4>
                              <p className="text-gray-400 mb-4">
                                    No spam. No tracking. No hidden goals.
                              </p>
                              <p className="text-gray-500 text-sm leading-relaxed">
                                    This is where we’d ask for your email, but we respect your privacy too much to do that.  
                                    We don’t want your email. So, no newsletter signup here.
                              </p>
                        </div>
                  </div>

                  <div className="mt-12 border-t border-gray-800 pt-6 text-center text-gray-500 text-sm">
                        © {new Date().getFullYear()} ComplyAge. All rights reserved. Built by normal people for normal people.
                  </div>
            </footer>
      );
}
