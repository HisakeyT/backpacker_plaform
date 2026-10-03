import { apiFetch } from "../../lib/apiFetch";
import type { TravelPlan, CreateTravelPlanInput } from "./types";

export const getTravelPlans = async (
  token: string,
  travelId: number,
): Promise<TravelPlan[]> => {
  return apiFetch<TravelPlan[]>(`/api/travels/${travelId}/plans`, { token, method: "GET" });
};

export const createTravelPlan = async (
  token: string,
  travelId: number,
  input: CreateTravelPlanInput,
): Promise<TravelPlan> => {
  return apiFetch<TravelPlan>(`/api/travels/${travelId}/plans`, {
    method: "POST",
    token,
    body: input,
  });
};
