
package model

import (
	"encoding/json"
	"fmt"
)

type SafeRelatedArtifact struct {
	Extension []*Extension `json:"extension,omitempty"`
	PublicationDate *string `json:"publicationDate,omitempty"`
	PublicationStatus *string `json:"publicationStatus,omitempty"`
	Type *string `json:"type,omitempty"`
	Citation *string `json:"citation,omitempty"`
	Document *Attachment `json:"document,omitempty"`
	Resource *string `json:"resource,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	ID *string `json:"id,omitempty"`
	ResourceReference TypedObject `json:"resourceReference"`
	Display *string `json:"display,omitempty"`
	Label *string `json:"label,omitempty"`
	Classifier []*CodeableConcept `json:"classifier,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *RelatedArtifact) UnmarshalJSON(b []byte) error {
	var safe SafeRelatedArtifact
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = RelatedArtifact{
		Extension: safe.Extension,
		PublicationDate: safe.PublicationDate,
		PublicationStatus: safe.PublicationStatus,
		Type: safe.Type,
		Citation: safe.Citation,
		Document: safe.Document,
		Resource: safe.Resource,
		ResourceType: safe.ResourceType,
		ID: safe.ID,
		Display: safe.Display,
		Label: safe.Label,
		Classifier: safe.Classifier,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "resourceReference", safe.ResourceReference.Typename, &o.ResourceReference); err != nil {
		return fmt.Errorf("failed to unmarshal ResourceReference: %w", err)
	}

	return nil
}
