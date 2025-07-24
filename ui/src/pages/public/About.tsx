import React from "react";
import NavMain from "../../components/nav/NavMain";
import FooterMain from "../../components/footer/FooterMain";

export default function About() {
   return (
      <main className="min-h-screen flex flex-col bg-base-100 text-base-content">
         <NavMain />

         {/* Hero Section */}
         <section className="relative min-h-[50vh] flex items-center justify-center bg-primary text-primary-content px-4 text-center">
            <div className="max-w-4xl mt-[50px]">
               <h1 className="text-5xl font-extrabold mb-4">About Us</h1>
               <p className="text-lg">
                  Learn more about our mission to protect privacy and empower platforms to stay compliant without compromise.
               </p>
            </div>
         </section>

         {/* Our Mission */}
         <section className="py-16 px-4 max-w-6xl mx-auto">
            <h2 className="text-4xl font-bold mb-6">Our Mission</h2>
            <p className="mb-4">
               At ComplyAge, we believe privacy is a fundamental right. Our goal is to help platforms around the world
               verify age requirements with maximum privacy, minimum hassle, and zero data exploitation.
            </p>
            <p>
               From open-source code to clear, simple policies, we’re building a better standard for age verification
               that keeps your users anonymous and your business secure.
            </p>
         </section>

         {/* Core Values */}
         <section className="py-16 px-4 bg-base-200">
            <div className="max-w-6xl mx-auto">
               <h2 className="text-4xl font-bold mb-12 text-center">Our Principles</h2>
               <div className="grid md:grid-cols-3 gap-8">
                  <div className="card bg-base-100 shadow-lg">
                     <div className="card-body">
                        <h3 className="card-title">Transparency</h3>
                        <p>Our code is open-source and our terms are clear. No hidden loopholes, ever.</p>
                     </div>
                  </div>
                  <div className="card bg-base-100 shadow-lg">
                     <div className="card-body">
                        <h3 className="card-title">Privacy First</h3>
                        <p>We never sell data or track you. Your information stays yours — encrypted, private, safe.</p>
                     </div>
                  </div>
                  <div className="card bg-base-100 shadow-lg">
                     <div className="card-body">
                        <h3 className="card-title">Freedom & Choice</h3>
                        <p>Users stay anonymous when they want to be. Platforms stay compliant without compromising freedom.</p>
                     </div>
                  </div>
               </div>
            </div>
         </section>

         {/* Call to Action */}
         <section className="py-16 text-center bg-primary text-primary-content">
            <h2 className="text-3xl md:text-4xl font-bold mb-4">
               Ready to get started?
            </h2>
            <p className="mb-6">Join us in building a freer, more private internet — for everyone.</p>
            <button className="btn btn-secondary text-xl px-8 py-4">Sign Up Now</button>
         </section>

         <FooterMain />
      </main>
   );
}
