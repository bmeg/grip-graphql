
package model

import (
	"encoding/json"
	"fmt"
)

type SafeResearchStudyType struct {
	StudyDesign []*CodeableConcept `json:"studyDesign,omitempty"`
	Recruitment TypedObject `json:"recruitment"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Note TypedObject `json:"note"`
	Site TypedObject `json:"site"`
	Classifier []*CodeableConcept `json:"classifier,omitempty"`
	Region []*CodeableConcept `json:"region,omitempty"`
	Objective []*ResearchStudyObjective `json:"objective,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	Version *string `json:"version,omitempty"`
	Description *string `json:"description,omitempty"`
	ComparisonGroup []*ResearchStudyComparisonGroup `json:"comparisonGroup,omitempty"`
	OutcomeMeasure []*ResearchStudyOutcomeMeasure `json:"outcomeMeasure,omitempty"`
	Language *string `json:"language,omitempty"`
	PrimaryPurposeType *CodeableConcept `json:"primaryPurposeType,omitempty"`
	ProgressStatus []*ResearchStudyProgressStatus `json:"progressStatus,omitempty"`
	Condition []*CodeableConcept `json:"condition,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	Title *string `json:"title,omitempty"`
	Period *Period `json:"period,omitempty"`
	Name *string `json:"name,omitempty"`
	AssociatedParty TypedObject `json:"associatedParty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Phase *CodeableConcept `json:"phase,omitempty"`
	Result TypedObject `json:"result"`
	Text *Narrative `json:"text,omitempty"`
	RelatedArtifact []*RelatedArtifact `json:"relatedArtifact,omitempty"`
	URL *string `json:"url,omitempty"`
	Status *string `json:"status,omitempty"`
	PartOf *ResearchStudyType `json:"partOf"`
	Date *string `json:"date,omitempty"`
	Keyword []*CodeableConcept `json:"keyword,omitempty"`
	ID *string `json:"id,omitempty"`
	Focus []*CodeableReference `json:"focus,omitempty"`
	Label []*ResearchStudyLabel `json:"label,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	WhyStopped *CodeableConcept `json:"whyStopped,omitempty"`
	DescriptionSummary *string `json:"descriptionSummary,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ResearchStudyType) UnmarshalJSON(b []byte) error {
	var safe SafeResearchStudyType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ResearchStudyType{
		StudyDesign: safe.StudyDesign,
		Identifier: safe.Identifier,
		Classifier: safe.Classifier,
		Region: safe.Region,
		Objective: safe.Objective,
		Meta: safe.Meta,
		Version: safe.Version,
		Description: safe.Description,
		ComparisonGroup: safe.ComparisonGroup,
		OutcomeMeasure: safe.OutcomeMeasure,
		Language: safe.Language,
		PrimaryPurposeType: safe.PrimaryPurposeType,
		ProgressStatus: safe.ProgressStatus,
		Condition: safe.Condition,
		Extension: safe.Extension,
		Title: safe.Title,
		Period: safe.Period,
		Name: safe.Name,
		ImplicitRules: safe.ImplicitRules,
		Phase: safe.Phase,
		Text: safe.Text,
		RelatedArtifact: safe.RelatedArtifact,
		URL: safe.URL,
		Status: safe.Status,
		PartOf: safe.PartOf,
		Date: safe.Date,
		Keyword: safe.Keyword,
		ID: safe.ID,
		Focus: safe.Focus,
		Label: safe.Label,
		WhyStopped: safe.WhyStopped,
		DescriptionSummary: safe.DescriptionSummary,
		ModifierExtension: safe.ModifierExtension,
		ResourceType: safe.ResourceType,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "recruitment", safe.Recruitment.Typename, &o.Recruitment); err != nil {
		return fmt.Errorf("failed to unmarshal Recruitment: %w", err)
	}
	if err := unmarshalUnion(b, "note", safe.Note.Typename, &o.Note); err != nil {
		return fmt.Errorf("failed to unmarshal Note: %w", err)
	}
	if err := unmarshalUnion(b, "site", safe.Site.Typename, &o.Site); err != nil {
		return fmt.Errorf("failed to unmarshal Site: %w", err)
	}
	if err := unmarshalUnion(b, "associatedParty", safe.AssociatedParty.Typename, &o.AssociatedParty); err != nil {
		return fmt.Errorf("failed to unmarshal AssociatedParty: %w", err)
	}
	if err := unmarshalUnion(b, "result", safe.Result.Typename, &o.Result); err != nil {
		return fmt.Errorf("failed to unmarshal Result: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}

	return nil
}
