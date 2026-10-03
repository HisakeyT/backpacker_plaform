export type TravelPlan = {
  id: number;
  travelId: number;
  date: string;
  place: string;
  content: string;
  sortOrder: number;
  createdAt: string;
  updatedAt: string;
};

export type CreateTravelPlanInput = {
  date: string; // YYYY-MM-DD
  place: string;
  content: string;
  sortOrder: number;
}

export type UpdateTravelPlanInput = {
  date?: string; // YYYY-MM-DD
  place?: string;
  content?: string;
}
