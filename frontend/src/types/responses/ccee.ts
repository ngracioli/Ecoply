export interface CCEEAgent {
  label: string;
  value: string;
  submarket: string;
}

export interface CCEEAgentsResponse {
  data: CCEEAgent[];
}
