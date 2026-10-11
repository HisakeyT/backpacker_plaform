export type User = {
  id: number;
  nickname: string;
  email: string;
};

export type LoginRequest = {
  email: string;
  password: string;
};

export type LoginResponse = {
  token: string;
  user: User;
};

export type RegisterRequest = {
  nickname: string;
  email: string;
  password: string;
  passwordConfirmation: string;
};

export type LoginLocationState = {
  from?: string; // ログイン後に戻るパス
} | null;
