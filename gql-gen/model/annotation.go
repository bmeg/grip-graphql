
package model

import (
	"encoding/json"
	"fmt"
)

type SafeAnnotation struct {
	Text *string `json:"text,omitempty"`
	Time *string `json:"time,omitempty"`
	AuthorReference TypedObject `json:"authorReference"`
	AuthorString *string `json:"authorString,omitempty"`
	Extension []*Extension `json:"extension,omitempty"`
	ID *string `json:"id,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	AuthResourcePath *string `json:"auth_resource_path,omitempty"`
}

func (o *Annotation) UnmarshalJSON(b []byte) error {
	var safe SafeAnnotation
	if err := json.Unmarshal(b, &safe); err != nil {
		return err
	}

	*o = Annotation{
		Text: safe.Text,
		Time: safe.Time,
		AuthorString: safe.AuthorString,
		Extension: safe.Extension,
		ID: safe.ID,
		ResourceType: safe.ResourceType,
		AuthResourcePath: safe.AuthResourcePath,
	}
	if err := unmarshalUnion(b, "authorReference", safe.AuthorReference.Typename, &o.AuthorReference); err != nil {
		return fmt.Errorf("failed to unmarshal AuthorReference: %w", err)
	}

	return nil
}
