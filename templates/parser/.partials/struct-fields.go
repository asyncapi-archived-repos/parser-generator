{{#ifHasExtensions def}}
Extensions map[string]json.RawMessage `json:"-"`
{{/ifHasExtensions}}
{{#if def.properties}}
{{#each def.properties as |prop propName|}}
{{GoPublicName propName}} {{GoTypeFrom prop ../def propName}} `json:"{{propName}},omitempty"`
{{/each}}
{{/if}}
