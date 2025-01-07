
package model

import (
	"encoding/json"
	"fmt"
)

type SafeMedicationStatementType struct {
	DateAsserted *string `json:"dateAsserted,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	RelatedClinicalInformation TypedObject `json:"relatedClinicalInformation"`
	Subject TypedObject `json:"subject"`
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	PartOf TypedObject `json:"partOf"`
	Extension []*Extension `json:"extension,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Dosage []*Dosage `json:"dosage,omitempty"`
	InformationSource TypedObject `json:"informationSource"`
	Note []*Annotation `json:"note,omitempty"`
	EffectiveTiming *Timing `json:"effectiveTiming,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Adherence *MedicationStatementAdherence `json:"adherence,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`
	Status *string `json:"status,omitempty"`
	ID *string `json:"id,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	DerivedFrom TypedObject `json:"derivedFrom"`
	RenderedDosageInstruction *string `json:"renderedDosageInstruction,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Language *string `json:"language,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Medication *CodeableReference `json:"medication,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *MedicationStatementType) UnmarshalJSON(b []byte) error {
	var safe SafeMedicationStatementType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = MedicationStatementType{
		DateAsserted: safe.DateAsserted,
		Meta: safe.Meta,
		EffectivePeriod: safe.EffectivePeriod,
		Extension: safe.Extension,
		Dosage: safe.Dosage,
		Note: safe.Note,
		EffectiveTiming: safe.EffectiveTiming,
		Reason: safe.Reason,
		Text: safe.Text,
		ImplicitRules: safe.ImplicitRules,
		Adherence: safe.Adherence,
		Identifier: safe.Identifier,
		EffectiveDateTime: safe.EffectiveDateTime,
		Status: safe.Status,
		ID: safe.ID,
		ModifierExtension: safe.ModifierExtension,
		RenderedDosageInstruction: safe.RenderedDosageInstruction,
		ResourceType: safe.ResourceType,
		Language: safe.Language,
		Category: safe.Category,
		Medication: safe.Medication,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "relatedClinicalInformation", safe.RelatedClinicalInformation.Typename, &o.RelatedClinicalInformation); err != nil {
		return fmt.Errorf("failed to unmarshal RelatedClinicalInformation: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "partOf", safe.PartOf.Typename, &o.PartOf); err != nil {
		return fmt.Errorf("failed to unmarshal PartOf: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "informationSource", safe.InformationSource.Typename, &o.InformationSource); err != nil {
		return fmt.Errorf("failed to unmarshal InformationSource: %w", err)
	}
	if err := unmarshalUnion(b, "derivedFrom", safe.DerivedFrom.Typename, &o.DerivedFrom); err != nil {
		return fmt.Errorf("failed to unmarshal DerivedFrom: %w", err)
	}

	return nil
}
