import React, {useEffect, useState} from "react";
import {MapContainer, TileLayer, Marker, Popup} from "react-leaflet";
import NavMain from "../../components/nav/NavMain";
import FooterMain from "../../components/footer/FooterMain";
import {ZoneRequirementPlain} from "../../interfaces/zoneRequirements";
import "leaflet/dist/leaflet.css";
import { useCountryFullName } from "../../hooks/useCountry";

export default function EnforcementZones() {

      const [zones, setZones] = useState([]); // State to hold zones data
      useEffect(() => {
            // Fetch zones data from the API
            const fetchZones = async () => {
                try {
                    const response = await fetch("/v1/api/zones");
                    if (!response.ok) {
                        throw new Error("Network response was not ok");
                    }
                    const data = await response.json();
                    console.log(data.data);
                    setZones(data.data || []);
                } catch (error) {
                    console.error("Failed to fetch zones:", error);
                }
            };
            fetchZones();
      }, []);

	return (
		<main className="min-h-screen bg-gray-700 text-base-content">
			<NavMain />

			<section className="py-5 text-center mt-[80px]">
				<h1 className="text-5xl font-extrabold mb-4">
					Enforcement Zones
				</h1>
				<p className="text-lg max-w-2xl mx-auto">
					Interactive map of countries and states with
					active or pending age verification laws.
				</p>
			</section>

			{/* Map */}
			<section className="px-4 mb-12 max-w-7xl mx-auto">
				<MapContainer
					center={[20, 0]}
					zoom={2}
					scrollWheelZoom={false} // Disables zooming with scroll wheel
					doubleClickZoom={false} // Disables zooming with double click
					boxZoom={false} // Disables zooming by dragging a box
					dragging={false} // Disables dragging the map (often goes hand-in-hand with no zoom)
					zoomControl={false} // Hides the +/- zoom controls
					style={{height: "500px", width: "100%"}}
					className="rounded-lg shadow-lg">
					<TileLayer
						url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
						attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
					/>

					{zones.map((zone, idx) => (
                                    zone.coords = [zone.lat, zone.long],
						<Marker
							key={idx}
							position={
								zone.coords as [number, number]
							}
							className="text-xs">
							<Popup>
								<span className="text-xs">
									<strong>
										{zone.name}
									</strong>{" "}
									- {zone.law}
									<br />
									<strong>
										Effective:
									</strong>{" "}
									{zone.effective}
									<br />
									<strong>
										Requirements:
									</strong>
									<ul>
										{" "}
										{zone.requirements.split(',').map(
											(item) => {
												return (
													<li
														className="ml-4 list-disc"
														key={
															item
														}>
														{ZoneRequirementPlain(
															item
														)}
													</li>
												);
											}
										)}
									</ul>
								</span>
							</Popup>
						</Marker>
					))}
				</MapContainer>
			</section>

			{/* List */}
			<section className="px-4 mb-20 max-w-4xl mx-auto">
				<h2 className="text-3xl font-bold mb-6 text-center">
					Zone Details
				</h2>
				<ul className="space-y-4">
					{zones.map((zone, idx) => (                                    
						<li
							key={idx}
							className="border border-base-content/20 p-4 rounded-lg shadow flex"
                                    >
                                          <div className="flex flex-col w-36 mr-4 justify-center items-center">
                                                <div className="rounded-full bg-black/30">
                                                      <img
                                                            src={`https://flagcdn.com/96x72/${zone.country.toLowerCase()}.png`}
                                                            alt={`${zone.country} Flag`}
                                                            className="w-24 h-24 object-contain rounded-full"
                                                      />
                                                </div>
                                          </div>
                                          <div className="pt-2 text-xs leading-5">
                                                <h3 className="text-lg font-bold">
                                                      {zone.state ? `${zone.state}, ${useCountryFullName(zone.country)}` : zone.country}
                                                </h3>
                                                <div>
                                                      <strong>Law:</strong> {zone.law}{" "}
                                                      <br />
                                                      <strong>Effective:</strong>{" "}
                                                      {new Date(zone.effective).toLocaleDateString() } <br />
                                                      <strong>
                                                            Requirements:
                                                      </strong>{" "}
                                                      <ul>
                                                            {" "}
                                                            {zone.requirements.split(',').map(
                                                                  (item) => {
                                                                        return (
                                                                              <li
                                                                                    className="ml-4 list-disc"
                                                                                    key={
                                                                                          item
                                                                                    }>
                                                                                    {ZoneRequirementPlain(
                                                                                          item
                                                                                    )}
                                                                              </li>
                                                                        );
                                                                  }
                                                            )}
                                                      </ul>
                                                      <strong>Penalties:</strong> { zone.penalties ? zone.penalties : "N/A" }
                                                </div>
                                          </div>
						</li>
					))}
				</ul>
			</section>

			<FooterMain />
		</main>
	);
}
