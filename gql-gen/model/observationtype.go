
package model

import (
	"encoding/json"
	"fmt"
)

type SafeObservationType struct {
	Note []*Annotation `json:"note,omitempty"`
	Focus TypedObject `json:"focus"`
	BasedOn TypedObject `json:"basedOn"`
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	ValueRange *Range `json:"valueRange,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	ValueString *string `json:"valueString,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Status *string `json:"status,omitempty"`
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	BodySite *CodeableConcept `json:"bodySite,omitempty"`
	HasMember TypedObject `json:"hasMember"`
	ID *string `json:"id,omitempty"`
	Component []*ObservationComponent `json:"component,omitempty"`
	Interpretation []*CodeableConcept `json:"interpretation,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	ReferenceRange []*ObservationReferenceRange `json:"referenceRange,omitempty"`
	EffectiveInstant *string `json:"effectiveInstant,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ValueSampledData *SampledData `json:"valueSampledData,omitempty"`
	Method *CodeableConcept `json:"method,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	DerivedFrom TypedObject `json:"derivedFrom"`
	Category []*CodeableConcept `json:"category,omitempty"`
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	DataAbsentReason *CodeableConcept `json:"dataAbsentReason,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Specimen TypedObject `json:"specimen"`
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	Issued *string `json:"issued,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	Language *string `json:"language,omitempty"`
	ValueTime *string `json:"valueTime,omitempty"`
	BodyStructure *BodyStructureType `json:"bodyStructure"`
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`
	EffectiveTiming *Timing `json:"effectiveTiming,omitempty"`
	Performer TypedObject `json:"performer"`
	Subject TypedObject `json:"subject"`
	ValueInteger *int `json:"valueInteger,omitempty"`
	PartOf TypedObject `json:"partOf"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	ValueDateTime *string `json:"valueDateTime,omitempty"`
	TriggeredBy []*ObservationTriggeredBy `json:"triggeredBy,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ObservationType) UnmarshalJSON(b []byte) error {
	var safe SafeObservationType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ObservationType{
		Note: safe.Note,
		EffectivePeriod: safe.EffectivePeriod,
		ValueRange: safe.ValueRange,
		ResourceType: safe.ResourceType,
		ValueString: safe.ValueString,
		Extension: safe.Extension,
		Status: safe.Status,
		ValuePeriod: safe.ValuePeriod,
		BodySite: safe.BodySite,
		ID: safe.ID,
		Component: safe.Component,
		Interpretation: safe.Interpretation,
		Text: safe.Text,
		ReferenceRange: safe.ReferenceRange,
		EffectiveInstant: safe.EffectiveInstant,
		ModifierExtension: safe.ModifierExtension,
		ValueSampledData: safe.ValueSampledData,
		Method: safe.Method,
		Meta: safe.Meta,
		Category: safe.Category,
		EffectiveDateTime: safe.EffectiveDateTime,
		ValueBoolean: safe.ValueBoolean,
		DataAbsentReason: safe.DataAbsentReason,
		ValueAttachment: safe.ValueAttachment,
		Issued: safe.Issued,
		ValueCodeableConcept: safe.ValueCodeableConcept,
		Language: safe.Language,
		ValueTime: safe.ValueTime,
		BodyStructure: safe.BodyStructure,
		InstantiatesCanonical: safe.InstantiatesCanonical,
		EffectiveTiming: safe.EffectiveTiming,
		ValueInteger: safe.ValueInteger,
		ImplicitRules: safe.ImplicitRules,
		ValueQuantity: safe.ValueQuantity,
		ValueDateTime: safe.ValueDateTime,
		TriggeredBy: safe.TriggeredBy,
		Code: safe.Code,
		Identifier: safe.Identifier,
		ValueRatio: safe.ValueRatio,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "focus", safe.Focus.Typename, &o.Focus); err != nil {
		return fmt.Errorf("failed to unmarshal Focus: %w", err)
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "hasMember", safe.HasMember.Typename, &o.HasMember); err != nil {
		return fmt.Errorf("failed to unmarshal HasMember: %w", err)
	}
	if err := unmarshalUnion(b, "derivedFrom", safe.DerivedFrom.Typename, &o.DerivedFrom); err != nil {
		return fmt.Errorf("failed to unmarshal DerivedFrom: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "specimen", safe.Specimen.Typename, &o.Specimen); err != nil {
		return fmt.Errorf("failed to unmarshal Specimen: %w", err)
	}
	if err := unmarshalUnion(b, "performer", safe.Performer.Typename, &o.Performer); err != nil {
		return fmt.Errorf("failed to unmarshal Performer: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "partOf", safe.PartOf.Typename, &o.PartOf); err != nil {
		return fmt.Errorf("failed to unmarshal PartOf: %w", err)
	}

	return nil
}
