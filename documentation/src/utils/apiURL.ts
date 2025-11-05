

      //||------------------------------------------------------------------------------------------------||
      //|| Handle Production / Dev URL Differences
      //||------------------------------------------------------------------------------------------------||

      export default function apiURL(path : string): string {
            const url = import.meta.env.VITE_COMPLYAGE_API_URL + path;
            console.log("API URL:", url);
            return url;
      }
