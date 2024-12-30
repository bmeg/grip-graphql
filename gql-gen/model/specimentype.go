
package model

import (
	"encoding/json"
	"fmt"
)

type SafeSpecimenType struct {
	ResourceType *string `json:"resourceType,omitempty"`
	Subject TypedObject `json:"subject"`
	Combined *string `json:"combined,omitempty"`
	AccessionIdentifier *Identifier `json:"accessionIdentifier,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Type *CodeableConcept `json:"type,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Container []*SpecimenContainer `json:"container,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Parent *SpecimenType `json:"parent"`
	Meta *Meta `json:"meta,omitempty"`
	Condition []*CodeableConcept `json:"condition,omitempty"`
	Feature []*SpecimenFeature `json:"feature,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	ID *string `json:"id,omitempty"`
	Processing []*SpecimenProcessing `json:"processing,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Collection TypedObject `json:"collection"`
	ReceivedTime *string `json:"receivedTime,omitempty"`
	Role []*CodeableConcept `json:"role,omitempty"`
	Language *string `json:"language,omitempty"`
	Note TypedObject `json:"note"`
	Status *string `json:"status,omitempty"`
}

func (o *SpecimenType) UnmarshalJSON(b []byte) error {
	var safe SafeSpecimenType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = SpecimenType{
		ResourceType: safe.ResourceType,
		Combined: safe.Combined,
		AccessionIdentifier: safe.AccessionIdentifier,
		ModifierExtension: safe.ModifierExtension,
		Type: safe.Type,
		Text: safe.Text,
		Container: safe.Container,
		ImplicitRules: safe.ImplicitRules,
		Parent: safe.Parent,
		Meta: safe.Meta,
		Condition: safe.Condition,
		Feature: safe.Feature,
		ID: safe.ID,
		Processing: safe.Processing,
		Extension: safe.Extension,
		Identifier: safe.Identifier,
		ReceivedTime: safe.ReceivedTime,
		Role: safe.Role,
		Language: safe.Language,
		Status: safe.Status,
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "collection", safe.Collection.Typename, &o.Collection); err != nil {
		return fmt.Errorf("failed to unmarshal Collection: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}

	return nil
}
