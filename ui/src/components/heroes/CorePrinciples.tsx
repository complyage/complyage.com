//||------------------------------------------------------------------------------------------------||
//|| CorePrinciples Hero
//|| src/components/heroes/CorePrinciples.tsx
//||------------------------------------------------------------------------------------------------||

import React from "react";

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function CorePrinciplesHero() {
      return (
            <section className="py-24 px-6 bg-base-200">
                  <div className="max-w-6xl mx-auto text-center">
                        <h2 className="text-4xl md:text-5xl font-extrabold mb-6 tracking-tight">
                              Our Core Principles
                        </h2>
                        <p className="max-w-3xl mx-auto mb-14 text-lg md:text-xl text-base-content/70 leading-relaxed">
                              These are the pillars that guide everything we build, publish, and stand for.  
                              They are more than words — they are commitments we live by every single day.
                        </p>
                  </div>

                  <div className="grid md:grid-cols-3 gap-10">
                        <div className="card bg-base-100 shadow-xl border border-base-300">
                              <div className="card-body">
                                    <h3 className="card-title text-2xl mb-4">Transparency</h3>
                                    <p className="leading-relaxed text-base-content/80">
                                          Every line of code is open-source. Every rule we follow is clear and public.  
                                          No backdoors, no hidden agendas — trust is earned by showing the world how we operate.
                                    </p>
                              </div>
                        </div>

                        <div className="card bg-base-100 shadow-xl border border-base-300">
                              <div className="card-body">
                                    <h3 className="card-title text-2xl mb-4">Privacy First</h3>
                                    <p className="leading-relaxed text-base-content/80">
                                          Data is never for sale, and tracking is never an option.  
                                          Verification should protect your rights — not put them at risk.  
                                          Your identity belongs to you, always.
                                    </p>
                              </div>
                        </div>

                        <div className="card bg-base-100 shadow-xl border border-base-300">
                              <div className="card-body">
                                    <h3 className="card-title text-2xl mb-4">Freedom & Choice</h3>
                                    <p className="leading-relaxed text-base-content/80">
                                          Freedom online is not a privilege — it is essential.  
                                          Platforms stay compliant without betraying their communities,  
                                          and users stay anonymous without losing access to what matters.
                                    </p>
                              </div>
                        </div>
                  </div>
            </section>
      );
}
