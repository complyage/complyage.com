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
import NavMembers                                     from "../components/nav/NavMembers";
import HealthBanner                                   from "../components/dynamic/HealthBanner";

//||------------------------------------------------------------------------------------------------||
//|| Members Layout
//||------------------------------------------------------------------------------------------------||

export default function MembersLayout({title, children}: {title : string, children: ReactNode}) {
	return (
            <>
                  <div className="min-h-screen flex bg-gray-700">
                        <NavMembers />
                        <div className="min-h-screen flex bg-gray-700 mt-4 w-full">
                              <Sidebar />
                              <main className="flex-1 flex flex-col p-5 ml-54">
                                    <section className="p-10 flex-1 bg-gray-700">
                                          <h1 className="text-white text-2xl mb-3 border-b border-white/20 p-3 font-bold">{ title }</h1>
                                          {children}
                                    </section>
                              </main>
                        </div>
                  </div>
                  <HealthBanner />
            </>
	);
}
