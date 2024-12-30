
package model

import (
	"encoding/json"
	"fmt"
)

type SafeMedicationStatementType struct {
	Text *Narrative `json:"text,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Medication *CodeableReference `json:"medication,omitempty"`
	Note TypedObject `json:"note"`
	Category []*CodeableConcept `json:"category,omitempty"`
	ID *string `json:"id,omitempty"`
	InformationSource TypedObject `json:"informationSource"`
	RelatedClinicalInformation TypedObject `json:"relatedClinicalInformation"`
	EffectiveTiming *Timing `json:"effectiveTiming,omitempty"`
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`
	Subject TypedObject `json:"subject"`
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	DerivedFrom TypedObject `json:"derivedFrom"`
	Adherence *MedicationStatementAdherence `json:"adherence,omitempty"`
	Language *string `json:"language,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	RenderedDosageInstruction *string `json:"renderedDosageInstruction,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Dosage []*Dosage `json:"dosage,omitempty"`
	Status *string `json:"status,omitempty"`
	DateAsserted *string `json:"dateAsserted,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	PartOf TypedObject `json:"partOf"`
}

func (o *MedicationStatementType) UnmarshalJSON(b []byte) error {
	var safe SafeMedicationStatementType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = MedicationStatementType{
		Text: safe.Text,
		Identifier: safe.Identifier,
		Reason: safe.Reason,
		ImplicitRules: safe.ImplicitRules,
		Medication: safe.Medication,
		Category: safe.Category,
		ID: safe.ID,
		EffectiveTiming: safe.EffectiveTiming,
		EffectiveDateTime: safe.EffectiveDateTime,
		EffectivePeriod: safe.EffectivePeriod,
		ModifierExtension: safe.ModifierExtension,
		Adherence: safe.Adherence,
		Language: safe.Language,
		ResourceType: safe.ResourceType,
		RenderedDosageInstruction: safe.RenderedDosageInstruction,
		Meta: safe.Meta,
		Dosage: safe.Dosage,
		Status: safe.Status,
		DateAsserted: safe.DateAsserted,
		Extension: safe.Extension,
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "informationSource", safe.InformationSource.Typename, &o.InformationSource); err != nil {
		return fmt.Errorf("failed to unmarshal InformationSource: %w", err)
	}
	if err := unmarshalUnion(b, "relatedClinicalInformation", safe.RelatedClinicalInformation.Typename, &o.RelatedClinicalInformation); err != nil {
		return fmt.Errorf("failed to unmarshal RelatedClinicalInformation: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "derivedFrom", safe.DerivedFrom.Typename, &o.DerivedFrom); err != nil {
		return fmt.Errorf("failed to unmarshal DerivedFrom: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "partOf", safe.PartOf.Typename, &o.PartOf); err != nil {
		return fmt.Errorf("failed to unmarshal PartOf: %w", err)
	}

	return nil
}
