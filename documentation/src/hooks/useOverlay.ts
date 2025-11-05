//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import { useNavigate, useLocation } from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Hook: useOverlayNavigate
//||------------------------------------------------------------------------------------------------||

export function useOverlayNavigate() {
      const navigate = useNavigate();
      const location = useLocation();

      const n = (to: string, options?: Parameters<typeof navigate>[1]) => {
            if (location.pathname.startsWith("/overlay")) {
                  if (!to.startsWith("/overlay")) {
                        to = "/overlay" + (to.startsWith("/") ? to : `/${to}`);
                  }
                  if (to.includes("?refresh")) {
                        if (window.self !== window.top) {
                              window.parent.postMessage({ action: "refreshParent" }, "*");
                        } else {
                              navigate(to);
                        }
                  }
            }
            navigate(to, options);
      };

      return n;
}
