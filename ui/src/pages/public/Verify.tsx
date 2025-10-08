//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React                           from "react";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import NavMain                         from "../../components/nav/NavMain";
import FooterMain                      from "../../components/footer/FooterMain";
import AuthVerify                      from "../auth/AuthVerify";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Verify() {

      return (
            <main className="min-h-screen flex flex-col">
                  <NavMain />

                  <div className="relative flex-1">
                        {/* Background */}
                        <img
                              src="/img/hero/verify.webp"
                              alt="Background"
                              className="absolute inset-0 w-full h-full object-cover"
                        />
                        <div className="absolute inset-0 bg-black/70"></div>

                        {/* Content */}
                        <div className="relative z-10 flex flex-col items-center justify-center min-h-screen pt-24 pb-12 px-4">
                              <AuthVerify />
                        </div>
                  </div>

                  <FooterMain />
            </main>
      );
}
