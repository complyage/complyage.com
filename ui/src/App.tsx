//||------------------------------------------------------------------------------------------------||
//|| Dependencies
//||------------------------------------------------------------------------------------------------||

import React                  from 'react'
import { Routes, Route }      from 'react-router-dom'

//||------------------------------------------------------------------------------------------------||
//|| Public
//||------------------------------------------------------------------------------------------------||

import Home                   from './pages/public/Home'
import Donate                 from './pages/public/Donate'
import Contact                from './pages/public/Contact'
import TOS                    from './pages/public/TOS'
import PrivacyPolicy          from './pages/public/PrivacyPolicy'
import About                  from './pages/public/About'
import Pricing                from './pages/public/Pricing'
import EnforcementZones       from './pages/public/EnforcementZones'
import Vendors                from './pages/public/Vendors'
import Exit                   from './pages/public/Exit'

//||------------------------------------------------------------------------------------------------||
//|| Auth
//||------------------------------------------------------------------------------------------------||

import Forgot                 from './pages/public/Forgot'
import Signup                 from './pages/public/Signup'
import TwoFactorVerify        from './pages/public/TwoFactorVerify'
import SignupComplete         from './pages/public/SignupComplete'
import Login                  from './pages/public/Login'

//||------------------------------------------------------------------------------------------------||
//|| Private
//||------------------------------------------------------------------------------------------------||

import Dashboard              from './pages/members/Dashboard'
import Sites                  from './pages/members/Sites'
import Identity               from './pages/members/Identity'
import Shared                 from './pages/members/Shared'
import Settings               from './pages/members/Settings'
import Verification           from './pages/members/Verification'
import Logout                 from './pages/members/Logout'
import VPNs                   from './pages/members/VPNs'
import Quit                   from './pages/members/Quit'

//||------------------------------------------------------------------------------------------------||
//|| Verification
//||------------------------------------------------------------------------------------------------||

import VerificationInit       from "./pages/verification/Init";
import VerificationCheck      from './pages/verification/check/VerificationCheck'
import VerificationStatus     from './pages/verification/status/VerificationStatus'

//||------------------------------------------------------------------------------------------------||
//|| Verification
//||------------------------------------------------------------------------------------------------||

import CCVerification         from './pages/verification/card/CCVerification'
import AddressVerification    from './pages/verification/address/AddressVerification'
import IDVerification         from './pages/verification/id/IDVerification'
import PhoneVerification      from './pages/verification/phone/PhoneVerification'
import FaceVerification       from './pages/verification/face/FaceVerification'

//||------------------------------------------------------------------------------------------------||
//|| Test
//||------------------------------------------------------------------------------------------------||

import SelfiePage            from './pages/verification/Selfie'

//||------------------------------------------------------------------------------------------------||
//|| App
//||------------------------------------------------------------------------------------------------||

export default function App() {

      //||------------------------------------------------------------------------------------------------||
      //|| Extract oauth param from URL if exists
      //||------------------------------------------------------------------------------------------------||

      const params            = new URLSearchParams(location.search);
	const oauthParam        = params.get("oauth");
	const appendOauth       = (path: string) => (oauthParam ? `${path}?oauth=${oauthParam}` : path);

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
		<Routes>
			<Route path="/" element={<Home />} />

			<Route path="/about" element={<About />} />
			<Route path="/pricing" element={<Pricing />} />
			<Route path="/vendors" element={<Vendors />} />
                  <Route path="/donate" element={<Donate />} />
                  <Route path="/terms" element={<TOS />} />
                  <Route path="/privacy" element={<PrivacyPolicy />} />
                  <Route path="/contact" element={<Contact />} />


			<Route path="/signup" element={<Signup />} />
                  <Route path="/forgot" element={<Forgot />} />
			<Route path="/gilead" element={<EnforcementZones />} />
			<Route path="/verify" element={<TwoFactorVerify />} />
			<Route path="/complete" element={<SignupComplete />} />
			<Route path="/login" element={<Login />} />
			<Route path="/logout" element={<Logout />} />
			<Route path="/exit" element={<Exit />} />
                  
                  <Route path="/selfie" element={<SelfiePage />} />

			<Route path="/members" element={<Dashboard />} />
			<Route path="/members/sites" element={<Sites />} />
			<Route path="/members/settings" element={<Settings />} />
			<Route path="/members/shared" element={<Shared />} />
			<Route path="/members/identity" element={<Identity />} />                  
			<Route path="/members/verification" element={<Verification />} />
                  <Route path="/members/vpns" element={<VPNs />} />                  
			<Route path="/members/quit" element={<Quit />} />

                  <Route path="/verification/init" element={<VerificationInit />} />
                  <Route path="/verification/check" element={<VerificationCheck />} />
                  <Route path="/verification/status" element={<VerificationStatus />} />
                  
                  <Route path="/verification/id" element={<IDVerification />} />
                  <Route path="/verification/face" element={<FaceVerification />} />
                  <Route path="/verification/card" element={<CCVerification />} />
                  <Route path="/verification/phone" element={<PhoneVerification />} />
                  <Route path="/verification/address" element={<AddressVerification />} />
                  
		</Routes>
	);
}