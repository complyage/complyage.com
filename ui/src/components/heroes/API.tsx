//||------------------------------------------------------------------------------------------------||
//|| SupportHero.tsx
//|| src/components/heroes/SupportHero.tsx
//||------------------------------------------------------------------------------------------------||

import React from "react";

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function APIHero() {
      return (
            <section className="py-20 px-6 max-w-4xl mx-auto text-center">
                  <h3 className="text-2xl md:text-3xl font-bold mb-4">Need help setting up?</h3>
                  <p className="mb-6 text-lg leading-relaxed">
                        Our documentation is available on GitHub Pages for quick reference and examples.
                  </p>
                  <a
                        className="btn btn-xl bg-orange-200 text-black hover:bg-orange-300"
                        href="https://complyage.github.io/complyage.com/"
                        target="_blank"
                        rel="noopener noreferrer"
                  >
                        API Documentation
                  </a>
            </section>
      );
}
