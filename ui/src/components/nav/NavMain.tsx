import React, { useEffect, useState }           from "react";
import { Link, useLocation }                    from "react-router-dom";
import LinkQuery          from "../../components/dynamic/LinkQuery";

export default function NavMain() {
	const location = useLocation();
      const [loggedIn, setLoggedIn] = useState(false);

      // Check for session cookie on mount

      useEffect(() => {
            const cookies = document.cookie.split(";").map(c => c.trim());
            const hasSessionUI = cookies.some(c => c.startsWith("session_ui="));
            setLoggedIn(hasSessionUI);
      }, []);      

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
                              <LinkQuery to="/about" className={`btn btn-ghost text-xl ${location.pathname === "/about" ? "text-orange-500" : ""}`}>About</LinkQuery>
                              <LinkQuery to="/vendors" className={`btn btn-ghost text-xl ${location.pathname === "/vendors" ? "text-orange-500" : ""}`}>Vendors</LinkQuery>
                              <LinkQuery to="/gilead" className={`btn btn-ghost text-xl ${location.pathname === "/gilead" ? "text-orange-500" : ""}`}>Enforcement</LinkQuery>
                              <LinkQuery to="/pricing" className={`btn btn-ghost text-xl ${location.pathname === "/pricing" ? "text-orange-500" : ""}`}>Pricing</LinkQuery>

                              {/* Conditional Links */}
                              {loggedIn ? (
                                    <>
                                          <LinkQuery to="/members" className={`btn btn-ghost text-xl ${location.pathname === "/members" ? "text-orange-500" : ""}`}>Members</LinkQuery>
                                          <LinkQuery to="/logout" className="btn btn-secondary">Logout</LinkQuery>
                                    </>
                              ) : (
                                    <>
                                          <LinkQuery to="/login" className={`btn btn-ghost text-xl ${location.pathname === "/login" ? "text-orange-500" : ""}`}>Login</LinkQuery>
                                          <LinkQuery to="/signup" className="btn btn-primary">Sign Up</LinkQuery>
                                    </>
                              )}
                        </div>
                  </div>
            </header>
      );
}
