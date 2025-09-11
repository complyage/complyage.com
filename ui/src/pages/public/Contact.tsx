/*||------------------------------------------------------------------------------------------------||
//|| Contact Page
//|| src/routes/contact.tsx
//||------------------------------------------------------------------------------------------------||*/

import React, { useState } from "react";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import NavMain      from "../../components/nav/NavMain";
import FooterMain   from "../../components/footer/FooterMain";
import InlineAlert  from "../../components/base/InlineAlert";
import Turnstile    from "../../components/base/Turnstile";
import MainHero     from "../../components/heroes/Main";

//||------------------------------------------------------------------------------------------------||
//|| Page
//||------------------------------------------------------------------------------------------------||

export default function ContactPage() {
      const [name, setName]       = useState("");
      const [email, setEmail]     = useState("");
      const [message, setMessage] = useState("");
      const [loading, setLoading] = useState(false);
      const [success, setSuccess] = useState<string | null>(null);      
      const [captcha, setCaptcha] = useState<string | null>(null);
      const [error, setError]     = useState<string | null>(null);

      //||------------------------------------------------------------------------------------------------||
      //|| Handle Submit
      //||------------------------------------------------------------------------------------------------||

      const handleSubmit = async (e: React.FormEvent) => {
            e.preventDefault();
            setSuccess(null);
            setError(null);

            if (!name.trim() || !email.trim() || !message.trim()) {
                  setError("Please fill out all fields.");
                  return;
            }
            setLoading(true);

            try {
                  const res = await fetch("/public/contact", {
                        method: "POST",
                        headers: {"Content-Type": "application/json"},
                        body: JSON.stringify({ name, email, message, turnstileToken: captcha }),
                  });

                  const payload = await res.json();
                  if (res.ok && payload.success) {
                        setSuccess(payload.message || "Message sent.");
                        setName("");
                        setEmail("");
                        setMessage("");
                  } else {
                        setError(payload?.message || "Failed to send message.");
                  }
            } catch {
                  setError("Network error. Please try again.");
            }
            setLoading(false);
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Success
      //||------------------------------------------------------------------------------------------------||

      if (success) {
            return (
                  <main className="min-h-screen flex flex-col bg-base-100 text-base-content">
                        <NavMain />
                        <section className="flex flex-col items-center justify-center min-h-[60vh] bg-primary text-primary-content px-4 text-center">
                              <div className="max-w-2xl mt-[50px]">
                                    <h1 className="text-5xl font-extrabold mb-4">Thank You!</h1>
                                    <p className="text-lg">{success}</p>
                              </div>
                        </section>
                        <FooterMain />
                  </main>
            );
      }

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <main className="min-h-screen flex flex-col bg-base-100 text-base-content">
                  <NavMain />

                  {/* Hero */}
                  <MainHero
                        image="/img/hero/contact.webp"
                        title="Contact Us"
                        description="Questions? Feedback? Reach out and we’ll get back to you as soon as possible."
                  />

                  {/* 2 Column Section */}
                  <section className="py-16 px-4 max-w-6xl mx-auto w-full grid grid-cols-1 md:grid-cols-2 gap-12">
                        
                        {/* Left: Response Info */}
                        <div className="bg-base-200 rounded-lg shadow-lg p-8">
                              <h2 className="text-3xl font-bold mb-4">When will we respond?</h2>
                              <p className="mb-4">
                                    We try to respond to all inquiries within <b>24–48 hours</b> during weekdays.
                                    Urgent compliance or support issues are prioritized immediately.
                              </p>
                              <ul className="list-disc list-inside space-y-2">
                                    <li><b>Support hours:</b> Monday – Friday, 9am – 6pm (PST)</li>
                                    <li><b>Average reply time:</b> Under 1 business day</li>
                                    <li><b>Priority cases:</b> Verified partners and compliance emergencies</li>
                              </ul>
                              <p className="mt-4 text-sm opacity-80">
                                    Please include as much detail as possible in your message so we can assist quickly.
                              </p>
                        </div>

                        {/* Right: Form */}
                        <form className="bg-base-200 rounded-lg shadow-lg p-8 space-y-5" onSubmit={handleSubmit}>
                              <div>
                                    <label className="block text-lg font-bold mb-1">Name</label>
                                    <input
                                          type="text"
                                          className="input input-bordered w-full"
                                          value={name}
                                          onChange={(e) => setName(e.target.value)}
                                          required
                                    />
                              </div>
                              <div>
                                    <label className="block text-lg font-bold mb-1">Email</label>
                                    <input
                                          type="email"
                                          className="input input-bordered w-full"
                                          value={email}
                                          onChange={(e) => setEmail(e.target.value)}
                                          required
                                    />
                              </div>
                              <div>
                                    <label className="block text-lg font-bold mb-1">Message</label>
                                    <textarea
                                          className="textarea textarea-bordered w-full h-32"
                                          value={message}
                                          onChange={(e) => setMessage(e.target.value)}
                                          required
                                    />
                              </div>
                              <Turnstile siteKey={ import.meta.env.VITE_TURNSTILE_PUBLIC } onVerify={ (token) => setCaptcha(token) } />
                              {error && <InlineAlert isError message={error} />}
                              <div className="flex justify-end">
                                    <button 
                                          type="submit" 
                                          className={`btn btn-secondary btn-lg px-10${loading ? " loading" : ""}`} 
                                          disabled={ captcha === null || loading }
                                    >
                                          Send
                                    </button>
                              </div>
                        </form>
                  </section>

                  <FooterMain />
            </main>
      );
}
