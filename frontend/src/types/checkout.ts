export type PaymentMethod = "pix" | "credit_card" | "boleto";

export interface CheckoutFormData {
  quantity_mwh: number;
  payment_method: PaymentMethod;
  installments?: number;
}

export interface CardDetails {
  number: string;
  holder_name: string;
  expiry_date: string;
  cvv: string;
}
