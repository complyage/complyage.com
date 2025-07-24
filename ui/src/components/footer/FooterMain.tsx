import React from "react";
import {Link} from "react-router-dom";
import {ShieldCheck, EyeOff, UserX} from "lucide-react";

export default function FooterMain() {
	return (
		<footer className="bg-black text-white py-12 px-4">
			<div className="max-w-6xl mx-auto grid md:grid-cols-4 gap-8">
				{/* Logo & Mission */}
				<div>
					<img src="/img/logo-white.png" alt="ComplyAge Logo" className="w-20 mb-4" />
					<p className="text-gray-400">Taking back your privacy & protecting your freedom since 2024.</p>
				</div>

				{/* Quick Links */}
				<div>
					<h4 className="font-bold text-lg mb-4">Quick Links</h4>
					<ul className="space-y-2">
						<li>
							<Link to="/about" className="hover:underline">
								About Us
							</Link>
						</li>
						<li>
							<Link to="/terms" className="hover:underline">
								Terms of Service
							</Link>
						</li>
						<li>
							<Link to="/privacy" className="hover:underline">
								Privacy Policy
							</Link>
						</li>
						<li>
							<Link to="/support" className="hover:underline">
								Support
							</Link>
						</li>
					</ul>
				</div>

				{/* Trust Badges */}
				<div>
					<h4 className="font-bold text-lg mb-4">We Stand For</h4>
					<div className="flex flex-col gap-3">
						<div className="flex items-center gap-2">
							<ShieldCheck className="w-5 h-5 text-success" />
							<span>Open Source</span>
						</div>
						<div className="flex items-center gap-2">
							<EyeOff className="w-5 h-5 text-success" />
							<span>No Tracking</span>
						</div>
						<div className="flex items-center gap-2">
							<UserX className="w-5 h-5 text-success" />
							<span>Stay Anonymous</span>
						</div>
					</div>
				</div>

				{/* Newsletter */}
				<div>
					<h4 className="font-bold text-lg mb-4">Stay in the Loop</h4>
					<p className="text-gray-400 mb-4">No spam. No tracking. Just updates on your rights.</p>
					<form className="flex">
						<input
							type="email"
							placeholder="Your email"
							className="p-2 rounded-l bg-gray-800 border border-gray-600 text-white flex-1"
						/>
						<button type="submit" className="p-2 bg-primary text-black rounded-r font-bold hover:bg-primary-focus">
							Join
						</button>
					</form>
				</div>
			</div>

			<div className="mt-12 border-t border-gray-800 pt-6 text-center text-gray-500 text-sm">
				© {new Date().getFullYear()} ComplyAge. All rights reserved. Built by rebels for rebels.
			</div>
		</footer>
	);
}
