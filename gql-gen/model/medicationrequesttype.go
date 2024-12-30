
package model

import (
	"encoding/json"
	"fmt"
)

type SafeMedicationRequestType struct {
	Subject TypedObject `json:"subject"`
	DosageInstruction []*Dosage `json:"dosageInstruction,omitempty"`
	Note TypedObject `json:"note"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Medication *CodeableReference `json:"medication,omitempty"`
	Priority *string `json:"priority,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	CourseOfTherapyType *CodeableConcept `json:"courseOfTherapyType,omitempty"`
	DispenseRequest *MedicationRequestDispenseRequest `json:"dispenseRequest,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	EffectiveDosePeriod *Period `json:"effectiveDosePeriod,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	PerformerType *CodeableConcept `json:"performerType,omitempty"`
	GroupIdentifier *Identifier `json:"groupIdentifier,omitempty"`
	Substitution *MedicationRequestSubstitution `json:"substitution,omitempty"`
	Status *string `json:"status,omitempty"`
	SupportingInformation TypedObject `json:"supportingInformation"`
	InformationSource TypedObject `json:"informationSource"`
	DoNotPerform *string `json:"doNotPerform,omitempty"`
	Reported *string `json:"reported,omitempty"`
	Performer TypedObject `json:"performer"`
	BasedOn TypedObject `json:"basedOn"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	StatusReason *CodeableConcept `json:"statusReason,omitempty"`
	RenderedDosageInstruction *string `json:"renderedDosageInstruction,omitempty"`
	Recorder TypedObject `json:"recorder"`
	AuthoredOn *string `json:"authoredOn,omitempty"`
	Intent *string `json:"intent,omitempty"`
	Language *string `json:"language,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Requester TypedObject `json:"requester"`
	PriorPrescription *MedicationRequestType `json:"priorPrescription"`
	Device []*CodeableReference `json:"device,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	StatusChanged *string `json:"statusChanged,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	ID *string `json:"id,omitempty"`
}

func (o *MedicationRequestType) UnmarshalJSON(b []byte) error {
	var safe SafeMedicationRequestType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = MedicationRequestType{
		DosageInstruction: safe.DosageInstruction,
		Category: safe.Category,
		Medication: safe.Medication,
		Priority: safe.Priority,
		ModifierExtension: safe.ModifierExtension,
		CourseOfTherapyType: safe.CourseOfTherapyType,
		DispenseRequest: safe.DispenseRequest,
		Text: safe.Text,
		EffectiveDosePeriod: safe.EffectiveDosePeriod,
		Extension: safe.Extension,
		PerformerType: safe.PerformerType,
		GroupIdentifier: safe.GroupIdentifier,
		Substitution: safe.Substitution,
		Status: safe.Status,
		DoNotPerform: safe.DoNotPerform,
		Reported: safe.Reported,
		Identifier: safe.Identifier,
		StatusReason: safe.StatusReason,
		RenderedDosageInstruction: safe.RenderedDosageInstruction,
		AuthoredOn: safe.AuthoredOn,
		Intent: safe.Intent,
		Language: safe.Language,
		Reason: safe.Reason,
		Meta: safe.Meta,
		PriorPrescription: safe.PriorPrescription,
		Device: safe.Device,
		ResourceType: safe.ResourceType,
		StatusChanged: safe.StatusChanged,
		ImplicitRules: safe.ImplicitRules,
		ID: safe.ID,
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "supportingInformation", safe.SupportingInformation.Typename, &o.SupportingInformation); err != nil {
		return fmt.Errorf("failed to unmarshal SupportingInformation: %w", err)
	}
	if err := unmarshalUnion(b, "informationSource", safe.InformationSource.Typename, &o.InformationSource); err != nil {
		return fmt.Errorf("failed to unmarshal InformationSource: %w", err)
	}
	if err := unmarshalUnion(b, "performer", safe.Performer.Typename, &o.Performer); err != nil {
		return fmt.Errorf("failed to unmarshal Performer: %w", err)
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "recorder", safe.Recorder.Typename, &o.Recorder); err != nil {
		return fmt.Errorf("failed to unmarshal Recorder: %w", err)
	}
	if err := unmarshalUnion(b, "requester", safe.Requester.Typename, &o.Requester); err != nil {
		return fmt.Errorf("failed to unmarshal Requester: %w", err)
	}

	return nil
}
