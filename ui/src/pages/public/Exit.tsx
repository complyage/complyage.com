
//||------------------------------------------------------------------------------------------------||
//|| Hooks
//||------------------------------------------------------------------------------------------------||

import { useOverlayNavigate }                        from "../../hooks/useOverlay";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import NavMain                                        from "../../components/nav/NavMain";
import FooterMain                                     from "../../components/footer/FooterMain";

//||------------------------------------------------------------------------------------------------||
//|| Exit
//||------------------------------------------------------------------------------------------------||

export default function ExitVerification() {

      //||------------------------------------------------------------------------------------------------||
      //|| Navigate
      //||------------------------------------------------------------------------------------------------||

      const navigate                     = useOverlayNavigate();  

      //||------------------------------------------------------------------------------------------------||
      //|| Back
      //||------------------------------------------------------------------------------------------------||

      const handleGoBack = () => {
		window.history.back();
	};

      //||------------------------------------------------------------------------------------------------||
      //|| Leave
      //||------------------------------------------------------------------------------------------------||

	const handleLeave = () => {
		navigate("/");
	};

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

	return (
		<main className="min-h-screen flex flex-col bg-base-100 text-base-content">
			<NavMain />

			{/* Hero Section */}
			<section className="flex-1 flex flex-col items-center justify-center bg-error text-error-content px-4 text-center pt-10 mt-[50px]">
				<div className="max-w-lg">
					<h1 className="text-5xl font-extrabold mb-4">Access Denied</h1>
					<p className="text-lg mb-8">
						We’re sorry, but you’re not eligible to view this content. If you believe this is an error, please contact support.
					</p>
					<div className="flex flex-col sm:flex-row gap-4 justify-center">
						<button onClick={handleGoBack} className="btn btn-outline btn-neutral px-6 py-3">
							Go Back
						</button>
						<button onClick={handleLeave} className="btn btn-primary px-6 py-3">
							Leave Site
						</button>
					</div>
				</div>
			</section>

			<FooterMain />
		</main>
	);
}
