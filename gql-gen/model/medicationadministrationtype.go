
package model

import (
	"encoding/json"
	"fmt"
)

type SafeMedicationAdministrationType struct {
	Reason []*CodeableReference `json:"reason,omitempty"`
	Status *string `json:"status,omitempty"`
	SubPotentReason []*CodeableConcept `json:"subPotentReason,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Subject TypedObject `json:"subject"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Medication *CodeableReference `json:"medication,omitempty"`
	PartOf TypedObject `json:"partOf"`
	Note TypedObject `json:"note"`
	OccurenceDateTime *string `json:"occurenceDateTime,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Performer []*MedicationAdministrationPerformer `json:"performer,omitempty"`
	SupportingInformation TypedObject `json:"supportingInformation"`
	Contained TypedObject `json:"contained,omitempty"`
	Language *string `json:"language,omitempty"`
	Dosage *MedicationAdministrationDosage `json:"dosage,omitempty"`
	ID *string `json:"id,omitempty"`
	Device []*CodeableReference `json:"device,omitempty"`
	StatusReason []*CodeableConcept `json:"statusReason,omitempty"`
	Request *MedicationRequestType `json:"request"`
	OccurencePeriod *Period `json:"occurencePeriod,omitempty"`
	OccurenceTiming *Timing `json:"occurenceTiming,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Recorded *string `json:"recorded,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	IsSubPotent *string `json:"isSubPotent,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *MedicationAdministrationType) UnmarshalJSON(b []byte) error {
	var safe SafeMedicationAdministrationType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = MedicationAdministrationType{
		Reason: safe.Reason,
		Status: safe.Status,
		SubPotentReason: safe.SubPotentReason,
		Category: safe.Category,
		Identifier: safe.Identifier,
		Medication: safe.Medication,
		OccurenceDateTime: safe.OccurenceDateTime,
		ModifierExtension: safe.ModifierExtension,
		Performer: safe.Performer,
		Language: safe.Language,
		Dosage: safe.Dosage,
		ID: safe.ID,
		Device: safe.Device,
		StatusReason: safe.StatusReason,
		Request: safe.Request,
		OccurencePeriod: safe.OccurencePeriod,
		OccurenceTiming: safe.OccurenceTiming,
		ImplicitRules: safe.ImplicitRules,
		Recorded: safe.Recorded,
		Text: safe.Text,
		ResourceType: safe.ResourceType,
		Extension: safe.Extension,
		Meta: safe.Meta,
		IsSubPotent: safe.IsSubPotent,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "partOf", safe.PartOf.Typename, &o.PartOf); err != nil {
		return fmt.Errorf("failed to unmarshal PartOf: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "supportingInformation", safe.SupportingInformation.Typename, &o.SupportingInformation); err != nil {
		return fmt.Errorf("failed to unmarshal SupportingInformation: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
