import type { Travel } from "./types";
import { apiFetch } from "../../lib/apiFetch";

export const getTravels = async (token: string): Promise<Travel[]> => {
  return apiFetch<Travel[]>("/api/travels", { token, method: "GET" })
}

