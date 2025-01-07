
package model

import (
	"encoding/json"
	"fmt"
)

type SafeRelatedArtifact struct {
	Citation *string `json:"citation,omitempty"`
	ID *string `json:"id,omitempty"`
	Display *string `json:"display,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ResourceReference TypedObject `json:"resourceReference"`
	PublicationStatus *string `json:"publicationStatus,omitempty"`
	Type *string `json:"type,omitempty"`
	Document *Attachment `json:"document,omitempty"`
	PublicationDate *string `json:"publicationDate,omitempty"`
	Label *string `json:"label,omitempty"`
	Resource *string `json:"resource,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	Classifier []*CodeableConcept `json:"classifier,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *RelatedArtifact) UnmarshalJSON(b []byte) error {
	var safe SafeRelatedArtifact
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = RelatedArtifact{
		Citation: safe.Citation,
		ID: safe.ID,
		Display: safe.Display,
		Extension: safe.Extension,
		PublicationStatus: safe.PublicationStatus,
		Type: safe.Type,
		Document: safe.Document,
		PublicationDate: safe.PublicationDate,
		Label: safe.Label,
		Resource: safe.Resource,
		ResourceType: safe.ResourceType,
		Classifier: safe.Classifier,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "resourceReference", safe.ResourceReference.Typename, &o.ResourceReference); err != nil {
		return fmt.Errorf("failed to unmarshal ResourceReference: %w", err)
	}

	return nil
}
