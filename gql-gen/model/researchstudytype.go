
package model

import (
	"encoding/json"
	"fmt"
)

type SafeResearchStudyType struct {
	Title *string `json:"title,omitempty"`
	Classifier []*CodeableConcept `json:"classifier,omitempty"`
	AssociatedParty []*ResearchStudyAssociatedParty `json:"associatedParty,omitempty"`
	Label []*ResearchStudyLabel `json:"label,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Status *string `json:"status,omitempty"`
	Recruitment *ResearchStudyRecruitment `json:"recruitment,omitempty"`
	Keyword []*CodeableConcept `json:"keyword,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Objective []*ResearchStudyObjective `json:"objective,omitempty"`
	Identifier []*Identifier `json:"identifier,omitempty"`
	Region []*CodeableConcept `json:"region,omitempty"`
	Phase *CodeableConcept `json:"phase,omitempty"`
	PrimaryPurposeType *CodeableConcept `json:"primaryPurposeType,omitempty"`
	Site TypedObject `json:"site"`
	Name *string `json:"name,omitempty"`
	DescriptionSummary *string `json:"descriptionSummary,omitempty"`
	ComparisonGroup []*ResearchStudyComparisonGroup `json:"comparisonGroup,omitempty"`
	Language *string `json:"language,omitempty"`
	Note []*Annotation `json:"note,omitempty"`
	Date *string `json:"date,omitempty"`
	Condition []*CodeableConcept `json:"condition,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
	ModifierExtension []*Extension `json:"modifierExtension,omitempty"`
	Period *Period `json:"period,omitempty"`
	PartOf *ResearchStudyType `json:"partOf"`
	Extension []*Extension `json:"extension,omitempty"`
	Contained TypedObject `json:"contained,omitempty"`
	StudyDesign []*CodeableConcept `json:"studyDesign,omitempty"`
	Text *Narrative `json:"text,omitempty"`
	OutcomeMeasure []*ResearchStudyOutcomeMeasure `json:"outcomeMeasure,omitempty"`
	ID *string `json:"id,omitempty"`
	WhyStopped *CodeableConcept `json:"whyStopped,omitempty"`
	URL *string `json:"url,omitempty"`
	Description *string `json:"description,omitempty"`
	RelatedArtifact []*RelatedArtifact `json:"relatedArtifact,omitempty"`
	Focus []*CodeableReference `json:"focus,omitempty"`
	Result TypedObject `json:"result"`
	Version *string `json:"version,omitempty"`
	ProgressStatus []*ResearchStudyProgressStatus `json:"progressStatus,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *ResearchStudyType) UnmarshalJSON(b []byte) error {
	var safe SafeResearchStudyType
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = ResearchStudyType{
		Title: safe.Title,
		Classifier: safe.Classifier,
		AssociatedParty: safe.AssociatedParty,
		Label: safe.Label,
		ImplicitRules: safe.ImplicitRules,
		Status: safe.Status,
		Recruitment: safe.Recruitment,
		Keyword: safe.Keyword,
		ResourceType: safe.ResourceType,
		Objective: safe.Objective,
		Identifier: safe.Identifier,
		Region: safe.Region,
		Phase: safe.Phase,
		PrimaryPurposeType: safe.PrimaryPurposeType,
		Name: safe.Name,
		DescriptionSummary: safe.DescriptionSummary,
		ComparisonGroup: safe.ComparisonGroup,
		Language: safe.Language,
		Note: safe.Note,
		Date: safe.Date,
		Condition: safe.Condition,
		Meta: safe.Meta,
		ModifierExtension: safe.ModifierExtension,
		Period: safe.Period,
		PartOf: safe.PartOf,
		Extension: safe.Extension,
		StudyDesign: safe.StudyDesign,
		Text: safe.Text,
		OutcomeMeasure: safe.OutcomeMeasure,
		ID: safe.ID,
		WhyStopped: safe.WhyStopped,
		URL: safe.URL,
		Description: safe.Description,
		RelatedArtifact: safe.RelatedArtifact,
		Focus: safe.Focus,
		Version: safe.Version,
		ProgressStatus: safe.ProgressStatus,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "site", safe.Site.Typename, &o.Site); err != nil {
		return fmt.Errorf("failed to unmarshal Site: %w", err)
	}
	if err := unmarshalUnion(b, "contained", safe.Contained.Typename, &o.Contained); err != nil {
		return fmt.Errorf("failed to unmarshal Contained: %w", err)
	}
	if err := unmarshalUnion(b, "result", safe.Result.Typename, &o.Result); err != nil {
		return fmt.Errorf("failed to unmarshal Result: %w", err)
	}

	return nil
}
