import React from "react";
import {Link} from "react-router-dom";
import {ShieldCheck, EyeOff, FileText, Lock, Ban, UserX} from "lucide-react";
import NavMain          from "../../components/nav/NavMain";
import FooterMain       from "../../components/footer/FooterMain";
import NewsStream       from "../../components/dynamic/NewsStream";

export default function Home() {
	return (
		<main className="min-h-screen bg-base-100 text-base-content">
			{/* Navbar */}
			<NavMain />

			{/* Hero Section with Video Background */}
			<section className="relative min-h-[100vh] flex items-center justify-center overflow-hidden">
				{/* Video background */}
				<img
					className="absolute top-0 left-0 w-full h-full object-cover"
					src="/vid/take.webp"
                        />

				{/* Overlay */}
				<div className="absolute inset-0 bg-black/60"></div>

				{/* Hero Content */}
				<div className="relative z-10 text-white px-4 text-center w-full mt-[100px]">
					<div className="max-w-6xl mx-auto">
						<img
							src="/public/img/logo-white.png"
							alt="ComplyAge Logo"
							className="mx-auto mb-4 w-24 h-24"
						/>
						<h1 className="text-5xl font-extrabold mb-8">
							Take Back Your Privacy. Keep Your
							Freedom.
						</h1>

						{/* 2-Row Grid: 6 Trust Items + 3 Buttons */}
						<div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-5">
							{/* 6 Features */}
							<div className="flex flex-col items-center text-center bg-gray-950/50 p-10 rounded-2xl hover:bg-black cursor-pointer">
								<ShieldCheck className="w-14 h-10 text-success mb-4" />
								<h3 className="text-xl font-bold mb-2">
									Open Source
								</h3>
								<p className="text-white text-opacity-80">
									Trust our transparent
									code.
								</p>
							</div>

							<div className="flex flex-col items-center text-center bg-gray-950/50 p-10 rounded-2xl hover:bg-black cursor-pointer">
								<FileText className="w-14 h-10 text-success mb-4" />
								<h3 className="text-xl font-bold mb-2">
									Clear Terms
								</h3>
								<p className="text-white text-opacity-80">
									No hidden clauses, ever.
								</p>
							</div>

							<div className="flex flex-col items-center text-center bg-gray-950/50 p-10 rounded-2xl hover:bg-black cursor-pointer">
								<EyeOff className="w-14 h-10 text-success mb-4" />
								<h3 className="text-xl font-bold mb-2">
									No Tracking
								</h3>
								<p className="text-white text-opacity-80">
									We never track you.
								</p>
							</div>

							<div className="flex flex-col items-center text-center bg-gray-950/50 p-10 rounded-2xl hover:bg-black cursor-pointer">
								<Lock className="w-14 h-10 text-success mb-4" />
								<h3 className="text-xl font-bold mb-2">
									Your Data
								</h3>
								<p className="text-white text-opacity-80">
									Encrypted. Not for sale.
								</p>
							</div>

							<div className="flex flex-col items-center text-center bg-gray-950/50 p-10 rounded-2xl hover:bg-black cursor-pointer">
								<Ban className="w-14 h-10 text-success mb-4" />
								<h3 className="text-xl font-bold mb-2">
									No Overreach
								</h3>
								<p className="text-white text-opacity-80">
									We push back on invasive
									laws.
								</p>
							</div>

							<div className="flex flex-col items-center text-center bg-gray-950/50 p-10 rounded-2xl hover:bg-black cursor-pointer">
								<UserX className="w-14 h-10 text-success mb-4" />
								<h3 className="text-xl font-bold mb-2">
									Stay Anonymous
								</h3>
								<p className="text-white text-opacity-80">
									We don’t want to know you.
								</p>
							</div>

							{/* 3 Buttons in their own row */}
							<div className="flex flex-col items-center">
                                                <button className="btn btn-primary text-3xl p-10 w-full">
									Get Started
								</button>
							</div>

							<div className="flex flex-col items-center">
								<button className="btn btn-secondary text-3xl p-10 w-full">
									View Our Code
								</button>
							</div>

							<div className="flex flex-col items-center">
                                                <button className="btn btn-tertiary text-3xl p-10 w-full">
									View Our Terms
								</button>
							</div>
						</div>
					</div>
				</div>
			</section>


			{/* Call To Action */}
			<section className="py-16 bg-primary text-primary-content text-center">
				<h2 className="text-3xl md:text-4xl font-bold mb-4">
					Ready to protect your platform?
				</h2>
				<p className="mb-6">
					Start verifying ages in under 5 minutes.
				</p>
				<button className="btn btn-secondary text-2xl p-5 rounded-lg">
					Sign Up Now
				</button>
			</section>

			<section className="py-12 bg-black">
                        <NewsStream />
                  </section>

			{/* Carousel */}
			<section className="py-12">
				<div className="max-w-4xl mx-auto">
					<div className="carousel w-full rounded-box">
						<div
							id="slide1"
							className="carousel-item relative w-full">
							<img
								src="https://picsum.photos/id/237/800/400"
								className="w-full object-cover"
								alt="Security"
							/>
							<div className="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
								<a
									href="#slide3"
									className="btn btn-circle">
									❮
								</a>
								<a
									href="#slide2"
									className="btn btn-circle">
									❯
								</a>
							</div>
						</div>
						<div
							id="slide2"
							className="carousel-item relative w-full">
							<img
								src="https://picsum.photos/id/238/800/400"
								className="w-full object-cover"
								alt="Verification"
							/>
							<div className="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
								<a
									href="#slide1"
									className="btn btn-circle">
									❮
								</a>
								<a
									href="#slide3"
									className="btn btn-circle">
									❯
								</a>
							</div>
						</div>
						<div
							id="slide3"
							className="carousel-item relative w-full">
							<img
								src="https://picsum.photos/id/239/800/400"
								className="w-full object-cover"
								alt="Compliance"
							/>
							<div className="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
								<a
									href="#slide2"
									className="btn btn-circle">
									❮
								</a>
								<a
									href="#slide1"
									className="btn btn-circle">
									❯
								</a>
							</div>
						</div>
					</div>
				</div>
			</section>

			{/* Features Section */}
			<section className="py-16 px-4 max-w-6xl mx-auto">
				<h2 className="text-4xl font-bold text-center mb-12">
					Why Choose ComplyAge?
				</h2>
				<div className="grid md:grid-cols-3 gap-8">
					<div className="card bg-base-100 shadow-lg">
						<div className="card-body">
							<h3 className="card-title">
								Easy Integration
							</h3>
							<p>
								Plug & play API and SDKs for any
								stack. Get up and running in
								minutes.
							</p>
						</div>
					</div>
					<div className="card bg-base-100 shadow-lg">
						<div className="card-body">
							<h3 className="card-title">
								Global Compliance
							</h3>
							<p>
								Stay ahead of local and
								international age restriction
								laws automatically.
							</p>
						</div>
					</div>
					<div className="card bg-base-100 shadow-lg">
						<div className="card-body">
							<h3 className="card-title">
								Rock-Solid Security
							</h3>
							<p>
								All data encrypted,
								privacy-first. We never store
								sensitive user IDs longer than
								needed.
							</p>
						</div>
					</div>
				</div>
			</section>

                  <FooterMain />
		</main>
	);
}
