export interface BuyerAnalytics {
  purchases_count: number;
  energy_transacted: number;
  advantage_offer_uuid: string;
}

export interface SupplierAnalytics {
  money_earned: number;
  purchases_count: number;
  active_offers: number;
  almost_expiring_offers: number;
  user_price_avg: number;
  platform_price_avg: number;
}

export interface AnalyticsData {
  buyer_info?: BuyerAnalytics;
  supplier_info?: SupplierAnalytics;
}

export interface AnalyticsResponse {
  data: AnalyticsData;
}
