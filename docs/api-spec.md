# API Reference

## General
#### Access control
Any request must contain a valid session token
```json
{
    "session_token": "<SessionTokenID: valid and active JWT>", // Usually carried out by HTTP title from cookies
    ... // rest of request
}
```
#### Meta wrapper
Every response contain metadata about the request. The requested data is stored in the `data` property. Any request may get rejected or the response contains no data whereof `data` will be `null`.
```json
{
    "data": "<Object: requested data>",
    "status_code": "<Number: HTTP status code",
    "message": "<String>",
    "processing_time": "<Number: Processing time of the request in ms>",
    "node_respondant": "<NodeID: ID of the node handling the request>"
}
```

---
## Posts

#### Create a post `posts.create`


Request
```json
{
    "post": {
        "title": "<String: title of the post>",
        "body": "<String: post body>"
    }
}
```

#### Get post `posts.get`
Request
```json
{
    "post_id": "<PostID: ID of the post>"
}
```
Response
```json
{
    "post_id": "<PostID: ID of the post>",
    "created_at": "<ISO8601 timestamp>",
    "author_id": "<UserID: ID of the author>",
    "post": {
        "title": "<title of the post>",
        "body": "<String: post body>"
    }
}
```

#### Update post `posts.update`
Request
```json
{
    "post_id": "<PostID: ID of the post>",
    "post": {
        "title": "<String: title of the post>",     // optional
        "body": "<String: post body>",              // optional
    }
}
```


#### Delete post `posts.delete`
Request
```json
{
    "post_id": "<PostID: ID of the post>"
}
```


#### Get many posts `posts.get.many`
Request
```json
{
    "post_ids": ["<PostID>", "<PostID>", ...],                                  // optional
    "paging": {                                                                 // optional
        "start": "<Number|ISO8601 timestamp: start of the current page>",       // default=0
        "end": "<Number|ISO8601 timestamp: end of the current page>",           // default=9
        "limit": "<Number: max number of posts in current page (-1 for all)>"   // optional
    }
}
```
Response
```json
[
    {
        "post_id": "<PostID: ID of the post>",
        "created_at": "<ISO8601 timestamp>",
        "author_id": "<UserID: ID of the author>",
        "post": {
            "title": "<String: title of the post>",
            "body": "<String: post body>"
        }
    },
    ...
]
```
#### Get history for post `posts.get.history`
Request
```json
{
    "post_id": "<PostID: ID of the post>",
    "paging": {                                                                 // optional
        "start": "<Number|ISO8601 timestamp: start of the current page>",       // default=0
        "end": "<Number|ISO8601 timestamp: end of the current page>",           // default=9
        "limit": "<Number: max number of posts in current page (-1 for all)>"   // optional
    }
}
```
Response
```json
{
    "post_id": "<PostID: ID of the post>",
    "author_id": "<UserID: ID of the author>",
    "history": [
        {
            "created_at": "<ISO8601 timestamp>",
            "title": "<String: title of the post>",
            "body": "<String: post body>",
        },
        ...
    ]
}
```

---
## Examples of usage


**Creating a new post**
Request/produce to `posts.create`
```json
{
    "post": {
        "title": "Hello, World! 🌎",
        "body": "This is my first post.."
    }
}
```
Response
```json
{
    "data": null,
    "status_code": 200,
    "message": "OK",
    "processing_time": 420,
    "node_respondant": "node-123"
}
```

---
**Getting a post:**
Request/produce to `posts.get`
```json
{
    "post_id": 123
}
```
Response
```json
{
    "data": {
        "post_id": 123,
        "created_at": "2020-11-12T14:29:59+01:00",
        "author_id": 123,
        "post": {
            "title": "Hello, World! 🌎",
            "body": "This is my first post.."
        }
    },
    "status_code": 200,
    "message": "OK",
    "processing_time": 42,
    "respondent_node": "node-456"
}
```
