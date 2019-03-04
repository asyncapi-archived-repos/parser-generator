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

  Handlebars.registerHelper('isRequired', (obj, key) => {
    return obj && obj.required && !!(obj.required.includes(key));
  });

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

  Handlebars.registerHelper('log', (something) => {
    console.log(require('util').inspect(something, { depth: null }));
  });

  function GoPublicName (something) {
    return _.upperFirst(_.camelCase(something));
  }

  Handlebars.registerHelper('GoPublicName', (something) => {
    return GoPublicName(something);
  });

  function GoTypeFrom (jsonSchemaObj) {
    if (jsonSchemaObj.$ref) {
      return `*${GoPublicName(jsonSchemaObj.$ref.split('/').slice(-1)[0])}`;
    }

    if (Object.keys(jsonSchemaObj).length === 0) {
      return 'map[string]json.RawMessage';
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
        return 'map[string]json.RawMessage';
      default:
        console.error('Unsupported JSONSchema type:', jsonSchemaObj.type, 'in object', jsonSchemaObj);
        return 'map[string]json.RawMessage';
    }
  }

  Handlebars.registerHelper('GoTypeFrom', (jsonSchemaObj) => {
    return GoTypeFrom(jsonSchemaObj);
  });

  Handlebars.registerHelper('ifHasExtensions', (obj, options) => {
    if (obj && obj.patternProperties && obj.patternProperties['^x-']) {
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
    if (obj && (obj.oneOf || obj.anyOf || obj.allOf || obj.additionalProperties === true)) {
      return options.fn(this);
    }

    return options.inverse(this);
  });
};
