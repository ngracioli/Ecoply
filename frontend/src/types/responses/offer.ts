import type { EnergyType } from "../offer";

export interface Offer {
  uuid: string;
  price_per_mwh: number;
  initial_quantity_mwh: number;
  remaining_quantity_mwh: number;
  description: string;
  period_start: string;
  period_end: string;
  status: number;
  energy_type: EnergyType;
  submarket: string;
  seller_agent_uuid: string;
  created_at: string;
}

export interface CreateOfferResponse {
  data: Offer;
}
