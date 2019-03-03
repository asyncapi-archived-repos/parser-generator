// {{GoPublicName defName}} maps AsyncAPI "{{defName}}" object
{{#ifNotBoolean def.additionalProperties}}
type {{GoPublicName ../defName}} map[string]{{GoTypeFrom ../def.additionalProperties}}
{{else}}
type {{GoPublicName ../defName}} struct {
  {{>structFields def=../def defName=../defName}}
}

// UnmarshalJSON unmarshals JSON
func (value *{{GoPublicName ../defName}}) UnmarshalJSON(data []byte) error {
	type Alias {{GoPublicName ../defName}}
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  {{#each ../def.properties as |prop propName|}}
  value.{{GoPublicName propName}} = jsonMap.{{GoPublicName propName}}
  {{/each}}

  {{#ifHasExtensions ../def}}
  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts
  {{/ifHasExtensions}}

	return nil
}

// MarshalJSON marshals JSON
func (value {{GoPublicName ../defName}}) MarshalJSON() ([]byte, error) {
	type Alias {{GoPublicName ../defName}}
	jsonByteArray, err := json.Marshal(&Alias{
    {{#each ../def.properties as |prop propName|}}
    {{GoPublicName propName}}: value.{{GoPublicName propName}},
    {{/each}}
	})
	if err != nil {
		return nil, err
  }
  {{#ifHasExtensions ../def}}
  return MergeExtensions(jsonByteArray, value.Extensions)
  {{else}}
  return jsonByteArray, nil
  {{/ifHasExtensions}}
}
{{/ifNotBoolean}}

