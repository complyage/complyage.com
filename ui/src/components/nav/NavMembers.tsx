import { User } from "lucide-react";
import React from "react";
import {Link, useLocation} from "react-router-dom";

export default function NavMain() {
	const location = useLocation();

	return (
		<header className="fixed top-0 left-0 right-0 z-50 bg-gray-900 shadow-md">
			<div className="navbar px-6 max-w-7xl mx-auto">
				{/* Logo */}
				<div className="flex-1">
					<Link to="/" className="flex items-center">
						<img
							src="/img/logow.png"
							alt="ComplyAge Logo"
							className="h-12 w-auto p-2.5"
						/>
					</Link>
				</div>

				{/* Navigation Links */}
				<div className="flex-none flex gap-4 items-center">

					<Link to="/members" className="btn btn-secondary text-lg">
						<User className="w-5 h-5" />
					</Link>                              

				</div>
			</div>
		</header>
	);
}
