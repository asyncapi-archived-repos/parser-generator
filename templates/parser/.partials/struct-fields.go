{{#if def.properties}}
{{#ifHasExtensions def}}
Extensions map[string]json.RawMessage `json:"-"`
{{/ifHasExtensions}}
{{#each def.properties as |prop propName|}}
{{GoPublicName propName}} {{GoTypeFrom prop}} `json:"{{propName}},omitempty"`
{{/each}}
{{/if}}
