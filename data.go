package ipo

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type IpoMessage struct {
	// 종목명
	ItemName string `json:"item_name"`
	// 종목코드
	TickerCode string `json:"ticker_code"`
	// 공모 청약 시작일
	StartDate string `json:"start_date"`
	// 공모 청약 종료일
	EndDate string `json:"end_date"`
	// 확정공모가
	FinalOfferingPrice *int64 `json:"final_offering_price"`
	// 희망공모가 하한
	OfferingRangeLowerLimit int64 `json:"offering_range_lower_limit"`
	// 희망공모가 상한
	OfferingRangeHigherLimit int64 `json:"offering_range_higher_limit"`
	// 경쟁률
	CompetitiveRate2f *int64 `json:"competitive_rate_2f"`
	// 주간사
	LeadManagers []string `json:"lead_manager"`
}

const fileName = "ipos.json"

func GetIpoData() ([]*IpoMessage, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var ipos []*IpoMessage
	if err := json.Unmarshal(byteValue, &ipos); err != nil {
		return nil, err
	}

	return ipos, nil
}
