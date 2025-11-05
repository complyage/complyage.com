/*||------------------------------------------------------------------------------------------------||
//|| Gate Documentation Main Page
//|| src/routes/docs/gate.tsx
//||------------------------------------------------------------------------------------------------||*/

/*||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||*/

import React, { useState } from "react";
import NavDocs from "../../components/nav/NavDocs";

import { Eye } from "lucide-react";

// Section Components
import GateOverview         from "../../sections/gate/GateOverview";
import GateAutomated        from "../../sections/gate/GateAutomated";
import GateManual           from "../../sections/gate/GateManual";
import GateWebhook          from "../../sections/gate/GateWebhook";

/*||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||*/

export default function Gate() {

      /*||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||*/
      const [activeSection, setActiveSection] = useState("overview");

      const sections = [
            { key: "overview",         label: "1. Overview" },
            { key: "automated",        label: "2. Automated" },
            { key: "manual",           label: "3. Manual" },
            { key: "webhook",          label: "4. Webhook - Optional" },
      ];

      /*||------------------------------------------------------------------------------------------------||
      //|| Render Section Content
      //||------------------------------------------------------------------------------------------------||*/
      const renderSection = () => {
            switch (activeSection) {
                  case "overview":         return <GateOverview />;
                  case "automated":        return <GateAutomated />;
                  case "manual":           return <GateManual />;
                  case "webhook":          return <GateWebhook />;
                  default:                 return <GateOverview />;
            }
      };

      /*||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||*/
      return (
            <main className="min-h-screen flex flex-col bg-base-100 text-base-content">
                  <NavDocs />

                  <div className="flex flex-1 pt-[60px]">
                        {/* Sidebar */}
                        <aside className="w-64 bg-base-200 border-r border-base-300 p-6 hidden md:block">
                              <h3 className="text-xs py-4 px-1 font-bold text-orange-500 mb-3"><Eye size="24" className="inline mr-2" />Age Gate Documentation</h3>
                              <nav className="flex flex-col gap-1">
                                    {sections.map(({ key, label }) => (
                                          <button
                                                key={key}
                                                onClick={() => setActiveSection(key)}
                                                className={`text-left px-3 py-2 rounded-md text-sm transition
                                                      ${
                                                            activeSection === key
                                                                  ? "bg-base-300 text-orange-400 font-semibold"
                                                                  : "hover:bg-base-300"
                                                      }`}
                                          >
                                                {label}
                                          </button>
                                    ))}
                              </nav>
                        </aside>

                        {/* Main Content Area */}
                        <section className="flex-1 p-8 overflow-y-auto">{renderSection()}</section>
                  </div>
            </main>
      );
}
