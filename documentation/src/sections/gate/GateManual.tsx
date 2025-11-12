/*||------------------------------------------------------------------------------------------------||
//|| Gate Manual Section
//|| src/sections/gate/GateManual.tsx
//||------------------------------------------------------------------------------------------------||*/

import React from "react";
import PageHeader from "../../components/base/PageHeader";

export default function GateManual() {
      return (
            <div className="prose prose-invert max-w-none space-y-10 leading-relaxed text-sm">
                  <PageHeader title="3. Age Gate - Manual Integration" />

                  {/* Section: Manual */}
                  <section className="space-y-4 border-t border-base-300">
                        <p className="">
                              The <strong>Manual</strong> method gives you full control over when and how users verify.
                              Use the API endpoints directly in your app or backend service.
                        </p>

                        <div className="bg-base-200 rounded-xl overflow-hidden">
                              <pre className="text-sm  p-4 overflow-x-auto">
{`fetch("https://api.complyage.com/v1/gate/start", {
   method: "POST",
   headers: { "Content-Type": "application/json" },
   body: JSON.stringify({ site: "YOUR_SITE_KEY" })
});`}
                              </pre>
                        </div>

                        <p className="">
                              Perfect for single-page apps, membership systems, or custom flows.
                              You decide how and when to display the verification prompt — with or without the default overlay.
                        </p>
                  </section>

                  {/* Manual Handler: /v1/manual */}
                  <section className="space-y-3">
                        <h2 className="text-xl text-gray-400 font-bold mb-1">Endpoint</h2>
                        <p><code>POST /v1/manual</code></p>
                        <p>
                              The handler in <code>gate/handlers/manual.go</code> accepts a JSON body and returns a
                              keeper (session) record. If <code>session_id</code> is known, the existing record is
                              returned; otherwise a new one is created.
                        </p>
                  </section>

                  {/* Request */}
                  <section className="space-y-3">
                        <h2 className="text-xl text-gray-400 font-bold mb-1">Request Body</h2>
                        <ul className="list-disc list-inside space-y-1">
                              <li><code>ip_address</code> string — user IP</li>
                              <li><code>session_id</code> string — existing keeper/session ID (optional)</li>
                              <li><code>client_id</code> string — your ComplyAge client ID</li>
                        </ul>

                        <div className="bg-base-200 rounded-xl overflow-hidden">
                              <pre className="text-sm p-4 overflow-x-auto">{`
POST https://gate.complyage.com/v1/manual
Content-Type: application/json

{
  "ip_address": "203.0.113.24",
  "session_id": "KEEPER-ABC-123",
  "client_id": "CLID-XXX-YOUR-CLIENT-ID"
}`}</pre>
                        </div>

                        <div className="bg-base-200 rounded-xl overflow-hidden">
                              <pre className="text-sm p-4 overflow-x-auto">{`// fetch() example
await fetch("https://gate.complyage.com/v1/manual", {
  method: "POST",
  headers: { "Content-Type": "application/json" },
  credentials: "include",
  body: JSON.stringify({
    ip_address: userIp,
    session_id: existingKeeperId,
    client_id: "CLID-XXX-YOUR-CLIENT-ID"
  })
});`}</pre>
                        </div>
                  </section>

                  {/* Response */}
                  <section className="space-y-3">
                        <h2 className="text-xl text-gray-400 font-bold mb-1">Response Body</h2>
                        <p>Returns a keeper record (<code>base/keeper/struct.go</code>):</p>
                        <ul className="list-disc list-inside space-y-1">
                              <li><code>keeperId</code> string — session identifier</li>
                              <li><code>enforced</code> boolean — whether enforcement applies</li>
                              <li><code>verified</code> boolean — whether user is verified</li>
                              <li><code>age</code> number — verified age (0 if unknown)</li>
                              <li><code>userId</code> number — associated user ID, if any</li>
                              <li><code>ipAddress</code> string — recorded IP address</li>
                              <li><code>clientId</code> string — your site client ID</li>
                              <li><code>status</code> string — e.g., "INIT"</li>
                              <li><code>returnUrl</code> string (optional)</li>
                        </ul>

                        <div className="bg-base-200 rounded-xl overflow-hidden">
                              <pre className="text-sm p-4 overflow-x-auto">{`{
  "keeperId": "KEEPER-ABC-123",
  "enforced": true,
  "verified": false,
  "age": 0,
  "userId": 0,
  "ipAddress": "203.0.113.24",
  "clientId": "CLID-XXX-YOUR-CLIENT-ID",
  "status": "INIT",
  "returnUrl": ""
}`}</pre>
                        </div>
                  </section>

            </div>
      );
}
