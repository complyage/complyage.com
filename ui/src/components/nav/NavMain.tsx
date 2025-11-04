//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||
import React, { useEffect, useState }           from "react";
import { useLocation }                          from "react-router-dom";
import LinkQuery                                from "../../components/dynamic/LinkQuery";
import { Menu, X, Home, Info, Store, Shield, DollarSign, LogIn, UserPlus, LogOut, Users } from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Default Component
//||------------------------------------------------------------------------------------------------||
export default function NavMain() {
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
                              <LinkQuery to="/" className="flex items-center">
                                    <img
                                          src="/img/logow.png"
                                          alt="ComplyAge Logo"
                                          className="h-12 w-auto p-2.5"
                                    />
                              </LinkQuery>
                        </div>

                        {/* Desktop Links */}
                        <div className="hidden md:flex flex-none gap-4 items-center">
                              <LinkQuery to="/donate"   className={linkClass("/donate")}><DollarSign size={18}/>Donate</LinkQuery>
                              <LinkQuery to="/about"    className={linkClass("/about")}><Info size={18}/>About</LinkQuery>
                              <LinkQuery to="/vendors"  className={linkClass("/vendors")}><Store size={18}/>Vendors</LinkQuery>
                              <LinkQuery to="/gilead"   className={linkClass("/gilead")}><Shield size={18}/>Enforcement</LinkQuery>
                              <LinkQuery to="/pricing"  className={linkClass("/pricing")}><Home size={18}/>Pricing</LinkQuery>

                              <div className="min-w-[200px] flex justify-end transition-opacity duration-300 ease-in-out"
                                   style={{ opacity: loggedIn === null ? 0 : 1 }}>
                                    { loggedIn ? (
                                          <>
                                                <LinkQuery to="/members" className={linkClass("/members")}><Users size={18}/>Members</LinkQuery>
                                                <LinkQuery to="/logout"  className="btn btn-primary btn-md ml-4 flex items-center gap-2"><LogOut size={16}/>Logout</LinkQuery>
                                          </>
                                    ) : (
                                          <>
                                                <LinkQuery to="/login"  className={linkClass("/login")}><LogIn size={18}/>Login</LinkQuery>
                                                <LinkQuery to="/signup" className={`btn btn-secondary ml-4 btn-md flex items-center gap-2 ${pathname === "/signup" ? "text-white" : ""}`}><UserPlus size={16}/>Sign Up</LinkQuery>
                                          </>
                                    )}
                              </div>
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
                              <LinkQuery to="/donate"   className={linkClass("/donate")}   onClick={() => setMenuOpen(false)}><DollarSign size={20}/>Donate</LinkQuery>
                              <LinkQuery to="/about"    className={linkClass("/about")}    onClick={() => setMenuOpen(false)}><Info size={20}/>About</LinkQuery>
                              <LinkQuery to="/vendors"  className={linkClass("/vendors")}  onClick={() => setMenuOpen(false)}><Store size={20}/>Vendors</LinkQuery>
                              <LinkQuery to="/gilead"   className={linkClass("/gilead")}   onClick={() => setMenuOpen(false)}><Shield size={20}/>Enforcement</LinkQuery>
                              <LinkQuery to="/pricing"  className={linkClass("/pricing")}  onClick={() => setMenuOpen(false)}><Home size={20}/>Pricing</LinkQuery>

                              <div className="border-t border-base-300 mt-4 pt-3">
                                    { loggedIn ? (
                                          <>
                                                <LinkQuery to="/members" className={linkClass("/members")} onClick={() => setMenuOpen(false)}><Users size={20}/>Members</LinkQuery>
                                                <LinkQuery to="/logout"  className="btn btn-primary btn-sm mt-3 flex items-center gap-2" onClick={() => setMenuOpen(false)}><LogOut size={18}/>Logout</LinkQuery>
                                          </>
                                    ) : (
                                          <>
                                                <LinkQuery to="/login"  className={linkClass("/login")} onClick={() => setMenuOpen(false)}><LogIn size={20}/>Login</LinkQuery>
                                                <LinkQuery to="/signup" className={`btn btn-secondary mt-3 btn-sm flex items-center gap-2 ${pathname === "/signup" ? "text-orange-500" : ""}`} onClick={() => setMenuOpen(false)}><UserPlus size={18}/>Sign Up</LinkQuery>
                                          </>
                                    )}
                              </div>
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
