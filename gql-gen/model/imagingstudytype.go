
package model

import (
	"encoding/json"
	"fmt"
)

type SafeImagingStudyType struct {
	Description *string `json:"description,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	ID *string `json:"id,omitempty"`
	Subject TypedObject `json:"subject"`
	Status *string `json:"status,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Series []*ImagingStudySeries `json:"series,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	Modality []*CodeableConcept `json:"modality,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	PartOf *ProcedureType `json:"partOf"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Referrer TypedObject `json:"referrer"`
	Procedure []*CodeableReference `json:"procedure,omitempty"`
	Started *string `json:"started,omitempty"`
	NumberOfSeries *string `json:"numberOfSeries,omitempty"`
	Note TypedObject `json:"note"`
	BasedOn TypedObject `json:"basedOn"`
	Language *string `json:"language,omitempty"`
	NumberOfInstances *string `json:"numberOfInstances,omitempty"`
}

func (o *ImagingStudyType) UnmarshalJSON(b []byte) error {
	var safe SafeImagingStudyType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ImagingStudyType{
		Description: safe.Description,
		Identifier: safe.Identifier,
		ID: safe.ID,
		Status: safe.Status,
		Text: safe.Text,
		Meta: safe.Meta,
		Series: safe.Series,
		Extension: safe.Extension,
		ImplicitRules: safe.ImplicitRules,
		Reason: safe.Reason,
		Modality: safe.Modality,
		ResourceType: safe.ResourceType,
		PartOf: safe.PartOf,
		ModifierExtension: safe.ModifierExtension,
		Procedure: safe.Procedure,
		Started: safe.Started,
		NumberOfSeries: safe.NumberOfSeries,
		Language: safe.Language,
		NumberOfInstances: safe.NumberOfInstances,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "referrer", safe.Referrer.Typename, &o.Referrer); err != nil {
		return fmt.Errorf("failed to unmarshal Referrer: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}

	return nil
}
