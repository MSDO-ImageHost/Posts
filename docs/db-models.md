# Database models


## Posts collection

```json
[{
    "object_id": "ObjectId",
    "creator_id": "<UserI: ID of the user>",
    "recent_change": "<ISO8601 timestamp: (post or latest revision) created_at>",
    "header_history": [
        {
            "rev": "<Number: revision index>", // probably not needed as date can infer revision
            "header_id": "<ObjectId: reference to internal header object"
        },
        ...
    ],
    "body_history": [
        {
            "rev": "<Number: revision index>", // probably not needed as date can infer revision
            "body_id": "<ObjectId: reference to internal post body object"
        },
        ...
    ]
}, ... ]
```

## Post headers collection
```json
[{
    "object_id": "ObjectId",
    "creator_id": "<UserID: id of the user>",
    "created_at": "<String: ISO8601 timestamp>", // Implied by MongoDB ObjectID
    "data": "String: post header data"
}, ... ]
```

## Post bodies collection
```json
[{
    "object_id": "ObjectId",
    "creator_id": "<UserID: id of the user>",
    "created_at": "<String: ISO8601 timestamp>", // Implied by MongoDB ObjectID
    "data": "String: post body data"
}, ... ]
```