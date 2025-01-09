
package model

import (
	"encoding/json"
	"fmt"
)

type SafeResearchSubjectType struct {
	Language *string `json:"language,omitempty"`
	Progress []*ResearchSubjectProgress `json:"progress,omitempty"`
	ActualComparisonGroup *string `json:"actualComparisonGroup,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	ID *string `json:"id,omitempty"`
	Period *Period `json:"period,omitempty"`
	Status *string `json:"status,omitempty"`
	Subject TypedObject `json:"subject"`
	Contained TypedObject `json:"contained,omitempty"`
	Study *ResearchStudyType `json:"study"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	AssignedComparisonGroup *string `json:"assignedComparisonGroup,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ResearchSubjectType) UnmarshalJSON(b []byte) error {
	var safe SafeResearchSubjectType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ResearchSubjectType{
		Language: safe.Language,
		Progress: safe.Progress,
		ActualComparisonGroup: safe.ActualComparisonGroup,
		ImplicitRules: safe.ImplicitRules,
		ResourceType: safe.ResourceType,
		Text: safe.Text,
		ID: safe.ID,
		Period: safe.Period,
		Status: safe.Status,
		Study: safe.Study,
		Identifier: safe.Identifier,
		Meta: safe.Meta,
		Extension: safe.Extension,
		ModifierExtension: safe.ModifierExtension,
		AssignedComparisonGroup: safe.AssignedComparisonGroup,
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
