import React from "react";

type CompleteSliderProps = {
      value: number;                        // current level from parent
      setValue?: (level: number) => void;   // callback to update parent
};

export default function CompleteSlider({ value, setValue }: CompleteSliderProps) {
      const labels: Record<number, string> = {
            1: "Standard",
            2: "BIP39 (6)",
            3: "BIP39 (12)",
            4: "BIP39 (18)",
            5: "BIP39 (24)",
            6: "Key Pair",
      };

      const handleChange = (n: number) => {
            if (setValue) setValue(n);
      };

      return (
            <div className="w-full max-w-xl mx-auto">
                  {/* Track with progress fill */}
                  <div className="relative w-full h-2 bg-gray-700 rounded-full">
                        <div
                              className="absolute top-0 left-0 h-2 bg-blue-500 rounded-full transition-all duration-300"
                              style={{ width: `${((value - 1) / 5) * 100}%` }}
                        />
                        {/* Step dots */}
                        <div className="absolute top-1/2 -translate-y-1/2 w-full flex justify-between px-[2px]">
                              {[1, 2, 3, 4, 5, 6].map((n) => (
                                    <button
                                          key={n}
                                          onClick={() => handleChange(n)}
                                          className={`w-5 h-5 rounded-full border-2 transition ${
                                                value === n
                                                      ? "bg-blue-500 border-blue-400 scale-110"
                                                      : "bg-gray-900 border-gray-500 hover:border-blue-300"
                                          }`}
                                    />
                              ))}
                        </div>
                  </div>

                  {/* Labels under each step */}
                  <div className="flex justify-between mt-6 text-xs text-gray-400">
                        {[1, 2, 3, 4, 5, 6].map((n) => (
                              <span
                                    key={n}
                                    className={`w-16 text-center transition ${
                                          value === n ? "text-blue-400 font-bold" : "opacity-50"
                                    }`}
                              >
                                    {labels[n]}
                              </span>
                        ))}
                  </div>
            </div>
      );
}
