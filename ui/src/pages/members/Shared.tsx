//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { ModelShared } from "../../interfaces/model/model.shared";

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

import MembersLayout from "../../layouts/MembersLayout";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Shared() {
    const navigate = useNavigate();
    const [searchQuery, setSearchQuery] = useState("");
    const [sharedItems, setSharedItems] = useState<ModelShared[]>([]);
    const [loading, setLoading] = useState(true);

    //||------------------------------------------------------------------------------------------------||
    //|| Fetch shared data from API
    //||------------------------------------------------------------------------------------------------||

    const fetchSharedItems = async () => {
        setLoading(true);
        try {
            const res = await fetch("/user/shared", { credentials: "include" });
            const json = await res.json();

            if (json.success && Array.isArray(json.data?.shared)) {
                setSharedItems(json.data.shared);
            } else {
                console.warn("Unexpected API response", json);
                setSharedItems([]);
            }
        } catch (err) {
            console.error("❌ Failed to fetch shared data:", err);
            setSharedItems([]);
        } finally {
            setLoading(false);
        }
    };

    //||------------------------------------------------------------------------------------------------||
    //|| Initial Load
    //||------------------------------------------------------------------------------------------------||

    useEffect(() => {
        fetchSharedItems();
    }, []);

    //||------------------------------------------------------------------------------------------------||
    //|| Group by Site
    //||------------------------------------------------------------------------------------------------||

    const groupedBySite = sharedItems
        .filter((item) => item.site_name.toLowerCase().includes(searchQuery.toLowerCase()))
        .reduce((groups: Record<string, ModelShared[]>, item) => {
            const key = `${item.site_name}||${item.site_url}`;
            if (!groups[key]) groups[key] = [];
            groups[key].push(item);
            return groups;
        }, {});

    const siteKeys = Object.keys(groupedBySite).sort((a, b) => a.localeCompare(b));

    //||------------------------------------------------------------------------------------------------||
    //|| JSX
    //||------------------------------------------------------------------------------------------------||

    return (
        <MembersLayout title="Currently Shared Data">
            <>
                <input
                    type="text"
                    placeholder="Search sites..."
                    className="input input-bordered w-full max-w-md mb-6"
                    value={searchQuery}
                    onChange={(e) => setSearchQuery(e.target.value)}
                />

                {loading ? (
                    <p className="text-center text-base-content/70">Loading shared data...</p>
                ) : siteKeys.length === 0 ? (
                    <p className="text-sm text-base-content/70 text-center bg-gray-800 p-5 text-yellow-500">No shared verifications found.</p>
                ) : (
                    siteKeys.map((siteKey) => {
                        const [siteName, siteURL] = siteKey.split("||");
                        const verifications = groupedBySite[siteKey];

                        return (
                            <div key={siteKey} className="bg-base-100 shadow rounded-lg mb-8">
                                <div className="p-4 border-b border-base-content/10">
                                    <h2 className="text-xl font-bold">{siteName}</h2>
                                    <a
                                        href={siteURL}
                                        target="_blank"
                                        rel="noopener noreferrer"
                                        className="text-primary text-sm"
                                    >
                                        {siteURL}
                                    </a>
                                </div>
                                <div className="overflow-x-auto">
                                    <table className="table w-full">
                                        <thead>
                                            <tr>
                                                <th>ID</th>
                                                <th>Type</th>
                                                <th>Status</th>
                                                <th>Shared On</th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            {verifications.map((v) => (
                                                <tr key={v.id_shared}>
                                                    <td>{v.id_shared}</td>
                                                    <td>{v.verification_type}</td>
                                                    <td>{v.verification_status}</td>
                                                    <td>{new Date(v.shared_timestamp).toLocaleString()}</td>
                                                </tr>
                                            ))}
                                        </tbody>
                                    </table>
                                </div>
                            </div>
                        );
                    })
                )}
            </>
        </MembersLayout>
    );
}
