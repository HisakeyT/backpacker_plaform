import type { Travel, CreateTravelInput } from "./types";
import { apiFetch } from "../../lib/apiFetch";

export const getTravels = async (token: string): Promise<Travel[]> => {
  return apiFetch<Travel[]>("/api/travels", { token, method: "GET" })
}

export const getTravel = async (token: string, travelId: number): Promise<Travel> => {
  return apiFetch<Travel>(`/api/travels/${travelId}`, { token, method: "GET" });
};

export const createTravel = async (
  token: string,
  input: CreateTravelInput,
): Promise<Travel> => {
  return apiFetch<Travel>("/api/travels", {
    method: "POST",
    token,
    body: input,
  });
};
