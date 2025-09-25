export interface RegisterFirstStepProps {
  name: string;
  email: string;
  password: string;
  confirmPassword: string;
}

export interface RegisterSecondStepProps {
  address: string;
  city: string;
  zip: string;
}

export interface RegisterRequest {
    name: string;
    email: string;
    password: string;
    address: string;
    city: string;
    zip: string;
}
