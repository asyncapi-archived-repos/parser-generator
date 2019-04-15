const RefParser = require('json-schema-ref-parser-sync');

module.exports = (Handlebars, _) => {

  Handlebars.registerHelper('concat', (str1, str2, separator) => {
    return `${str1 || ''}${separator || ''}${str2 || ''}`;
  });

  Handlebars.registerHelper('tree', path => {
    const filteredPaths = path.split('.').filter(Boolean);
    if (!filteredPaths.length) return;
    const dottedPath = filteredPaths.join('.');

    return `${dottedPath}.`;
  });

  Handlebars.registerHelper('buildPath', (propName, path, key) => {
    if (!path) return propName;
    return `${path}.${propName}`;
  });

  function isRequired (obj, key) {
    return obj && Array.isArray(obj.required) && !!(obj.required.includes(key));
  }

  Handlebars.registerHelper('isRequired', isRequired);

  Handlebars.registerHelper('acceptedValues', items => {
    if (!items) return '<em>Any</em>';

    return items.map(i => `<code>${i}</code>`).join(', ');
  });

  Handlebars.registerHelper('equal', (lvalue, rvalue, options) => {
    if (arguments.length < 3)
      throw new Error('Handlebars Helper equal needs 2 parameters');
    if (lvalue!==rvalue) {
      return options.inverse(this);
    }

    return options.fn(this);
  });

  Handlebars.registerHelper('inc', (number) => {
    return number + 1;
  });

  function GoPublicName (something) {
    return _.upperFirst(_.camelCase(something));
  }

  Handlebars.registerHelper('GoPublicName', (something) => {
    return GoPublicName(something);
  });

  function nameFromRef(url) {
    if (!url) return;
    return url.split('/').slice(-1)[0];
  }

  Handlebars.registerHelper('nameFromRef', nameFromRef);

  function GoTypeFrom (jsonSchemaObj, ctxt, name) {
    if (jsonSchemaObj['x-go-parser-type']) return jsonSchemaObj['x-go-parser-type'];

    let required = false;
    if (ctxt && name) required = isRequired(ctxt, name);

    if (jsonSchemaObj.$ref) {
      return `${required ? '' : '*'}${GoPublicName(nameFromRef(jsonSchemaObj.$ref))}`;
    }

    if (jsonSchemaObj.patternProperties && jsonSchemaObj.patternProperties[''] && jsonSchemaObj.patternProperties[''].$ref) {
      return `${required ? '' : '*'}${GoPublicName(nameFromRef(jsonSchemaObj.patternProperties[''].$ref))}`;
    }

    if (Object.keys(jsonSchemaObj).length === 0) {
      return 'json.RawMessage';
    }

    switch (jsonSchemaObj.type) {
      case 'string':
        return 'string';
      case 'boolean':
        return 'bool';
      case "integer":
        return 'int';
      case 'number':
        return 'float64';
      case 'array':
        if (jsonSchemaObj.items) {
          return `[]${GoTypeFrom(jsonSchemaObj.items)}`;
        }
      case 'object':
        return 'json.RawMessage';
      default:
        console.error('Unsupported JSONSchema type:', jsonSchemaObj.type, 'in object', jsonSchemaObj);
        return 'json.RawMessage';
    }
  }

  Handlebars.registerHelper('GoTypeFrom', GoTypeFrom);

  Handlebars.registerHelper('ifHasExtensions', (obj, options) => {
    if (obj && obj.patternProperties && obj.patternProperties['^x-[\\w\\d\\.\\-\\_]+$']) {
      return options.fn(this);
    }

    return options.inverse(this);
  });

  Handlebars.registerHelper('ifNotBoolean', (obj, options) => {
    if (typeof obj !== 'boolean' && obj !== undefined && obj !== null) {
      return options.fn(this);
    }

    return options.inverse(this);
  });

  Handlebars.registerHelper('ifTypeOf', (obj, type, options) => {
    if (obj && typeof obj === type) {
      return options.fn(this);
    }

    return options.inverse(this);
  });

  Handlebars.registerHelper('ifCanNotGuessStruct', (obj, options) => {
    if (
      obj && (
        obj.allOf || obj.additionalProperties === true ||
        (obj.patternProperties && obj.patternProperties[''] && obj.patternProperties[''].$ref)
      )
    ) {
      return options.fn(this);
    }

    return options.inverse(this);
  });

  Handlebars.registerHelper('shouldAvoidSchema', (name) => {
    return ['SecurityScheme', 'HTTPSecurityScheme', 'message'].includes(name);
  });

  Handlebars.registerHelper('deref', (obj) => {
    return RefParser.dereference(obj, {
      dereference: {
        circular: 'ignore'
      }
    });
  });

  Handlebars.registerHelper('isPrimitive', (type) => {
    return ['string', 'boolean', 'integer', 'number'].includes(type);
  });
};
