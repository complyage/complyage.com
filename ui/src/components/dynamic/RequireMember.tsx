import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

export type User = {
   user_id: number;
   email: string;
};

type RequireMemberProps = {
   children: (user: User) => React.ReactNode;
};

export default function RequireMember({ children }: RequireMemberProps) {
   const navigate = useNavigate();
   const [loading, setLoading] = useState(true);
   const [user, setUser] = useState<User | null>(null);

   useEffect(() => {
      const checkAuth = async () => {
         // TEMP: hardcoded user
         setUser({ user_id: 1, email: "test@example.com" });
         setLoading(false);
         return;

         // Uncomment for real check:
         // try {
         //    const res = await fetch("/auth/me", { method: "GET", credentials: "include" });
         //    if (res.ok) {
         //       const data = await res.json();
         //       if (data.success) {
         //          setUser(data.data);
         //       } else {
         //          navigate("/login");
         //       }
         //    } else {
         //       navigate("/login");
         //    }
         // } catch (err) {
         //    console.error(err);
         //    navigate("/login");
         // } finally {
         //    setLoading(false);
         // }
      };

      checkAuth();
   }, [navigate]);

   if (loading) {
      return (
         <main className="min-h-screen flex items-center justify-center">
            <p>Checking your login status...</p>
         </main>
      );
   }

   if (!user) return null;

   return <>{children(user)}</>; // ✅ This is key
}
