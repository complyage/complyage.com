import React, {useEffect, useState} from "react";

type NewsArticle = {
	title: string;
	link: string;
	pubDate: string;
	source?: string;
};

const ITEMS_PER_PAGE = 5;

export default function NewsStream() {
	const [articles, setArticles] = useState<NewsArticle[]>([]);
	const [visibleCount, setVisibleCount] = useState(ITEMS_PER_PAGE);
	const [loading, setLoading] = useState(false);

	useEffect(() => {
		const fetchNews = async () => {
			const res = await fetch("/v1/api/news");
			const data = await res.json();
			setArticles(data.articles || []);
		};
		fetchNews();
	}, []);

	const loadMore = () => {
		setLoading(true);
		setTimeout(() => {
			setVisibleCount((prev) => prev + ITEMS_PER_PAGE);
			setLoading(false);
		}, 300); // Small delay for UX
	};

	return (
		<section className="py-16 bg-black">
			<div className="max-w-[80vw] mx-auto px-4 grid md:grid-cols-2 gap-12 text-white">
				{/* Left column: description */}
                        <div className="flex flex-col justify-center">
   <h2 className="text-4xl font-bold mb-4 text-white">
      Why We Do This
   </h2>
   <p className="text-lg mb-4 leading-loose">
      <b>The internet was built on the promise of freedom and anonymity</b> 
      &nbsp; a place where everyone could speak, share, and connect without constant surveillance.
      Today, new age verification laws and intrusive requirements threaten to chip away at that freedom,
      forcing people to trade privacy for access. We track the latest news to expose these overreaches,
      keep you informed, and help you protect your right to stay anonymous online.
      No matter your beliefs, no matter who’s in charge. <b>Freedom online belongs to everyone.</b>
   </p>
</div>


				{/* Right column: scrollable feed with lazy load */}
				<div className="h-[50vh] overflow-y-auto pr-4">
					<ul className="space-y-4">
						{articles
							.slice(0, visibleCount)
							.map((article, idx) => (
								<li
									key={idx}
									className="bg-gray-700 border p-4 rounded-lg shadow hover:shadow-md transition-shadow">
									<a
										href={article.link}
										target="_blank"
										rel="noreferrer"
										className="text-xl font-semibold hover:underline">
										{article.title}
									</a>
									<p className="text-sm text-white mt-1">
										{article.source ||
											"Source"}{" "}
										—{" "}
										{new Date(
											article.pubDate
										).toLocaleDateString()}
									</p>
								</li>
							))}
					</ul>

					{visibleCount < articles.length && (
						<div className="mt-6 text-center">
							<button
								onClick={loadMore}
								disabled={loading}
								className="btn btn-primary">
								{loading
									? "Loading..."
									: "Load More"}
							</button>
						</div>
					)}
				</div>
			</div>
		</section>
	);
}
