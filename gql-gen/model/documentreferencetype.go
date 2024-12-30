
package model

import (
	"encoding/json"
	"fmt"
)

type SafeDocumentReferenceType struct {
	Modality []*CodeableConcept `json:"modality,omitempty"`
	SecurityLabel []*CodeableConcept `json:"securityLabel,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Subject TypedObject `json:"subject"`
	ID *string `json:"id,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Event []*CodeableReference `json:"event,omitempty"`
	PracticeSetting *CodeableConcept `json:"practiceSetting,omitempty"`
	BodySite []*CodeableReference `json:"bodySite,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Date *string `json:"date,omitempty"`
	Language *string `json:"language,omitempty"`
	Author TypedObject `json:"author"`
	Type *CodeableConcept `json:"type,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Period *Period `json:"period,omitempty"`
	FacilityType *CodeableConcept `json:"facilityType,omitempty"`
	Attester TypedObject `json:"attester"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	DocStatus *string `json:"docStatus,omitempty"`
	Custodian *OrganizationType `json:"custodian"`
	Description *string `json:"description,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	RelatesTo []*DocumentReferenceRelatesTo `json:"relatesTo,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Status *string `json:"status,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Content []*DocumentReferenceContent `json:"content,omitempty"`
	Version *string `json:"version,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
}

func (o *DocumentReferenceType) UnmarshalJSON(b []byte) error {
	var safe SafeDocumentReferenceType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = DocumentReferenceType{
		Modality: safe.Modality,
		SecurityLabel: safe.SecurityLabel,
		Meta: safe.Meta,
		ID: safe.ID,
		ModifierExtension: safe.ModifierExtension,
		Event: safe.Event,
		PracticeSetting: safe.PracticeSetting,
		BodySite: safe.BodySite,
		Identifier: safe.Identifier,
		Date: safe.Date,
		Language: safe.Language,
		Type: safe.Type,
		Category: safe.Category,
		Period: safe.Period,
		FacilityType: safe.FacilityType,
		ImplicitRules: safe.ImplicitRules,
		DocStatus: safe.DocStatus,
		Custodian: safe.Custodian,
		Description: safe.Description,
		RelatesTo: safe.RelatesTo,
		ResourceType: safe.ResourceType,
		Status: safe.Status,
		Text: safe.Text,
		Content: safe.Content,
		Version: safe.Version,
		Extension: safe.Extension,
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "author", safe.Author.Typename, &o.Author); err != nil {
		return fmt.Errorf("failed to unmarshal Author: %w", err)
	}
	if err := unmarshalUnion(b, "attester", safe.Attester.Typename, &o.Attester); err != nil {
		return fmt.Errorf("failed to unmarshal Attester: %w", err)
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}

	return nil
}
