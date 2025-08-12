//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useEffect, useState, ReactNode}        from "react";
import {useNavigate, Link}                            from "react-router-dom";
import {Home, User, Settings, CheckCircle, LogOut}    from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import Sidebar                                        from "../components/nav/Sidebar";
import NavMain                                        from "../components/nav/NavMain";
import HealthBanner                                   from "../components/dynamic/HealthBanner";

//||------------------------------------------------------------------------------------------------||
//|| Members Layout
//||------------------------------------------------------------------------------------------------||

export default function MembersLayout({title, children}: {title : string, children: ReactNode}) {
      return (
            <>
                  <div className="min-h-screen flex bg-gray-700">
                        <NavMain />

                        {/* Sidebar fixed width */}
                        <Sidebar />

                        {/* Main Content pushed right */}
                        <main className="flex flex-col p-5 ml-54 w-full mt-5">
                              <section className="p-10 flex-1 bg-gray-700 w-full">
                                    <h1 className="text-white text-2xl mb-3 border-b border-white/20 p-3 font-bold">
                                          {title}
                                    </h1>
                                    {children}
                              </section>
                        </main>
                  </div>
                  <HealthBanner />
            </>
      );
}
