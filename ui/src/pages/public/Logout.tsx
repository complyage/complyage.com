//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React                           from "react";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import NavMain                         from "../../components/nav/NavMain";
import FooterMain                      from "../../components/footer/FooterMain";
import AuthLogout                      from "../auth/AuthLogout";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Logout() {
      return (
            <main className="min-h-screen flex flex-col">
                  <NavMain />
                  <img
                        src="/img/hero/logout.webp"
                        alt="Background"
                        className="absolute inset-0 w-full h-full object-cover"
                  />
                  <div className="relative flex-1">
                        <AuthLogout />
                  </div>

                  <FooterMain />
            </main>
      );
}
