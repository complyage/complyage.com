/*||------------------------------------------------------------------------------------------------||
//|| TOS
//||------------------------------------------------------------------------------------------------||*/

/*||------------------------------------------------------------------------------------------------||
//|| React
//||------------------------------------------------------------------------------------------------||*/

import React from "react";

/*||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||*/

import NavMain                      from "../../components/nav/NavMain";
import FooterMain                   from "../../components/footer/FooterMain";
import MainHero                     from "../../components/heroes/Main";
import ContactHero                  from "../../components/heroes/Contact";

/*||------------------------------------------------------------------------------------------------||
//|| TOS
//||------------------------------------------------------------------------------------------------||*/

export default function TermsOfService() {
	return (
		<main className="min-h-screen flex flex-col bg-base-100 text-base-content">
			<NavMain />

			{/* Hero */}

			<MainHero
				image="/img/hero/tos.webp"
				title="Terms of Service"
				description="These terms are for your protection and ours. Please read them carefully before using ComplyAge."
			/>

			{/* Terms Content */}
			<section className="py-16 px-4 max-w-3xl mx-auto">
				<h2 className="text-3xl font-bold mb-6">1. About ComplyAge</h2>
				<p className="mb-4">
					ComplyAge is a 501(c)(3) nonprofit dedicated to privacy-first compliance and verification. Our mission is to protect your data,
					empower platform owners, and make the internet safer for everyone.
				</p>

				<h2 className="text-3xl font-bold mt-12 mb-6">2. Privacy and Data Use</h2>
				<ul className="mb-4 space-y-2 list-disc list-inside">
					<li>
						<b>No selling, sharing, or trading:</b> We do <b>not</b> sell, share, trade, or give away your data, except as required by
						law or if you explicitly request it (such as exporting your own data).
					</li>
					<li>
						<b>Encryption by default:</b> All sensitive data is stored encrypted, and the strength of encryption depends on the options
						you select when using our platform.
					</li>
					<li>
						<b>Anonymous by design:</b> We make it possible for you to use our platform without providing personally identifying
						information, wherever the law allows.
					</li>
					<li>
						<b>We can’t access what we can’t decrypt:</b> In some cases, only you hold the keys to your encrypted data. If we can’t
						access it, we can’t share it—even if someone asks.
					</li>
				</ul>
				<p className="mb-4">
					For more on our privacy practices, please review our{" "}
					<a href="/privacy" className="link link-hover underline">
						Privacy Policy
					</a>
					.
				</p>

				<h2 className="text-3xl font-bold mt-12 mb-6">3. Legal Compliance</h2>
				<p className="mb-4">
					We fully comply with all valid legal processes, including subpoenas, court orders, and lawful government requests. If we are
					legally required to provide data, we will do so—but only to the extent we can access it and only what the law requires.
				</p>
				<p className="mb-4">If permitted, we will attempt to notify affected users before disclosing any information.</p>

				<h2 className="text-3xl font-bold mt-12 mb-6">4. Data Security</h2>
				<p className="mb-4">
					We use industry-standard encryption and security measures. The actual encryption level (such as key strength or storage options)
					may depend on the privacy options you select when verifying or submitting information.
				</p>
				<p className="mb-4">
					Some encrypted data may be accessible only to you; other data may be accessible to our staff for support, compliance, or account
					recovery purposes.
				</p>

				<h2 className="text-3xl font-bold mt-12 mb-6">5. Service Usage</h2>
				<ul className="mb-4 space-y-2 list-disc list-inside">
					<li>Do not use ComplyAge for illegal purposes or to circumvent lawful compliance requirements.</li>
					<li>You are responsible for the data you submit and for keeping your login information secure.</li>
					<li>You may not attempt to reverse engineer, disrupt, or abuse our services.</li>
				</ul>

				<h2 className="text-3xl font-bold mt-12 mb-6">6. Limitation of Liability</h2>
				<p className="mb-4">
					ComplyAge is provided “as is.” We do our best, but cannot guarantee uninterrupted service or absolute security. We are not liable
					for losses, damages, or claims arising from use of our services, to the fullest extent permitted by law.
				</p>

				<h2 className="text-3xl font-bold mt-12 mb-6">7. Changes to Terms</h2>
				<p className="mb-4">
					We may update these terms from time to time. You’ll always find the latest version on this page. Continued use of ComplyAge means
					you accept the updated terms.
				</p>

				<h2 className="text-3xl font-bold mt-12 mb-6">8. Contact</h2>
				<p>
					For questions, requests, or legal concerns, email us at{" "}
					<a href="mailto:support@complyage.com" className="link link-hover underline">
						support@complyage.com
					</a>
					.
				</p>
			</section>

			{/* Call to Action */}
			<ContactHero />

			<FooterMain />
		</main>
	);
}
