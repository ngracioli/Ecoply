import type { User } from "../user";

export interface AuthResponse {
  data: {
    token: string;
    user: User;
  };
}

export interface UserMeResponse {
  data: User;
}
