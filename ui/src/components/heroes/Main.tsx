//||------------------------------------------------------------------------------------------------||
//|| MainHero.tsx
//|| src/components/heroes/MainHero.tsx
//||------------------------------------------------------------------------------------------------||

import React from "react";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface MainHeroProps {
      image?       : string;
      title        : string;
      description  : string;
      ctaLabel?    : string;
      ctaHref?     : string;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function MainHero({ image, title, description, ctaLabel, ctaHref }: MainHeroProps) {
      return (
            <section
                  className="relative min-h-[60vh] flex items-center justify-center text-center px-6"
                  style={{
                        backgroundImage: image ? `url(${image})` : undefined,
                        backgroundSize: "cover",
                        backgroundPosition: "center",
                  }}
            >
                  {/* Overlay for readability */}
                  <div className="absolute inset-0 bg-black/40" />

                  <div className="relative max-w-5xl mt-[60px] text-white">
                        <h1 className="text-6xl md:text-7xl font-extrabold mb-6 leading-tight drop-shadow-lg">
                              {title}
                        </h1>
                        <p className="text-xl md:text-2xl leading-relaxed max-w-4xl mx-auto opacity-95 drop-shadow font-bold text-shadow-2xs">
                              {description}
                        </p>
                        {ctaLabel && ctaHref && (
                              <div className="mt-8">
                                    <a 
                                          href={ctaHref}
                                          className="btn btn-secondary btn-lg px-10 py-4 font-bold shadow-md"
                                    >
                                          {ctaLabel}
                                    </a>
                              </div>
                        )}
                  </div>
            </section>
      );
}
