export type UserType = "buyer" | "supplier";

export interface UserAddress {
  cep: string;
  state_initial: string;
  state: string;
  city: string;
  neighborhood: string;
  street: string;
  number: string;
  complement: string;
}

export interface UserAgent {
  cnpj: string;
  ccee_code: string;
  submarket_name: string;
  company_name: string;
}

export interface User {
  name: string;
  email: string;
  user_type: UserType;
  address: UserAddress;
  agent: UserAgent;
}
