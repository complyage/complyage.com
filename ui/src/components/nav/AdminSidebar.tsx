//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React                                          from "react";
import {Link, useLocation}                            from "react-router-dom";
import { 
      Home, User as UserIcon, CheckCircle, 
      KeyRound, Settings, Share, 
      Globe, LogOut, EarthLock
}   from "lucide-react";
import RequireAdmin, { User }                        from "../dynamic/RequireMember";
import LinkQuery                                      from "../../components/dynamic/LinkQuery";

//||------------------------------------------------------------------------------------------------||
//|| Sidebar
//||------------------------------------------------------------------------------------------------||

export default function AdminSidebar() {
      //||------------------------------------------------------------------------------------------------||
      //|| Hooks
      //||------------------------------------------------------------------------------------------------||
      const location = useLocation();
      //||------------------------------------------------------------------------------------------------||
      //|| Is Active
      //||------------------------------------------------------------------------------------------------||
	const isActive = (path: string) => location.pathname === path;
      //||------------------------------------------------------------------------------------------------||
      //|| Return
      //||------------------------------------------------------------------------------------------------||
	return (
            <RequireAdmin>
                  {(user: User) => (            
                        <aside className="w-64 bg-gray-800 fixed left-0 bottom-0 top-[80px] z-100">
                              <nav className="flex flex-col gap-2 flex-1">                        
                                    <div className="text-xs p-4 text-center bg-gray-800 border-b border-gray-400"> 
                                          <span className="border-b border-yellow-400 p-4">Welcome <b>{ user.username }</b></span>
                                    </div>
                                    <LinkQuery
                                          to="/admin/"
                                          className={`btn justify-start py-5 ${
                                                isActive("/members/")
                                                      ? "btn-primary"
                                                      : "btn-ghost"
                                          }`}>
                                          <Home className="w-5 h-5 mr-2" /> Dashboard
                                    </LinkQuery>
                                    <LinkQuery
                                          to="/admin/verifications"
                                          className={`btn justify-start ${
                                                isActive("/members/identity")
                                                      ? "btn-primary"
                                                      : "btn-ghost"
                                          }`}>
                                          <CheckCircle className="w-5 h-5 mr-2" /> Verifications
                                    </LinkQuery>                                    

                                    <button className="btn btn-ghost fixed bottom-0 w-64 py-4 pb-10 text-sm">
                                          <LogOut className="w-4 h-4" /><span className="text-xs">Logout</span>
                                    </button>
                              </nav>
                        </aside>
                  )}
            </RequireAdmin>            
	);
}
