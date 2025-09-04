/*||------------------------------------------------------------------------------------------------||
//|| Contact Page
//|| src/routes/contact.tsx
//||------------------------------------------------------------------------------------------------||*/

/*||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||*/

import React, { useState }        from "react";
import NavMain                    from "../../components/nav/NavMain";
import FooterMain                 from "../../components/footer/FooterMain";
import InlineAlert                from "../../components/base/InlineAlert";

/*||------------------------------------------------------------------------------------------------||
//|| Page
//||------------------------------------------------------------------------------------------------||*/

export default function ContactPage() {

   // State
   const [name, setName]             = useState("");
   const [email, setEmail]           = useState("");
   const [message, setMessage]       = useState("");
   const [loading, setLoading]       = useState(false);
   const [success, setSuccess]       = useState<string | null>(null);
   const [error, setError]           = useState<string | null>(null);

   // Submit handler
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
         const res = await fetch("/v1/api/contact", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ name, email, message })
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
      } catch (err) {
         setError("Network error. Please try again.");
      }
      setLoading(false);
   };

   // Render
   return (
      <main className="min-h-screen flex flex-col bg-base-100 text-base-content">
         <NavMain />

         {/* Hero */}
         <section className="flex flex-col items-center justify-center min-h-[40vh] bg-primary text-primary-content px-4 text-center">
            <div className="max-w-2xl mt-[50px]">
               <h1 className="text-5xl font-extrabold mb-4">Contact Us</h1>
               <p className="text-lg">
                  Questions? Feedback? Reach out and we’ll get back to you as soon as possible.
               </p>
            </div>
         </section>

         {/* Form Section */}
         <section className="py-16 px-4 max-w-2xl mx-auto w-full">
            <form className="bg-base-200 rounded-lg shadow-lg p-8 space-y-5" onSubmit={handleSubmit}>
               <div>
                  <label className="block text-lg font-bold mb-1">Name</label>
                  <input
                     type="text"
                     className="input input-bordered w-full"
                     value={name}
                     onChange={e => setName(e.target.value)}
                     required
                  />
               </div>
               <div>
                  <label className="block text-lg font-bold mb-1">Email</label>
                  <input
                     type="email"
                     className="input input-bordered w-full"
                     value={email}
                     onChange={e => setEmail(e.target.value)}
                     required
                  />
               </div>
               <div>
                  <label className="block text-lg font-bold mb-1">Message</label>
                  <textarea
                     className="textarea textarea-bordered w-full h-32"
                     value={message}
                     onChange={e => setMessage(e.target.value)}
                     required
                  />
               </div>
               {error && <InlineAlert isError message={error} />}
               {success && <InlineAlert isSuccess message={success} />}
               <div className="flex justify-end">
                  <button
                     type="submit"
                     className={`btn btn-primary px-10 text-lg${loading ? " loading" : ""}`}
                     disabled={loading}
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
