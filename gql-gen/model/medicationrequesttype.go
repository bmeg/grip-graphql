
package model

import (
	"encoding/json"
	"fmt"
)

type SafeMedicationRequestType struct {
	Requester TypedObject `json:"requester"`
	EffectiveDosePeriod *Period `json:"effectiveDosePeriod,omitempty"`
	PerformerType *CodeableConcept `json:"performerType,omitempty"`
	AuthoredOn *string `json:"authoredOn,omitempty"`
	StatusChanged *string `json:"statusChanged,omitempty"`
	GroupIdentifier *Identifier `json:"groupIdentifier,omitempty"`
	SupportingInformation TypedObject `json:"supportingInformation"`
	Text *Narrative `json:"text,omitempty"`
	StatusReason *CodeableConcept `json:"statusReason,omitempty"`
	DispenseRequest *MedicationRequestDispenseRequest `json:"dispenseRequest,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Substitution *MedicationRequestSubstitution `json:"substitution,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	DosageInstruction []*Dosage `json:"dosageInstruction,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	CourseOfTherapyType *CodeableConcept `json:"courseOfTherapyType,omitempty"`
	Intent *string `json:"intent,omitempty"`
	PriorPrescription *MedicationRequestType `json:"priorPrescription"`
	Medication *CodeableReference `json:"medication,omitempty"`
	Subject TypedObject `json:"subject"`
	Recorder TypedObject `json:"recorder"`
	Contained TypedObject `json:"contained,omitempty"`
	RenderedDosageInstruction *string `json:"renderedDosageInstruction,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Language *string `json:"language,omitempty"`
	Priority *string `json:"priority,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	DoNotPerform *string `json:"doNotPerform,omitempty"`
	Device []*CodeableReference `json:"device,omitempty"`
	ID *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	Reported *string `json:"reported,omitempty"`
	InformationSource TypedObject `json:"informationSource"`
	Note TypedObject `json:"note"`
	Extension []*Extension `json:"extension,omitempty"`
	Performer TypedObject `json:"performer"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *MedicationRequestType) UnmarshalJSON(b []byte) error {
	var safe SafeMedicationRequestType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = MedicationRequestType{
		EffectiveDosePeriod: safe.EffectiveDosePeriod,
		PerformerType: safe.PerformerType,
		AuthoredOn: safe.AuthoredOn,
		StatusChanged: safe.StatusChanged,
		GroupIdentifier: safe.GroupIdentifier,
		Text: safe.Text,
		StatusReason: safe.StatusReason,
		DispenseRequest: safe.DispenseRequest,
		Meta: safe.Meta,
		Substitution: safe.Substitution,
		ResourceType: safe.ResourceType,
		DosageInstruction: safe.DosageInstruction,
		Category: safe.Category,
		CourseOfTherapyType: safe.CourseOfTherapyType,
		Intent: safe.Intent,
		PriorPrescription: safe.PriorPrescription,
		Medication: safe.Medication,
		RenderedDosageInstruction: safe.RenderedDosageInstruction,
		ModifierExtension: safe.ModifierExtension,
		ImplicitRules: safe.ImplicitRules,
		Identifier: safe.Identifier,
		Language: safe.Language,
		Priority: safe.Priority,
		Reason: safe.Reason,
		DoNotPerform: safe.DoNotPerform,
		Device: safe.Device,
		ID: safe.ID,
		Status: safe.Status,
		Reported: safe.Reported,
		Extension: safe.Extension,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "requester", safe.Requester.Typename, &o.Requester); err != nil {
		return fmt.Errorf("failed to unmarshal Requester: %w", err)
	}
	if err := unmarshalUnion(b, "supportingInformation", safe.SupportingInformation.Typename, &o.SupportingInformation); err != nil {
		return fmt.Errorf("failed to unmarshal SupportingInformation: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "recorder", safe.Recorder.Typename, &o.Recorder); err != nil {
		return fmt.Errorf("failed to unmarshal Recorder: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "informationSource", safe.InformationSource.Typename, &o.InformationSource); err != nil {
		return fmt.Errorf("failed to unmarshal InformationSource: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "performer", safe.Performer.Typename, &o.Performer); err != nil {
		return fmt.Errorf("failed to unmarshal Performer: %w", err)
	}

	return nil
}
