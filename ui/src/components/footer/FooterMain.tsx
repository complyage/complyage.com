import React from "react";
import {Link} from "react-router-dom";
import {ShieldCheck, EyeOff, UserX} from "lucide-react";
import LinkQuery          from "../../components/dynamic/LinkQuery";

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
							<LinkQuery to="/about" className="hover:underline">
								About Us
							</LinkQuery>
						</li>
						<li>
							<LinkQuery to="/contact" className="hover:underline">
								Contact Us
							</LinkQuery>
						</li>
						<li>
							<LinkQuery to="/terms" className="hover:underline">
								Terms of Service
							</LinkQuery>
						</li>
						<li>
							<LinkQuery to="/privacy" className="hover:underline">
								Privacy Policy
							</LinkQuery>
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
					<h4 className="font-bold text-lg mb-4">We don't want to know</h4>
					<p className="text-gray-400 mb-4">No spam. No tracking. No hidden goals.</p>
					<div className="flex text-gray-500">
                                    This is where we'd ask for your email, but we respect your privacy too much to do that.
                                    We don't want your email.
                                    So, no newsletter signup here.
                              </div>
				</div>
			</div>

			<div className="mt-12 border-t border-gray-800 pt-6 text-center text-gray-500 text-sm">
				© {new Date().getFullYear()} ComplyAge. All rights reserved. Built by normal people for normal people.
			</div>
		</footer>
	);
}
