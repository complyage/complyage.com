import React from "react";
import {Link, useLocation} from "react-router-dom";

export default function NavMain() {
	const location = useLocation();

	return (
		<header className="fixed top-0 left-0 right-0 z-50 bg-base-100 shadow-md">
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
					<Link
						to="/about"
						className={`btn btn-ghost hover:text-orange-300 text-xl ${
							location.pathname === "/about"
								? "text-orange-500"
								: ""
						}`}>
						About
					</Link>
					<Link
						to="/vendors"
						className={`btn btn-ghost hover:text-orange-300 text-xl ${
							location.pathname === "/vendors"
								? "text-orange-500"
								: ""
						}`}>
						Vendors
					</Link>
					<Link
						to="/gilead"
						className={`btn btn-ghost hover:text-orange-300 text-xl ${
							location.pathname === "/gilead"
								? "text-orange-500"
								: ""
						}`}>
						Enforcement
					</Link>                              
					<Link
						to="/pricing"
						className={`btn btn-ghost hover:text-orange-300 text-xl ${
							location.pathname === "/pricing"
								? "text-orange-500"
								: ""
						}`}>
						Pricing
					</Link>
					<Link
						to="/login"
						className={`btn btn-ghost hover:text-orange-300 text-xl ${
							location.pathname === "/login"
								? "text-orange-500"
								: ""
						}`}>
						Login
					</Link>
					<Link to="/signup" className="btn btn-primary">
						Sign Up
					</Link>
				</div>
			</div>
		</header>
	);
}
