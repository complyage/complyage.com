import React, {useEffect, useState} from "react";
import {MapContainer, TileLayer, Marker, Popup} from "react-leaflet";
import NavMain from "../../components/nav/NavMain";
import FooterMain from "../../components/footer/FooterMain";
import {ZoneRequirementPlain} from "../../interfaces/zoneRequirements";
import "leaflet/dist/leaflet.css";
import {useCountryFullName} from "../../hooks/useCountry";

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
                                                Enforcement Zones/Laws
                                          </h1>
                                          <p className="text-lg max-w-2xl mx-auto">
                                                Interactive map of countries and states with
                                                active or pending age verification laws.
                                          </p>
                                    </section>

                                    {/* Map */}
                                    <section className="px-4 mb-12 max-w-7xl mx-auto">
                                    <MapContainer
                  center={[20, 0]} zoom={2}
                  scrollWheelZoom={false} doubleClickZoom={false}
                  boxZoom={false} dragging={false} zoomControl={false}
                  style={{ height: "500px", width: "100%" }}
                  className="rounded-lg shadow-lg"
                  >
                  <TileLayer
                  url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                  attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
                  />

                  {zones.map((zone) => (
                  // Ensure zone.coords is valid before rendering Marker
                  zone.coords && zone.coords.length === 2 && (
                        <Marker key={zone.id} position={zone.coords} className="text-xs">
                        <Popup>
                        <span className="text-xs">
                              <strong>{zone.state || zone.country || "Unknown Zone"}</strong> - {zone.law || "No Law Specified"}
                              <br />
                              <strong>Effective:</strong> {zone.effective || "N/A"}
                              <br />
                              <strong>Requirements:</strong>
                              <ul>
                              {zone.requirements && zone.requirements.split(",").map((item, idx) => (
                              <li className="ml-4 list-disc" key={idx}>
                                    {ZoneRequirementPlain(item.trim())}
                              </li>
                              ))}
                              {!zone.requirements && <li>No specific requirements.</li>}
                              </ul>
                              {zone.penalties && (
                              <>
                              <strong>Penalties:</strong> {zone.penalties}
                              </>
                              )}
                        </span>
                        </Popup>
                        </Marker>
                  )
                  ))}
                  </MapContainer>

			</section>

			{/* List */}
                  <section className="px-4 mb-20 max-w-7xl mx-auto"> {/* Adjusted max-w to match map section for consistency */}
                        <h2 className="text-3xl font-bold mb-6 text-center">
                              Zone Details
                        </h2>
                        {/* This is the new grid container */}
                        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                              {zones.map((zone, idx) => (
                                    <div key={zone.id || idx} className="card bg-base-100 shadow-sm"> {/* Removed w-[50%] from here */}
                                    <figure className="bg-cover bg-center h-32" style={{ backgroundImage: `url(https://flagcdn.com/256x192/${zone.country?.toLowerCase()}.png)` }}>
                                          <div className="bg-black/70 w-full h-full flex justify-center items-center font-bold text-2xl text-white"> {/* Added text-white for visibility */}
                                                {zone.state ? `${zone.state}, ${useCountryFullName(zone.country)}` : zone.country}
                                          </div>
                                    </figure>
                                    <div className="card-body">
                                          <h2 className="card-title border-b border-gray-500 p-2">
                                                <b className="text-gray-200">{zone.law}</b>
                                                { new Date(zone.effective || '') > new Date(new Date().setMonth(new Date().getMonth() - 6)) && (<div className="ml-auto badge badge-secondary justify-end">NEW</div>) }
                                          </h2>
                                          <p>
                                                <strong className="text-gray-400 mr-2 inline-block">Effective:</strong><b>{zone.effective ? new Date(zone.effective).toLocaleDateString() : 'N/A'}</b> <br />
                                                <strong className="text-gray-400 mr-2 inline-block">Requirements:</strong>{" "}
                                                <ul className="leading-5">
                                                {zone.requirements && zone.requirements.split(',').map(
                                                      (item, reqIdx) => {
                                                            return (
                                                            <li
                                                                  className="ml-4 list-disc text-xs leading-5"
                                                                  key={reqIdx}> {/* Use reqIdx for key within inner map */}
                                                                  {ZoneRequirementPlain(item.trim())}
                                                            </li>
                                                            );
                                                      }
                                                )}
                                                {!zone.requirements && <li>No specific requirements.</li>}
                                                </ul>
                                          </p>
                                          {zone.penalties && ( <div><b className="text-gray-400 mr-2">Penalties:</b> {zone.penalties} </div> )}
                                    </div>
                                    </div>
                              ))}
                        </div>
                        </section>

			<FooterMain />
		</main>
	);
}
