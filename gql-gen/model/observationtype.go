
package model

import (
	"encoding/json"
	"fmt"
)

type SafeObservationType struct {
	DataAbsentReason *CodeableConcept `json:"dataAbsentReason,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	ValueSampledData *SampledData `json:"valueSampledData,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	TriggeredBy []*ObservationTriggeredBy `json:"triggeredBy,omitempty"`
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	ValueRange *Range `json:"valueRange,omitempty"`
	ValueInteger *string `json:"valueInteger,omitempty"`
	ValueString *string `json:"valueString,omitempty"`
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	BodySite *CodeableConcept `json:"bodySite,omitempty"`
	Focus TypedObject `json:"focus"`
	ReferenceRange []*ObservationReferenceRange `json:"referenceRange,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	ValueDateTime *string `json:"valueDateTime,omitempty"`
	Specimen TypedObject `json:"specimen"`
	ID *string `json:"id,omitempty"`
	Issued *string `json:"issued,omitempty"`
	HasMember TypedObject `json:"hasMember"`
	Note TypedObject `json:"note"`
	Interpretation []*CodeableConcept `json:"interpretation,omitempty"`
	Status *string `json:"status,omitempty"`
	BodyStructure *BodyStructureType `json:"bodyStructure"`
	ValueTime *string `json:"valueTime,omitempty"`
	Language *string `json:"language,omitempty"`
	Method *CodeableConcept `json:"method,omitempty"`
	EffectiveInstant *string `json:"effectiveInstant,omitempty"`
	EffectiveTiming *Timing `json:"effectiveTiming,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Performer TypedObject `json:"performer"`
	Component []*ObservationComponent `json:"component,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	ValueBoolean *string `json:"valueBoolean,omitempty"`
	Subject TypedObject `json:"subject"`
	ResourceType *string `json:"resourceType,omitempty"`
	PartOf TypedObject `json:"partOf"`
	Extension []*Extension `json:"extension,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`
	DerivedFrom TypedObject `json:"derivedFrom"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ObservationType) UnmarshalJSON(b []byte) error {
	var safe SafeObservationType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ObservationType{
		DataAbsentReason: safe.DataAbsentReason,
		Text: safe.Text,
		Code: safe.Code,
		ValuePeriod: safe.ValuePeriod,
		ValueRatio: safe.ValueRatio,
		ValueSampledData: safe.ValueSampledData,
		ModifierExtension: safe.ModifierExtension,
		TriggeredBy: safe.TriggeredBy,
		ValueAttachment: safe.ValueAttachment,
		ValueQuantity: safe.ValueQuantity,
		EffectivePeriod: safe.EffectivePeriod,
		ValueRange: safe.ValueRange,
		ValueInteger: safe.ValueInteger,
		ValueString: safe.ValueString,
		EffectiveDateTime: safe.EffectiveDateTime,
		ValueCodeableConcept: safe.ValueCodeableConcept,
		BodySite: safe.BodySite,
		ReferenceRange: safe.ReferenceRange,
		ValueDateTime: safe.ValueDateTime,
		ID: safe.ID,
		Issued: safe.Issued,
		Interpretation: safe.Interpretation,
		Status: safe.Status,
		BodyStructure: safe.BodyStructure,
		ValueTime: safe.ValueTime,
		Language: safe.Language,
		Method: safe.Method,
		EffectiveInstant: safe.EffectiveInstant,
		EffectiveTiming: safe.EffectiveTiming,
		ImplicitRules: safe.ImplicitRules,
		Component: safe.Component,
		Meta: safe.Meta,
		ValueBoolean: safe.ValueBoolean,
		ResourceType: safe.ResourceType,
		Extension: safe.Extension,
		Category: safe.Category,
		Identifier: safe.Identifier,
		InstantiatesCanonical: safe.InstantiatesCanonical,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "focus", safe.Focus.Typename, &o.Focus); err != nil {
		return fmt.Errorf("failed to unmarshal Focus: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "specimen", safe.Specimen.Typename, &o.Specimen); err != nil {
		return fmt.Errorf("failed to unmarshal Specimen: %w", err)
	}
	if err := unmarshalUnion(b, "hasMember", safe.HasMember.Typename, &o.HasMember); err != nil {
		return fmt.Errorf("failed to unmarshal HasMember: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
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
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "derivedFrom", safe.DerivedFrom.Typename, &o.DerivedFrom); err != nil {
		return fmt.Errorf("failed to unmarshal DerivedFrom: %w", err)
	}

	return nil
}
