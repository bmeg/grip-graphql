
package model

import (
	"encoding/json"
	"fmt"
)

type SafeImagingStudyType struct {
	PartOf *ProcedureType `json:"partOf"`
	Text *Narrative `json:"text,omitempty"`
	Procedure []*CodeableReference `json:"procedure,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	Extension []*Extension `json:"extension,omitempty"`
	NumberOfInstances *int `json:"numberOfInstances,omitempty"`
	Referrer TypedObject `json:"referrer"`
	Series []*ImagingStudySeries `json:"series,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Description *string `json:"description,omitempty"`
	Started *string `json:"started,omitempty"`
	Status *string `json:"status,omitempty"`
	Modality []*CodeableConcept `json:"modality,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Language *string `json:"language,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Subject TypedObject `json:"subject"`
	ID *string `json:"id,omitempty"`
	NumberOfSeries *int `json:"numberOfSeries,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ImagingStudyType) UnmarshalJSON(b []byte) error {
	var safe SafeImagingStudyType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ImagingStudyType{
		PartOf: safe.PartOf,
		Text: safe.Text,
		Procedure: safe.Procedure,
		Extension: safe.Extension,
		NumberOfInstances: safe.NumberOfInstances,
		Series: safe.Series,
		Meta: safe.Meta,
		Description: safe.Description,
		Started: safe.Started,
		Status: safe.Status,
		Modality: safe.Modality,
		Identifier: safe.Identifier,
		Language: safe.Language,
		Note: safe.Note,
		ModifierExtension: safe.ModifierExtension,
		Reason: safe.Reason,
		ResourceType: safe.ResourceType,
		ImplicitRules: safe.ImplicitRules,
		ID: safe.ID,
		NumberOfSeries: safe.NumberOfSeries,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "referrer", safe.Referrer.Typename, &o.Referrer); err != nil {
		return fmt.Errorf("failed to unmarshal Referrer: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}

	return nil
}
