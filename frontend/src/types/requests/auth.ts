export interface RegisterRequest {
  name: string;
  email: string;
  password: string;
  confirm_password: string;
  user_type: "buyer" | "supplier";
  address: {
    cep: string;
    state_initials: string;
    state: string;
    city: string;
    neighborhood: string;
    street: string;
    number: string;
    complement?: string;
  };
  agent: {
    cnpj: string;
    ccee_code: string;
    submarket_name: string | null;
    company_name: string;
  };
}
