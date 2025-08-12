import { User as UserIcon } from "lucide-react";
import React from "react";
import {Link, useLocation} from "react-router-dom";
import RequireMember, { User } from "../dynamic/RequireMember";
import LinkQuery          from "../../components/dynamic/LinkQuery";

export default function NavMain() {
	const location = useLocation();

	return (
            <RequireMember>
                  {(user: User) => (
                  <header className="fixed top-0 left-0 right-0 z-50 bg-gray-900 shadow-md">
                        <div className="navbar px-6 max-w-7xl mx-auto">
                              {/* Logo */}
                              <div className="flex-1">
                                    <LinkQuery to="/" className="flex items-center">
                                          <img
                                                src="/img/logow.png"
                                                alt="ComplyAge Logo"
                                                className="h-12 w-auto p-2.5"
                                          />
                                    </LinkQuery
                              </div>

                              {/* Navigation Links */}
                              <div className="flex-none flex gap-4 items-center">
                                    <LinkQuery
                                          to="/about"
                                          className={`btn btn-ghost hover:text-orange-300 text-xl ${
                                                location.pathname === "/about"
                                                      ? "text-orange-500"
                                                      : ""
                                          }`}>
                                          About
                                    </LinkQuery
                                    <LinkQuery
                                          to="/vendors"
                                          className={`btn btn-ghost hover:text-orange-300 text-xl ${
                                                location.pathname === "/vendors"
                                                      ? "text-orange-500"
                                                      : ""
                                          }`}>
                                          Vendors
                                    </LinkQuery
                                    <LinkQuery
                                          to="/gilead"
                                          className={`btn btn-ghost hover:text-orange-300 text-xl ${
                                                location.pathname === "/gilead"
                                                      ? "text-orange-500"
                                                      : ""
                                          }`}>
                                          Enforcement
                                    </LinkQuery                              
                                    <LinkQuery
                                          to="/pricing"
                                          className={`btn btn-ghost hover:text-orange-300 text-xl ${
                                                location.pathname === "/pricing"
                                                      ? "text-orange-500"
                                                      : ""
                                          }`}>
                                          Pricing
                                    </LinkQuery
                              </div>                              

                              {/* Navigation Links */}
                              <div className="flex-none flex gap-4 items-center">

                                    <LinkQuery to="/members" className="btn btn-secondary text-sm ml-5">
                                          <UserIcon className="w-5 h-5" />Logout
                                    </LinkQuery                              

                              </div>
                        </div>
                  </header>
                  )}
            </RequireMember>
	);
}
