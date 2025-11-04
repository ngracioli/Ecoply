import type { AuthState } from "./modules/auth";
import type { UserState } from "./modules/user";

export interface RootState {
  auth: AuthState;
  user: UserState;
}
