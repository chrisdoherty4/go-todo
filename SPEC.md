# Specification

The project has 2 components: a RESTful service and a CLI tool that interacts with the RESTful service. The RESTful service is the core component of the project as is a dependency of the CLI tool. You are at liberty to write the service and tooling from scratch, or to use third-party packages.

If you wish to target writing a RESTful service and a CLI tool you are at liberty to implement parts of the RESTful service and CLI tool in the allotted time.

## RESTful Service

### Retrieve all items

> GET /api/v1/items

##### Request body

None

##### Responses

    200 Successful
    {
      "items": [
        {
          "title": "string",
          "description": "string",
          "complete": boolean
        },
        ...
      ]
    }

### Retrieve item details

> GET /api/v1/items/{title}

##### Request body

None

##### Responses

    200 Successful
    {
      "title": "string",
      "description": "string",
      "complete": boolean
    }


    404 Not found

### Create a new item

> POST /api/v1/items/{title}

##### Request body

    {
      "description": "string"
    }

##### Responses

    200 Successful

    403 Already exists
    {
      "title": "string",
      "description": "string",
      "complete": boolean
    }

### Update an item

> PUT /api/v1/items/{title}

##### Request body

    {
      "description": "string",
      "complete": boolean
    }

##### Responses

    200 Successful
    {
      "old": {
        "title": "string",
        "description": "string",
        "complete": boolean
      },
      "new": {
        "title": "string",
        "description": "string",
        "complete": boolean
      }
    }

    404 Not found

### Remove an item

> DELETE /api/v1/items/{title}

##### Request body

None

##### Responses

    200 Successful
    {
      "title": "string",
      "description": "string",
      "complete": boolean
    }

    404 Not found

## CLI Tool

The implementation of the CLI tool is up to the developer. You can integrate the CLI tool as a front-end to the REST service described above, or you can create a completely independent tool that manages todo items in some other way.

The CLI tool should be capable of listing, creating, marking complete and deleting todo items. 

### List items

    todo list [-a]

##### Options

    -a : List all items, complete and incomplete.

### List specific item.

    todo list <title>

##### Arguments

    <title> : A unique title for the new todo item.

### Create item
 
    todo create [-f <file-path>] <title> [<description>]

##### Arguments

    <title> : A unique title for the new todo item.
    <description> : [Optional] A detailed description for the todo item.

##### Options

    -f <file-path> : [Optional] A file containing the <description> content. Takes precedence over <description>.


### Update an item

    todo update [-f <file-path>] <title> [<description>]

##### Arguments

    <title> : A unique title for the new todo item.
    <description> : [Optional] A detailed description for the todo item.

##### Options

    -f <file-path> : [Optional] A file containing the <description> content. Takes precedence over <description>.

### Mark item complete

    todo complete <title>

##### Arguments

    <title> : A unique title for the new todo item.

### Delete an item

    todo delete <title>

##### Arguments

    <title> : A unique title for the new todo item.
