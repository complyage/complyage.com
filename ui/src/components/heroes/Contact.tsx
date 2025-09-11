//||------------------------------------------------------------------------------------------------||
//|| ContactHero.tsx
//|| src/components/heroes/ContactHero.tsx
//||------------------------------------------------------------------------------------------------||

import React from "react";

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function ContactHero() {
      return (
            <section className="py-20 px-4 text-center bg-gray-800">
                  <div className="w-4xl mx-auto">
                        <h2 className="text-4xl md:text-5xl font-extrabold mb-10">
                              We’re building privacy and compliance for everyone.
                        </h2>
                        <p className="text-lg md:text-xl mb-8 max-w-3xl mx-auto">
                              Your feedback makes us better. Get in touch if you have ideas or concerns.
                        </p>
                        <a 
                              href="/contact" 
                              className="btn btn-secondary text-lg md:text-xl px-10 py-4 font-bold shadow-md"
                        >
                              Contact Us
                        </a>
                  </div>
            </section>
      );
}
