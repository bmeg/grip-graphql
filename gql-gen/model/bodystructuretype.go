
package model

import (
	"encoding/json"
	"fmt"
)

type SafeBodyStructureType struct {
	Meta *Meta `json:"meta,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Morphology *CodeableConcept `json:"morphology,omitempty"`
	ExcludedStructure []*BodyStructureIncludedStructure `json:"excludedStructure,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ID *string `json:"id,omitempty"`
	Language *string `json:"language,omitempty"`
	Active *string `json:"active,omitempty"`
	Image []*Attachment `json:"image,omitempty"`
	Patient *PatientType `json:"patient"`
	IncludedStructure []*BodyStructureIncludedStructure `json:"includedStructure,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Description *string `json:"description,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *BodyStructureType) UnmarshalJSON(b []byte) error {
	var safe SafeBodyStructureType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = BodyStructureType{
		Meta: safe.Meta,
		Text: safe.Text,
		Extension: safe.Extension,
		ModifierExtension: safe.ModifierExtension,
		Morphology: safe.Morphology,
		ExcludedStructure: safe.ExcludedStructure,
		Identifier: safe.Identifier,
		ID: safe.ID,
		Language: safe.Language,
		Active: safe.Active,
		Image: safe.Image,
		Patient: safe.Patient,
		IncludedStructure: safe.IncludedStructure,
		Description: safe.Description,
		ImplicitRules: safe.ImplicitRules,
		ResourceType: safe.ResourceType,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
