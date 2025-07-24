import React from "react";

interface ProgressBarProps {
  current: number;
  goal: number;
}

const ProgressBar: React.FC<ProgressBarProps> = ({ current, goal }) => {
  const pct = Math.min(100, Math.round((current / goal) * 100));
  return (
    <div className="w-full max-w-xl mx-auto my-8">
      <div className="w-full h-4 bg-base-300 rounded-full overflow-hidden">
        <div
          className="h-full bg-primary transition-all"
          style={{ width: `${pct}%` }}
        />
      </div>
      <div className="flex justify-between text-sm mt-2">
        <span>${current.toLocaleString()}</span>
        <span>${goal.toLocaleString()}</span>
      </div>
      <div className="text-center text-sm mt-1">{pct}% of goal</div>
    </div>
  );
};

export default ProgressBar;
