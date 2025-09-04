import { useState } from "react";
import { Menu, User, LogOut } from "lucide-react";

const navLinks = [
   { label: "Dashboard", icon: <Menu className="w-5 h-5" />, href: "#" },
   { label: "Users",      icon: <User className="w-5 h-5" />, href: "#" },
   // Add more as you want
];

export default function AdminLayout() {
   const [sidebarOpen, setSidebarOpen] = useState(false);

   return (
      <div className="min-h-screen flex bg-gray-100">
         {/* Sidebar */}
         <aside className={`fixed md:static top-0 left-0 h-full z-30 bg-white shadow-lg
            transition-transform duration-200 w-64
            ${sidebarOpen ? "translate-x-0" : "-translate-x-full md:translate-x-0"}`}>
            <div className="h-16 flex items-center justify-between px-6 border-b">
               <span className="font-bold text-xl text-primary">Admin</span>
               <button className="md:hidden" onClick={() => setSidebarOpen(false)}>
                  <Menu className="w-6 h-6" />
               </button>
            </div>
            <nav className="py-4">
               {navLinks.map(link => (
                  <a key={link.label} href={link.href}
                     className="flex items-center gap-3 px-6 py-3 text-gray-700 hover:bg-gray-100 font-medium transition rounded">
                     {link.icon} {link.label}
                  </a>
               ))}
            </nav>
         </aside>

         {/* Overlay (for mobile) */}
         {sidebarOpen && (
            <div className="fixed inset-0 bg-black bg-opacity-30 z-20 md:hidden"
               onClick={() => setSidebarOpen(false)} />
         )}

         {/* Main Area */}
         <div className="flex-1 flex flex-col min-h-screen ml-0 md:ml-64 transition-all duration-200">
            {/* Topbar */}
            <header className="h-16 flex items-center justify-between px-8 bg-white shadow-sm border-b">
               <button className="md:hidden" onClick={() => setSidebarOpen(true)}>
                  <Menu className="w-6 h-6" />
               </button>
               <div className="flex items-center gap-3">
                  <User className="w-7 h-7 text-gray-400" />
                  <span className="font-semibold">Admin User</span>
                  <button className="ml-3 p-2 rounded hover:bg-gray-200">
                     <LogOut className="w-5 h-5 text-gray-400" />
                  </button>
               </div>
            </header>
            {/* Main Content */}
            <main className="flex-1 p-8">
               <div className="text-2xl font-bold mb-4">Dashboard</div>
               <div className="bg-white p-8 rounded shadow text-gray-500 text-lg">
                  Welcome to your admin panel. Build whatever you want here.
               </div>
            </main>
         </div>
      </div>
   );
}
