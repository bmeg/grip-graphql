
package model

import (
	"encoding/json"
	"fmt"
)

type SafeObservationType struct {
	BodyStructure *BodyStructureType `json:"bodyStructure"`
	DataAbsentReason *CodeableConcept `json:"dataAbsentReason,omitempty"`
	EffectiveInstant *string `json:"effectiveInstant,omitempty"`
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Issued *string `json:"issued,omitempty"`
	Method *CodeableConcept `json:"method,omitempty"`
	ValueTime *string `json:"valueTime,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	ValueInteger *string `json:"valueInteger,omitempty"`
	ReferenceRange []*ObservationReferenceRange `json:"referenceRange,omitempty"`
	Note TypedObject `json:"note"`
	Text *Narrative `json:"text,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	ValueDateTime *string `json:"valueDateTime,omitempty"`
	ValueString *string `json:"valueString,omitempty"`
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`
	HasMember TypedObject `json:"hasMember"`
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`
	Focus TypedObject `json:"focus"`
	Performer TypedObject `json:"performer"`
	ValueSampledData *SampledData `json:"valueSampledData,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Language *string `json:"language,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	PartOf TypedObject `json:"partOf"`
	Status *string `json:"status,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	TriggeredBy []*ObservationTriggeredBy `json:"triggeredBy,omitempty"`
	Subject TypedObject `json:"subject"`
	DerivedFrom TypedObject `json:"derivedFrom"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	BodySite *CodeableConcept `json:"bodySite,omitempty"`
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	Component []*ObservationComponent `json:"component,omitempty"`
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	ValueRange *Range `json:"valueRange,omitempty"`
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	Specimen TypedObject `json:"specimen"`
	EffectiveTiming *Timing `json:"effectiveTiming,omitempty"`
	ID *string `json:"id,omitempty"`
	Interpretation []*CodeableConcept `json:"interpretation,omitempty"`
	ValueBoolean *string `json:"valueBoolean,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
}

func (o *ObservationType) UnmarshalJSON(b []byte) error {
	var safe SafeObservationType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ObservationType{
		BodyStructure: safe.BodyStructure,
		DataAbsentReason: safe.DataAbsentReason,
		EffectiveInstant: safe.EffectiveInstant,
		EffectivePeriod: safe.EffectivePeriod,
		Identifier: safe.Identifier,
		Meta: safe.Meta,
		Issued: safe.Issued,
		Method: safe.Method,
		ValueTime: safe.ValueTime,
		ValueInteger: safe.ValueInteger,
		ReferenceRange: safe.ReferenceRange,
		Text: safe.Text,
		ResourceType: safe.ResourceType,
		ValueDateTime: safe.ValueDateTime,
		ValueString: safe.ValueString,
		ValueQuantity: safe.ValueQuantity,
		EffectiveDateTime: safe.EffectiveDateTime,
		InstantiatesCanonical: safe.InstantiatesCanonical,
		ValueSampledData: safe.ValueSampledData,
		ImplicitRules: safe.ImplicitRules,
		Language: safe.Language,
		ValueCodeableConcept: safe.ValueCodeableConcept,
		Code: safe.Code,
		Status: safe.Status,
		Category: safe.Category,
		TriggeredBy: safe.TriggeredBy,
		ModifierExtension: safe.ModifierExtension,
		BodySite: safe.BodySite,
		ValueRatio: safe.ValueRatio,
		Component: safe.Component,
		ValueAttachment: safe.ValueAttachment,
		ValueRange: safe.ValueRange,
		ValuePeriod: safe.ValuePeriod,
		EffectiveTiming: safe.EffectiveTiming,
		ID: safe.ID,
		Interpretation: safe.Interpretation,
		ValueBoolean: safe.ValueBoolean,
		Extension: safe.Extension,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "hasMember", safe.HasMember.Typename, &o.HasMember); err != nil {
		return fmt.Errorf("failed to unmarshal HasMember: %w", err)
	}
	if err := unmarshalUnion(b, "focus", safe.Focus.Typename, &o.Focus); err != nil {
		return fmt.Errorf("failed to unmarshal Focus: %w", err)
	}
	if err := unmarshalUnion(b, "performer", safe.Performer.Typename, &o.Performer); err != nil {
		return fmt.Errorf("failed to unmarshal Performer: %w", err)
	}
	if err := unmarshalUnion(b, "partOf", safe.PartOf.Typename, &o.PartOf); err != nil {
		return fmt.Errorf("failed to unmarshal PartOf: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "derivedFrom", safe.DerivedFrom.Typename, &o.DerivedFrom); err != nil {
		return fmt.Errorf("failed to unmarshal DerivedFrom: %w", err)
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "specimen", safe.Specimen.Typename, &o.Specimen); err != nil {
		return fmt.Errorf("failed to unmarshal Specimen: %w", err)
	}

	return nil
}
