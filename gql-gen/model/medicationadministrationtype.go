
package model

import (
	"encoding/json"
	"fmt"
)

type SafeMedicationAdministrationType struct {
	Note TypedObject `json:"note"`
	Dosage *MedicationAdministrationDosage `json:"dosage,omitempty"`
	Language *string `json:"language,omitempty"`
	Request *MedicationRequestType `json:"request"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	OccurencePeriod *Period `json:"occurencePeriod,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Medication *CodeableReference `json:"medication,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	IsSubPotent *string `json:"isSubPotent,omitempty"`
	Recorded *string `json:"recorded,omitempty"`
	Performer []*MedicationAdministrationPerformer `json:"performer,omitempty"`
	StatusReason []*CodeableConcept `json:"statusReason,omitempty"`
	Subject TypedObject `json:"subject"`
	Contained TypedObject `json:"contained,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	PartOf TypedObject `json:"partOf"`
	OccurenceTiming *Timing `json:"occurenceTiming,omitempty"`
	SubPotentReason []*CodeableConcept `json:"subPotentReason,omitempty"`
	SupportingInformation TypedObject `json:"supportingInformation"`
	Status *string `json:"status,omitempty"`
	ID *string `json:"id,omitempty"`
	OccurenceDateTime *string `json:"occurenceDateTime,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Device []*CodeableReference `json:"device,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
}

func (o *MedicationAdministrationType) UnmarshalJSON(b []byte) error {
	var safe SafeMedicationAdministrationType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = MedicationAdministrationType{
		Dosage: safe.Dosage,
		Language: safe.Language,
		Request: safe.Request,
		Reason: safe.Reason,
		OccurencePeriod: safe.OccurencePeriod,
		ResourceType: safe.ResourceType,
		Text: safe.Text,
		Medication: safe.Medication,
		ImplicitRules: safe.ImplicitRules,
		Category: safe.Category,
		Meta: safe.Meta,
		IsSubPotent: safe.IsSubPotent,
		Recorded: safe.Recorded,
		Performer: safe.Performer,
		StatusReason: safe.StatusReason,
		Extension: safe.Extension,
		OccurenceTiming: safe.OccurenceTiming,
		SubPotentReason: safe.SubPotentReason,
		Status: safe.Status,
		ID: safe.ID,
		OccurenceDateTime: safe.OccurenceDateTime,
		ModifierExtension: safe.ModifierExtension,
		Device: safe.Device,
		Identifier: safe.Identifier,
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "partOf", safe.PartOf.Typename, &o.PartOf); err != nil {
		return fmt.Errorf("failed to unmarshal PartOf: %w", err)
	}
	if err := unmarshalUnion(b, "supportingInformation", safe.SupportingInformation.Typename, &o.SupportingInformation); err != nil {
		return fmt.Errorf("failed to unmarshal SupportingInformation: %w", err)
	}

	return nil
}
