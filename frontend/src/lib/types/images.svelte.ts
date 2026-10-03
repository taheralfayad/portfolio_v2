import { api } from "$lib/utils/api.svelte.js";

type ImageSite = "home" | "books" | "blog" | "coffee_hero" | null;
type UploadResult =
  | { success: true; resp: Image }
  | { success: false; error: string };

type GetResult =
  | { success: true; resp: Image[] }
  | { success: false; error: string };

export interface Image {
  id: string;
  title: string;
  caption?: string;
  site: string;
  imageLink: string;
}

export const getImages = async (site: ImageSite): Promise<GetResult> => {
  let url = "";

  if (!site) {
    url = "/images";
  } else {
    url = `/images?site=${site}`;
  }

  try {
    const resp: Image[] = await api.get(url);
    return { success: true, resp };
  } catch (err) {
    console.error(err);
    const error = err instanceof Error ? err.message : String(err);
    return { success: false, error };
  }
};

export const uploadImage = async (
  title: string,
  image: string,
  site: ImageSite,
  caption?: string,
): Promise<UploadResult> => {
  try {
    const resp: Image = await api.post("/images", {
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
