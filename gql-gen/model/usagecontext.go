
package model

import (
	"encoding/json"
	"fmt"
)

type SafeUsageContext struct {
	ID *string `json:"id,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	ValueRange *Range `json:"valueRange,omitempty"`
	ValueReference TypedObject `json:"valueReference"`
	Code *Coding `json:"code,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *UsageContext) UnmarshalJSON(b []byte) error {
	var safe SafeUsageContext
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = UsageContext{
		ID: safe.ID,
		ResourceType: safe.ResourceType,
		ValueCodeableConcept: safe.ValueCodeableConcept,
		ValueQuantity: safe.ValueQuantity,
		ValueRange: safe.ValueRange,
		Code: safe.Code,
		Extension: safe.Extension,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "valueReference", safe.ValueReference.Typename, &o.ValueReference); err != nil {
		return fmt.Errorf("failed to unmarshal ValueReference: %w", err)
	}

	return nil
}
