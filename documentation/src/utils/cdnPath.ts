
export const CDN_URL = import.meta.env.VITE_CDN_URL || "/";

export default function cdnPath(path: string): string {
      return `${CDN_URL}/${path.replace(/^\/+/, "")}`;
}