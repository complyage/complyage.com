/*||------------------------------------------------------------------------------------------------||
//|| DonateCheck Component
//|| src/components/donate/DonateCheck.tsx
//||------------------------------------------------------------------------------------------------||*/

import React, { useState, useEffect } from "react";
import Call from "../../classes/call";
import type { DonateApiResponse, DonateAddress } from "../../interfaces/donate/donate";

export default function DonateCheck() {
      const [address, setAddress] = useState<DonateAddress | null>(null);

      //||------------------------------------------------------------------------------------------------||
      //|| Fetch mailing address
      //||------------------------------------------------------------------------------------------------||
      useEffect(() => {
            (async () => {
                  const chirp = new Call("/v1/api/donate");
                  chirp.method = "GET";
                  await chirp.execute();
                  if (chirp.ok()) {
                        setAddress((chirp.responsePayload.data as DonateApiResponse).address || null);
                  }
            })();
      }, []);

      if (!address) return null;

      return (
            <div className="w-full mt-7">
                  <div className="flex flex-col items-center p-6 w-full ">
                        <div className="mb-3 text-xl font-bold tracking-wide uppercase border-b w-full text-center pb-2">
                              Mail a Cash/Check Donation
                        </div>
                        <div className="mt-2 mb-2 w-full max-w-2xl pl-5 font-bold">
                              <div className="text-gray-100 text-2xl text-center">{address.name}</div>
                              <div className="text-gray-300 text-xl text-center">
                                    {address.address1}<br />
                                    {address.address2 && <> {address.address2}</>}<br />
                                    {address.city}, {address.state} {address.postal}<br />
                                    {address.country}<br />
                              </div>
                        </div>
                        <div className="text-xl mt-2 opacity-90 bg-black/20 p-3 rounded text-center">
                              Please make checks payable to <b>{address.name}</b>.<br />
                              <span className="opacity-70">Include your email if you’d like a receipt for deductions.</span>
                        </div>
                  </div>
            </div>
      );
}
