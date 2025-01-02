
package model

import (
	"encoding/json"
	"fmt"
)

type SafeTaskType struct {
	Input []*TaskInput `json:"input,omitempty"`
	Owner TypedObject `json:"owner"`
	Meta *Meta `json:"meta,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Requester TypedObject `json:"requester"`
	Restriction TypedObject `json:"restriction"`
	PartOf *TaskType `json:"partOf"`
	ExecutionPeriod *Period `json:"executionPeriod,omitempty"`
	Priority *string `json:"priority,omitempty"`
	ID *string `json:"id,omitempty"`
	Performer TypedObject `json:"performer"`
	ResourceType *string `json:"resourceType,omitempty"`
	DoNotPerform *string `json:"doNotPerform,omitempty"`
	Output []*TaskOutput `json:"output,omitempty"`
	StatusReason *CodeableReference `json:"statusReason,omitempty"`
	Note TypedObject `json:"note"`
	LastModified *string `json:"lastModified,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	Description *string `json:"description,omitempty"`
	Intent *string `json:"intent,omitempty"`
	Status *string `json:"status,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	BusinessStatus *CodeableConcept `json:"businessStatus,omitempty"`
	Focus TypedObject `json:"focus"`
	RequestedPeriod *Period `json:"requestedPeriod,omitempty"`
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`
	GroupIdentifier *Identifier `json:"groupIdentifier,omitempty"`
	Language *string `json:"language,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	AuthoredOn *string `json:"authoredOn,omitempty"`
	InstantiatesURI *string `json:"instantiatesUri,omitempty"`
	RequestedPerformer []*CodeableReference `json:"requestedPerformer,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *TaskType) UnmarshalJSON(b []byte) error {
	var safe SafeTaskType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = TaskType{
		Input: safe.Input,
		Meta: safe.Meta,
		Code: safe.Code,
		Extension: safe.Extension,
		ModifierExtension: safe.ModifierExtension,
		PartOf: safe.PartOf,
		ExecutionPeriod: safe.ExecutionPeriod,
		Priority: safe.Priority,
		ID: safe.ID,
		ResourceType: safe.ResourceType,
		DoNotPerform: safe.DoNotPerform,
		Output: safe.Output,
		StatusReason: safe.StatusReason,
		LastModified: safe.LastModified,
		Reason: safe.Reason,
		Description: safe.Description,
		Intent: safe.Intent,
		Status: safe.Status,
		Identifier: safe.Identifier,
		BusinessStatus: safe.BusinessStatus,
		RequestedPeriod: safe.RequestedPeriod,
		InstantiatesCanonical: safe.InstantiatesCanonical,
		GroupIdentifier: safe.GroupIdentifier,
		Language: safe.Language,
		Text: safe.Text,
		AuthoredOn: safe.AuthoredOn,
		InstantiatesURI: safe.InstantiatesURI,
		RequestedPerformer: safe.RequestedPerformer,
		ImplicitRules: safe.ImplicitRules,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "owner", safe.Owner.Typename, &o.Owner); err != nil {
		return fmt.Errorf("failed to unmarshal Owner: %w", err)
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "requester", safe.Requester.Typename, &o.Requester); err != nil {
		return fmt.Errorf("failed to unmarshal Requester: %w", err)
	}
	if err := unmarshalUnion(b, "restriction", safe.Restriction.Typename, &o.Restriction); err != nil {
		return fmt.Errorf("failed to unmarshal Restriction: %w", err)
	}
	if err := unmarshalUnion(b, "performer", safe.Performer.Typename, &o.Performer); err != nil {
		return fmt.Errorf("failed to unmarshal Performer: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "focus", safe.Focus.Typename, &o.Focus); err != nil {
		return fmt.Errorf("failed to unmarshal Focus: %w", err)
	}

	return nil
}
