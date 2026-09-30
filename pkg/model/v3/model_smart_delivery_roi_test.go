package model

import (
	"encoding/json"
	"strings"
	"testing"
)

// Standard AIM bidding fields, without deprecated scene-level bidding.
func TestSmartDeliveryDay30RoundTrip(t *testing.T) {
	const input = `{"smart_delivery_platform":"SMART_DELIVERY_PLATFORM_EDITION_MINI_GAME_PROMOTION","conversion_id":10018,"bid_amount":123,"deep_conversion_worth_rate":1.234}`
	var request AdgroupsAddRequest
	decoder := json.NewDecoder(strings.NewReader(input))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		t.Fatal(err)
	}
	if request.ConversionId == nil || *request.ConversionId != 10018 {
		t.Fatal("incorrect goal")
	}
	encoded, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	var want, got interface{}
	_ = json.Unmarshal([]byte(input), &want)
	_ = json.Unmarshal(encoded, &got)
	wantJSON, _ := json.Marshal(want)
	gotJSON, _ := json.Marshal(got)
	if string(wantJSON) != string(gotJSON) {
		t.Fatalf("ROI fields lost: %s", encoded)
	}
}
