//||------------------------------------------------------------------------------------------------||
//|| PasswordField.tsx
//||------------------------------------------------------------------------------------------------||

import React, { useState } from "react";
import { Eye, EyeOff } from "lucide-react"; // if lucide-react is available, otherwise replace with emoji or SVG

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface PasswordFieldProps {
      id?            : string;
      placeholder?   : string;
      onChange?      : (e: React.ChangeEvent<HTMLInputElement>) => void;
      className?     : string;
      required?      : boolean;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function PasswordField({
      id = "password",
      placeholder = "Enter your password",
      onChange,
      className = "",
      required = false,
}: PasswordFieldProps) {

      //||------------------------------------------------------------------------------------------------||
      //|| Var
      //||------------------------------------------------------------------------------------------------||

      const [visible, setVisible] = useState(false);
      const [password, setPassword] = useState("");
      const toggleVisibility = () => setVisible(!visible);

      //||------------------------------------------------------------------------------------------------||
      //|| OnChange
      //||------------------------------------------------------------------------------------------------||

      const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
            setPassword(e.target.value);
            onChange?.(e);
      };

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <div className={`w-full ${className}`}>

                  <div className="relative">
                        <input
                              id={id}
                              type={visible ? "text" : "password"}
                              placeholder={placeholder}
                              value={password}
                              onChange={handleChange}
                              required={required}
                              className="input input-bordered w-full py-5 text-xl h-12"
                        />

                        <button
                              type="button"
                              onClick={toggleVisibility}
                              className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-200 p-3 rounded"
                              aria-label={visible ? "Hide password" : "Show password"}
                        >
                              {visible ? <EyeOff size={20} /> : <Eye size={20} />}
                        </button>
                  </div>
            </div>
      );
}
