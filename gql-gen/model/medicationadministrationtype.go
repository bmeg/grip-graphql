
package model

import (
	"encoding/json"
	"fmt"
)

type SafeMedicationAdministrationType struct {
	SubPotentReason []*CodeableConcept `json:"subPotentReason,omitempty"`
	Status *string `json:"status,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	Language *string `json:"language,omitempty"`
	OccurenceDateTime *string `json:"occurenceDateTime,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Medication *CodeableReference `json:"medication,omitempty"`
	Recorded *string `json:"recorded,omitempty"`
	IsSubPotent *bool `json:"isSubPotent,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	OccurenceTiming *Timing `json:"occurenceTiming,omitempty"`
	StatusReason []*CodeableConcept `json:"statusReason,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	SupportingInformation TypedObject `json:"supportingInformation"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Subject TypedObject `json:"subject"`
	Performer []*MedicationAdministrationPerformer `json:"performer,omitempty"`
	Request *MedicationRequestType `json:"request"`
	Contained TypedObject `json:"contained,omitempty"`
	Dosage *MedicationAdministrationDosage `json:"dosage,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	OccurencePeriod *Period `json:"occurencePeriod,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	PartOf TypedObject `json:"partOf"`
	Extension []*Extension `json:"extension,omitempty"`
	ID *string `json:"id,omitempty"`
	Device []*CodeableReference `json:"device,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *MedicationAdministrationType) UnmarshalJSON(b []byte) error {
	var safe SafeMedicationAdministrationType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = MedicationAdministrationType{
		SubPotentReason: safe.SubPotentReason,
		Status: safe.Status,
		Note: safe.Note,
		Language: safe.Language,
		OccurenceDateTime: safe.OccurenceDateTime,
		ModifierExtension: safe.ModifierExtension,
		ResourceType: safe.ResourceType,
		Medication: safe.Medication,
		Recorded: safe.Recorded,
		IsSubPotent: safe.IsSubPotent,
		Text: safe.Text,
		OccurenceTiming: safe.OccurenceTiming,
		StatusReason: safe.StatusReason,
		Reason: safe.Reason,
		Category: safe.Category,
		Performer: safe.Performer,
		Request: safe.Request,
		Dosage: safe.Dosage,
		Identifier: safe.Identifier,
		ImplicitRules: safe.ImplicitRules,
		OccurencePeriod: safe.OccurencePeriod,
		Meta: safe.Meta,
		Extension: safe.Extension,
		ID: safe.ID,
		Device: safe.Device,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "supportingInformation", safe.SupportingInformation.Typename, &o.SupportingInformation); err != nil {
		return fmt.Errorf("failed to unmarshal SupportingInformation: %w", err)
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

	return nil
}
