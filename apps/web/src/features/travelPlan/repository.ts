import { apiFetch } from "../../lib/apiFetch";
import type { TravelPlan } from "./types";

export const getTravelPlans = async (
  token: string,
  travelId: number,
): Promise<TravelPlan[]> => {
  return apiFetch<TravelPlan[]>(`/api/travels/${travelId}/plans`, { token, method: "GET" });
};
