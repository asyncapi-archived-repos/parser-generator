{{#if (isPrimitive def.type)}}
type {{GoPublicName defName}} {{GoTypeFrom def}}
{{else}}
// {{GoPublicName defName}} maps AsyncAPI "{{defName}}" object
type {{GoPublicName defName}} struct {
  {{>structFields def=def asyncapi=asyncapi}}
  {{#each def.oneOf as |item|}}
  {{>structFields def=(lookup ../asyncapi.definitions (nameFromRef item.$ref))}}
  {{/each}}
}

{{>marshal def=def defName=defName}}
{{/if}}
