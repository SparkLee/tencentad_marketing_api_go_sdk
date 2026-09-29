package model

import (
	"encoding/json"
	"strings"
	"testing"
)

// Uses the official goal enum; only checks local serialization of the fork's ROI field.
func TestSmartDeliveryDay30RoundTrip(t *testing.T) {
	const input = `{"smart_delivery_platform":"SMART_DELIVERY_PLATFORM_EDITION_MINI_GAME_PROMOTION","smart_delivery_scene_spec":{"smart_delivery_goal":"SMART_DELIVERY_GOAL_APP_REGISTER_30DAY_MONETIZATION_ROAS","conversion_id_list":[10018],"smart_delivery_goal_spec":{"mini_game_promotion_spec":{"register_cost":123,"day30_monetization_roi":1.234}}}}`
	var request AdgroupsAddRequest
	decoder := json.NewDecoder(strings.NewReader(input))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		t.Fatal(err)
	}
	if request.SmartDeliverySceneSpec.SmartDeliveryGoal != SmartDeliveryGoal_APP_REGISTER_30_DAY_MONETIZATION_ROAS {
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
