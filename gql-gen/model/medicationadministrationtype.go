
package model

import (
	"encoding/json"
	"fmt"
)

type SafeMedicationAdministrationType struct {
	PartOf TypedObject `json:"partOf"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	Status *string `json:"status,omitempty"`
	Dosage *MedicationAdministrationDosage `json:"dosage,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Language *string `json:"language,omitempty"`
	OccurenceDateTime *string `json:"occurenceDateTime,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	ID *string `json:"id,omitempty"`
	Device []*CodeableReference `json:"device,omitempty"`
	Recorded *string `json:"recorded,omitempty"`
	Request *MedicationRequestType `json:"request"`
	SupportingInformation TypedObject `json:"supportingInformation"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Subject TypedObject `json:"subject"`
	OccurencePeriod *Period `json:"occurencePeriod,omitempty"`
	IsSubPotent *string `json:"isSubPotent,omitempty"`
	Medication *CodeableReference `json:"medication,omitempty"`
	SubPotentReason []*CodeableConcept `json:"subPotentReason,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	OccurenceTiming *Timing `json:"occurenceTiming,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Performer []*MedicationAdministrationPerformer `json:"performer,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	StatusReason []*CodeableConcept `json:"statusReason,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
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
		Dosage: safe.Dosage,
		Text: safe.Text,
		Extension: safe.Extension,
		Language: safe.Language,
		OccurenceDateTime: safe.OccurenceDateTime,
		ID: safe.ID,
		Device: safe.Device,
		Recorded: safe.Recorded,
		Request: safe.Request,
		Identifier: safe.Identifier,
		OccurencePeriod: safe.OccurencePeriod,
		IsSubPotent: safe.IsSubPotent,
		Medication: safe.Medication,
		SubPotentReason: safe.SubPotentReason,
		Category: safe.Category,
		Meta: safe.Meta,
		ModifierExtension: safe.ModifierExtension,
		OccurenceTiming: safe.OccurenceTiming,
		ResourceType: safe.ResourceType,
		Performer: safe.Performer,
		ImplicitRules: safe.ImplicitRules,
		StatusReason: safe.StatusReason,
		Note: safe.Note,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "partOf", safe.PartOf.Typename, &o.PartOf); err != nil {
		return fmt.Errorf("failed to unmarshal PartOf: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "supportingInformation", safe.SupportingInformation.Typename, &o.SupportingInformation); err != nil {
		return fmt.Errorf("failed to unmarshal SupportingInformation: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}

	return nil
}
