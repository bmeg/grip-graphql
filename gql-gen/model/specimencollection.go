
package model

import (
	"encoding/json"
	"fmt"
)

type SafeSpecimenCollection struct {
	Method *CodeableConcept `json:"method,omitempty"`
	Collector TypedObject `json:"collector"`
	CollectedDateTime *string `json:"collectedDateTime,omitempty"`
	BodySite *CodeableReference `json:"bodySite,omitempty"`
	CollectedPeriod *Period `json:"collectedPeriod,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Procedure *ProcedureType `json:"procedure"`
	Quantity *Quantity `json:"quantity,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	FastingStatusDuration *Duration `json:"fastingStatusDuration,omitempty"`
	FastingStatusCodeableConcept *CodeableConcept `json:"fastingStatusCodeableConcept,omitempty"`
	ID *string `json:"id,omitempty"`
	Device *CodeableReference `json:"device,omitempty"`
	Duration *Duration `json:"duration,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *SpecimenCollection) UnmarshalJSON(b []byte) error {
	var safe SafeSpecimenCollection
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = SpecimenCollection{
		Method: safe.Method,
		CollectedDateTime: safe.CollectedDateTime,
		BodySite: safe.BodySite,
		CollectedPeriod: safe.CollectedPeriod,
		Extension: safe.Extension,
		Procedure: safe.Procedure,
		Quantity: safe.Quantity,
		ModifierExtension: safe.ModifierExtension,
		FastingStatusDuration: safe.FastingStatusDuration,
		FastingStatusCodeableConcept: safe.FastingStatusCodeableConcept,
		ID: safe.ID,
		Device: safe.Device,
		Duration: safe.Duration,
		ResourceType: safe.ResourceType,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "collector", safe.Collector.Typename, &o.Collector); err != nil {
		return fmt.Errorf("failed to unmarshal Collector: %w", err)
	}

	return nil
}
