
package model

import (
	"encoding/json"
	"fmt"
)

type SafeImagingStudyType struct {
	Identifier []*Identifier `json:"identifier,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	Procedure []*CodeableReference `json:"procedure,omitempty"`
	Series []*ImagingStudySeries `json:"series,omitempty"`
	PartOf *ProcedureType `json:"partOf"`
	Text *Narrative `json:"text,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Modality []*CodeableConcept `json:"modality,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Started *string `json:"started,omitempty"`
	Status *string `json:"status,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	Language *string `json:"language,omitempty"`
	Note TypedObject `json:"note"`
	Extension []*Extension `json:"extension,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Referrer TypedObject `json:"referrer"`
	NumberOfInstances *string `json:"numberOfInstances,omitempty"`
	NumberOfSeries *string `json:"numberOfSeries,omitempty"`
	ID *string `json:"id,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Subject TypedObject `json:"subject"`
	Description *string `json:"description,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ImagingStudyType) UnmarshalJSON(b []byte) error {
	var safe SafeImagingStudyType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ImagingStudyType{
		Identifier: safe.Identifier,
		Reason: safe.Reason,
		Procedure: safe.Procedure,
		Series: safe.Series,
		PartOf: safe.PartOf,
		Text: safe.Text,
		Meta: safe.Meta,
		ModifierExtension: safe.ModifierExtension,
		Modality: safe.Modality,
		Started: safe.Started,
		Status: safe.Status,
		Language: safe.Language,
		Extension: safe.Extension,
		ImplicitRules: safe.ImplicitRules,
		NumberOfInstances: safe.NumberOfInstances,
		NumberOfSeries: safe.NumberOfSeries,
		ID: safe.ID,
		ResourceType: safe.ResourceType,
		Description: safe.Description,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "referrer", safe.Referrer.Typename, &o.Referrer); err != nil {
		return fmt.Errorf("failed to unmarshal Referrer: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}

	return nil
}
