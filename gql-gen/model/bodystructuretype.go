
package model

import (
	"encoding/json"
	"fmt"
)

type SafeBodyStructureType struct {
	Description *string `json:"description,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Patient *PatientType `json:"patient"`
	Language *string `json:"language,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	IncludedStructure []*BodyStructureIncludedStructure `json:"includedStructure,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	ExcludedStructure []*BodyStructureIncludedStructure `json:"excludedStructure,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ID *string `json:"id,omitempty"`
	Active *string `json:"active,omitempty"`
	Morphology *CodeableConcept `json:"morphology,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Image []*Attachment `json:"image,omitempty"`
}

func (o *BodyStructureType) UnmarshalJSON(b []byte) error {
	var safe SafeBodyStructureType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = BodyStructureType{
		Description: safe.Description,
		ImplicitRules: safe.ImplicitRules,
		Patient: safe.Patient,
		Language: safe.Language,
		Meta: safe.Meta,
		ModifierExtension: safe.ModifierExtension,
		IncludedStructure: safe.IncludedStructure,
		Text: safe.Text,
		ExcludedStructure: safe.ExcludedStructure,
		Identifier: safe.Identifier,
		ID: safe.ID,
		Active: safe.Active,
		Morphology: safe.Morphology,
		Extension: safe.Extension,
		ResourceType: safe.ResourceType,
		Image: safe.Image,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
