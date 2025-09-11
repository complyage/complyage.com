//||------------------------------------------------------------------------------------------------||
//|| SignupHero.tsx
//||------------------------------------------------------------------------------------------------||

import React from "react";

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function SignupHero() {
      return (
            <section className="py-20 px-6 text-center bg-primary text-primary-content">
                  <div className="max-w-3xl mx-auto">
                        <h2 className="text-4xl md:text-5xl font-extrabold mb-6 leading-tight">
                              Ready to get started?
                        </h2>
                        <p className="text-lg md:text-xl mb-8 opacity-90">
                              The fight for digital freedom starts with small choices.  
                              By joining ComplyAge, you’re choosing a future where privacy is preserved,  
                              compliance is simple, and online communities remain open and safe.  
                              Together, we can set a new standard for trust, transparency, and freedom.
                        </p>
                        <button
                              className="btn btn-secondary btn-xl"
                              onClick={() => window.location.href = "/signup"}
                        >
                              Sign Up Now
                        </button>
                        <p className="mt-5 text-sm opacity-75">
                              It only takes a minute to join — your freedom lasts a lifetime.
                        </p>
                  </div>
            </section>
      );
}
