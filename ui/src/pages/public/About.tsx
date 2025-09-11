/*||------------------------------------------------------------------------------------------------||
//|| About Page
//|| src/routes/about.tsx
//||------------------------------------------------------------------------------------------------||*/

/*||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||*/

import React                     from "react";

/*||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||*/

import NavMain                   from "../../components/nav/NavMain";
import MainHero                  from "../../components/heroes/Main";
import FooterMain                from "../../components/footer/FooterMain";
import SignupHero                from "../../components/heroes/Signup";
import CorePrinciplesHero        from "../../components/heroes/CorePrinciples";
import VisionHero                from "../../components/heroes/Vision";
import MissionHero               from "../../components/heroes/Mission";

/*||------------------------------------------------------------------------------------------------||
//|| Page
//||------------------------------------------------------------------------------------------------||*/

export default function About() {

      /*||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||*/
      return (
            <main className="min-h-screen flex flex-col bg-base-100 text-base-content">
                  <NavMain />

                  {/* Hero */}
                  <MainHero
                        image="/img/hero/liberty.webp"
                        title="About ComplyAge"
                        description="We are not just building software. We are building a movement. A future where compliance never means surrendering privacy, and freedom is the default state of the internet."
                        ctaLabel="Join the Movement"
                        ctaHref="/signup"
                  />

                  {/* Mission */}
                  <MissionHero />

                  {/* Principles */}
                  <CorePrinciplesHero />

                  {/* Vision */}
                  <VisionHero />

                  {/* Signup */}
                  <SignupHero />

                  <FooterMain />
            </main>
      );
}
