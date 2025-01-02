
package model

import (
	"encoding/json"
	"fmt"
)

type SafeDocumentReferenceType struct {
	BodySite []*CodeableReference `json:"bodySite,omitempty"`
	Author TypedObject `json:"author"`
	Language *string `json:"language,omitempty"`
	Version *string `json:"version,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Type *CodeableConcept `json:"type,omitempty"`
	Status *string `json:"status,omitempty"`
	Modality []*CodeableConcept `json:"modality,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Attester TypedObject `json:"attester"`
	Date *string `json:"date,omitempty"`
	Event []*CodeableReference `json:"event,omitempty"`
	DocStatus *string `json:"docStatus,omitempty"`
	Content []*DocumentReferenceContent `json:"content,omitempty"`
	Custodian *OrganizationType `json:"custodian"`
	Description *string `json:"description,omitempty"`
	ID *string `json:"id,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Subject TypedObject `json:"subject"`
	SecurityLabel []*CodeableConcept `json:"securityLabel,omitempty"`
	Period *Period `json:"period,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	PracticeSetting *CodeableConcept `json:"practiceSetting,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	RelatesTo []*DocumentReferenceRelatesTo `json:"relatesTo,omitempty"`
	Category []*CodeableConcept `json:"category,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	FacilityType *CodeableConcept `json:"facilityType,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *DocumentReferenceType) UnmarshalJSON(b []byte) error {
	var safe SafeDocumentReferenceType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = DocumentReferenceType{
		BodySite: safe.BodySite,
		Language: safe.Language,
		Version: safe.Version,
		ImplicitRules: safe.ImplicitRules,
		Type: safe.Type,
		Status: safe.Status,
		Modality: safe.Modality,
		Meta: safe.Meta,
		Date: safe.Date,
		Event: safe.Event,
		DocStatus: safe.DocStatus,
		Content: safe.Content,
		Custodian: safe.Custodian,
		Description: safe.Description,
		ID: safe.ID,
		ResourceType: safe.ResourceType,
		ModifierExtension: safe.ModifierExtension,
		SecurityLabel: safe.SecurityLabel,
		Period: safe.Period,
		Extension: safe.Extension,
		PracticeSetting: safe.PracticeSetting,
		Text: safe.Text,
		RelatesTo: safe.RelatesTo,
		Category: safe.Category,
		Identifier: safe.Identifier,
		FacilityType: safe.FacilityType,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "author", safe.Author.Typename, &o.Author); err != nil {
		return fmt.Errorf("failed to unmarshal Author: %w", err)
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "attester", safe.Attester.Typename, &o.Attester); err != nil {
		return fmt.Errorf("failed to unmarshal Attester: %w", err)
	}
	if err := unmarshalUnion(b, "subject", safe.Subject.Typename, &o.Subject); err != nil {
		return fmt.Errorf("failed to unmarshal Subject: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
