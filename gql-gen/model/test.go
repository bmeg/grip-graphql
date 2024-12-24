package model

import (
	"encoding/json"
	"fmt"
)

type TypedObject struct {
	Typename string `json:"__typename"`
}

type SafeObservationType struct {
	EffectivePeriod       *Period                      `json:"effectivePeriod,omitempty"`
	EffectiveInstant      *string                      `json:"effectiveInstant,omitempty"`
	ValuePeriod           *Period                      `json:"valuePeriod,omitempty"`
	InstantiatesCanonical *string                      `json:"instantiatesCanonical,omitempty"`
	ReferenceRange        []*ObservationReferenceRange `json:"referenceRange,omitempty"`
	Code                  *CodeableConcept             `json:"code,omitempty"`
	Category              []*CodeableConcept           `json:"category,omitempty"`
	ValueBoolean          *string                      `json:"valueBoolean,omitempty"`
	BodySite              *CodeableConcept             `json:"bodySite,omitempty"`
	Text                  *Narrative                   `json:"text,omitempty"`
	ValueQuantity         *Quantity                    `json:"valueQuantity,omitempty"`
	ValueString           *string                      `json:"valueString,omitempty"`
	ID                    *string                      `json:"id,omitempty"`
	ModifierExtension     []*Extension                 `json:"modifierExtension,omitempty"`
	ValueAttachment       *Attachment                  `json:"valueAttachment,omitempty"`
	Issued                *string                      `json:"issued,omitempty"`
	ImplicitRules         *string                      `json:"implicitRules,omitempty"`
	Meta                  *Meta                        `json:"meta,omitempty"`
	Method                *CodeableConcept             `json:"method,omitempty"`
	BodyStructure         []*BodyStructureType         `json:"bodyStructure,omitempty"`
	EffectiveDateTime     *string                      `json:"effectiveDateTime,omitempty"`
	ValueRatio            *Ratio                       `json:"valueRatio,omitempty"`
	Interpretation        []*CodeableConcept           `json:"interpretation,omitempty"`
	ValueDateTime         *string                      `json:"valueDateTime,omitempty"`
	ValueTime             *string                      `json:"valueTime,omitempty"`
	DataAbsentReason      *CodeableConcept             `json:"dataAbsentReason,omitempty"`
	ResourceType          *string                      `json:"resourceType,omitempty"`
	ValueCodeableConcept  *CodeableConcept             `json:"valueCodeableConcept,omitempty"`
	ValueInteger          *string                      `json:"valueInteger,omitempty"`
	Language              *string                      `json:"language,omitempty"`
	Identifier            []*Identifier                `json:"identifier,omitempty"`
	ValueSampledData      *SampledData                 `json:"valueSampledData,omitempty"`
	Extension             []*Extension                 `json:"extension,omitempty"`
	TriggeredBy           []*ObservationTriggeredBy    `json:"triggeredBy,omitempty"`
	Component             []*ObservationComponent      `json:"component,omitempty"`
	EffectiveTiming       *Timing                      `json:"effectiveTiming,omitempty"`
	ValueRange            *Range                       `json:"valueRange,omitempty"`
	Status                *string                      `json:"status,omitempty"`
	Focus                 TypedObject                  `json:"focus,omitempty"`
	Subject               TypedObject                  `json:"subject,omitempty"`
}

func (o *ObservationType) UnmarshalJSON(b []byte) error {
	var safe SafeObservationType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}
	o.ID = safe.ID
	fmt.Printf("SAFE: %#v\n", safe.ID)

	switch safe.Focus.Typename {
	case "PatientType":
		var partial struct {
			Focus PatientType `json:"focus"`
		}
		if err := json.Unmarshal(b, &partial); err != nil {
			return err
		}
		o.Focus = partial.Focus
	case "SpecimenType":
		var partial struct {
			Focus SpecimenType `json:"focus"`
		}
		if err := json.Unmarshal(b, &partial); err != nil {
			return err
		}
		o.Focus = partial.Focus
	}

	switch safe.Subject.Typename {
	case "PatientType":
		var partial struct {
			Subject PatientType `json:"subject"`
		}
		if err := json.Unmarshal(b, &partial); err != nil {
			return err
		}
		o.Subject = partial.Subject
	case "MedicationType":
		var partial struct {
			Subject MedicationType `json:"subject"`
		}
		if err := json.Unmarshal(b, &partial); err != nil {
			return err
		}
		o.Subject = partial.Subject
	}

	return nil
}
