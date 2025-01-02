
package model

import (
	"encoding/json"
	"fmt"
)

type SafeSpecimenType struct {
	Meta *Meta `json:"meta,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Container []*SpecimenContainer `json:"container,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Parent *SpecimenType `json:"parent"`
	Type *CodeableConcept `json:"type,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	ID *string `json:"id,omitempty"`
	Collection TypedObject `json:"collection"`
	Language *string `json:"language,omitempty"`
	Combined *string `json:"combined,omitempty"`
	Subject TypedObject `json:"subject"`
	Processing []*SpecimenProcessing `json:"processing,omitempty"`
	AccessionIdentifier *Identifier `json:"accessionIdentifier,omitempty"`
	Note TypedObject `json:"note"`
	Role []*CodeableConcept `json:"role,omitempty"`
	Condition []*CodeableConcept `json:"condition,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	ReceivedTime *string `json:"receivedTime,omitempty"`
	Feature []*SpecimenFeature `json:"feature,omitempty"`
	Status *string `json:"status,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *SpecimenType) UnmarshalJSON(b []byte) error {
	var safe SafeSpecimenType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = SpecimenType{
		Meta: safe.Meta,
		Text: safe.Text,
		Container: safe.Container,
		ModifierExtension: safe.ModifierExtension,
		Parent: safe.Parent,
		Type: safe.Type,
		ImplicitRules: safe.ImplicitRules,
		ID: safe.ID,
		Language: safe.Language,
		Combined: safe.Combined,
		Processing: safe.Processing,
		AccessionIdentifier: safe.AccessionIdentifier,
		Role: safe.Role,
		Condition: safe.Condition,
		Extension: safe.Extension,
		ResourceType: safe.ResourceType,
		Identifier: safe.Identifier,
		ReceivedTime: safe.ReceivedTime,
		Feature: safe.Feature,
		Status: safe.Status,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "collection", safe.Collection.Typename, &o.Collection); err != nil {
		return fmt.Errorf("failed to unmarshal Collection: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}

	return nil
}
