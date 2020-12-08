# Database models


## Scaffold collection
```json
[{
    "_id": "<ObjectID>",
    "author_id": "<UserID: ID of the user who created this>",
    "created_at": "<ISO8601 timestamp>", // Also implied by MongoDB ObjectID
    "updated_at": "<ISO8601 timestamp>",
    "marked_deleted": "<Boolean>"
    "header": [
        "<ObjectID>",
        "<ObjectID>",
        "<ObjectID>",
        ...
    ],
    "body": [
        "<ObjectID>",
        "<ObjectID>",
        "<ObjectID>",
        ...
    ]
}, ... ]
```

## Headers collection
```json
[{
    "_id": "<ObjectID>",
    "author_id": "<UserID: id of the user who created this>",
    "created_at": "<Time: ISO8601 timestamp>", // Also implied by MongoDB ObjectID
    "data": "String: post header data",
    "marked_deleted": "<Boolean>"
}, ... ]
```

## Bodies collection
```json
[{
    "_id": "<ObjectID>",
    "author_id": "<UserID: id of the user who created this>",
    "created_at": "<Time: ISO8601 timestamp>", // Also implied by MongoDB ObjectID
    "data": "String: post body data",
    "marked_deleted": "<Boolean>"
}, ... ]
```