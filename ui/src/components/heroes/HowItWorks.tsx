//||------------------------------------------------------------------------------------------------||
//|| HowItWorksHero.tsx
//|| src/components/heroes/HowItWorksHero.tsx
//||------------------------------------------------------------------------------------------------||

import React from "react";

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function HowItWorksHero() {
      return (
            <section className="py-20 px-4 max-w-6xl mx-auto">
                  <h2 className="text-4xl font-bold mb-12 text-center">How to Set Up ComplyAge</h2>
                  <div className="grid md:grid-cols-2 gap-12">
                        
                        {/* Location-Based Tracking */}
                        <div className="bg-base-200 rounded-lg shadow-lg p-8">
                              <h3 className="text-3xl font-bold mb-4">1. Location-Based Tracking</h3>
                              <p className="mb-4">
                                    Only verify visitors in regions that have age verification laws.  
                                    Our system automatically checks the visitor’s jurisdiction and enforces 
                                    compliance only where required.
                              </p>
                              <ul className="list-disc list-inside mb-4">
                                    <li>Lightweight script — easy to install</li>
                                    <li>Automatic region detection</li>
                                    <li>No invasive tracking or fingerprinting</li>
                              </ul>
                              <p>
                                    Ideal for sites that want minimum friction and regional compliance without impacting all users globally.
                              </p>
                        </div>

                        {/* OAuth Registration */}
                        <div className="bg-base-200 rounded-lg shadow-lg p-8">
                              <h3 className="text-3xl font-bold mb-4">2. OAuth Registration</h3>
                              <p className="mb-4">
                                    Add ComplyAge as an OAuth provider — just like Google or Facebook Login.  
                                    During registration, users securely share only the data you need to comply 
                                    with age and jurisdiction requirements.
                              </p>
                              <ul className="list-disc list-inside mb-4">
                                    <li>Standard OAuth 2.0 — works with any modern stack</li>
                                    <li>You control what data to request (age, region, verification status)</li>
                                    <li>Fully user-consented and privacy-friendly</li>
                              </ul>
                              <p>
                                    Ideal for platforms with user accounts that need continuous, verified access.
                              </p>
                        </div>
                  </div>
            </section>
      );
}
