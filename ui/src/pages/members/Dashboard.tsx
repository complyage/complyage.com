//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React                        from "react";
import { useNavigate }              from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { VerificationStatus }       from "../../interfaces/VerificationStatus";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import MembersDashboardContent      from "../../components/members/MembersDashboardContent";
import MembersLayout                from "../../layouts/MembersLayout";
import {
	CheckCircle,
	ArrowRight,
	CreditCard,
	MapPin,
	Mail,
	Phone,
	IdCard,
	Image as ImageIcon,
} from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Dashboard() {

      //||------------------------------------------------------------------------------------------------||
      //|| Data
      //||------------------------------------------------------------------------------------------------||

      const verifications: VerificationStatus[] = [
		{
			type: "MAIL",
			label: "Email",
			blurb: "Confirm your email to receive account updates.",
			complete: true,
			icon: <Mail />,
		},
		{
			type: "PHNE",
			label: "Phone",
			blurb: "Add and confirm your phone number.",
			complete: false,
			icon: <Phone />,
		},
		{
			type: "AGE",
			label: "ID / Age",
			blurb: "Verify your age with a government ID.",
			complete: true,
			icon: <IdCard />,
		},
		{
			type: "ADDR",
			label: "Address",
			blurb: "Verify your billing or home address.",
			complete: true,
			icon: <MapPin />,
		},
		{
			type: "CARD",
			label: "Credit Card",
			blurb: "Secure your account with a valid card on file.",
			complete: true,
			icon: <CreditCard />,
		},
		{
			type: "PROF",
			label: "Profile Photo",
			blurb: "Upload a clear profile picture.",
			complete: false,
			icon: <ImageIcon />,
		},
	];
      
      //||------------------------------------------------------------------------------------------------||
      //|| Status
      //||------------------------------------------------------------------------------------------------||

	return (
		<MembersLayout title="Your Verification Checklist">

                  <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
				{verifications.map((v, idx) => (
					<div
						key={idx}
						className="flex flex-col justify-between text-left p-6 bg-base-100 rounded-xl shadow gap-3 h-full relative">
						<div className="flex flex-col gap-2 flex-grow">
							<div className="text-primary mb-2">
								{React.cloneElement(
									v.icon as React.ReactElement,
									{size: 48}
								)}
							</div>

							<h3 className="text-xl font-bold">
								{v.label}
							</h3>

							<p className="text-base-content/70 text-sm">
								{v.blurb}
							</p>
						</div>

						<div className="flex justify-end mt-4">
							{v.complete ? (
								<button className="btn btn-success btn-sm">
									<CheckCircle className="w-4 h-4 mr-1" />{" "}
									Verified
								</button>
							) : (
								<button className="btn btn-primary btn-sm">
									Complete{" "}
									<ArrowRight className="w-4 h-4 ml-1" />
								</button>
							)}
						</div>
					</div>
				))}
			</div>

		</MembersLayout>
	);
}
