import { api } from "$lib/utils/api.svelte";

export interface Project {
  id: string;
  name: string;
  description: string;
  githubLink: string;
  blogLink: string;
  type: string;
  image: string;
}

type ProjectResponse = Omit<
  Project,
  "githubLink" | "blogLink" | "imageLink"
> & {
  github_link: string;
  blog_link: string | null;
  image_link: string;
};

type GetResult =
  | { success: true; resp: Project[] }
  | { success: false; error: string };

type ProjectType = "work" | "personal";
export const getProjects = async (
  type: ProjectType,
  limit: number,
): Promise<GetResult> => {
  let url = `/projects?type=${type}`;

  if (limit > 0) {
    url += `&limit=${limit}`;
  } else {
    url += "&limit=5";
  }

  try {
    const resp: ProjectResponse[] = await api.get(url);

    const formattedResp: Project[] = resp.map(
      ({ github_link, image_link, blog_link, ...rest }) => ({
        ...rest,
        githubLink: github_link,
        blogLink: blog_link,
      }),
    );

    return { success: true, resp: formattedResp };
  } catch (err) {
    console.error(err);
    return { success: false, error: err };
  }
};
