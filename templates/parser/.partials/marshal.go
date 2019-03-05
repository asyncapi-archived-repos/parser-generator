// UnmarshalJSON unmarshals JSON
func (value *{{GoPublicName defName}}) UnmarshalJSON(data []byte) error {
	type {{GoPublicName defName}}Alias {{GoPublicName defName}}
	jsonMap := {{GoPublicName defName}}Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
  }

  {{#each def.properties as |prop propName|}}
  value.{{GoPublicName propName}} = jsonMap.{{GoPublicName propName}}
  {{/each}}

  {{#ifHasExtensions def}}
  exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts
  {{/ifHasExtensions}}

	return nil
}

// MarshalJSON marshals JSON
func (value {{GoPublicName defName}}) MarshalJSON() ([]byte, error) {
	type {{GoPublicName defName}}Alias {{GoPublicName defName}}
	jsonByteArray, err := json.Marshal(&{{GoPublicName defName}}Alias{
    {{#each def.properties as |prop propName|}}
    {{GoPublicName propName}}: value.{{GoPublicName propName}},
    {{/each}}
	})
	if err != nil {
		return nil, err
  }
  {{#ifHasExtensions def}}
  return MergeExtensions(jsonByteArray, value.Extensions)
  {{else}}
  return jsonByteArray, nil
  {{/ifHasExtensions}}
}
