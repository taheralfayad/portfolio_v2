import { api } from "$lib/utils/api.svelte";
import { normalizeDate } from "$lib/utils/utils.svelte";

export interface WorkExperience {
  id: number;
  title: string;
  workplace: string;
  description: string;
  startDate: string;
  endDate?: string;
}

type WorkExperienceResponse = Omit<WorkExperience, "startDate" | "endDate"> & {
  start_date: string;
  end_date?: string | null;
};

type GetResult =
  | { success: true; resp: WorkExperience[] }
  | { success: false; error: string };

export const getWorkExperiences = async (limit: number): Promise<GetResult> => {
  let url = "";

  if (limit > 0) {
    url = `/work-experiences?limit=${limit}`;
  } else {
    url = "/work-experiences?limit=3";
  }

  try {
    const resp: WorkExperienceResponse[] = await api.get(url);

    const formattedResp: WorkExperience[] = resp.map(
      ({ start_date, end_date, ...rest }) => ({
        ...rest,
        startDate: normalizeDate(start_date),
        endDate: end_date ? normalizeDate(end_date) : undefined,
      }),
    );

    return { success: true, resp: formattedResp };
  } catch (err) {
    console.error(err);
    const error = err instanceof Error ? err.message : String(err);
    return { success: false, error };
  }
};
