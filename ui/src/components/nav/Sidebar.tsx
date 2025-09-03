//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React                                          from "react";
import {Link, useLocation}                            from "react-router-dom";
import { 
      Home, User as UserIcon, 
      KeyRound, Settings, Share, 
      Globe, LogOut, EarthLock
}   from "lucide-react";
import RequireMember, { User }                        from "../dynamic/RequireMember";
import LinkQuery                                      from "../../components/dynamic/LinkQuery";

//||------------------------------------------------------------------------------------------------||
//|| Sidebar
//||------------------------------------------------------------------------------------------------||

export default function Sidebar() {
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
            <RequireMember>
                  {(user: User) => (            
                        <aside className="w-64 bg-gray-800 fixed left-0 bottom-0 top-[80px] z-100">
                              <nav className="flex flex-col gap-2 flex-1">                        
                                    <div className="text-xs p-4 text-center bg-gray-800 border-b border-gray-400"> 
                                          <span className="border-b border-yellow-400 p-4">Welcome <b>{ user.username }</b></span>
                                    </div>
                                    <LinkQuery
                                          to="/members/"
                                          className={`btn justify-start py-5 ${
                                                isActive("/members/")
                                                      ? "btn-primary"
                                                      : "btn-ghost"
                                          }`}>
                                          <Home className="w-5 h-5 mr-2" /> Dashboard
                                    </LinkQuery>
                                    <LinkQuery
                                          to="/members/identity"
                                          className={`btn justify-start ${
                                                isActive("/members/identity")
                                                      ? "btn-primary"
                                                      : "btn-ghost"
                                          }`}>
                                          <KeyRound className="w-5 h-5 mr-2" /> My Identity
                                    </LinkQuery>                                    
                                    <LinkQuery
                                          to="/members/sites"
                                          className={`btn justify-start ${
                                                isActive("/members/sites")
                                                      ? "btn-primary"
                                                      : "btn-ghost"
                                          }`}>
                                          <Globe className="w-5 h-5 mr-2" /> Sites
                                    </LinkQuery>
                                    <LinkQuery
                                          to="/members/shared"
                                          className={`btn justify-start ${
                                                isActive("/members/shared")
                                                      ? "btn-primary"
                                                      : "btn-ghost"
                                          }`}>
                                          <Share className="w-5 h-5 mr-2" /> Shared Data
                                    </LinkQuery>
                                    <LinkQuery
                                          to="/members/settings"
                                          className={`btn justify-start ${
                                                isActive("/members/settings")
                                                      ? "btn-primary"
                                                      : "btn-ghost"
                                          }`}>
                                          <Settings className="w-5 h-5 mr-2" /> Settings
                                    </LinkQuery>

                                   <LinkQuery
                                          to="/members/vpns"
                                          className={`btn justify-start text-orange-400 ${
                                                isActive("/members/vpns")
                                                      ? "btn-primary"
                                                      : "btn-ghost"
                                          }`}>
                                          <EarthLock className="w-5 h-5 mr-2" /> Recommended VPNs
                                    </LinkQuery>                                    

                                    <button className="btn btn-ghost fixed bottom-0 w-64 py-4 pb-10 text-sm">
                                          <LogOut className="w-4 h-4" /><span className="text-xs">Logout</span>
                                    </button>
                              </nav>
                        </aside>
                  )}
            </RequireMember>            
	);
}
