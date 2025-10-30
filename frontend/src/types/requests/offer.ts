import type { EnergyType } from "../offer";

export interface CreateOfferRequest {
  price_per_mwh: number;
  quantity_mwh: number;
  period_start: string;
  period_end: string;
  description: string;
  energy_type: EnergyType;
}
