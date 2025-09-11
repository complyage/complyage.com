//||------------------------------------------------------------------------------------------------||
//|| CTAHero.tsx
//|| src/components/heroes/CTAHero.tsx
//||------------------------------------------------------------------------------------------------||

import React from "react";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface CTAHeroProps {
      title       : string;
      description : string;
      ctaLabel    : string;
      ctaHref     : string;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function CTAHero({ title, description, ctaLabel, ctaHref }: CTAHeroProps) {
      return (
            <section className="py-20 text-center bg-primary text-primary-content">
                  <div className="max-w-3xl mx-auto">
                        <h2 className="text-3xl md:text-4xl font-bold mb-6">{title}</h2>
                        <p className="mb-8 text-lg leading-relaxed">{description}</p>
                        <a href={ctaHref} className="btn btn-secondary text-xl px-10 py-4">
                              {ctaLabel}
                        </a>
                  </div>
            </section>
      );
}
