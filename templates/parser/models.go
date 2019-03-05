package models

import "encoding/json"

{{>struct def=asyncapi defName='AsyncapiDocument'}}

{{#each asyncapi.definitions as |def defName| }}
{{#unless (shouldAvoidSchema defName)}}
{{#ifCanNotGuessStruct def}}
{{>alias typeName=defName type=(GoTypeFrom def)}}
{{else}}
{{#ifNotBoolean def.additionalProperties}}
{{>alias typeName=defName type=(GoTypeFrom def.additionalProperties)}}
{{else}}
{{>struct def=def defName=defName asyncapi=../../asyncapi}}
{{/ifNotBoolean}}
{{/ifCanNotGuessStruct}}

{{/unless}}
{{/each}}
