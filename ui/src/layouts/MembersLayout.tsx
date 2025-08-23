//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useEffect, useState, ReactNode}        from "react";
import {useNavigate, Link}                            from "react-router-dom";
import type { LucideIcon }                            from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import Sidebar                                        from "../components/nav/Sidebar";
import NavMain                                        from "../components/nav/NavMain";
import HealthBanner                                   from "../components/dynamic/HealthBanner";

//||------------------------------------------------------------------------------------------------||
//|| Members Layout
//||------------------------------------------------------------------------------------------------||

export default function MembersLayout({title, icon : Icon, children}: {title? : string, icon? : LucideIcon, children: ReactNode}) {
      return (
            <>
                  <div className="min-h-screen flex bg-gray-700">
                        <NavMain />

                        {/* Sidebar fixed width */}
                        <Sidebar />

                        {/* Main Content pushed right */}
                        <main className="flex flex-col p-5 ml-54 w-full mt-5">
                              <section className="p-10 flex-1 bg-gray-700 w-full">
                                    {(title || Icon) && (<h1 className="text-white text-2xl mb-3 border-b border-white/20 p-3 font-bold flex items-center">
                                          {Icon && <Icon className="inline-block mr-2 w-6 h-6" />}
                                          {title}
                                    </h1>)}
                                    {children}
                              </section>
                        </main>
                  </div>
                  <HealthBanner />
            </>
      );
}
