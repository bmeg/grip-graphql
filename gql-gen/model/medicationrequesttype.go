
package model

import (
	"encoding/json"
	"fmt"
)

type SafeMedicationRequestType struct {
	Medication *CodeableReference `json:"medication,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Device []*CodeableReference `json:"device,omitempty"`
	StatusReason *CodeableConcept `json:"statusReason,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	DispenseRequest *MedicationRequestDispenseRequest `json:"dispenseRequest,omitempty"`
	Status *string `json:"status,omitempty"`
	PriorPrescription *MedicationRequestType `json:"priorPrescription"`
	StatusChanged *string `json:"statusChanged,omitempty"`
	Substitution *MedicationRequestSubstitution `json:"substitution,omitempty"`
	Requester TypedObject `json:"requester"`
	ID *string `json:"id,omitempty"`
	AuthoredOn *string `json:"authoredOn,omitempty"`
	Language *string `json:"language,omitempty"`
	Intent *string `json:"intent,omitempty"`
	SupportingInformation TypedObject `json:"supportingInformation"`
	RenderedDosageInstruction *string `json:"renderedDosageInstruction,omitempty"`
	DoNotPerform *string `json:"doNotPerform,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	PerformerType *CodeableConcept `json:"performerType,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Priority *string `json:"priority,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	InformationSource TypedObject `json:"informationSource"`
	BasedOn TypedObject `json:"basedOn"`
	Subject TypedObject `json:"subject"`
	Recorder TypedObject `json:"recorder"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	DosageInstruction []*Dosage `json:"dosageInstruction,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	CourseOfTherapyType *CodeableConcept `json:"courseOfTherapyType,omitempty"`
	EffectiveDosePeriod *Period `json:"effectiveDosePeriod,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Reported *string `json:"reported,omitempty"`
	GroupIdentifier *Identifier `json:"groupIdentifier,omitempty"`
	Performer TypedObject `json:"performer"`
	Extension []*Extension `json:"extension,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *MedicationRequestType) UnmarshalJSON(b []byte) error {
	var safe SafeMedicationRequestType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = MedicationRequestType{
		Medication: safe.Medication,
		Meta: safe.Meta,
		Device: safe.Device,
		StatusReason: safe.StatusReason,
		ResourceType: safe.ResourceType,
		DispenseRequest: safe.DispenseRequest,
		Status: safe.Status,
		PriorPrescription: safe.PriorPrescription,
		StatusChanged: safe.StatusChanged,
		Substitution: safe.Substitution,
		ID: safe.ID,
		AuthoredOn: safe.AuthoredOn,
		Language: safe.Language,
		Intent: safe.Intent,
		RenderedDosageInstruction: safe.RenderedDosageInstruction,
		DoNotPerform: safe.DoNotPerform,
		Reason: safe.Reason,
		PerformerType: safe.PerformerType,
		ImplicitRules: safe.ImplicitRules,
		Priority: safe.Priority,
		ModifierExtension: safe.ModifierExtension,
		Text: safe.Text,
		DosageInstruction: safe.DosageInstruction,
		Identifier: safe.Identifier,
		Note: safe.Note,
		CourseOfTherapyType: safe.CourseOfTherapyType,
		EffectiveDosePeriod: safe.EffectiveDosePeriod,
		Category: safe.Category,
		Reported: safe.Reported,
		GroupIdentifier: safe.GroupIdentifier,
		Extension: safe.Extension,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "requester", safe.Requester.Typename, &o.Requester); err != nil {
		return fmt.Errorf("failed to unmarshal Requester: %w", err)
	}
	if err := unmarshalUnion(b, "supportingInformation", safe.SupportingInformation.Typename, &o.SupportingInformation); err != nil {
		return fmt.Errorf("failed to unmarshal SupportingInformation: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "informationSource", safe.InformationSource.Typename, &o.InformationSource); err != nil {
		return fmt.Errorf("failed to unmarshal InformationSource: %w", err)
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "recorder", safe.Recorder.Typename, &o.Recorder); err != nil {
		return fmt.Errorf("failed to unmarshal Recorder: %w", err)
	}
	if err := unmarshalUnion(b, "performer", safe.Performer.Typename, &o.Performer); err != nil {
		return fmt.Errorf("failed to unmarshal Performer: %w", err)
	}

	return nil
}
