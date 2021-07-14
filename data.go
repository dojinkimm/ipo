package ipo

type IpoMessage struct {
	// 종목명
	ItemName string `json:"item_name"`
	// 종목코드
	TickerCode string `json:"ticker_code"`
	// 공모 청약 시작일
	StartDateMs int64 `json:"start_date_ms"`
	// 공모 청약 종료일
	EndDateMs int64 `json:"end_date_ms"`
	// 환불일
	RefundDateMs int64 `json:"refund_date_ms"`
	// 상장일
	ListingDateMs int64 `json:"listing_date_ms"`
	// 확정공모가
	FinalOfferingPrice int64 `json:"final_offering_price"`
	// 희망공모가 하한
	OfferingRangeLowerLimit int64 `json:"offering_range_lower_limit"`
	// 희망공모가 상한
	OfferingRangeHigherLimit int64 `json:"offering_range_higher_limit"`
	// 경쟁률
	CompetitiveRate2f int64 `json:"competitive_rate_2f"`
	// 주간사
	LeadManagers []string `json:"lead_manager"`
}

var data = []IpoMessage{
	{
		ItemName:                 "에스디바이오센서",
		TickerCode:               "137310",
		StartDateMs:              1625702400000, // 2021.07.08
		EndDateMs:                1625788800000, // 2021.07.09
		RefundDateMs:             1626134400000, // 2021.07.13
		ListingDateMs:            1626393600000, // 2021.07.16
		FinalOfferingPrice:       52000,
		OfferingRangeLowerLimit:  52000,
		OfferingRangeHigherLimit: 45000,
		CompetitiveRate2f:        27402,
		LeadManagers: []string{
			"NH투자증권", "한국투자증권", "삼성증권", "KB증권",
		},
	},
}
