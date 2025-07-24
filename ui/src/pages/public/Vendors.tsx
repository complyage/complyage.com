import React from "react";
import NavMain from "../../components/nav/NavMain";
import FooterMain from "../../components/footer/FooterMain";

export default function Vendors() {
	return (
		<main className="min-h-screen flex flex-col bg-base-100 text-base-content">
			<NavMain />

			{/* Hero */}
			<section className="relative flex flex-col items-center justify-center text-center py-20 bg-primary text-primary-content px-4">
				<h1 className="text-5xl font-extrabold mb-4 mt-[120px]">
					For Vendors & Site Owners
				</h1>
				<p className="text-lg max-w-2xl">
					Easily add age verification to your site. Stay
					compliant with local laws — without tracking your
					users.
				</p>
			</section>

			{/* How It Works */}
			<section className="py-20 px-4 max-w-6xl mx-auto">
				<h2 className="text-4xl font-bold mb-12 text-center">
					How to Set Up ComplyAge
				</h2>
				<div className="grid md:grid-cols-2 gap-12">
					{/* Location-Based Tracking */}
					<div className="bg-base-200 rounded-lg shadow-lg p-8">
						<h3 className="text-3xl font-bold mb-4">
							1. Location-Based Tracking
						</h3>
						<p className="mb-4">
							Only verify visitors in regions that
							have age verification laws. Our
							location-based system automatically
							checks the visitor’s jurisdiction and
							enforces compliance only where
							required.
						</p>
						<ul className="list-disc list-inside mb-4">
							<li>
								Lightweight script — easy to
								install
							</li>
							<li>Automatic region detection</li>
							<li>
								No invasive tracking or
								fingerprinting
							</li>
						</ul>
						<p>
							Ideal for sites that want minimum
							friction and regional compliance
							without impacting all users globally.
						</p>
					</div>

					{/* OAuth Registration */}
					<div className="bg-base-200 rounded-lg shadow-lg p-8">
						<h3 className="text-3xl font-bold mb-4">
							2. OAuth Registration
						</h3>
						<p className="mb-4">
							Add ComplyAge as an OAuth provider —
							just like you would with Google or
							Facebook Login. During registration,
							your users securely share only the
							data you need to comply with age and
							jurisdiction requirements.
						</p>
						<ul className="list-disc list-inside mb-4">
							<li>
								Standard OAuth 2.0 — works with
								any modern stack
							</li>
							<li>
								You control what data to request
								(age, region, verification
								status)
							</li>
							<li>
								Fully user-consented and
								privacy-friendly
							</li>
						</ul>
						<p>
							Ideal for platforms with user accounts
							that need continuous, verified access.
						</p>
					</div>
				</div>
			</section>

			{/* CTA */}
			<section className="py-16 text-center bg-primary text-primary-content">
				<h2 className="text-3xl md:text-4xl font-bold mb-4">
					Ready to integrate age verification?
				</h2>
				<p className="mb-6">
					It takes just minutes to stay compliant and
					protect your business — without sacrificing user
					trust.
				</p>
				<button className="btn btn-secondary text-xl px-8 py-4">
					Get Started Now
				</button>
			</section>

			{/* Support Section */}
			<section className="py-12 px-4 max-w-4xl mx-auto text-center">
				<h3 className="text-2xl font-bold mb-2">
					Need help setting up?
				</h3>
				<p>
					Our documentation can be found in our github pages
					<div className="block p-4">
						<a
							className="btn bg-orange-200 h-16 p-5 text-black text-2xl"
							href="https://complyage.github.io/complyage.com/"
							target="_blank"
							rel="noopener noreferrer">
							API Documentation
						</a>
					</div>
				</p>
			</section>

			<FooterMain />
		</main>
	);
}
