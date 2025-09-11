//||------------------------------------------------------------------------------------------------||
//|| VisionHero.tsx
//|| src/components/heroes/VisionHero.tsx
//||------------------------------------------------------------------------------------------------||

import React from "react";

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function VisionHero() {
      return (
            <section className="relative py-24 px-6 bg-base-100 text-center">
                  <div className="max-w-6xl mx-auto">
                        <h2 className="text-4xl md:text-5xl font-extrabold mb-8 tracking-tight">
                              Our Vision
                        </h2>
                        <p className="text-lg md:text-xl leading-relaxed max-w-4xl mx-auto text-base-content/80">
                              We envision a digital world where <span className="font-bold text-accent">compliance protects instead of exploits</span>.  
                              Where businesses can stay safe without being forced into the role of data brokers.  
                              Where an age check is just that — a check, not a permanent record, not another surveillance tool.  
                              <br /><br />
                              We believe every person, no matter where they live, deserves the right to step online  
                              <span className="font-bold"> freely, privately, and without compromise.</span>  
                              Your personal details should never be currency. Your privacy should never be negotiable.  
                              Compliance should never mean control.  
                              <br /><br />
                              <span className="text-xl font-bold text-accent">
                                    This is not just compliance. This is digital freedom — rebuilt from the ground up,  
                                    with honesty, transparency, and accountability at its core.
                              </span>
                        </p>
                  </div>
            </section>
      );
}
