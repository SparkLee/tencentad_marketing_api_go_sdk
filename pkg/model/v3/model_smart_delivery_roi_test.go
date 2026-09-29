package model

import (
	"encoding/json"
	"strings"
	"testing"
)

// This checks local serialization, not remote acceptance of the DAY30 extension.
func TestSmartDeliveryDay30RoundTrip(t *testing.T) {
	const input = `{"smart_delivery_platform":"SMART_DELIVERY_PLATFORM_EDITION_MINI_GAME_PROMOTION","smart_delivery_scene_spec":{"smart_delivery_goal":"SMART_DELIVERY_GOAL_DAY30_MONETIZATION","conversion_id_list":[10018],"smart_delivery_goal_spec":{"mini_game_promotion_spec":{"register_cost":123,"day30_monetization_roi":1.234}}}}`
	var request AdgroupsAddRequest
	decoder := json.NewDecoder(strings.NewReader(input))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		t.Fatal(err)
	}
	if request.SmartDeliverySceneSpec.SmartDeliveryGoal != SmartDeliveryGoal_DAY30_MONETIZATION {
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
