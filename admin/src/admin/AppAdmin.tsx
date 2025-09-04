//||-------------------------------------------------------------------------------------||
//|| Admin App Shell (React-Admin)
//||-------------------------------------------------------------------------------------||

import { Admin, Resource, ListGuesser } from "react-admin";
import simpleRestProvider from "ra-data-simple-rest";

// Demo: JSONPlaceholder
const dataProvider = simpleRestProvider("https://jsonplaceholder.typicode.com");

export default function AppAdmin() {
   return (
      <Admin dataProvider={dataProvider}>
         <Resource name="users" list={ListGuesser} />
         <Resource name="posts" list={ListGuesser} />
      </Admin>
   );
}