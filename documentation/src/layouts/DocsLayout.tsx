/*||------------------------------------------------------------------------------------------------||
//|| Docs Layout
//|| src/layouts/DocsLayout.tsx
//||------------------------------------------------------------------------------------------------||*/

import React from "react";
import { NavLink } from "react-router-dom";
import NavDocs from "../components/nav/NavDocs";

/*||------------------------------------------------------------------------------------------------||
//|| Layout
//||------------------------------------------------------------------------------------------------||*/

export default function DocsLayout({ children }: { children: React.ReactNode }) {
      return (
            <main className="min-h-screen flex flex-col bg-base-100 text-base-content">
                  <NavDocs />

                  <div className="flex flex-1 pt-[60px]">
                        {/* Sidebar */}
                        <aside className="w-64 bg-base-200 border-r border-base-300 p-6 hidden md:block">
                              <nav className="flex flex-col gap-2">
                                    <h3 className="text-sm uppercase text-gray-500 mb-3">Documentation</h3>

                                    <NavLink
                                          to="/docs/oauth"
                                          className={({ isActive }) =>
                                                `px-3 py-2 rounded-md text-sm transition ${
                                                      isActive
                                                            ? "bg-base-300 text-primary font-semibold"
                                                            : "hover:bg-base-300"
                                                }`
                                          }
                                    >
                                          OAuth
                                    </NavLink>

                                    <NavLink
                                          to="/docs/gate"
                                          className={({ isActive }) =>
                                                `px-3 py-2 rounded-md text-sm transition ${
                                                      isActive
                                                            ? "bg-base-300 text-primary font-semibold"
                                                            : "hover:bg-base-300"
                                                }`
                                          }
                                    >
                                          Age Gate
                                    </NavLink>

                                    <NavLink
                                          to="/docs/agent"
                                          className={({ isActive }) =>
                                                `px-3 py-2 rounded-md text-sm transition ${
                                                      isActive
                                                            ? "bg-base-300 text-primary font-semibold"
                                                            : "hover:bg-base-300"
                                                }`
                                          }
                                    >
                                          Agent
                                    </NavLink>
                              </nav>
                        </aside>

                        {/* Main Content */}
                        <section className="flex-1 p-8 overflow-y-auto">{children}</section>
                  </div>
            </main>
      );
}
