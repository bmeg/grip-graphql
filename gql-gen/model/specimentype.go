
package model

import (
	"encoding/json"
	"fmt"
)

type SafeSpecimenType struct {
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Parent *SpecimenType `json:"parent"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	ID *string `json:"id,omitempty"`
	Language *string `json:"language,omitempty"`
	Feature []*SpecimenFeature `json:"feature,omitempty"`
	Role []*CodeableConcept `json:"role,omitempty"`
	ReceivedTime *string `json:"receivedTime,omitempty"`
	Collection *SpecimenCollection `json:"collection,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Combined *string `json:"combined,omitempty"`
	Processing []*SpecimenProcessing `json:"processing,omitempty"`
	Type *CodeableConcept `json:"type,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	AccessionIdentifier *Identifier `json:"accessionIdentifier,omitempty"`
	Container []*SpecimenContainer `json:"container,omitempty"`
	Condition []*CodeableConcept `json:"condition,omitempty"`
	Status *string `json:"status,omitempty"`
	Subject TypedObject `json:"subject"`
	Text *Narrative `json:"text,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *SpecimenType) UnmarshalJSON(b []byte) error {
	var safe SafeSpecimenType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = SpecimenType{
		ModifierExtension: safe.ModifierExtension,
		Parent: safe.Parent,
		Identifier: safe.Identifier,
		Note: safe.Note,
		ID: safe.ID,
		Language: safe.Language,
		Feature: safe.Feature,
		Role: safe.Role,
		ReceivedTime: safe.ReceivedTime,
		Collection: safe.Collection,
		Extension: safe.Extension,
		Combined: safe.Combined,
		Processing: safe.Processing,
		Type: safe.Type,
		Meta: safe.Meta,
		ResourceType: safe.ResourceType,
		AccessionIdentifier: safe.AccessionIdentifier,
		Container: safe.Container,
		Condition: safe.Condition,
		Status: safe.Status,
		Text: safe.Text,
		ImplicitRules: safe.ImplicitRules,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
