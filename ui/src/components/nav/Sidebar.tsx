//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React                                          from "react";
import {Link, useLocation}                            from "react-router-dom";
import {Home, Lock, Settings, Share, Globe, LogOut}   from "lucide-react";

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
		<aside className="w-64 bg-gray-800 fixed left-0 bottom-0 top-[80px] z-100">
			<nav className="flex flex-col gap-2 flex-1=">
				<Link
					to="/members/"
					className={`btn justify-start py-5 ${
						isActive("/members/")
							? "btn-primary"
							: "btn-ghost"
					}`}>
					<Home className="w-5 h-5 mr-2" /> Dashboard
				</Link>
				<Link
					to="/members/sites"
					className={`btn justify-start ${
						isActive("/members/sites")
							? "btn-primary"
							: "btn-ghost"
					}`}>
					<Globe className="w-5 h-5 mr-2" /> Sites
				</Link>
				<Link
					to="/members/encrypted"
					className={`btn justify-start ${
						isActive("/members/encrypted")
							? "btn-primary"
							: "btn-ghost"
					}`}>
					<Lock className="w-5 h-5 mr-2" /> Encrypted Data
				</Link>
				<Link
					to="/members/shared"
					className={`btn justify-start ${
						isActive("/members/shared")
							? "btn-primary"
							: "btn-ghost"
					}`}>
					<Share className="w-5 h-5 mr-2" /> Shared Data
				</Link>
				<Link
					to="/members/settings"
					className={`btn justify-start ${
						isActive("/members/settings")
							? "btn-primary"
							: "btn-ghost"
					}`}>
					<Settings className="w-5 h-5 mr-2" /> Settings
				</Link>

				<button className="btn btn-ghost fixed bottom-0 w-64 py-4 pb-10 text-sm">
					<LogOut className="w-4 h-4" /><span className="text-xs">Logout</span>
				</button>
			</nav>
		</aside>
	);
}
