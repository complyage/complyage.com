//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React                           from "react";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import NavMain                         from "../../components/nav/NavMain";
import FooterMain                      from "../../components/footer/FooterMain";
import AuthComplete                    from "../auth/AuthComplete";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Complete() {
      return (
            <main className="min-h-screen flex flex-col">
                  <NavMain />

                  <div className="relative flex-1">
                        <img
                              src="/img/hero/complete.webp"
                              alt="Background"
                              className="absolute inset-0 w-full h-full object-cover"
                        />
                        <div className="absolute inset-0 bg-black/70"></div>

                        <AuthComplete />
                  </div>

                  <FooterMain />
            </main>
      );
}
