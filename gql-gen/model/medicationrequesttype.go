
package model

import (
	"encoding/json"
	"fmt"
)

type SafeMedicationRequestType struct {
	Subject TypedObject `json:"subject"`
	Recorder TypedObject `json:"recorder"`
	SupportingInformation TypedObject `json:"supportingInformation"`
	StatusReason *CodeableConcept `json:"statusReason,omitempty"`
	Substitution *MedicationRequestSubstitution `json:"substitution,omitempty"`
	ID *string `json:"id,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Intent *string `json:"intent,omitempty"`
	Requester TypedObject `json:"requester"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	DosageInstruction []*Dosage `json:"dosageInstruction,omitempty"`
	Priority *string `json:"priority,omitempty"`
	PriorPrescription *MedicationRequestType `json:"priorPrescription"`
	InformationSource TypedObject `json:"informationSource"`
	Device []*CodeableReference `json:"device,omitempty"`
	DoNotPerform *bool `json:"doNotPerform,omitempty"`
	EffectiveDosePeriod *Period `json:"effectiveDosePeriod,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	PerformerType *CodeableConcept `json:"performerType,omitempty"`
	Reported *bool `json:"reported,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	GroupIdentifier *Identifier `json:"groupIdentifier,omitempty"`
	StatusChanged *string `json:"statusChanged,omitempty"`
	RenderedDosageInstruction *string `json:"renderedDosageInstruction,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Performer TypedObject `json:"performer"`
	DispenseRequest *MedicationRequestDispenseRequest `json:"dispenseRequest,omitempty"`
	Language *string `json:"language,omitempty"`
	Medication *CodeableReference `json:"medication,omitempty"`
	AuthoredOn *string `json:"authoredOn,omitempty"`
	Status *string `json:"status,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	CourseOfTherapyType *CodeableConcept `json:"courseOfTherapyType,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *MedicationRequestType) UnmarshalJSON(b []byte) error {
	var safe SafeMedicationRequestType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = MedicationRequestType{
		StatusReason: safe.StatusReason,
		Substitution: safe.Substitution,
		ID: safe.ID,
		Meta: safe.Meta,
		Intent: safe.Intent,
		ModifierExtension: safe.ModifierExtension,
		ImplicitRules: safe.ImplicitRules,
		DosageInstruction: safe.DosageInstruction,
		Priority: safe.Priority,
		PriorPrescription: safe.PriorPrescription,
		Device: safe.Device,
		DoNotPerform: safe.DoNotPerform,
		EffectiveDosePeriod: safe.EffectiveDosePeriod,
		PerformerType: safe.PerformerType,
		Reported: safe.Reported,
		ResourceType: safe.ResourceType,
		GroupIdentifier: safe.GroupIdentifier,
		StatusChanged: safe.StatusChanged,
		RenderedDosageInstruction: safe.RenderedDosageInstruction,
		Note: safe.Note,
		Reason: safe.Reason,
		Category: safe.Category,
		DispenseRequest: safe.DispenseRequest,
		Language: safe.Language,
		Medication: safe.Medication,
		AuthoredOn: safe.AuthoredOn,
		Status: safe.Status,
		Extension: safe.Extension,
		Identifier: safe.Identifier,
		Text: safe.Text,
		CourseOfTherapyType: safe.CourseOfTherapyType,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "recorder", safe.Recorder.Typename, &o.Recorder); err != nil {
		return fmt.Errorf("failed to unmarshal Recorder: %w", err)
	}
	if err := unmarshalUnion(b, "supportingInformation", safe.SupportingInformation.Typename, &o.SupportingInformation); err != nil {
		return fmt.Errorf("failed to unmarshal SupportingInformation: %w", err)
	}
	if err := unmarshalUnion(b, "requester", safe.Requester.Typename, &o.Requester); err != nil {
		return fmt.Errorf("failed to unmarshal Requester: %w", err)
	}
	if err := unmarshalUnion(b, "informationSource", safe.InformationSource.Typename, &o.InformationSource); err != nil {
		return fmt.Errorf("failed to unmarshal InformationSource: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "performer", safe.Performer.Typename, &o.Performer); err != nil {
		return fmt.Errorf("failed to unmarshal Performer: %w", err)
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}

	return nil
}
