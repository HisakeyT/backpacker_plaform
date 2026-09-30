import camelcaseKeys from "camelcase-keys";
import snakecaseKeys from "snakecase-keys";


type HTTPMethod = "GET" | "POST" | "PUT" | "DELETE" | "PATCH";

type ApiFetchOptions = {
  method: HTTPMethod;
  token?: string;
  body?: Record<string, unknown>;
};

export const apiFetch = async <T>(
  path: string,
  options: ApiFetchOptions,
): Promise<T> => {
  const { method, token, body } = options;

  const headers: Record<string, string> = {};
  if (token) {
    headers.Authorization = `Bearer ${token}`;
  }
  if (body) {
    headers["Content-Type"] = "application/json";
  }

  const response = await fetch(path, {
    method,
    headers,
    body: body ? JSON.stringify(snakecaseKeys(body, { deep: true })) : undefined,
  });

  if (!response.ok) {
    throw new Error(`Request failed: ${method} ${path} (${response.status})`);
  }

  if (response.status === 204) {
    return undefined as T;
  }

  const data = await response.json();
  return camelcaseKeys(data, { deep: true }) as T;
};
