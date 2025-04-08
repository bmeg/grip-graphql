
package model

import (
	"encoding/json"
	"fmt"
)

type SafeMedicationStatementType struct {
	Status *string `json:"status,omitempty"`
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	RelatedClinicalInformation TypedObject `json:"relatedClinicalInformation"`
	Note []*Annotation `json:"note,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Medication *CodeableReference `json:"medication,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	ID *string `json:"id,omitempty"`
	EffectiveTiming *Timing `json:"effectiveTiming,omitempty"`
	Adherence *MedicationStatementAdherence `json:"adherence,omitempty"`
	RenderedDosageInstruction *string `json:"renderedDosageInstruction,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	PartOf TypedObject `json:"partOf"`
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Subject TypedObject `json:"subject"`
	InformationSource TypedObject `json:"informationSource"`
	DerivedFrom TypedObject `json:"derivedFrom"`
	Dosage []*Dosage `json:"dosage,omitempty"`
	Language *string `json:"language,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	DateAsserted *string `json:"dateAsserted,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *MedicationStatementType) UnmarshalJSON(b []byte) error {
	var safe SafeMedicationStatementType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = MedicationStatementType{
		Status: safe.Status,
		EffectivePeriod: safe.EffectivePeriod,
		Note: safe.Note,
		ModifierExtension: safe.ModifierExtension,
		Medication: safe.Medication,
		Extension: safe.Extension,
		ResourceType: safe.ResourceType,
		ID: safe.ID,
		EffectiveTiming: safe.EffectiveTiming,
		Adherence: safe.Adherence,
		RenderedDosageInstruction: safe.RenderedDosageInstruction,
		Identifier: safe.Identifier,
		Text: safe.Text,
		EffectiveDateTime: safe.EffectiveDateTime,
		Meta: safe.Meta,
		Category: safe.Category,
		Dosage: safe.Dosage,
		Language: safe.Language,
		Reason: safe.Reason,
		ImplicitRules: safe.ImplicitRules,
		DateAsserted: safe.DateAsserted,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "relatedClinicalInformation", safe.RelatedClinicalInformation.Typename, &o.RelatedClinicalInformation); err != nil {
		return fmt.Errorf("failed to unmarshal RelatedClinicalInformation: %w", err)
	}
	if err := unmarshalUnion(b, "partOf", safe.PartOf.Typename, &o.PartOf); err != nil {
		return fmt.Errorf("failed to unmarshal PartOf: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "informationSource", safe.InformationSource.Typename, &o.InformationSource); err != nil {
		return fmt.Errorf("failed to unmarshal InformationSource: %w", err)
	}
	if err := unmarshalUnion(b, "derivedFrom", safe.DerivedFrom.Typename, &o.DerivedFrom); err != nil {
		return fmt.Errorf("failed to unmarshal DerivedFrom: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
