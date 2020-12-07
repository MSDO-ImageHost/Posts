# API Reference

## Message parameters
The following table displays fields that messages can or must contain
| Key               | Value                                         | Type      | Required          | Action    |
|-------------------|-----------------------------------------------|-----------|-------------------|-----------|
| ContentType       | "application/json"                            | Property  | Yes               | Req/Res   |
| CorrelationId     | "\<String: something id\>"                    | Property  | Yes               | Req/Res   |
| ReplyTo           | "<String: reply queue>"                       | Property  | No                | Req       |
| JWT               | "\<String:xxx.yyy.zzz\>"                      | Header    | For write only    | Req       |
| StatusCode        | "\<Number: status code\>"                     | Header    |                   | Res       |
| StatusMessage     | "\<String: status message\>"                  | Header    |                   | Res       |
| ProcessingTimeNs  | "\<Number: processing time in nano seconds\>" | Header    |                   | Res       |



-----
## Routing
| Event                 | Queue                 | API   |
|-----------------------|-----------------------|-------|
| Create one post       | posts.create.one      |       |
| Read a single post    | posts.read.one        |       |
| Read post history     | posts.read.history    |       |
| Read many posts       | posts.read.many       |       |
| Read user posts       | posts.read.userposts  |       |
| Update one post       | posts.update.one      |       |
| Delete one post       | posts.delete.one      |       |
| Delete many posts     | posts.delete.many     |       |




#### Submit a create message on queue `posts.new`
Payload of request message
```json
{
    "header": "<String: title of the post>",
    "body": "<String: body text of the post>"
}
```
Response is published in queue `posts.return.new` or specified by `reply_to` in the message
```json
{
    "post_id": "<PostID: ID of the created post>",
    "created_at": "<ISO8601 timestamp>",
    "author_id": "<UserID: ID of the author>",
    "header": {
        "author_id": "<UserID: ID of the author>",
        "created_at": "<ISO8601 timestamp>",
        "data":"<String: title of the post>",
    },
    "body": {
        "author_id": "<UserID: ID of the author>",
        "created_at": "<ISO8601 timestamp>",
        "data":"<String: body of the post>",
    }
}
```
#### Get post `posts.read.one`
Payload of request message
```json
{
    "post_id": "<PostID: ID of the post>"
}
```
Response
```json
{
    "post_id": "<PostID: ID of the created post>",
    "created_at": "<ISO8601 timestamp>",
    "author_id": "<UserID: ID of the author>",
    "header": {
        "author_id": "<UserID: ID of the author>",
        "created_at": "<ISO8601 timestamp>",
        "data":"<String: title of the post>",
    },
    "body": {
        "author_id": "<UserID: ID of the author>",
        "created_at": "<ISO8601 timestamp>",
        "data":"<String: body of the post>",
    }
}
```

#### Update post `posts.update`
Request
```json
{
    "post_id": "<PostID: ID of the post to update>",
    "header": "<String: updated title of the post>",     // optional
    "body": "<String: updated text of the post>",              // optional
}
```


#### Delete post `posts.delete`
Request
```json
{
    "post_id": "<PostID: ID of the updated post>"
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
        "header": "<String: title of the post>",
        "body": "<String: post body>"

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
            "header": "<String: title of the post>",
            "body": "<String: body text of the post>",
        },
        ...
    ]
}
```

