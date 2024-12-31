package graph

import (
	"encoding/json"
	"fmt"
)

// Custom JSON type handler for json filter query arguments
type JSON map[string]interface{}

func (j *JSON) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, j); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return nil
}

func (j JSON) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}(j))
}
