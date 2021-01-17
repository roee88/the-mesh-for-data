// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"encoding/json"
	"fmt"

	externalRef0 "github.com/ibm/the-mesh-for-data/connectors/katalog/pkg/taxonomy"
	"github.com/pkg/errors"
)

// Asset defines model for Asset.
type Asset struct {
	Spec AssetSpec `json:"spec"`
}

// AssetDetails defines model for AssetDetails.
type AssetDetails struct {

	// Connection information
	Connection externalRef0.Connection `json:"connection"`
	DataFormat *string                 `json:"dataFormat,omitempty"`
}

// AssetMetadata defines model for AssetMetadata.
type AssetMetadata struct {

	// metadata for each component in asset (e.g., column)
	ComponentsMetadata *AssetMetadata_ComponentsMetadata `json:"componentsMetadata,omitempty"`
	Geography          *string                           `json:"geography,omitempty"`
	NamedMetadata      *AssetMetadata_NamedMetadata      `json:"namedMetadata,omitempty"`
	Owner              *string                           `json:"owner,omitempty"`

	// Tags associated with the asset
	Tags *[]string `json:"tags,omitempty"`
}

// AssetMetadata_ComponentsMetadata defines model for AssetMetadata.ComponentsMetadata.
type AssetMetadata_ComponentsMetadata struct {
	AdditionalProperties map[string]ComponentMetadata `json:"-"`
}

// AssetMetadata_NamedMetadata defines model for AssetMetadata.NamedMetadata.
type AssetMetadata_NamedMetadata struct {
	AdditionalProperties map[string]string `json:"-"`
}

// AssetSpec defines model for AssetSpec.
type AssetSpec struct {

	// Asset details
	Details AssetDetails `json:"details"`

	// Reference to a Secret resource holding credentials for this asset
	SecretRef SecretRef     `json:"secretRef"`
	Security  AssetMetadata `json:"security"`
}

// ComponentMetadata defines model for ComponentMetadata.
type ComponentMetadata struct {
	ComponentType *string `json:"componentType,omitempty"`

	// Named terms, that exist in Catalog toxonomy and the values for these terms for columns we will have "SchemaDetails" key, that will include technical schema details for this column TODO: Consider create special field for schema outside of metadata
	NamedMetadata *ComponentMetadata_NamedMetadata `json:"namedMetadata,omitempty"`

	// Tags - can be any free text added to a component (no taxonomy)
	Tags *[]string `json:"tags,omitempty"`
}

// ComponentMetadata_NamedMetadata defines model for ComponentMetadata.NamedMetadata.
type ComponentMetadata_NamedMetadata struct {
	AdditionalProperties map[string]string `json:"-"`
}

// SecretRef defines model for SecretRef.
type SecretRef struct {

	// Name of the Secret resource (must exist in the same namespace)
	Name string `json:"name"`
}

// Getter for additional properties for AssetMetadata_ComponentsMetadata. Returns the specified
// element and whether it was found
func (a AssetMetadata_ComponentsMetadata) Get(fieldName string) (value ComponentMetadata, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AssetMetadata_ComponentsMetadata
func (a *AssetMetadata_ComponentsMetadata) Set(fieldName string, value ComponentMetadata) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]ComponentMetadata)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AssetMetadata_ComponentsMetadata to handle AdditionalProperties
func (a *AssetMetadata_ComponentsMetadata) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]ComponentMetadata)
		for fieldName, fieldBuf := range object {
			var fieldVal ComponentMetadata
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AssetMetadata_ComponentsMetadata to handle AdditionalProperties
func (a AssetMetadata_ComponentsMetadata) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AssetMetadata_NamedMetadata. Returns the specified
// element and whether it was found
func (a AssetMetadata_NamedMetadata) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AssetMetadata_NamedMetadata
func (a *AssetMetadata_NamedMetadata) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AssetMetadata_NamedMetadata to handle AdditionalProperties
func (a *AssetMetadata_NamedMetadata) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AssetMetadata_NamedMetadata to handle AdditionalProperties
func (a AssetMetadata_NamedMetadata) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for ComponentMetadata_NamedMetadata. Returns the specified
// element and whether it was found
func (a ComponentMetadata_NamedMetadata) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ComponentMetadata_NamedMetadata
func (a *ComponentMetadata_NamedMetadata) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ComponentMetadata_NamedMetadata to handle AdditionalProperties
func (a *ComponentMetadata_NamedMetadata) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ComponentMetadata_NamedMetadata to handle AdditionalProperties
func (a ComponentMetadata_NamedMetadata) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}
