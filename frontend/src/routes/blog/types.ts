export interface Blog {
  id?: number;
  content: string;
  date: string;
  metadata: {
    id: string;
    title: string;
    creator?: string;
    rating?: number;
    clickbait: string;
    tags: string[];
  };
}

export interface BlogResponse {
  id: number;
  created_at: string;
  metadata: string;
  content: string;
}
