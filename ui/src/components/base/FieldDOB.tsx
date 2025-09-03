/*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
//|| ~/components/forms/FieldDOB.tsx
//|| FieldDOB Component
//||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

      /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
      //|| React
      //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

      import React, {useState, useEffect}             from "react";
      import { CircleCheckBig, CircleX }              from "lucide-react";

      /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
      //|| DOB
      //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

      import { DOB }                                  from "../../interfaces/base/user";

      /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
      //|| Props
      //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

      interface FieldDOBProps {
            id          : string;
            value?      : string;
            minage      : number;
            onChange    : (DOB: DOB) => void;
      }

      /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
      //|| Component
      //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/


      const FieldDOB: React.FC<FieldDOBProps> = ({id, value, minage, onChange }) => {

            /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
            //|| State
            //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

            const initialDate = value && !isNaN(new Date(value).getTime()) ? new Date(value) : new Date();
            const defaultY    = initialDate.getFullYear() - 18;
            const defaultM    = initialDate.getMonth() + 1;
            const defaultD    = initialDate.getDate();

            /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
            //|| State
            //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/
                        
            const [dob, setDOB]         = useState<DOB>({
                  year  : defaultY,
                  month : defaultM,
                  day   : defaultD
            });
            const [age, setAge]           = useState<number>(0);
            const [isValid, setIsValid]   = useState<boolean>(true);

            /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
            //|| Calculate Age
            //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

            const calculateAge = (birthDate: string): number => {
                  const today = new Date();
                  const birthDateObj = new Date(birthDate);
                  let calculatedAge =
                        today.getFullYear() - birthDateObj.getFullYear();
                  const monthDifference = today.getMonth() - birthDateObj.getMonth();
                  if (
                        monthDifference < 0 ||
                        (monthDifference === 0 &&
                              today.getDate() < birthDateObj.getDate())
                  ) {
                        calculatedAge--;
                  }
                  return calculatedAge;
            };

            /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
            //|| Handle changes in day, month, or year
            //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

            const handleChange = (type: string, value: string) => {
                  const intValue = parseInt(value);
                  if (type === "day") {
                        setDOB({ ...dob, day: intValue });
                  } else if (type === "month") {
                        setDOB({ ...dob, month: intValue });
                  } else if (type === "year") {
                        setDOB({ ...dob, year: intValue });
                  }                  
            };

            /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
            //|| Update the date and validate age
            //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

            const updateDate = () => {
                  const formattedDate = `${dob.year}-${String(dob.month).padStart(2,"0")}-${String(dob.day).padStart(2, "0")}`;
                  const calculatedAge = calculateAge(formattedDate);
                  setAge(calculatedAge);
                  if (calculatedAge < minage) {
                        setIsValid(false);
                  } else {
                        setIsValid(true);
                  }
            };

            /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
            //|| Generate days for the selected month and year
            //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

            const generateDays = () => {
                  const daysInMonth = new Date(dob.year || 1981, dob.month || 6, 0).getDate();
                  console.log("DAYS IN MONTH", daysInMonth, dob.year, dob.month);
                  return Array.from({ length: daysInMonth }, (_, i) => i + 1);
            };

            /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
            //|| Generate Months
            //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

            const generateMonths = () => {
                  return Array.from({length: 12}, (_, i) => i + 1);
            };

            /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
            //|| Generate Years
            //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

            const generateYears = () => {
                  const currentYear = new Date().getFullYear();
                  return Array.from({length: 110}, (_, i) => currentYear - i);
            };

            /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
            //|| Update Date
            //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

            useEffect(() => {
                  updateDate();
                  onChange(dob);
            }, [dob]);

            /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
            //|| CSS
            //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

            const selectClass = "flex-1 select select-bordered text-white text-2xl h-auto bg-black p-4";

            /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
            //|| Component
            //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

            return (
                  <>
                  <div className="form-group flex max-w-xl mx-auto" minage={minage}>
                        <div className="flex space-x-4 flex-grow">
                              
                              {/* Month */}

                              <select
                                    className={selectClass}
                                    defaultValue={String(dob.month)}
                                    onChange={(e) =>
                                          handleChange("month", e.target.value)
                                    }>
                                    {generateMonths().map((m) => (
                                          <option
                                                key={m}
                                                className="text-white"
                                                value={String(m)}
                                          >{new Date(0, m - 1).toLocaleString("default", { month: "long" })}</option>
                                    ))}
                              </select>

                              {/* Day */}

                              <select
                                    className={selectClass}
                                    defaultValue={String(dob.day)}
                                    onChange={(e) =>
                                          handleChange("day", e.target.value)
                                    }>
                                    {generateDays().map((d) => (
                                          <option key={d} value={String(d)}>{String(d)}</option>
                                    ))}
                              </select>

                              {/* Year */}

                              <select
                                    className={selectClass}
                                    defaultValue={String(dob.year)}
                                    onChange={(e) =>
                                          handleChange("year", e.target.value)
                                    }>
                                    {generateYears().map((y) => (
                                          <option key={y} value={String(y)} >{String(y)}</option>
                                    ))}
                              </select>
                        </div>
                              
                        <input
                              type="date"
                              id={id}
                              name="dob"
                              value={`${dob.year}-${String(dob.month).padStart(
                                    2,
                                    "0"
                              )}-${String(dob.day).padStart(2, "0")}`}
                              className="hidden"
                              readOnly
                        />
                  </div>
                  <div className="mt-2 mx-auto max-w-xl rounded bg-black/20 p-2">
                        {isValid ? (
                              <span className="inline-flex items-center gap-2 text-xl text-green-600 leading-none">
                                    <CircleCheckBig className="h-6 w-6" />
                                    <b>{age} years old.</b>
                              </span>
                        ) : (
                              <span className="inline-flex items-center gap-2 text-xl text-red-600 leading-none">
                                    <CircleX className="h-6 w-6" />
                                    <b>You must be at least {minage} years old.</b>
                              </span>
                        )}
                  </div>
            </>

            );
      };

      export default FieldDOB;
