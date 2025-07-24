import React from "react";
import NavMain from "../../components/nav/NavMain";
import FooterMain from "../../components/footer/FooterMain";
import ProgressBar from "../../components/base/ProgressBar";
export default function Pricing() {
	return (
		<main className="min-h-screen flex flex-col bg-base-100 text-base-content">
			<NavMain />

			{/* Hero */}
			<section className="relative flex flex-col items-center justify-center text-center py-20 bg-primary text-primary-content px-4">
				<h1 className="text-5xl font-extrabold mb-4 mt-[120px]">Simple, Transparent Pricing</h1>
				<p className="text-lg max-w-2xl">
					Protect your privacy and stay compliant — always free for platforms and free for individuals, with one small optional cost if you
					choose physical address verification.
				</p>
			</section>

			{/* Pricing Details */}
			<section className="py-20 px-4 max-w-6xl mx-auto grid md:grid-cols-2 gap-12">
				{/* For Platforms */}
				<div className="bg-base-200 rounded-lg shadow-lg p-8 flex flex-col justify-between">
					<h2 className="text-3xl font-bold mb-4">For Platforms & Businesses</h2>
					<p className="mb-6 text-lg">
						<strong>Free.</strong>
					</p>
					<p className="mb-6">
						We don't charge platforms or businesses for our core verification service. Stay compliant, protect user privacy, and scale
						freely - F*** contracts, F*** hidden costs. Lets weather this storm together.
					</p>
					<button className="btn btn-primary w-full">Start Verifying</button>
				</div>

				{/* For Consumers */}
				<div className="bg-base-200 rounded-lg shadow-lg p-8 flex flex-col justify-between">
					<h2 className="text-3xl font-bold mb-4">For Consumers</h2>
					<p className="mb-6 text-lg">
						<strong>100% Free.</strong>
					</p>
					<p className="mb-6">
						Age or ID verifications are free - forever. The only thing we charge for is optional address verification, which costs a
						small fee to cover postage and materials - no markup, full transparency. We're not making money. We're just not holding the
						bag for the postal service.
					</p>
					<button className="btn btn-secondary w-full">Create Free Account</button>
				</div>
			</section>

			{/* Donations Progress */}
			<section className="py-8 px-4 bg-black text-base-content text-center">
				<h3 className="text-2xl font-bold mb-4">Monthly Support Raised</h3>
				<ProgressBar current={0} goal={4314} />
			</section>
                  
			{/* How It's Possible */}
			<section className="py-20 px-4 max-w-4xl mx-auto text-center">
				<h2 className="text-3xl md:text-4xl font-bold mb-4">How Are We Free?</h2>
				<p className="mb-4">
					Our mission is to keep privacy and compliance accessible for everyone. We partner with carefully vetted, non-tracking sponsors who
					support us with respectful ads from reputable companies.
				</p>
				<p className="mb-4">
					We also rely on our community. Donations help cover our operational costs so we can remain independent, open source, and
					transparent.
				</p>
				<p className="mb-6">If you’d like to support our mission, you can donate securely via PayPal or Bitcoin.</p>
				<div className="flex flex-col md:flex-row justify-center gap-4">
					<a href="https://www.paypal.com/donate" target="_blank" rel="noopener noreferrer" className="btn btn-primary px-8 py-4">
						Donate with PayPal
					</a>
					<a href="bitcoin:YOUR-BITCOIN-ADDRESS" className="btn btn-secondary px-8 py-4">
						Donate with Bitcoin
					</a>
				</div>
			</section>



			{/* Call to Action */}
			<section className="py-16 text-center bg-primary text-primary-content">
				<h2 className="text-3xl md:text-4xl font-bold mb-4">Ready to verify and stay private?</h2>
				<p className="mb-6">Join thousands protecting their platforms and their privacy — for free.</p>
				<button className="btn btn-secondary text-xl px-8 py-4">Sign Up Now</button>
			</section>

			<FooterMain />
		</main>
	);
}
