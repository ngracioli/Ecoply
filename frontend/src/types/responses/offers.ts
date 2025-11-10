import type { EnergyType } from "../offer";

export interface OfferListItem {
  uuid: string;
  price_per_mwh: number;
  initial_quantity_mwh: number;
  remaining_quantity_mwh: number;
  description: string;
  period_start: string;
  period_end: string;
  status: string;
  energy_type: EnergyType;
  submarket: string;
  seller_agent_uuid: string;
  created_at: string;
}

export interface OfferDetailResponse {
  data: OfferListItem;
}

export interface PaginatedResponse<T> {
  page: number;
  page_size: number;
  has_next: boolean;
  has_prev: boolean;
  data: T[];
}

export interface OffersListResponse extends PaginatedResponse<OfferListItem> {}
