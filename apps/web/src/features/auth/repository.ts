import { apiFetch } from "../../lib/apiFetch";
import type { LoginRequest, LoginResponse, User } from "./types";

export const login = async (input: LoginRequest): Promise<LoginResponse> => {
  return apiFetch<LoginResponse>("/api/login", {
    method: "POST",
    body: input,
  });
};

export const getMe = async (token: string): Promise<User> => {
  return apiFetch<User>("/api/users/me", { token, method: "GET" })
};
