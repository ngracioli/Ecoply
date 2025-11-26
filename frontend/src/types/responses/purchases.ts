export interface PurchaseListItem {
  uuid: string;
  quantity_mwh: number;
  price_per_mwh: number;
  status: "completed" | "waiting" | "canceled";
  payment_method: "pix" | "card" | "billet";
  offer_uuid: string;
  buyer_name: string;
  seller_name: string;
  seller_uuid: string;
  created_at: string;
}

export interface PaginatedResponse<T> {
  page: number;
  page_size: number;
  has_next: boolean;
  has_prev: boolean;
  data: T[];
}

export interface PurchasesListResponse
  extends PaginatedResponse<PurchaseListItem> {}

export interface ContractSupplier {
  uuid: string;
  cnpj: string;
  ccee_code: string;
  submarket_name: string;
  company_name: string;
}

export interface ContractOffer {
  uuid: string;
  price_per_mwh: number;
  initial_quantity_mwh: number;
  contracted_quantity_mwh: number;
  description: string;
  period_start: string;
  period_end: string;
  energy_type: string;
  submarket: string;
  created_at: string;
}

export interface ContractBuyer {
  uuid: string;
  cnpj: string;
  ccee_code: string;
  submarket_name: string;
  company_name: string;
}

export interface ContractData {
  supplier: ContractSupplier;
  offer: ContractOffer;
  buyer: ContractBuyer;
}

export interface PurchaseContractResponse {
  data: ContractData;
}
