# Database models


## Scaffold collection

```json
[{
    "_id": "<ObjectID>",
    "author_id": "<UserI: ID of the user who created this>",
    "created_at": "<Time: ISO8601 timestamp>", // Implied by MongoDB ObjectID
    "header_ids": [
        "<ObjectID>",
        "<ObjectID>",
        "<ObjectID>",
        ...
    ],
    "body_ids": [
        "<ObjectID>",
        "<ObjectID>",
        "<ObjectID>",
        ...
    ]
}, ... ]
```

## Post headers collection
```json
[{
    "_id": "<ObjectID>",
    "author_id": "<UserID: id of the user who created this>",
    "created_at": "<Time: ISO8601 timestamp>", // Implied by MongoDB ObjectID
    "data": "String: post header data"
}, ... ]
```

## Post bodies collection
```json
[{
    "_id": "<ObjectID>",
    "author_id": "<UserID: id of the user who created this>",
    "created_at": "<Time: ISO8601 timestamp>", // Implied by MongoDB ObjectID
    "data": "String: post body data"
}, ... ]
```