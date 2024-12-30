
package model

import (
	"encoding/json"
	"fmt"
)

type SafeTaskType struct {
	InstantiatesURI *string `json:"instantiatesUri,omitempty"`
	Owner TypedObject `json:"owner"`
	ID *string `json:"id,omitempty"`
	Intent *string `json:"intent,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	RequestedPeriod *Period `json:"requestedPeriod,omitempty"`
	Restriction TypedObject `json:"restriction"`
	AuthoredOn *string `json:"authoredOn,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Note TypedObject `json:"note"`
	DoNotPerform *string `json:"doNotPerform,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Language *string `json:"language,omitempty"`
	Output []*TaskOutput `json:"output,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`
	Status *string `json:"status,omitempty"`
	BusinessStatus *CodeableConcept `json:"businessStatus,omitempty"`
	ExecutionPeriod *Period `json:"executionPeriod,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	PartOf *TaskType `json:"partOf"`
	StatusReason *CodeableReference `json:"statusReason,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	Performer TypedObject `json:"performer"`
	Input []*TaskInput `json:"input,omitempty"`
	Requester TypedObject `json:"requester"`
	Description *string `json:"description,omitempty"`
	GroupIdentifier *Identifier `json:"groupIdentifier,omitempty"`
	Focus TypedObject `json:"focus"`
	Priority *string `json:"priority,omitempty"`
	LastModified *string `json:"lastModified,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	RequestedPerformer []*CodeableReference `json:"requestedPerformer,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
}

func (o *TaskType) UnmarshalJSON(b []byte) error {
	var safe SafeTaskType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = TaskType{
		InstantiatesURI: safe.InstantiatesURI,
		ID: safe.ID,
		Intent: safe.Intent,
		Code: safe.Code,
		RequestedPeriod: safe.RequestedPeriod,
		AuthoredOn: safe.AuthoredOn,
		ModifierExtension: safe.ModifierExtension,
		DoNotPerform: safe.DoNotPerform,
		Text: safe.Text,
		Identifier: safe.Identifier,
		Language: safe.Language,
		Output: safe.Output,
		Extension: safe.Extension,
		InstantiatesCanonical: safe.InstantiatesCanonical,
		Status: safe.Status,
		BusinessStatus: safe.BusinessStatus,
		ExecutionPeriod: safe.ExecutionPeriod,
		Meta: safe.Meta,
		ImplicitRules: safe.ImplicitRules,
		PartOf: safe.PartOf,
		StatusReason: safe.StatusReason,
		Input: safe.Input,
		Description: safe.Description,
		GroupIdentifier: safe.GroupIdentifier,
		Priority: safe.Priority,
		LastModified: safe.LastModified,
		ResourceType: safe.ResourceType,
		RequestedPerformer: safe.RequestedPerformer,
		Reason: safe.Reason,
	}
	if err := unmarshalUnion(b, "owner", safe.Owner.Typename, &o.Owner); err != nil {
		return fmt.Errorf("failed to unmarshal Owner: %w", err)
	}
	if err := unmarshalUnion(b, "restriction", safe.Restriction.Typename, &o.Restriction); err != nil {
		return fmt.Errorf("failed to unmarshal Restriction: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "performer", safe.Performer.Typename, &o.Performer); err != nil {
		return fmt.Errorf("failed to unmarshal Performer: %w", err)
	}
	if err := unmarshalUnion(b, "requester", safe.Requester.Typename, &o.Requester); err != nil {
		return fmt.Errorf("failed to unmarshal Requester: %w", err)
	}
	if err := unmarshalUnion(b, "focus", safe.Focus.Typename, &o.Focus); err != nil {
		return fmt.Errorf("failed to unmarshal Focus: %w", err)
	}

	return nil
}
