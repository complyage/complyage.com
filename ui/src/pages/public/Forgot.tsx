//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React                           from "react";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import NavMain                         from "../../components/nav/NavMain";
import FooterMain                      from "../../components/footer/FooterMain";
import AuthForgot                      from "../auth/AuthForgot";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Forgot() {
      return (
            <main className="min-h-screen flex flex-col">
                  <NavMain />

                  <div className="relative flex-1 flex items-center justify-center">
                        {/* Background Image */}
                        <img 
                              src="/img/hero/forgot.webp"
                              alt="Background"
                              className="absolute inset-0 w-full h-full object-cover" 
                        />

                        {/* Centered Box */}
                        <div className="relative bg-black/80 p-8 rounded-lg max-w-2xl w-full m-10">
                              <AuthForgot />
                        </div>
                  </div>

                  <FooterMain />
            </main>
      );
}
