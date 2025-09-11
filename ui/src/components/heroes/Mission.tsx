//||------------------------------------------------------------------------------------------------||
//|| MissionHero
//|| src/components/heroes/MissionHero.tsx
//||------------------------------------------------------------------------------------------------||

import React from "react";

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function MissionHero() {
      return (
            <section className="py-24 px-6 bg-base-100 text-center">
                  <div className="max-w-5xl mx-auto">
                        <h2 className="text-4xl md:text-5xl font-extrabold mb-10">
                              Our Mission
                        </h2>
                        <div className="space-y-8 text-lg leading-relaxed text-base-content/90">
                              <p>
                                    Around the world, lawmakers demand that platforms enforce rules — age checks, identity verifications,
                                    compliance controls. But the tools offered today are broken. They harvest personal data, sell it,
                                    and lock it away in walled gardens. They force platforms to compromise trust, and they force
                                    users to give up their dignity.
                              </p>
                              <p>
                                    <span className="font-semibold">ComplyAge exists to break that cycle.</span>
                                    We are building privacy-first compliance infrastructure that empowers platforms to stay
                                    legal without ever exploiting their communities.
                              </p>
                              <p>
                                    Our mission is bigger than checkboxes.
                                    It’s about protecting <span className="font-bold text-accent">the right to exist online without surveillance,
                                    without data trafficking, and without fear</span>.
                                    We stand for compliance that serves people, not corporations.
                              </p>
                        </div>
                  </div>
            </section>
      );
}
