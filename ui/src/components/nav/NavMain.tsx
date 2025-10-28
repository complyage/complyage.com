//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useState }           from "react";
import { useLocation }                          from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||


import LinkQuery                                from "../../components/dynamic/LinkQuery";

//||------------------------------------------------------------------------------------------------||
//|| Default Component
//||------------------------------------------------------------------------------------------------||

export default function NavMain() {
      const { pathname }                       = useLocation();
      const [loggedIn, setLoggedIn]            = useState<boolean | null>(null);

      //||------------------------------------------------------------------------------------------------||
      //|| Check for session cookie on mount
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            fetch("/auth/quick", { credentials: "include" })
                  .then(res => setLoggedIn(res.status === 200))
                  .catch(() => setLoggedIn(false));
      }, []);

      //||------------------------------------------------------------------------------------------------||
      //|| Helper: active link style
      //||------------------------------------------------------------------------------------------------||

      const linkClass = (path: string, extra = "") =>
            `btn btn-ghost text-xl ${pathname === path || pathname.startsWith(path) ? "text-orange-500" : ""} ${extra}`;

      //||------------------------------------------------------------------------------------------------||
      //|| Render
      //||------------------------------------------------------------------------------------------------||

      return (
            <header className="fixed top-0 left-0 right-0 z-50 bg-base-100 shadow-md">
                  <div className="navbar px-6 max-w-7xl mx-auto">

                        {/* Logo */}
                        <div className="flex-1">
                              <LinkQuery to="/" className="flex items-center">
                                    <img
                                          src="/img/logow.png"
                                          alt="ComplyAge Logo"
                                          className="h-12 w-auto p-2.5"
                                    />
                              </LinkQuery>
                        </div>

                        {/* Navigation Links */}
                        <div className="flex-none flex gap-4 items-center">

                              <LinkQuery to="/donate"   className={linkClass("/donate")}>Donate</LinkQuery>
                              <LinkQuery to="/about"    className={linkClass("/about")}>About</LinkQuery>
                              <LinkQuery to="/vendors"  className={linkClass("/vendors")}>Vendors</LinkQuery>
                              <LinkQuery to="/gilead"   className={linkClass("/gilead")}>Enforcement</LinkQuery>
                              <LinkQuery to="/pricing"  className={linkClass("/pricing")}>Pricing</LinkQuery>

                              {/* Auth Links Area (reserved space) */}
                              <div className="min-w-[200px] flex justify-end transition-opacity duration-300 ease-in-out"
                                   style={{ opacity: loggedIn === null ? 0 : 1 }}>
                                    { loggedIn ? (
                                          <>
                                                <LinkQuery to="/members" className={linkClass("/members")}>Members</LinkQuery>
                                                <LinkQuery to="/logout"  className="btn btn-primary btn=md ml-4">Logout</LinkQuery>
                                          </>
                                    ) : (
                                          <>
                                                <LinkQuery to="/login"  className={linkClass("/login")}>Login</LinkQuery>
                                                <LinkQuery to="/signup" className={`btn btn-secondary ml-4 btn-md ${pathname === "/signup" ? "text-orange-500" : ""}`}>Sign Up</LinkQuery>
                                          </>
                                    )}
                              </div>
                        </div>
                  </div>
            </header>
      );
}
