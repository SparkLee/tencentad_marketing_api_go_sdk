package model

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestDynamicCreativeBrandsRoundTrip(t *testing.T) {
	const input = `{"creative_components":{"brand":[{"component_id":123}],"channels_brand":[{"value":{"jump_info":{"page_type":"PAGE_TYPE_WECHAT_CHANNELS_PROFILE","page_spec":{"wechat_channels_profile_spec":{"username":"test-channel"}}}}}],"wechat_channels":[{"value":{"username":"test-channel","finder_object_visibility":false}}]}}`
	var request DynamicCreativesAddRequest
	decoder := json.NewDecoder(strings.NewReader(input))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		t.Fatal(err)
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
		t.Fatalf("creative brand fields lost: %s", encoded)
	}
}
