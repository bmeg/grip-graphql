
package model

import (
	"encoding/json"
	"fmt"
)

type SafeMedicationStatementType struct {
	Category []*CodeableConcept `json:"category,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	ID *string `json:"id,omitempty"`
	RenderedDosageInstruction *string `json:"renderedDosageInstruction,omitempty"`
	DateAsserted *string `json:"dateAsserted,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	Language *string `json:"language,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	EffectiveTiming *Timing `json:"effectiveTiming,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Medication *CodeableReference `json:"medication,omitempty"`
	Note TypedObject `json:"note"`
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`
	Status *string `json:"status,omitempty"`
	Dosage []*Dosage `json:"dosage,omitempty"`
	PartOf TypedObject `json:"partOf"`
	ResourceType *string `json:"resourceType,omitempty"`
	RelatedClinicalInformation TypedObject `json:"relatedClinicalInformation"`
	DerivedFrom TypedObject `json:"derivedFrom"`
	InformationSource TypedObject `json:"informationSource"`
	Subject TypedObject `json:"subject"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Adherence *MedicationStatementAdherence `json:"adherence,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *MedicationStatementType) UnmarshalJSON(b []byte) error {
	var safe SafeMedicationStatementType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = MedicationStatementType{
		Category: safe.Category,
		Meta: safe.Meta,
		ID: safe.ID,
		RenderedDosageInstruction: safe.RenderedDosageInstruction,
		DateAsserted: safe.DateAsserted,
		Reason: safe.Reason,
		Language: safe.Language,
		Text: safe.Text,
		EffectiveTiming: safe.EffectiveTiming,
		Identifier: safe.Identifier,
		Extension: safe.Extension,
		EffectivePeriod: safe.EffectivePeriod,
		ImplicitRules: safe.ImplicitRules,
		Medication: safe.Medication,
		EffectiveDateTime: safe.EffectiveDateTime,
		Status: safe.Status,
		Dosage: safe.Dosage,
		ResourceType: safe.ResourceType,
		ModifierExtension: safe.ModifierExtension,
		Adherence: safe.Adherence,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "partOf", safe.PartOf.Typename, &o.PartOf); err != nil {
		return fmt.Errorf("failed to unmarshal PartOf: %w", err)
	}
	if err := unmarshalUnion(b, "relatedClinicalInformation", safe.RelatedClinicalInformation.Typename, &o.RelatedClinicalInformation); err != nil {
		return fmt.Errorf("failed to unmarshal RelatedClinicalInformation: %w", err)
	}
	if err := unmarshalUnion(b, "derivedFrom", safe.DerivedFrom.Typename, &o.DerivedFrom); err != nil {
		return fmt.Errorf("failed to unmarshal DerivedFrom: %w", err)
	}
	if err := unmarshalUnion(b, "informationSource", safe.InformationSource.Typename, &o.InformationSource); err != nil {
		return fmt.Errorf("failed to unmarshal InformationSource: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}

	return nil
}
