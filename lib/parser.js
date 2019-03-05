const fs = require('fs');
const path = require('path');
const ZSchema = require('z-schema');
const YAML = require('js-yaml');
const RefParser = require('json-schema-ref-parser');

const validator = new ZSchema();

async function getFileContent (filePath) {
  return new Promise((resolve, reject) => {
    fs.readFile(path.resolve(__dirname, filePath), (err, content) => {
      if (err) return reject(err);
      resolve(content);
    });
  });
}

function parseContent (content) {
  content = content.toString('utf8');
  try {
    return JSON.parse(content);
  } catch (e) {
    return YAML.safeLoad(content);
  }
}

async function dereference (json) {
  return RefParser.dereference(json, {
    dereference: {
      circular: 'ignore'
    }
  });
}

async function bundle (json) {
  return RefParser.bundle(json, {
    dereference: {
      circular: 'ignore'
    }
  });
}

async function validate (json, schema) {
  return new Promise((resolve, reject) => {
    validator.validate(json, schema, (err, valid) => {
      if (err) return reject(err);
      return resolve(json);
    });
  });
}

async function parse (filePath) {
  let content, parsedContent;

  try {
    if (typeof filePath === 'string') {
      content = await getFileContent(filePath);
    } else if (typeof filePath === 'object') {
      content = JSON.stringify(filePath);
    } else {
      console.error(`Could not find a valid asyncapi definition: ${filePath}`);
      throw new Error(`Could not find a valid asyncapi definition: ${filePath}`);
    }
  } catch (e) {
    console.error('Can not load the content of the AsyncAPI specification file');
    console.error(e);
    throw e;
  }

  try {
    parsedContent = parseContent(content);
  } catch (e) {
    console.error('Can not parse the content of the AsyncAPI specification file');
    console.error(e);
    throw e;
  }

  return JSON.parse(JSON.stringify(parsedContent));
};

async function parseText (content) {
  let parsedStringified;

  try {
    const parsedContent = parseContent(content);
    if (typeof parsedContent !== 'object' || parsedContent === null) {
      throw new Error('Invalid YAML content.');
    }
    const dereferencedJSON = await dereference(parsedContent);
    const bundledJSON = await bundle(dereferencedJSON);
    const asyncAPIschema = require('asyncapi')[bundledJSON.asyncapi];

    const parsed = await validate(bundledJSON, asyncAPIschema);

    parsedStringified = JSON.parse(JSON.stringify(parsed));
  } catch (e) {
    throw e;
  }

  return parsedStringified;
};

module.exports = parse;
module.exports.parseText = parseText;
