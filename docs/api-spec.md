# API Specification

## General
#### Access control
Any request must contain a valid session token
```json
{
    "session_token": "<SessionTokenID: valid and active token>", // Usually carried out by HTTP cookies
    ... // rest of request
}
```
#### Rejection message
Any request may get rejected with the following response
```json
{
    "status_code": "<Number: HTTP status code",
    "message": "<String>"
}
```

---
## Posts

#### Create a post `POST posts/post/crud`

Request
```json
{
    "post": {
        "header": "<title of the post>",
        "body": "<String: post body>"
    },
    "tags": ["<TagID", "<TagID", ...],
    },
    "image": {
        "image_id": "<ImageId: ID of the image>"
    }
}
```
Response
```json
{
    "created_at": "<ISO8601 timestamp>",
    "post_id": "<PostID: ID of the post>",
}
```
#### Read post `GET posts/post/crud`

Request
```json
{
    "post_id": "<PostId: ID of the post>"
}
```
Response
```json
{
    "post_id": "<PostID: ID of the post>",
    "created_at": "<ISO8601 timestamp>",
    "creator": "<UserID: ID of the author>",
    "post": {
        "header": "<title of the post>",
        "body": "<String: post body>"
    },
    "tags": ["<TagID", "<TagID", ...],
    },
    "image": {
        "image_id": "<ImageId: ID of the image>"
    }
}
```

#### Update post `PUT posts/post/crud`

Request
```json
{
    "post": {
        "header": "<title of the post>",
        "body": "<post body>",
    },
}
```
Response
```json
{
    "post_created": "<boolean>",
    "created_at": "<ISO8601 timestamp>",
    "post_id": "<PostID: ID of the post>",
    
}
```

#### Delete post `DELETE posts/post/crud`

Request
```json
{
    "post_id": "<PostId: ID of the post>"
}
```

Response
```json
{
    "post_deleted": "<Boolean>",
    "deleted_at": "<ISO8601 timestamp>",
    "post_id": "<PostID: ID of the post>"
}
```

#### Read many posts `GET posts/post/many`

Request
```json
{
    "pagning": {
        "start": "<Number|ISO8601 timestamp: start of the current page>",
        "end": "<Number|ISO8601 timestamp: end of the current page>",
        "limit": "<Number: max number of posts in current page (-1 for all)>"
    }
}
```
Response
```json
[
    {
        "post_id": "<PostID: ID of the post>",
        "created_at": "<ISO8601 timestamp>",
        "creator": "<UserID: ID of the author>",
        "post": {
            "header": "<title of the post>",
            "body": "<String: post body>"
        },
        "tags": ["<TagID", "<TagID", ...],
        },
        "image": {
            "image_id": "<ImageId: ID of the image>"
        }
    },
    ...
]
```

## Comments
#### Read comment `GET posts/comment/crud`
#### Update comment `GET posts/comment/crud`
#### Delete comment `GET posts/comment/crud`
#### Read many comments `GET posts/comment/many`

---

## Likes
#### Read like `GET posts/like/crud`
#### Update like `GET posts/like/crud`
#### Delete like `GET posts/like/crud`
---

## Tags
#### Read tag `GET posts/tag/crud`
#### Update tag `GET posts/tag/crud`
#### Delete tag `GET posts/tag/crud`
#### Read many tags `GET posts/tag/many`
---


## Images
#### Read image `GET posts/like/crud`
#### Update image `GET posts/like/crud`
#### Delete image `GET posts/like/crud`
#### Read many images `GET posts/like/many`
---
