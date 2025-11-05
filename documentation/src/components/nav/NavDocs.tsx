//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||
import React, { useEffect, useState }           from "react";
import { useLocation }                          from "react-router-dom";
import LinkQuery                                from "../dynamic/LinkQuery";
import { Menu, X, Home, Eye, Lock, Cpu } from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Default Component
//||------------------------------------------------------------------------------------------------||

export default function NavDocs() {
      const { pathname }                       = useLocation();
      const [loggedIn, setLoggedIn]            = useState<boolean | null>(null);
      const [menuOpen, setMenuOpen]            = useState<boolean>(false);

      //||------------------------------------------------------------------------------------------------||
      //|| Check for session cookie on mount
      //||------------------------------------------------------------------------------------------------||
      useEffect(() => {
            fetch("/auth/quick", { credentials: "include" })
                  .then(res => setLoggedIn(res.status === 200))
                  .catch(() => setLoggedIn(false));
      }, []);

      //||------------------------------------------------------------------------------------------------||
      //|| UI URL
      //||------------------------------------------------------------------------------------------------||

      const UI_URL = import.meta.env.VITE_COMPLYAGE_UI_URL || "https://complyage.com";

      //||------------------------------------------------------------------------------------------------||
      //|| Helper: active link style
      //||------------------------------------------------------------------------------------------------||

      const linkClass = (path: string, extra = "") =>
            `flex items-center gap-3 text-lg font-medium px-4 py-1 rounded-lg transition-colors duration-200
             ${pathname === path || pathname.startsWith(path)
                  ? "bg-black text-orange-300"
                  : "hover:bg-base-200 text-base-content"} ${extra}`;

      //||------------------------------------------------------------------------------------------------||
      //|| Render
      //||------------------------------------------------------------------------------------------------||
      return (
            <header className="fixed top-0 left-0 right-0 z-50 bg-base-100 shadow-md">
                  <div className="navbar px-6 max-w-7xl mx-auto">

                        {/* Logo */}
                        <div className="flex-1">
                              <LinkQuery to={UI_URL} className="flex items-center">
                                    <img
                                          src="/img/logow.png"
                                          alt="ComplyAge Logo"
                                          className="h-12 w-auto p-2.5"
                                    />
                              </LinkQuery>
                        </div>

                        {/* Desktop Links */}
                        <div className="hidden md:flex flex-none gap-4 items-center">
                              <LinkQuery to="/"   className={linkClass("//")}><Home /></LinkQuery>
                              <LinkQuery to="/oauth"   className={linkClass("/oauth")}>OAuth</LinkQuery>
                              <LinkQuery to="/gate"    className={linkClass("/gate")}>AgeGate</LinkQuery>
                              <LinkQuery to="/agent"    className={linkClass("/agent")}>Agent</LinkQuery>
                        </div>

                        {/* Mobile Menu Button */}
                        <div className="md:hidden flex-none">
                              <button
                                    className="btn btn-ghost text-xl"
                                    onClick={() => setMenuOpen(!menuOpen)}
                                    aria-label="Toggle menu"
                              >
                                    {menuOpen ? <X size={26}/> : <Menu size={26}/>}
                              </button>
                        </div>
                  </div>

                  {/* Mobile Full-Screen Drawer (Right-Side) */}
                  <div
                        className={`fixed top-0 right-0 h-screen w-4/5 max-w-sm bg-base-100 shadow-lg transform transition-transform duration-300 ease-in-out
                              ${menuOpen ? "translate-x-0" : "translate-x-full"} z-40`}
                  >
                        <div className="flex justify-between items-center p-4 border-b border-base-300">
                              <LinkQuery to="/" className="flex items-center" onClick={() => setMenuOpen(false)}>
                                    <img src="/img/logow.png" alt="ComplyAge Logo" className="h-10 w-auto" />
                              </LinkQuery>
                              <button onClick={() => setMenuOpen(false)} aria-label="Close menu">
                                    <X size={24} />
                              </button>
                        </div>

                        <nav className="flex flex-col gap-1 mt-4 p-2">
                              <LinkQuery to="/"   className={linkClass("/")}   onClick={() => setMenuOpen(false)}><Home size={20}/>Donate</LinkQuery>
                              <LinkQuery to="/oauth"   className={linkClass("/oauth")}   onClick={() => setMenuOpen(false)}><Lock size={20}/>Donate</LinkQuery>
                              <LinkQuery to="/gate"    className={linkClass("/gate")}    onClick={() => setMenuOpen(false)}><Eye size={20}/>Gate</LinkQuery>=
                              <LinkQuery to="/agent"   className={linkClass("/agent")}   onClick={() => setMenuOpen(false)}><Cpu size={20}/>Agent</LinkQuery>
                        </nav>
                  </div>

                  {/* Overlay */}
                  {menuOpen && (
                        <div
                              className="fixed inset-0 bg-black bg-opacity-40 z-30"
                              onClick={() => setMenuOpen(false)}
                        />
                  )}
            </header>
      );
}
