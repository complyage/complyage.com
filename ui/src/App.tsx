//||------------------------------------------------------------------------------------------------||
//|| Dependencies
//||------------------------------------------------------------------------------------------------||

import React                  from 'react'
import { Routes, Route }      from 'react-router-dom'

//||------------------------------------------------------------------------------------------------||
//|| Public
//||------------------------------------------------------------------------------------------------||

import Home                   from './pages/public/Home'
import About                  from './pages/public/About'
import Pricing                from './pages/public/Pricing'
import EnforcementZones       from './pages/public/EnforcementZones'
import Vendors                from './pages/public/Vendors'
import Exit                   from './pages/public/Exit'

//||------------------------------------------------------------------------------------------------||
//|| Auth
//||------------------------------------------------------------------------------------------------||

import Signup                 from './pages/public/Signup'
import TwoFactorVerify        from './pages/public/TwoFactorVerify'
import SignupComplete         from './pages/public/SignupComplete'
import Login                  from './pages/public/Login'

//||------------------------------------------------------------------------------------------------||
//|| Private
//||------------------------------------------------------------------------------------------------||

import Test                   from './pages/members/Test'
import Dashboard              from './pages/members/Dashboard'
import Sites                  from './pages/members/Sites'
import Encrypted              from './pages/members/Encrypted'
import Shared                 from './pages/members/Shared'
import Settings               from './pages/members/Settings'
import Verification           from './pages/members/Verification'

//||------------------------------------------------------------------------------------------------||
//|| OAuth
//||------------------------------------------------------------------------------------------------||

import OAuth                  from './pages/oauth/OAuth'

//||------------------------------------------------------------------------------------------------||
//|| App
//||------------------------------------------------------------------------------------------------||

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />

      <Route path="/about" element={<About />} />
      <Route path="/pricing" element={<Pricing />} />
      <Route path="/vendors" element={<Vendors />} />
      
      <Route path="/signup" element={<Signup />} />
      <Route path="/gilead" element={<EnforcementZones />} />
      <Route path="/verify" element={<TwoFactorVerify />} />
      <Route path="/complete" element={<SignupComplete />} />
      <Route path="/login" element={<Login />} />
      <Route path="/exit" element={<Exit />} />

      <Route path="/members/test" element={<Test />} />
      <Route path="/members" element={<Dashboard />} />
      <Route path="/members/sites" element={<Sites />} />
      <Route path="/members/settings" element={<Settings />} />
      <Route path="/members/shared" element={<Shared />} />
      <Route path="/members/encrypted" element={<Encrypted />} />
      <Route path="/members/verification" element={<Verification />} />
      
      <Route path="/oauth" element={<OAuth />} />
    </Routes>
  )
}