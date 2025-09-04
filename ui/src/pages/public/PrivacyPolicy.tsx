import React from "react";
import NavMain from "../../components/nav/NavMain";
import FooterMain from "../../components/footer/FooterMain";

export default function PrivacyPolicy() {
   return (
      <main className="min-h-screen flex flex-col bg-base-100 text-base-content">
         <NavMain />

         {/* Hero Section */}
         <section className="relative min-h-[45vh] flex items-center justify-center bg-primary text-primary-content px-4 text-center">
            <div className="max-w-4xl mt-[50px]">
               <h1 className="text-5xl font-extrabold mb-4">Privacy Policy</h1>
               <p className="text-lg">
                  Your privacy matters. This policy explains how ComplyAge collects, uses, and protects your information.
               </p>
            </div>
         </section>

         {/* Privacy Policy Content */}
         <section className="py-16 px-4 max-w-3xl mx-auto text-base">

            <h2 className="text-3xl font-bold mb-6">1. Who We Are</h2>
            <p className="mb-4">
               ComplyAge (“we”, “us”, “our”) is a 501(c)(3) nonprofit providing privacy-first compliance and age verification services. We are committed to protecting your data and your privacy rights.
            </p>

            <h2 className="text-3xl font-bold mt-12 mb-6">2. Information We Collect</h2>
            <p className="mb-2">We collect the minimum data necessary to provide our services, including:</p>
            <ul className="mb-4 space-y-2 list-disc list-inside">
               <li>Account information: Such as email address or username, if you create an account.</li>
               <li>Verification data: Information you submit for compliance or verification (e.g., ID, address, phone, etc.).</li>
               <li>Technical data: Such as IP address, device/browser info, and logs, only as required for security and legal compliance.</li>
            </ul>

            <h2 className="text-3xl font-bold mt-12 mb-6">3. How We Use Your Data</h2>
            <ul className="mb-4 space-y-2 list-disc list-inside">
               <li>To provide and improve our verification services.</li>
               <li>To fulfill legal obligations and ensure compliance with applicable laws.</li>
               <li>To communicate with you, if you request support or notifications.</li>
            </ul>
            <p className="mb-4">
               We do <b>not</b> use your data for advertising, marketing, or profiling, and we never sell, trade, or rent your information.
            </p>

            <h2 className="text-3xl font-bold mt-12 mb-6">4. Data Encryption & Security</h2>
            <ul className="mb-4 space-y-2 list-disc list-inside">
               <li><b>Encryption by default:</b> All sensitive data you provide is encrypted at rest and in transit.</li>
               <li><b>Encryption strength:</b> You may have the ability to choose the level of encryption and key management for certain types of data (e.g., self-encrypted or server-encrypted).</li>
               <li><b>Limited access:</b> In many cases, only you control the keys to your data. Where we can’t access it, we cannot share or decrypt it—even in response to legal requests.</li>
            </ul>
            <p className="mb-4">
               While we use industry-standard safeguards, no method of transmission or storage is 100% secure. We cannot guarantee absolute security, but we work hard to keep your information protected.
            </p>

            <h2 className="text-3xl font-bold mt-12 mb-6">5. Data Sharing & Legal Compliance</h2>
            <ul className="mb-4 space-y-2 list-disc list-inside">
               <li>
                  <b>No sale or marketing:</b> We do not sell, share, or otherwise distribute your data to advertisers, marketers, or other third parties.
               </li>
               <li>
                  <b>Legal compliance:</b> We may disclose your information only in response to valid legal processes (such as subpoenas, court orders, or government requests), and only to the extent required by law.
               </li>
               <li>
                  <b>Notice:</b> If legally permitted, we will attempt to notify you prior to any disclosure of your data due to legal process.
               </li>
               <li>
                  <b>Third-party processors:</b> We use reputable service providers (e.g., cloud storage, payment processors) who must adhere to strict privacy and security obligations.
               </li>
            </ul>

            <h2 className="text-3xl font-bold mt-12 mb-6">6. Data Retention</h2>
            <p className="mb-4">
               We retain your data only as long as necessary for verification, compliance, or legal reasons. You may request deletion of your account and associated data at any time, subject to legal and regulatory retention requirements.
            </p>

            <h2 className="text-3xl font-bold mt-12 mb-6">7. Your Rights</h2>
            <ul className="mb-4 space-y-2 list-disc list-inside">
               <li>Request access to your personal information.</li>
               <li>Request correction, deletion, or export of your data.</li>
               <li>Request information about how your data is used.</li>
               <li>Withdraw consent where applicable.</li>
            </ul>
            <p className="mb-4">
               For requests, email <a href="mailto:support@complyage.com" className="underline link link-hover">support@complyage.com</a>. We honor requests to the extent permitted by law and technical feasibility.
            </p>

            <h2 className="text-3xl font-bold mt-12 mb-6">8. Children’s Privacy</h2>
            <p className="mb-4">
               ComplyAge does not knowingly collect information from children under 13. If you believe a child has provided us personal data without consent, please contact us and we will remove the data.
            </p>

            <h2 className="text-3xl font-bold mt-12 mb-6">9. Changes to this Policy</h2>
            <p className="mb-4">
               We may update this policy from time to time. You will always find the latest version on this page. Material changes will be announced in advance.
            </p>

            <h2 className="text-3xl font-bold mt-12 mb-6">10. Contact Us</h2>
            <p>
               For privacy questions, data requests, or legal concerns, email <a href="mailto:support@complyage.com" className="underline link link-hover">support@complyage.com</a> or use our <a href="/contact" className="underline link link-hover">contact form</a>.
            </p>
         </section>

         {/* Call to Action */}
         <section className="py-16 text-center bg-primary text-primary-content">
            <h2 className="text-3xl md:text-4xl font-bold mb-4">
               Privacy is your right—ComplyAge is here to protect it.
            </h2>
            <a href="/donate" className="btn btn-secondary text-xl px-8 py-4 mt-6">Support Privacy</a>
         </section>

         <FooterMain />
      </main>
   );
}
