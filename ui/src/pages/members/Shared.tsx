//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useState }                 from "react";
import { useNavigate }                                from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { SharedItem }                                 from "../../interfaces/sharedItem";

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

import MembersLayout                                  from "../../layouts/MembersLayout";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Shared() {
	
      //||------------------------------------------------------------------------------------------------||
	//|| Var
	//||------------------------------------------------------------------------------------------------||

	const navigate = useNavigate();
	const [searchQuery, setSearchQuery] = useState("");
	const [sharedItems, setSharedItems] = useState<SharedItem[]>([]);

      //||------------------------------------------------------------------------------------------------||
	//|| Mock
	//||------------------------------------------------------------------------------------------------||

	const sharedData: SharedItem[] = [
		{
			id_shared: 1,
			fid_site: 1,
			fid_verification: 100,
			shared_timestamp: new Date().toISOString(),
			site_name: "Example Site A",
			site_url: "https://site-a.com",
			verification_type: "DOCS",
			verification_status: "ACTV",
		},
		{
			id_shared: 2,
			fid_site: 1,
			fid_verification: 101,
			shared_timestamp: new Date().toISOString(),
			site_name: "Example Site A",
			site_url: "https://site-a.com",
			verification_type: "PHOTO",
			verification_status: "QUEU",
		},
		{
			id_shared: 3,
			fid_site: 2,
			fid_verification: 200,
			shared_timestamp: new Date().toISOString(),
			site_name: "Zebra Media",
			site_url: "https://zebramedia.net",
			verification_type: "DOCS",
			verification_status: "DENY",
		},
		{
			id_shared: 4,
			fid_site: 3,
			fid_verification: 300,
			shared_timestamp: new Date().toISOString(),
			site_name: "Alpha Platform",
			site_url: "https://alpha.io",
			verification_type: "DOCS",
			verification_status: "ACTV",
		},
	];

      //||------------------------------------------------------------------------------------------------||
	//|| Initial Load
	//||------------------------------------------------------------------------------------------------||
      
	useEffect(() => {
            setSharedItems(sharedData);
      }, []);

      //||------------------------------------------------------------------------------------------------||
	//|| Group
	//||------------------------------------------------------------------------------------------------||

	const groupedBySite = sharedItems
		.filter((item) => item.site_name.toLowerCase().includes(searchQuery.toLowerCase()))
		.reduce((groups: Record<string, SharedItem[]>, item) => {
			const key = `${item.site_name}||${item.site_url}`;
			if (!groups[key]) groups[key] = [];
			groups[key].push(item);
			return groups;
	}, {});

            
      //||------------------------------------------------------------------------------------------------||
	//|| Get Keys
	//||------------------------------------------------------------------------------------------------||

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

                        {siteKeys.length === 0 ? (
                              <p className="text-sm text-base-content/70">No shared verifications found.</p>
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
                                                            className="text-primary text-sm">
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
