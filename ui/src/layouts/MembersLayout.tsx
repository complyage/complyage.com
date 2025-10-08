//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, ReactNode }      from "react";

import { useLocation }                      from "react-router-dom";
import type { LucideIcon }                  from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import Sidebar                                        from "../components/nav/Sidebar";
import NavMain                                        from "../components/nav/NavMain";
import HealthBanner                                   from "../components/dynamic/HealthBanner";
import RequireMember, { User }                        from "../components/dynamic/RequireMember";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface MembersLayoutProps {
      title?   : string,
      icon?    : LucideIcon,
      overlay? : boolean,
      content? : boolean,
      children : ReactNode
}

//||------------------------------------------------------------------------------------------------||
//|| Content
//||------------------------------------------------------------------------------------------------||

const ContentArea = ({ title, overlay = false, icon: Icon, children }: MembersLayoutProps) => {
      return (
            <>
                  {(title || Icon) && (
                        <h1 className="text-white text-2xl mb-3 border-b border-white/20 p-3 font-bold flex items-center">
                              {Icon && <Icon className="inline-block mr-2 w-6 h-6" />}
                              {title}
                        </h1>
                  )}
                  {children}
            </>
      );
};

//||------------------------------------------------------------------------------------------------||
//|| Members Layout
//||------------------------------------------------------------------------------------------------||

export default function MembersLayout({ title, overlay = false, content = false, icon: Icon, children }: MembersLayoutProps) {
      const location = useLocation();
      overlay = false;
      if (location.pathname.startsWith("/overlay")) {
            overlay = true;
      }
      useEffect(() => {
            if (!overlay) return;
            const original = document.body.style.background;
            document.body.style.background = "transparent";
            return () => { document.body.style.background = original };
      }, []);

      return (
            <RequireMember>
                  {(user: User) => (
                        <>
                              {overlay ? (
                                    <ContentArea title={title} icon={Icon} overlay={overlay}>
                                          {children}
                                    </ContentArea>
                              ) : (
                                    <>
                                          <div className="min-h-screen flex">
                                                <NavMain />

                                                {/* Sidebar fixed width */}
                                                <Sidebar username={user.username || ("Guest" as string)} />

                                                {/* Main Content pushed right */}
                                                <main className="flex flex-col p-5 ml-54 w-full mt-5">
                                                      <section className="p-10 flex-1 bg-gray-700 w-full">
                                                            <ContentArea title={title} icon={Icon} overlay={overlay}>
                                                                  {children}
                                                            </ContentArea>
                                                      </section>
                                                </main>
                                          </div>
                                          <HealthBanner />
                                    </>
                              )}
                        </>
                  )}
            </RequireMember>
      );
}
