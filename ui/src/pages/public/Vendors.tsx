/*||------------------------------------------------------------------------------------------------||
//|| Vendors Page
//|| src/routes/vendors.tsx
//||------------------------------------------------------------------------------------------------||*/

/*||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||*/

import React                     from "react";

/*||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||*/

import NavMain                   from "../../components/nav/NavMain";
import FooterMain                from "../../components/footer/FooterMain";
import MainHero                  from "../../components/heroes/Main";
import HowItWorksHero            from "../../components/heroes/HowItWorks";
import CTAHero                   from "../../components/heroes/CTA";
import APIHero                   from "../../components/heroes/API";

/*||------------------------------------------------------------------------------------------------||
//|| Page
//||------------------------------------------------------------------------------------------------||*/

export default function Vendors() {

      /*||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||*/
      return (
            <main className="min-h-screen flex flex-col bg-base-100 text-base-content">
                  <NavMain />

                  <MainHero
                        image="/img/hero/vendors.webp"
                        title="For Vendors & Site Owners"
                        description="Easily add age verification to your site. Stay compliant with local laws — without tracking your users."
                        ctaLabel="Get Started"
                        ctaHref="/signup"
                  />
                  
                  <HowItWorksHero />

                  <CTAHero
                        title="Ready to integrate age verification?"
                        description="It takes just minutes to stay compliant and protect your business — without sacrificing user trust."
                        ctaLabel="Get Started Now"
                        ctaHref="/signup"
                  />

                  <APIHero />

                  <FooterMain />
            </main>
      );
}
