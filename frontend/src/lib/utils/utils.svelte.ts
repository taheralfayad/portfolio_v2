import { api } from "$lib/utils/api.svelte.js";

export const isLoggedIn = async () => {
  try {
    const resp = await api.post("/me");
  } catch (e) {
    console.error(e);
    return false;
  }

  return true;
};

export const normalizeDate = (date) => {
  return date ? date.split("T")[0] : null;
};

export const handleImageChange = (file) => {
  return new Promise((resolve) => {
    if (!file) {
      resolve("");
      return;
    }

    const reader = new FileReader();

    reader.onload = (e) => {
      resolve(e.target.result);
    };

    reader.readAsDataURL(file);
  });
};

type UploadResult =
  | { success: true; resp: uploadImageResponse }
  | { success: false; error: string };
type ImageSite = "home" | "books" | "blog";
interface uploadImageResponse {
  id: string;
  title: string;
  caption?: string;
  site: string;
  image: string;
}
export const uploadImage = async (
  title: string,
  image: string,
  site: ImageSite,
  caption?: string,
): Promise<UploadResult> => {
  try {
    const resp: uploadImageResponse = await api.post("/images", {
      title,
      caption,
      image,
      site,
    });
    return { success: true, resp };
  } catch (err) {
    console.error(err);
    const error = err instanceof Error ? err.message : String(err);
    return { success: false, error };
  }
};

export const formatDate = (dateStr) => {
  const date = new Date(dateStr);
  return date.toLocaleDateString("en-US", {
    year: "numeric",
    month: "long",
    day: "numeric",
  });
};
