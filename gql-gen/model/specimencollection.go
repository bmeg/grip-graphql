
package model

import (
	"encoding/json"
	"fmt"
)

type SafeSpecimenCollection struct {
	BodySite *CodeableReference `json:"bodySite,omitempty"`
	Method *CodeableConcept `json:"method,omitempty"`
	ID *string `json:"id,omitempty"`
	Procedure *ProcedureType `json:"procedure"`
	CollectedPeriod *Period `json:"collectedPeriod,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Quantity *Quantity `json:"quantity,omitempty"`
	FastingStatusDuration *Duration `json:"fastingStatusDuration,omitempty"`
	Duration *Duration `json:"duration,omitempty"`
	Collector TypedObject `json:"collector"`
	Device *CodeableReference `json:"device,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	FastingStatusCodeableConcept *CodeableConcept `json:"fastingStatusCodeableConcept,omitempty"`
	CollectedDateTime *string `json:"collectedDateTime,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *SpecimenCollection) UnmarshalJSON(b []byte) error {
	var safe SafeSpecimenCollection
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = SpecimenCollection{
		BodySite: safe.BodySite,
		Method: safe.Method,
		ID: safe.ID,
		Procedure: safe.Procedure,
		CollectedPeriod: safe.CollectedPeriod,
		ModifierExtension: safe.ModifierExtension,
		Quantity: safe.Quantity,
		FastingStatusDuration: safe.FastingStatusDuration,
		Duration: safe.Duration,
		Device: safe.Device,
		Extension: safe.Extension,
		ResourceType: safe.ResourceType,
		FastingStatusCodeableConcept: safe.FastingStatusCodeableConcept,
		CollectedDateTime: safe.CollectedDateTime,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "collector", safe.Collector.Typename, &o.Collector); err != nil {
		return fmt.Errorf("failed to unmarshal Collector: %w", err)
	}

	return nil
}
