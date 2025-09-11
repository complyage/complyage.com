//||------------------------------------------------------------------------------------------------||
//|| MonthlySupport
//|| src/components/widgets/MonthlySupport.tsx
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useState } from "react";
import Call                           from "../../classes/call";
import ProgressBar                    from "../base/ProgressBar";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

interface MonthlySupportHeroProps {
      cta?: boolean;
}

interface MonthlyTotalsResponse {
      donated : number;
      bills   : number;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function MonthlySupportHero({ cta = true }: MonthlySupportHeroProps) {

      //||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||

      const [donated, setDonated] = useState<number>(0);
      const [bills, setBills]     = useState<number>(0);
      const [loading, setLoading] = useState<boolean>(true);

      //||------------------------------------------------------------------------------------------------||
      //|| Fetch Data
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            (async () => { 
                  const chirp = new Call("/public/donated", {});
                  chirp.method = "GET";
                  chirp.debug  = true;
                  await chirp.execute();
                  if (chirp.ok()) {
                        setDonated(chirp.data("donated") || 0);
                        setBills(chirp.data("bills") || 0);
                  }
                  setLoading(false);
            })();
      }, []);

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      const goal = bills > 0 ? bills : .01; // fallback to avoid /0
      const current = donated;

      return (
            <section className="py-8 px-4 bg-gray-800 text-base-content text-center">
                  <h3 className="text-2xl font-bold mb-4">Monthly Support Raised</h3>
                  {loading ? (
                        <p className="text-gray-400">Loading...</p>
                  ) : (
                        <>
                              <ProgressBar
                                    current={current}
                                    goal={goal}
                                    prefix="$"
                                    label={` of goal \$${Math.floor(bills)}`}
                              />
                              {cta && (
                                    <div className="mt-6">
                                          <a
                                                href="/donate"
                                                className="btn btn-secondary btn-lg px-10 font-bold shadow-md"
                                          >
                                                Donate
                                          </a>
                                    </div>
                              )}
                        </>
                  )}
            </section>
      );
}
