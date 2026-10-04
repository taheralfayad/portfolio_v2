import { api } from "$lib/utils/api.svelte";

export interface Skill {
  id: number;
  name: string;
  category: string;
  blogLink: string;
  createdAt: string;
}

type SkillResponse = Omit<Skill, "blogLink" | "createdAt"> & {
  blog_link: string;
  created_at: string;
};

type GetResult =
  | { success: true; resp: Skill[] }
  | { success: false; error: string };

export const getSkills = async (): Promise<GetResult> => {
  let url = "/skills";

  try {
    const resp: SkillResponse[] = await api.get(url);

    const formattedResp: Skill[] = resp.map(
      ({ blog_link, created_at, ...rest }) => ({
        ...rest,
        blogLink: blog_link,
        createdAt: created_at,
      }),
    );

    return { success: true, resp: formattedResp };
  } catch (err) {
    console.error(err);
    return { success: false, error: err };
  }
};
