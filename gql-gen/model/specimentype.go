
package model

import (
	"encoding/json"
	"fmt"
)

type SafeSpecimenType struct {
	Parent *SpecimenType `json:"parent"`
	Language *string `json:"language,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Processing []*SpecimenProcessing `json:"processing,omitempty"`
	Type *CodeableConcept `json:"type,omitempty"`
	Subject TypedObject `json:"subject"`
	Role []*CodeableConcept `json:"role,omitempty"`
	Combined *string `json:"combined,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Collection *SpecimenCollection `json:"collection,omitempty"`
	Feature []*SpecimenFeature `json:"feature,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	ID *string `json:"id,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Condition []*CodeableConcept `json:"condition,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	AccessionIdentifier *Identifier `json:"accessionIdentifier,omitempty"`
	Container []*SpecimenContainer `json:"container,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ReceivedTime *string `json:"receivedTime,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Status *string `json:"status,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *SpecimenType) UnmarshalJSON(b []byte) error {
	var safe SafeSpecimenType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = SpecimenType{
		Parent: safe.Parent,
		Language: safe.Language,
		Extension: safe.Extension,
		Text: safe.Text,
		Processing: safe.Processing,
		Type: safe.Type,
		Role: safe.Role,
		Combined: safe.Combined,
		Meta: safe.Meta,
		Collection: safe.Collection,
		Feature: safe.Feature,
		Note: safe.Note,
		ID: safe.ID,
		Identifier: safe.Identifier,
		Condition: safe.Condition,
		ImplicitRules: safe.ImplicitRules,
		ResourceType: safe.ResourceType,
		AccessionIdentifier: safe.AccessionIdentifier,
		Container: safe.Container,
		ModifierExtension: safe.ModifierExtension,
		ReceivedTime: safe.ReceivedTime,
		Status: safe.Status,
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
