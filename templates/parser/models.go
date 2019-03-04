package models

import "encoding/json"

{{>struct def=asyncapi defName='AsyncapiDocument'}}

{{#each asyncapi.definitions as |def defName| }}
{{#ifCanNotGuessStruct def}}
{{>alias typeName=defName type='json.RawMessage'}}
{{else}}
{{#ifNotBoolean def.additionalProperties}}
{{>alias typeName=defName type=(GoTypeFrom def.additionalProperties)}}
{{else}}
{{>struct def=def defName=defName}}
{{/ifNotBoolean}}
{{/ifCanNotGuessStruct}}

{{/each}}
