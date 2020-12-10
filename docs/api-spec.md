# API Reference

## Message parameters
The following table displays fields that messages can or must contain
| Key               | Value                                         | Type      | Required          | Action                |
|-------------------|-----------------------------------------------|-----------|-------------------|-----------------------|
| ContentType       | "application/json"                            | Property  | Yes               | Request / Response    |
| CorrelationId     | "\<String: something id\>"                    | Property  | No                | Request / Response    |
| ReplyTo           | "<String: reply queue>"                       | Property  | No                | Request               |
| JWT               | "\<String:xxx.yyy.zzz\>"                      | Header    | For write only    | Request               |
| StatusCode        | "\<Number: status code\>"                     | Header    |                   | Response              |
| StatusMessage     | "\<String: status message\>"                  | Header    |                   | Response              |
| ProcessingTimeNs  | "\<Number: processing time in nano seconds\>" | Header    |                   | Response              |



-----
## Routing
| Intent                | Request queues        | Response queues           | Request RK on 'rapid'     | Response RK on 'Posts' queue  |
|-----------------------|-----------------------|---------------------------|---------------------------|-------------------------------|
| Create one post       | posts.create.one      | posts.return.one          | CreateOnePost             | ConfirmOnePostCreation        |
| Read a single post    | posts.read.one        | posts.return.one          | RequestOnePost            | ReturnOnePost                 |
| Read post history     | posts.read.history    | posts.return.one.history  | RequestPostHistory        | ReturnPostHistory             |
| Read many posts       | posts.read.many       | posts.return.many         | RequestManyPosts          | ReturnManyPosts               |
| Read user posts       | posts.read.userposts  | posts.return.many         | RequestUserPosts          | ReturnUserPosts               |
| Update one post       | posts.update.one      | posts.return.one          | UpdateOnePost             | ConfirmUpdateOnePost          |
| Delete one post       | posts.delete.one      | posts.return.one          | DeleteOnePost             | ConfirmDeleteOnePost          |
| Delete many posts     | posts.delete.many     | posts.return.many         | DeleteManyPosts           | ConfirmDeleteManyPosts        |

## Message structures
#### Create a new post
Request
```json
{
    "header": "<String: title of the post>",
    "body": "<String: body text of the post>"
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
    },
    "image_data": "<Base64 Byte Array>",
    "tags": ["Tag1", "Tag2", ..., "TagN"]
}
```
#### Get a single post
Request
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
#### Get many posts
Request
```json
{
    "post_ids": ["<PostID?>", "<PostID?>", ...],                                  // optional
    "paging": {                                                                 // optional
        "start": "<Number?|ISO8601 timestamp?: start of the current page>",       // default=0
        "end": "<Number?|ISO8601 timestamp?: end of the current page>",           // default=9
        "limit": "<Number?: max number of posts in current page (-1 for all)>"   // optional
    }
}
```
Response
```json
[
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
    },
    ...
]
```

#### Get history for a single post
Request
```json
{
    "post_id": "<PostID: ID of the post>",
}
```
Response
```json
{
    "post_id": "<PostID: ID of the post>",
    "author_id": "<UserID: ID of the author>",
    "created_at": "<ISO8601 timestamp>",
    "updated_at": "<ISO8601 timestamp>",
    "header": [
        {
            "author_id": "<UserID: ID of the author>",
            "created_at": "<ISO8601 timestamp>",
            "data":"<String: title of the post>",
        },
        ...
    ],
    "body": [
        {
            "author_id": "<UserID: ID of the author>",
            "created_at": "<ISO8601 timestamp>",
            "data":"<String: body content of the post>",
        },
        ...
    ]
}
```


#### Get user posts
Request
```json
{
    "author_id": "<UserID: ID of the user>",
}
```
Response
```json
[
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
    },
    ...
]
```

#### Update a single post
Request
```json
{
    "post_id": "<PostID: ID of the post to update>",
    "header": "<String?: updated title of the post>",
    "body": "<String?: updated text of the post>",
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

#### Delete a single post
Request
```json
{
    "post_id": "<PostID: ID of the updated post>"
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


#### Delete many posts
Request
```json
{
    "post_ids": ["<PostID?>", "<PostID?>", ...],
    "author_id": "<UserID?>",
}
```
Response
```json
[
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
    },
    ...
]
```



