
package model

import (
	"encoding/json"
	"fmt"
)

type SafeImagingStudyType struct {
	ResourceType *string `json:"resourceType,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	Subject TypedObject `json:"subject"`
	Note []*Annotation `json:"note,omitempty"`
	Procedure []*CodeableReference `json:"procedure,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Status *string `json:"status,omitempty"`
	NumberOfSeries *string `json:"numberOfSeries,omitempty"`
	Description *string `json:"description,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Language *string `json:"language,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Referrer TypedObject `json:"referrer"`
	NumberOfInstances *string `json:"numberOfInstances,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	Series []*ImagingStudySeries `json:"series,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Modality []*CodeableConcept `json:"modality,omitempty"`
	PartOf *ProcedureType `json:"partOf"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Started *string `json:"started,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ID *string `json:"id,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ImagingStudyType) UnmarshalJSON(b []byte) error {
	var safe SafeImagingStudyType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ImagingStudyType{
		ResourceType: safe.ResourceType,
		Note: safe.Note,
		Procedure: safe.Procedure,
		Text: safe.Text,
		Status: safe.Status,
		NumberOfSeries: safe.NumberOfSeries,
		Description: safe.Description,
		Meta: safe.Meta,
		Language: safe.Language,
		Extension: safe.Extension,
		NumberOfInstances: safe.NumberOfInstances,
		Reason: safe.Reason,
		Series: safe.Series,
		Identifier: safe.Identifier,
		Modality: safe.Modality,
		PartOf: safe.PartOf,
		ImplicitRules: safe.ImplicitRules,
		Started: safe.Started,
		ModifierExtension: safe.ModifierExtension,
		ID: safe.ID,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "referrer", safe.Referrer.Typename, &o.Referrer); err != nil {
		return fmt.Errorf("failed to unmarshal Referrer: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
