// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json vmaster
// - protoc             v3.9.1
// source: lorawan-stack/api/networkserver.proto

package ttnpb

import (
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the GetDefaultMACSettingsRequest message to JSON.
func (x *GetDefaultMACSettingsRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.FrequencyPlanId != "" {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("frequency_plan_id")
		s.WriteString(x.FrequencyPlanId)
	}
	if x.LorawanPhyVersion != 0 {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("lorawan_phy_version")
		x.LorawanPhyVersion.MarshalProtoJSON(s)
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the GetDefaultMACSettingsRequest message from JSON.
func (x *GetDefaultMACSettingsRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "frequency_plan_id", "frequencyPlanId":
			x.FrequencyPlanId = s.ReadString()
		case "lorawan_phy_version", "lorawanPhyVersion":
			x.LorawanPhyVersion.UnmarshalProtoJSON(s)
		}
	})
}
