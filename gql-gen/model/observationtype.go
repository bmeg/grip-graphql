
package model

import (
	"encoding/json"
	"fmt"
)

type SafeObservationType struct {
	Meta *Meta `json:"meta,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	PartOf TypedObject `json:"partOf"`
	ValueString *string `json:"valueString,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	Issued *string `json:"issued,omitempty"`
	DerivedFrom TypedObject `json:"derivedFrom"`
	EffectiveTiming *Timing `json:"effectiveTiming,omitempty"`
	Focus TypedObject `json:"focus"`
	ValueBoolean *string `json:"valueBoolean,omitempty"`
	Status *string `json:"status,omitempty"`
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	Method *CodeableConcept `json:"method,omitempty"`
	TriggeredBy []*ObservationTriggeredBy `json:"triggeredBy,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Subject TypedObject `json:"subject"`
	DataAbsentReason *CodeableConcept `json:"dataAbsentReason,omitempty"`
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	ValueInteger *string `json:"valueInteger,omitempty"`
	Interpretation []*CodeableConcept `json:"interpretation,omitempty"`
	ValueTime *string `json:"valueTime,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	ReferenceRange []*ObservationReferenceRange `json:"referenceRange,omitempty"`
	Component []*ObservationComponent `json:"component,omitempty"`
	ID *string `json:"id,omitempty"`
	EffectiveInstant *string `json:"effectiveInstant,omitempty"`
	Language *string `json:"language,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`
	HasMember TypedObject `json:"hasMember"`
	Category []*CodeableConcept `json:"category,omitempty"`
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	BodyStructure *BodyStructureType `json:"bodyStructure"`
	Performer TypedObject `json:"performer"`
	ResourceType *string `json:"resourceType,omitempty"`
	ValueRange *Range `json:"valueRange,omitempty"`
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Specimen TypedObject `json:"specimen"`
	Text *Narrative `json:"text,omitempty"`
	BodySite *CodeableConcept `json:"bodySite,omitempty"`
	ValueDateTime *string `json:"valueDateTime,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	ValueSampledData *SampledData `json:"valueSampledData,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ObservationType) UnmarshalJSON(b []byte) error {
	var safe SafeObservationType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ObservationType{
		Meta: safe.Meta,
		Identifier: safe.Identifier,
		ValueString: safe.ValueString,
		Extension: safe.Extension,
		ValueQuantity: safe.ValueQuantity,
		Issued: safe.Issued,
		EffectiveTiming: safe.EffectiveTiming,
		ValueBoolean: safe.ValueBoolean,
		Status: safe.Status,
		EffectivePeriod: safe.EffectivePeriod,
		Method: safe.Method,
		TriggeredBy: safe.TriggeredBy,
		ImplicitRules: safe.ImplicitRules,
		DataAbsentReason: safe.DataAbsentReason,
		ValueAttachment: safe.ValueAttachment,
		ValueInteger: safe.ValueInteger,
		Interpretation: safe.Interpretation,
		ValueTime: safe.ValueTime,
		ReferenceRange: safe.ReferenceRange,
		Component: safe.Component,
		ID: safe.ID,
		EffectiveInstant: safe.EffectiveInstant,
		Language: safe.Language,
		ModifierExtension: safe.ModifierExtension,
		EffectiveDateTime: safe.EffectiveDateTime,
		InstantiatesCanonical: safe.InstantiatesCanonical,
		Category: safe.Category,
		ValueRatio: safe.ValueRatio,
		BodyStructure: safe.BodyStructure,
		ResourceType: safe.ResourceType,
		ValueRange: safe.ValueRange,
		ValuePeriod: safe.ValuePeriod,
		Code: safe.Code,
		Text: safe.Text,
		BodySite: safe.BodySite,
		ValueDateTime: safe.ValueDateTime,
		ValueCodeableConcept: safe.ValueCodeableConcept,
		Note: safe.Note,
		ValueSampledData: safe.ValueSampledData,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "partOf", safe.PartOf.Typename, &o.PartOf); err != nil {
		return fmt.Errorf("failed to unmarshal PartOf: %w", err)
	}
	if err := unmarshalUnion(b, "derivedFrom", safe.DerivedFrom.Typename, &o.DerivedFrom); err != nil {
		return fmt.Errorf("failed to unmarshal DerivedFrom: %w", err)
	}
	if err := unmarshalUnion(b, "focus", safe.Focus.Typename, &o.Focus); err != nil {
		return fmt.Errorf("failed to unmarshal Focus: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "hasMember", safe.HasMember.Typename, &o.HasMember); err != nil {
		return fmt.Errorf("failed to unmarshal HasMember: %w", err)
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "performer", safe.Performer.Typename, &o.Performer); err != nil {
		return fmt.Errorf("failed to unmarshal Performer: %w", err)
	}
	if err := unmarshalUnion(b, "specimen", safe.Specimen.Typename, &o.Specimen); err != nil {
		return fmt.Errorf("failed to unmarshal Specimen: %w", err)
	}

	return nil
}
