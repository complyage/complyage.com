
import { DOB } from "../interfaces/base/user";

export function getAge(dob: DOB): number | null {
      if (!dob.year) return null;
      const today = new Date();
      let age = today.getFullYear() - dob.year;

      // Has birthday occurred yet this year?
      if (
            today.getMonth() + 1 < dob.month || // months are 0-based
            (today.getMonth() + 1 === dob.month && today.getDate() < dob.day)
      ) {
            age--;
      }
      return age;
}