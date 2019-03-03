package models

import "encoding/json"

// Reference maps AsyncAPI "Reference" object
type Reference struct {
  Ref string `json:"$ref,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Reference) UnmarshalJSON(data []byte) error {
	type Alias Reference
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Ref = jsonMap.Ref


	return nil
}

// MarshalJSON marshals JSON
func (value Reference) MarshalJSON() ([]byte, error) {
	type Alias Reference
	jsonByteArray, err := json.Marshal(&Alias{
    Ref: value.Ref,
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// Info maps AsyncAPI "info" object
type Info struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Title string `json:"title,omitempty"`
  Version string `json:"version,omitempty"`
  Description string `json:"description,omitempty"`
  TermsOfService string `json:"termsOfService,omitempty"`
  Contact *Contact `json:"contact,omitempty"`
  License *License `json:"license,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Info) UnmarshalJSON(data []byte) error {
	type Alias Info
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Title = jsonMap.Title
  value.Version = jsonMap.Version
  value.Description = jsonMap.Description
  value.TermsOfService = jsonMap.TermsOfService
  value.Contact = jsonMap.Contact
  value.License = jsonMap.License

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Info) MarshalJSON() ([]byte, error) {
	type Alias Info
	jsonByteArray, err := json.Marshal(&Alias{
    Title: value.Title,
    Version: value.Version,
    Description: value.Description,
    TermsOfService: value.TermsOfService,
    Contact: value.Contact,
    License: value.License,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// Contact maps AsyncAPI "contact" object
type Contact struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Name string `json:"name,omitempty"`
  Url string `json:"url,omitempty"`
  Email string `json:"email,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Contact) UnmarshalJSON(data []byte) error {
	type Alias Contact
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Name = jsonMap.Name
  value.Url = jsonMap.Url
  value.Email = jsonMap.Email

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Contact) MarshalJSON() ([]byte, error) {
	type Alias Contact
	jsonByteArray, err := json.Marshal(&Alias{
    Name: value.Name,
    Url: value.Url,
    Email: value.Email,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// License maps AsyncAPI "license" object
type License struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Name string `json:"name,omitempty"`
  Url string `json:"url,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *License) UnmarshalJSON(data []byte) error {
	type Alias License
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Name = jsonMap.Name
  value.Url = jsonMap.Url

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value License) MarshalJSON() ([]byte, error) {
	type Alias License
	jsonByteArray, err := json.Marshal(&Alias{
    Name: value.Name,
    Url: value.Url,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// Server maps AsyncAPI "server" object
type Server struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Url string `json:"url,omitempty"`
  Description string `json:"description,omitempty"`
  Scheme string `json:"scheme,omitempty"`
  SchemeVersion string `json:"schemeVersion,omitempty"`
  Variables *ServerVariables `json:"variables,omitempty"`
  BaseChannel string `json:"baseChannel,omitempty"`
  Security []*SecurityRequirement `json:"security,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Server) UnmarshalJSON(data []byte) error {
	type Alias Server
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Url = jsonMap.Url
  value.Description = jsonMap.Description
  value.Scheme = jsonMap.Scheme
  value.SchemeVersion = jsonMap.SchemeVersion
  value.Variables = jsonMap.Variables
  value.BaseChannel = jsonMap.BaseChannel
  value.Security = jsonMap.Security

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Server) MarshalJSON() ([]byte, error) {
	type Alias Server
	jsonByteArray, err := json.Marshal(&Alias{
    Url: value.Url,
    Description: value.Description,
    Scheme: value.Scheme,
    SchemeVersion: value.SchemeVersion,
    Variables: value.Variables,
    BaseChannel: value.BaseChannel,
    Security: value.Security,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// ServerVariables maps AsyncAPI "serverVariables" object
type ServerVariables map[string]*ServerVariable

// ServerVariable maps AsyncAPI "serverVariable" object
type ServerVariable struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Enum []string `json:"enum,omitempty"`
  Default string `json:"default,omitempty"`
  Description string `json:"description,omitempty"`
  Examples []string `json:"examples,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *ServerVariable) UnmarshalJSON(data []byte) error {
	type Alias ServerVariable
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Enum = jsonMap.Enum
  value.Default = jsonMap.Default
  value.Description = jsonMap.Description
  value.Examples = jsonMap.Examples

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value ServerVariable) MarshalJSON() ([]byte, error) {
	type Alias ServerVariable
	jsonByteArray, err := json.Marshal(&Alias{
    Enum: value.Enum,
    Default: value.Default,
    Description: value.Description,
    Examples: value.Examples,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// Channels maps AsyncAPI "channels" object
type Channels struct {
}

// UnmarshalJSON unmarshals JSON
func (value *Channels) UnmarshalJSON(data []byte) error {
	type Alias Channels
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value Channels) MarshalJSON() ([]byte, error) {
	type Alias Channels
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// Components maps AsyncAPI "components" object
type Components struct {
  Schemas *Schemas `json:"schemas,omitempty"`
  Messages *Messages `json:"messages,omitempty"`
  SecuritySchemes string `json:"securitySchemes,omitempty"`
  Parameters *Parameters `json:"parameters,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Components) UnmarshalJSON(data []byte) error {
	type Alias Components
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Schemas = jsonMap.Schemas
  value.Messages = jsonMap.Messages
  value.SecuritySchemes = jsonMap.SecuritySchemes
  value.Parameters = jsonMap.Parameters


	return nil
}

// MarshalJSON marshals JSON
func (value Components) MarshalJSON() ([]byte, error) {
	type Alias Components
	jsonByteArray, err := json.Marshal(&Alias{
    Schemas: value.Schemas,
    Messages: value.Messages,
    SecuritySchemes: value.SecuritySchemes,
    Parameters: value.Parameters,
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// Schemas maps AsyncAPI "schemas" object
type Schemas map[string]*Schema

// Messages maps AsyncAPI "messages" object
type Messages map[string]*Message

// Parameters maps AsyncAPI "parameters" object
type Parameters map[string]*Parameter

// Schema maps AsyncAPI "schema" object
type Schema struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Ref string `json:"$ref,omitempty"`
  Format string `json:"format,omitempty"`
  Title *Title `json:"title,omitempty"`
  Description *Description `json:"description,omitempty"`
  Default *Default `json:"default,omitempty"`
  MultipleOf *MultipleOf `json:"multipleOf,omitempty"`
  Maximum *Maximum `json:"maximum,omitempty"`
  ExclusiveMaximum *ExclusiveMaximum `json:"exclusiveMaximum,omitempty"`
  Minimum *Minimum `json:"minimum,omitempty"`
  ExclusiveMinimum *ExclusiveMinimum `json:"exclusiveMinimum,omitempty"`
  MaxLength *PositiveInteger `json:"maxLength,omitempty"`
  MinLength *PositiveIntegerDefault0 `json:"minLength,omitempty"`
  Pattern *Pattern `json:"pattern,omitempty"`
  MaxItems *PositiveInteger `json:"maxItems,omitempty"`
  MinItems *PositiveIntegerDefault0 `json:"minItems,omitempty"`
  UniqueItems *UniqueItems `json:"uniqueItems,omitempty"`
  MaxProperties *PositiveInteger `json:"maxProperties,omitempty"`
  MinProperties *PositiveIntegerDefault0 `json:"minProperties,omitempty"`
  Required *StringArray `json:"required,omitempty"`
  Enum *Enum `json:"enum,omitempty"`
  Deprecated bool `json:"deprecated,omitempty"`
  AdditionalProperties string `json:"additionalProperties,omitempty"`
  Type *Type `json:"type,omitempty"`
  Items string `json:"items,omitempty"`
  AllOf []*Schema `json:"allOf,omitempty"`
  OneOf []*Schema `json:"oneOf,omitempty"`
  AnyOf []*Schema `json:"anyOf,omitempty"`
  Not *Schema `json:"not,omitempty"`
  Properties string `json:"properties,omitempty"`
  Discriminator string `json:"discriminator,omitempty"`
  ReadOnly bool `json:"readOnly,omitempty"`
  Xml *Xml `json:"xml,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  Example string `json:"example,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Schema) UnmarshalJSON(data []byte) error {
	type Alias Schema
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Ref = jsonMap.Ref
  value.Format = jsonMap.Format
  value.Title = jsonMap.Title
  value.Description = jsonMap.Description
  value.Default = jsonMap.Default
  value.MultipleOf = jsonMap.MultipleOf
  value.Maximum = jsonMap.Maximum
  value.ExclusiveMaximum = jsonMap.ExclusiveMaximum
  value.Minimum = jsonMap.Minimum
  value.ExclusiveMinimum = jsonMap.ExclusiveMinimum
  value.MaxLength = jsonMap.MaxLength
  value.MinLength = jsonMap.MinLength
  value.Pattern = jsonMap.Pattern
  value.MaxItems = jsonMap.MaxItems
  value.MinItems = jsonMap.MinItems
  value.UniqueItems = jsonMap.UniqueItems
  value.MaxProperties = jsonMap.MaxProperties
  value.MinProperties = jsonMap.MinProperties
  value.Required = jsonMap.Required
  value.Enum = jsonMap.Enum
  value.Deprecated = jsonMap.Deprecated
  value.AdditionalProperties = jsonMap.AdditionalProperties
  value.Type = jsonMap.Type
  value.Items = jsonMap.Items
  value.AllOf = jsonMap.AllOf
  value.OneOf = jsonMap.OneOf
  value.AnyOf = jsonMap.AnyOf
  value.Not = jsonMap.Not
  value.Properties = jsonMap.Properties
  value.Discriminator = jsonMap.Discriminator
  value.ReadOnly = jsonMap.ReadOnly
  value.Xml = jsonMap.Xml
  value.ExternalDocs = jsonMap.ExternalDocs
  value.Example = jsonMap.Example

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Schema) MarshalJSON() ([]byte, error) {
	type Alias Schema
	jsonByteArray, err := json.Marshal(&Alias{
    Ref: value.Ref,
    Format: value.Format,
    Title: value.Title,
    Description: value.Description,
    Default: value.Default,
    MultipleOf: value.MultipleOf,
    Maximum: value.Maximum,
    ExclusiveMaximum: value.ExclusiveMaximum,
    Minimum: value.Minimum,
    ExclusiveMinimum: value.ExclusiveMinimum,
    MaxLength: value.MaxLength,
    MinLength: value.MinLength,
    Pattern: value.Pattern,
    MaxItems: value.MaxItems,
    MinItems: value.MinItems,
    UniqueItems: value.UniqueItems,
    MaxProperties: value.MaxProperties,
    MinProperties: value.MinProperties,
    Required: value.Required,
    Enum: value.Enum,
    Deprecated: value.Deprecated,
    AdditionalProperties: value.AdditionalProperties,
    Type: value.Type,
    Items: value.Items,
    AllOf: value.AllOf,
    OneOf: value.OneOf,
    AnyOf: value.AnyOf,
    Not: value.Not,
    Properties: value.Properties,
    Discriminator: value.Discriminator,
    ReadOnly: value.ReadOnly,
    Xml: value.Xml,
    ExternalDocs: value.ExternalDocs,
    Example: value.Example,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// Xml maps AsyncAPI "xml" object
type Xml struct {
  Name string `json:"name,omitempty"`
  Namespace string `json:"namespace,omitempty"`
  Prefix string `json:"prefix,omitempty"`
  Attribute bool `json:"attribute,omitempty"`
  Wrapped bool `json:"wrapped,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Xml) UnmarshalJSON(data []byte) error {
	type Alias Xml
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Name = jsonMap.Name
  value.Namespace = jsonMap.Namespace
  value.Prefix = jsonMap.Prefix
  value.Attribute = jsonMap.Attribute
  value.Wrapped = jsonMap.Wrapped


	return nil
}

// MarshalJSON marshals JSON
func (value Xml) MarshalJSON() ([]byte, error) {
	type Alias Xml
	jsonByteArray, err := json.Marshal(&Alias{
    Name: value.Name,
    Namespace: value.Namespace,
    Prefix: value.Prefix,
    Attribute: value.Attribute,
    Wrapped: value.Wrapped,
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// ExternalDocs maps AsyncAPI "externalDocs" object
type ExternalDocs struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Description string `json:"description,omitempty"`
  Url string `json:"url,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *ExternalDocs) UnmarshalJSON(data []byte) error {
	type Alias ExternalDocs
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Description = jsonMap.Description
  value.Url = jsonMap.Url

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value ExternalDocs) MarshalJSON() ([]byte, error) {
	type Alias ExternalDocs
	jsonByteArray, err := json.Marshal(&Alias{
    Description: value.Description,
    Url: value.Url,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// ChannelItem maps AsyncAPI "channelItem" object
type ChannelItem struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Ref string `json:"$ref,omitempty"`
  Parameters []*Parameter `json:"parameters,omitempty"`
  Publish *Operation `json:"publish,omitempty"`
  Subscribe *Operation `json:"subscribe,omitempty"`
  Deprecated bool `json:"deprecated,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *ChannelItem) UnmarshalJSON(data []byte) error {
	type Alias ChannelItem
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Ref = jsonMap.Ref
  value.Parameters = jsonMap.Parameters
  value.Publish = jsonMap.Publish
  value.Subscribe = jsonMap.Subscribe
  value.Deprecated = jsonMap.Deprecated

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value ChannelItem) MarshalJSON() ([]byte, error) {
	type Alias ChannelItem
	jsonByteArray, err := json.Marshal(&Alias{
    Ref: value.Ref,
    Parameters: value.Parameters,
    Publish: value.Publish,
    Subscribe: value.Subscribe,
    Deprecated: value.Deprecated,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// Parameter maps AsyncAPI "parameter" object
type Parameter struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Description string `json:"description,omitempty"`
  Name string `json:"name,omitempty"`
  Schema *Schema `json:"schema,omitempty"`
  Ref string `json:"$ref,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Parameter) UnmarshalJSON(data []byte) error {
	type Alias Parameter
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Description = jsonMap.Description
  value.Name = jsonMap.Name
  value.Schema = jsonMap.Schema
  value.Ref = jsonMap.Ref

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Parameter) MarshalJSON() ([]byte, error) {
	type Alias Parameter
	jsonByteArray, err := json.Marshal(&Alias{
    Description: value.Description,
    Name: value.Name,
    Schema: value.Schema,
    Ref: value.Ref,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// Operation maps AsyncAPI "operation" object
type Operation struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Summary string `json:"summary,omitempty"`
  Description string `json:"description,omitempty"`
  Tags []*Tag `json:"tags,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  OperationId string `json:"operationId,omitempty"`
  Message string `json:"message,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Operation) UnmarshalJSON(data []byte) error {
	type Alias Operation
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Summary = jsonMap.Summary
  value.Description = jsonMap.Description
  value.Tags = jsonMap.Tags
  value.ExternalDocs = jsonMap.ExternalDocs
  value.OperationId = jsonMap.OperationId
  value.Message = jsonMap.Message

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Operation) MarshalJSON() ([]byte, error) {
	type Alias Operation
	jsonByteArray, err := json.Marshal(&Alias{
    Summary: value.Summary,
    Description: value.Description,
    Tags: value.Tags,
    ExternalDocs: value.ExternalDocs,
    OperationId: value.OperationId,
    Message: value.Message,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// Stream maps AsyncAPI "stream" object
type Stream struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Framing string `json:"framing,omitempty"`
  Read []*Message `json:"read,omitempty"`
  Write []*Message `json:"write,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Stream) UnmarshalJSON(data []byte) error {
	type Alias Stream
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Framing = jsonMap.Framing
  value.Read = jsonMap.Read
  value.Write = jsonMap.Write

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Stream) MarshalJSON() ([]byte, error) {
	type Alias Stream
	jsonByteArray, err := json.Marshal(&Alias{
    Framing: value.Framing,
    Read: value.Read,
    Write: value.Write,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// Message maps AsyncAPI "message" object
type Message struct {
}

// UnmarshalJSON unmarshals JSON
func (value *Message) UnmarshalJSON(data []byte) error {
	type Alias Message
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value Message) MarshalJSON() ([]byte, error) {
	type Alias Message
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// VendorExtension maps AsyncAPI "vendorExtension" object
type VendorExtension struct {
}

// UnmarshalJSON unmarshals JSON
func (value *VendorExtension) UnmarshalJSON(data []byte) error {
	type Alias VendorExtension
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value VendorExtension) MarshalJSON() ([]byte, error) {
	type Alias VendorExtension
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// Tag maps AsyncAPI "tag" object
type Tag struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Name string `json:"name,omitempty"`
  Description string `json:"description,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Tag) UnmarshalJSON(data []byte) error {
	type Alias Tag
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Name = jsonMap.Name
  value.Description = jsonMap.Description
  value.ExternalDocs = jsonMap.ExternalDocs

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Tag) MarshalJSON() ([]byte, error) {
	type Alias Tag
	jsonByteArray, err := json.Marshal(&Alias{
    Name: value.Name,
    Description: value.Description,
    ExternalDocs: value.ExternalDocs,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// SecurityScheme maps AsyncAPI "SecurityScheme" object
type SecurityScheme struct {
}

// UnmarshalJSON unmarshals JSON
func (value *SecurityScheme) UnmarshalJSON(data []byte) error {
	type Alias SecurityScheme
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value SecurityScheme) MarshalJSON() ([]byte, error) {
	type Alias SecurityScheme
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// UserPassword maps AsyncAPI "userPassword" object
type UserPassword struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Type string `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *UserPassword) UnmarshalJSON(data []byte) error {
	type Alias UserPassword
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Type = jsonMap.Type
  value.Description = jsonMap.Description

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value UserPassword) MarshalJSON() ([]byte, error) {
	type Alias UserPassword
	jsonByteArray, err := json.Marshal(&Alias{
    Type: value.Type,
    Description: value.Description,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// ApiKey maps AsyncAPI "apiKey" object
type ApiKey struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Type string `json:"type,omitempty"`
  In string `json:"in,omitempty"`
  Description string `json:"description,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *ApiKey) UnmarshalJSON(data []byte) error {
	type Alias ApiKey
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Type = jsonMap.Type
  value.In = jsonMap.In
  value.Description = jsonMap.Description

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value ApiKey) MarshalJSON() ([]byte, error) {
	type Alias ApiKey
	jsonByteArray, err := json.Marshal(&Alias{
    Type: value.Type,
    In: value.In,
    Description: value.Description,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// X509 maps AsyncAPI "X509" object
type X509 struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Type string `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *X509) UnmarshalJSON(data []byte) error {
	type Alias X509
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Type = jsonMap.Type
  value.Description = jsonMap.Description

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value X509) MarshalJSON() ([]byte, error) {
	type Alias X509
	jsonByteArray, err := json.Marshal(&Alias{
    Type: value.Type,
    Description: value.Description,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// SymmetricEncryption maps AsyncAPI "symmetricEncryption" object
type SymmetricEncryption struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Type string `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *SymmetricEncryption) UnmarshalJSON(data []byte) error {
	type Alias SymmetricEncryption
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Type = jsonMap.Type
  value.Description = jsonMap.Description

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value SymmetricEncryption) MarshalJSON() ([]byte, error) {
	type Alias SymmetricEncryption
	jsonByteArray, err := json.Marshal(&Alias{
    Type: value.Type,
    Description: value.Description,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// AsymmetricEncryption maps AsyncAPI "asymmetricEncryption" object
type AsymmetricEncryption struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Type string `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *AsymmetricEncryption) UnmarshalJSON(data []byte) error {
	type Alias AsymmetricEncryption
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Type = jsonMap.Type
  value.Description = jsonMap.Description

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value AsymmetricEncryption) MarshalJSON() ([]byte, error) {
	type Alias AsymmetricEncryption
	jsonByteArray, err := json.Marshal(&Alias{
    Type: value.Type,
    Description: value.Description,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// HttpSecurityScheme maps AsyncAPI "HTTPSecurityScheme" object
type HttpSecurityScheme struct {
}

// UnmarshalJSON unmarshals JSON
func (value *HttpSecurityScheme) UnmarshalJSON(data []byte) error {
	type Alias HttpSecurityScheme
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value HttpSecurityScheme) MarshalJSON() ([]byte, error) {
	type Alias HttpSecurityScheme
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// NonBearerHttpSecurityScheme maps AsyncAPI "NonBearerHTTPSecurityScheme" object
type NonBearerHttpSecurityScheme struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Scheme string `json:"scheme,omitempty"`
  Description string `json:"description,omitempty"`
  Type string `json:"type,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *NonBearerHttpSecurityScheme) UnmarshalJSON(data []byte) error {
	type Alias NonBearerHttpSecurityScheme
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Scheme = jsonMap.Scheme
  value.Description = jsonMap.Description
  value.Type = jsonMap.Type

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value NonBearerHttpSecurityScheme) MarshalJSON() ([]byte, error) {
	type Alias NonBearerHttpSecurityScheme
	jsonByteArray, err := json.Marshal(&Alias{
    Scheme: value.Scheme,
    Description: value.Description,
    Type: value.Type,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// BearerHttpSecurityScheme maps AsyncAPI "BearerHTTPSecurityScheme" object
type BearerHttpSecurityScheme struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Scheme string `json:"scheme,omitempty"`
  BearerFormat string `json:"bearerFormat,omitempty"`
  Type string `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *BearerHttpSecurityScheme) UnmarshalJSON(data []byte) error {
	type Alias BearerHttpSecurityScheme
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Scheme = jsonMap.Scheme
  value.BearerFormat = jsonMap.BearerFormat
  value.Type = jsonMap.Type
  value.Description = jsonMap.Description

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value BearerHttpSecurityScheme) MarshalJSON() ([]byte, error) {
	type Alias BearerHttpSecurityScheme
	jsonByteArray, err := json.Marshal(&Alias{
    Scheme: value.Scheme,
    BearerFormat: value.BearerFormat,
    Type: value.Type,
    Description: value.Description,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// ApiKeyHttpSecurityScheme maps AsyncAPI "APIKeyHTTPSecurityScheme" object
type ApiKeyHttpSecurityScheme struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Type string `json:"type,omitempty"`
  Name string `json:"name,omitempty"`
  In string `json:"in,omitempty"`
  Description string `json:"description,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *ApiKeyHttpSecurityScheme) UnmarshalJSON(data []byte) error {
	type Alias ApiKeyHttpSecurityScheme
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Type = jsonMap.Type
  value.Name = jsonMap.Name
  value.In = jsonMap.In
  value.Description = jsonMap.Description

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value ApiKeyHttpSecurityScheme) MarshalJSON() ([]byte, error) {
	type Alias ApiKeyHttpSecurityScheme
	jsonByteArray, err := json.Marshal(&Alias{
    Type: value.Type,
    Name: value.Name,
    In: value.In,
    Description: value.Description,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// Oauth2Flows maps AsyncAPI "oauth2Flows" object
type Oauth2Flows struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Type string `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
  Flows string `json:"flows,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Oauth2Flows) UnmarshalJSON(data []byte) error {
	type Alias Oauth2Flows
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Type = jsonMap.Type
  value.Description = jsonMap.Description
  value.Flows = jsonMap.Flows

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Oauth2Flows) MarshalJSON() ([]byte, error) {
	type Alias Oauth2Flows
	jsonByteArray, err := json.Marshal(&Alias{
    Type: value.Type,
    Description: value.Description,
    Flows: value.Flows,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// Oauth2Flow maps AsyncAPI "oauth2Flow" object
type Oauth2Flow struct {
  Extensions map[string]json.RawMessage `json:"-"`
  AuthorizationUrl string `json:"authorizationUrl,omitempty"`
  TokenUrl string `json:"tokenUrl,omitempty"`
  RefreshUrl string `json:"refreshUrl,omitempty"`
  Scopes *Oauth2Scopes `json:"scopes,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Oauth2Flow) UnmarshalJSON(data []byte) error {
	type Alias Oauth2Flow
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.AuthorizationUrl = jsonMap.AuthorizationUrl
  value.TokenUrl = jsonMap.TokenUrl
  value.RefreshUrl = jsonMap.RefreshUrl
  value.Scopes = jsonMap.Scopes

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Oauth2Flow) MarshalJSON() ([]byte, error) {
	type Alias Oauth2Flow
	jsonByteArray, err := json.Marshal(&Alias{
    AuthorizationUrl: value.AuthorizationUrl,
    TokenUrl: value.TokenUrl,
    RefreshUrl: value.RefreshUrl,
    Scopes: value.Scopes,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// Oauth2Scopes maps AsyncAPI "oauth2Scopes" object
type Oauth2Scopes struct {
}

// UnmarshalJSON unmarshals JSON
func (value *Oauth2Scopes) UnmarshalJSON(data []byte) error {
	type Alias Oauth2Scopes
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value Oauth2Scopes) MarshalJSON() ([]byte, error) {
	type Alias Oauth2Scopes
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// OpenIdConnect maps AsyncAPI "openIdConnect" object
type OpenIdConnect struct {
  Extensions map[string]json.RawMessage `json:"-"`
  Type string `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
  OpenIdConnectUrl string `json:"openIdConnectUrl,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *OpenIdConnect) UnmarshalJSON(data []byte) error {
	type Alias OpenIdConnect
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  value.Type = jsonMap.Type
  value.Description = jsonMap.Description
  value.OpenIdConnectUrl = jsonMap.OpenIdConnectUrl

  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value OpenIdConnect) MarshalJSON() ([]byte, error) {
	type Alias OpenIdConnect
	jsonByteArray, err := json.Marshal(&Alias{
    Type: value.Type,
    Description: value.Description,
    OpenIdConnectUrl: value.OpenIdConnectUrl,
	})
	if err != nil {
		return nil, err
  }
  return MergeExtensions(jsonByteArray, value.Extensions)
}

// SecurityRequirement maps AsyncAPI "SecurityRequirement" object
type SecurityRequirement struct {
}

// UnmarshalJSON unmarshals JSON
func (value *SecurityRequirement) UnmarshalJSON(data []byte) error {
	type Alias SecurityRequirement
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value SecurityRequirement) MarshalJSON() ([]byte, error) {
	type Alias SecurityRequirement
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// Title maps AsyncAPI "title" object
type Title struct {
}

// UnmarshalJSON unmarshals JSON
func (value *Title) UnmarshalJSON(data []byte) error {
	type Alias Title
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value Title) MarshalJSON() ([]byte, error) {
	type Alias Title
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// Description maps AsyncAPI "description" object
type Description struct {
}

// UnmarshalJSON unmarshals JSON
func (value *Description) UnmarshalJSON(data []byte) error {
	type Alias Description
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value Description) MarshalJSON() ([]byte, error) {
	type Alias Description
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// Default maps AsyncAPI "default" object
type Default struct {
}

// UnmarshalJSON unmarshals JSON
func (value *Default) UnmarshalJSON(data []byte) error {
	type Alias Default
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value Default) MarshalJSON() ([]byte, error) {
	type Alias Default
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// MultipleOf maps AsyncAPI "multipleOf" object
type MultipleOf struct {
}

// UnmarshalJSON unmarshals JSON
func (value *MultipleOf) UnmarshalJSON(data []byte) error {
	type Alias MultipleOf
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value MultipleOf) MarshalJSON() ([]byte, error) {
	type Alias MultipleOf
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// Maximum maps AsyncAPI "maximum" object
type Maximum struct {
}

// UnmarshalJSON unmarshals JSON
func (value *Maximum) UnmarshalJSON(data []byte) error {
	type Alias Maximum
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value Maximum) MarshalJSON() ([]byte, error) {
	type Alias Maximum
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// ExclusiveMaximum maps AsyncAPI "exclusiveMaximum" object
type ExclusiveMaximum struct {
}

// UnmarshalJSON unmarshals JSON
func (value *ExclusiveMaximum) UnmarshalJSON(data []byte) error {
	type Alias ExclusiveMaximum
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value ExclusiveMaximum) MarshalJSON() ([]byte, error) {
	type Alias ExclusiveMaximum
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// Minimum maps AsyncAPI "minimum" object
type Minimum struct {
}

// UnmarshalJSON unmarshals JSON
func (value *Minimum) UnmarshalJSON(data []byte) error {
	type Alias Minimum
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value Minimum) MarshalJSON() ([]byte, error) {
	type Alias Minimum
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// ExclusiveMinimum maps AsyncAPI "exclusiveMinimum" object
type ExclusiveMinimum struct {
}

// UnmarshalJSON unmarshals JSON
func (value *ExclusiveMinimum) UnmarshalJSON(data []byte) error {
	type Alias ExclusiveMinimum
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value ExclusiveMinimum) MarshalJSON() ([]byte, error) {
	type Alias ExclusiveMinimum
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// MaxLength maps AsyncAPI "maxLength" object
type MaxLength struct {
}

// UnmarshalJSON unmarshals JSON
func (value *MaxLength) UnmarshalJSON(data []byte) error {
	type Alias MaxLength
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value MaxLength) MarshalJSON() ([]byte, error) {
	type Alias MaxLength
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// MinLength maps AsyncAPI "minLength" object
type MinLength struct {
}

// UnmarshalJSON unmarshals JSON
func (value *MinLength) UnmarshalJSON(data []byte) error {
	type Alias MinLength
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value MinLength) MarshalJSON() ([]byte, error) {
	type Alias MinLength
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// Pattern maps AsyncAPI "pattern" object
type Pattern struct {
}

// UnmarshalJSON unmarshals JSON
func (value *Pattern) UnmarshalJSON(data []byte) error {
	type Alias Pattern
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value Pattern) MarshalJSON() ([]byte, error) {
	type Alias Pattern
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// MaxItems maps AsyncAPI "maxItems" object
type MaxItems struct {
}

// UnmarshalJSON unmarshals JSON
func (value *MaxItems) UnmarshalJSON(data []byte) error {
	type Alias MaxItems
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value MaxItems) MarshalJSON() ([]byte, error) {
	type Alias MaxItems
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// MinItems maps AsyncAPI "minItems" object
type MinItems struct {
}

// UnmarshalJSON unmarshals JSON
func (value *MinItems) UnmarshalJSON(data []byte) error {
	type Alias MinItems
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value MinItems) MarshalJSON() ([]byte, error) {
	type Alias MinItems
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// UniqueItems maps AsyncAPI "uniqueItems" object
type UniqueItems struct {
}

// UnmarshalJSON unmarshals JSON
func (value *UniqueItems) UnmarshalJSON(data []byte) error {
	type Alias UniqueItems
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value UniqueItems) MarshalJSON() ([]byte, error) {
	type Alias UniqueItems
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

// Enum maps AsyncAPI "enum" object
type Enum struct {
}

// UnmarshalJSON unmarshals JSON
func (value *Enum) UnmarshalJSON(data []byte) error {
	type Alias Enum
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }



	return nil
}

// MarshalJSON marshals JSON
func (value Enum) MarshalJSON() ([]byte, error) {
	type Alias Enum
	jsonByteArray, err := json.Marshal(&Alias{
	})
	if err != nil {
		return nil, err
  }
  return jsonByteArray
}

