import React, {useEffect, useMemo, useState} from "react";
import {Star} from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Interactive Stars 
//||------------------------------------------------------------------------------------------------||

export default function Stars({
	n,
	avg,
	interactive = false,
	onRate,
	userRating,
	disabled,
}: {
	n                 : number;
	avg?              : number;
	interactive?      : boolean;
	onRate?           : (val: number) => void;
	userRating?       : number;
	disabled?         : boolean;
}) {
      //||------------------------------------------------------------------------------------------------||
      //|| Hover
      //||------------------------------------------------------------------------------------------------||

      const [hovered, setHovered] = useState<number | null>(null);

      //||------------------------------------------------------------------------------------------------||
      //|| Component
      //||------------------------------------------------------------------------------------------------||

      return (
		<div className="flex items-center gap-1 group">
			{[...Array(5)].map((_, i) => {
				const filled = hovered !== null ? i < hovered : typeof userRating === "number" ? i < userRating : i < n;
				return (
					<button
						key={i}
						type="button"
						disabled={!interactive || disabled}
						className={`p-0 m-0 bg-transparent border-none focus:outline-none transition-transform duration-100 ${
							interactive && !disabled ? "hover:scale-110 active:scale-125" : ""
						}`}
						style={{
							cursor: interactive && !disabled ? "pointer" : "default",
						}}
						tabIndex={interactive ? 0 : -1}
						onMouseEnter={() => interactive && setHovered(i + 1)}
						onMouseLeave={() => interactive && setHovered(null)}
						onClick={() => interactive && onRate?.(i + 1)}
						aria-label={`Rate ${i + 1} star${i > 0 ? "s" : ""}`}>
						<Star
							className={`w-5 h-5 transition-colors duration-100 ${
								filled ? "text-yellow-400 fill-yellow-400 drop-shadow-[0_0_5px_rgba(251,191,36,0.3)]" : "text-base-content/30"
							}`}
						/>
					</button>
				);
			})}
			{typeof avg === "number" && <span className="ml-1 text-xs text-base-content/60 tabular-nums font-semibold">{avg.toFixed(2)} / 5</span>}
		</div>
	);
}
