import React from "react";
import { useNavigate } from "react-router-dom";
import NavMain from "../../components/nav/NavMain";
import FooterMain from "../../components/footer/FooterMain";

export default function ExitVerification() {
  const navigate = useNavigate();

  const handleGoBack = () => {
    window.history.back();
  };

  const handleLeave = () => {
    // Redirect to external site or homepage
    window.location.href = "/";
  };

  return (
    <main className="min-h-screen flex flex-col bg-base-100 text-base-content">
      <NavMain />

      {/* Hero Section */}
      <section className="flex-1 flex flex-col items-center justify-center bg-error text-error-content px-4 text-center pt-10 mt-[50px]">
        <div className="max-w-lg">
          <h1 className="text-5xl font-extrabold mb-4">Access Denied</h1>
          <p className="text-lg mb-8">
            We’re sorry, but you’re not eligible to view this content.  
            If you believe this is an error, please contact support.
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <button
              onClick={handleGoBack}
              className="btn btn-outline btn-neutral px-6 py-3"
            >
              Go Back
            </button>
            <button
              onClick={handleLeave}
              className="btn btn-primary px-6 py-3"
            >
              Leave Site
            </button>
          </div>
        </div>
      </section>

      <FooterMain />
    </main>
  );
}
