/*||------------------------------------------------------------------------------------------------||
//|| Progress Steps Component
//|| ProgressSteps
//||------------------------------------------------------------------------------------------------||*/

/*||------------------------------------------------------------------------------------------------||
//|| React
//||------------------------------------------------------------------------------------------------||*/

import React from "react";

/*||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||*/

interface ProgressBarProps {
	current           : number;
	goal              : number;
      label?            : string;
      prefix?           : string;
      showStartEnd?     : boolean;
}

/*||------------------------------------------------------------------------------------------------||
//|| Progress Bar
//||------------------------------------------------------------------------------------------------||*/

const ProgressBar: React.FC<ProgressBarProps> = ({current, goal, label, prefix, showStartEnd}) => {
	const pct = Math.min(100, Math.round((current / goal) * 100));
	return (
		<div className="w-full max-w-xl mx-auto my-8">
			<div className="w-full h-4 bg-base-300 rounded-full overflow-hidden">
				<div className="h-full bg-primary transition-all" style={{width: `${pct}%`}} />
			</div>
			{showStartEnd  && ( <div className="flex justify-between text-sm mt-2">
				<span>{prefix}{current.toLocaleString()}</span>
				<span>{prefix}{goal.toLocaleString()}</span>
			</div>) }
			{ label && (<div className="text-center text-sm mt-1">{pct}% {label}</div>) }
		</div>
	);
};

export default ProgressBar;
