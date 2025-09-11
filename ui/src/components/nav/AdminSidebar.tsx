//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React                     from "react";
import { useLocation }           from "react-router-dom";
import { 
      Home, CheckCircle, LogOut
} from "lucide-react";

import RequireAdmin, { User }    from "../dynamic/RequireMember";
import LinkQuery                 from "../../components/dynamic/LinkQuery";

//||------------------------------------------------------------------------------------------------||
//|| Sidebar
//||------------------------------------------------------------------------------------------------||

export default function AdminSidebar() {
      //||------------------------------------------------------------------------------------------------||
      //|| Hooks
      //||------------------------------------------------------------------------------------------------||
      const location = useLocation();

      //||------------------------------------------------------------------------------------------------||
      //|| Nav Items
      //||------------------------------------------------------------------------------------------------||
      const navItems = [
            {
                  to: "/admin/",
                  label: "Dashboard",
                  icon: Home,
                  activePath: "/admin/",
            },
            {
                  to: "/admin/verifications",
                  label: "Verifications",
                  icon: CheckCircle,
                  activePath: "/admin/verifications",
            },
            {
                  to: "/admin/sites",
                  label: "Site Approvals",
                  icon: CheckCircle,
                  activePath: "/admin/verifications",
            },
            {
                  to: "/admin/payments/add",
                  label: "Add Payment",
                  icon: CheckCircle,
                  activePath: "/admin/verifications",
            },
            {
                  to: "/admin/payments/find",
                  label: "Add Payment",
                  icon: CheckCircle,
                  activePath: "/admin/verifications",
            },
      ];

      //||------------------------------------------------------------------------------------------------||
      //|| Is Active
      //||------------------------------------------------------------------------------------------------||
      const isActive = (path: string) => location.pathname === path;

      //||------------------------------------------------------------------------------------------------||
      //|| Return
      //||------------------------------------------------------------------------------------------------||
      return (
            <RequireAdmin>
                  {(user: User) => (
                        <aside className="w-64 bg-gray-800 fixed left-0 bottom-0 top-[80px] z-100">
                              <nav className="flex flex-col gap-2 flex-1">
                                    <div className="text-xs p-4 text-center bg-gray-800 border-b border-gray-400">
                                          <span className="border-b border-yellow-400 p-4">
                                                Welcome <b>{user.username}</b>
                                          </span>
                                    </div>

                                    {navItems.map(({ to, label, icon: Icon, activePath }) => (
                                          <LinkQuery
                                                key={to}
                                                to={to}
                                                className={`btn justify-start py-5 ${
                                                      isActive(activePath) ? "btn-primary" : "btn-ghost"
                                                }`}
                                          >
                                                <Icon className="w-5 h-5 mr-2" /> {label}
                                          </LinkQuery>
                                    ))}

                                    <button className="btn btn-ghost fixed bottom-0 w-64 py-4 pb-10 text-sm">
                                          <LogOut className="w-4 h-4" />
                                          <span className="text-xs">Logout</span>
                                    </button>
                              </nav>
                        </aside>
                  )}
            </RequireAdmin>
      );
}
