
package model

import (
	"encoding/json"
	"fmt"
)

type SafeTaskType struct {
	Meta *Meta `json:"meta,omitempty"`
	ExecutionPeriod *Period `json:"executionPeriod,omitempty"`
	BusinessStatus *CodeableConcept `json:"businessStatus,omitempty"`
	RequestedPerformer []*CodeableReference `json:"requestedPerformer,omitempty"`
	RequestedPeriod *Period `json:"requestedPeriod,omitempty"`
	Owner TypedObject `json:"owner"`
	ResourceType *string `json:"resourceType,omitempty"`
	Intent *string `json:"intent,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	PartOf *TaskType `json:"partOf"`
	Text *Narrative `json:"text,omitempty"`
	Restriction *TaskRestriction `json:"restriction,omitempty"`
	StatusReason *CodeableReference `json:"statusReason,omitempty"`
	AuthoredOn *string `json:"authoredOn,omitempty"`
	BasedOn TypedObject `json:"basedOn"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Output []*TaskOutput `json:"output,omitempty"`
	Priority *string `json:"priority,omitempty"`
	DoNotPerform *bool `json:"doNotPerform,omitempty"`
	LastModified *string `json:"lastModified,omitempty"`
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`
	Code *CodeableConcept `json:"code,omitempty"`
	Status *string `json:"status,omitempty"`
	GroupIdentifier *Identifier `json:"groupIdentifier,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Input []*TaskInput `json:"input,omitempty"`
	Focus TypedObject `json:"focus"`
	Requester TypedObject `json:"requester"`
	InstantiatesURI *string `json:"instantiatesUri,omitempty"`
	Language *string `json:"language,omitempty"`
	Reason []*CodeableReference `json:"reason,omitempty"`
	ID *string `json:"id,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Performer []*TaskPerformer `json:"performer,omitempty"`
	Description *string `json:"description,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *TaskType) UnmarshalJSON(b []byte) error {
	var safe SafeTaskType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = TaskType{
		Meta: safe.Meta,
		ExecutionPeriod: safe.ExecutionPeriod,
		BusinessStatus: safe.BusinessStatus,
		RequestedPerformer: safe.RequestedPerformer,
		RequestedPeriod: safe.RequestedPeriod,
		ResourceType: safe.ResourceType,
		Intent: safe.Intent,
		ModifierExtension: safe.ModifierExtension,
		PartOf: safe.PartOf,
		Text: safe.Text,
		Restriction: safe.Restriction,
		StatusReason: safe.StatusReason,
		AuthoredOn: safe.AuthoredOn,
		Identifier: safe.Identifier,
		Output: safe.Output,
		Priority: safe.Priority,
		DoNotPerform: safe.DoNotPerform,
		LastModified: safe.LastModified,
		InstantiatesCanonical: safe.InstantiatesCanonical,
		Code: safe.Code,
		Status: safe.Status,
		GroupIdentifier: safe.GroupIdentifier,
		ImplicitRules: safe.ImplicitRules,
		Input: safe.Input,
		InstantiatesURI: safe.InstantiatesURI,
		Language: safe.Language,
		Reason: safe.Reason,
		ID: safe.ID,
		Note: safe.Note,
		Extension: safe.Extension,
		Performer: safe.Performer,
		Description: safe.Description,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "owner", safe.Owner.Typename, &o.Owner); err != nil {
		return fmt.Errorf("failed to unmarshal Owner: %w", err)
	}
	if err := unmarshalUnion(b, "basedOn", safe.BasedOn.Typename, &o.BasedOn); err != nil {
		return fmt.Errorf("failed to unmarshal BasedOn: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "focus", safe.Focus.Typename, &o.Focus); err != nil {
		return fmt.Errorf("failed to unmarshal Focus: %w", err)
	}
	if err := unmarshalUnion(b, "requester", safe.Requester.Typename, &o.Requester); err != nil {
		return fmt.Errorf("failed to unmarshal Requester: %w", err)
	}

	return nil
}
