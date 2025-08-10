export type DateRange = {
  start: string; // ISO date or YYYY-MM
  end?: string; // ISO date, YYYY-MM, or undefined for present
};

export type Project = {
  title: string;
  bullets: string[];
};

export type Job = {
  title: string;
  period: DateRange;
  projects?: Project[];
};

export type Company = {
  name: string;
  url?: string;
  jobs: Job[];
};
