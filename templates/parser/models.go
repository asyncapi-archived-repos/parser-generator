package models

import "encoding/json"

{{>struct def=asyncapi defName='AsyncapiDocument'}}

{{#each asyncapi.definitions as |def defName| }}
{{>struct def=. defName=@key}}
{{/each}}
