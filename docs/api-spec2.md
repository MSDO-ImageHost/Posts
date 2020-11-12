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
Request
```json
{
    "post_id": "<PostID: ID of the post>",
    "pagning": {
        "start": "<Number|ISO8601 timestamp: start of the current page>",
        "end": "<Number|ISO8601 timestamp: end of the current page>",
        "limit": "<Number: max number of posts in current page (-1 for all)>"
    }
}
```
Response
```json
{
    "post_id": "<PostID: ID of the post>",
    "comments": ["<CommentID>","<CommentID>"]
}
```
#### Add comment `GET posts/comment/crud`
Request
```json
{
    "post_id": "<PostID: ID of the post>",
    "comment_id": "<CommentID: ID of the new comment>"
}
```
Response
```json
{
    "post_id": "<PostID: ID of the post>",
    "comment_id": "<CommentID: ID of the new comment>"
}
```
#### Delete comment `DELETE posts/comment/crud`
Request
```json
{
    "post_id": "<PostID: ID of the post>",
    "comment_id": "<CommentID: ID of the comment to be deleted>"
}
```
Response
```json
{
    "post_id": "<PostID: ID of the post>",
    "comment_id": "<CommentID: ID of the deleted comment>"
}
```
---

## Likes
#### Read like `GET posts/like/crud`
Request
```json
{
    "post_id": "<PostID: ID of the post>",
}
```
Response
```json
{
    "number_of_likes": "<Number>",
    "likes": ["<UserID>","<UserID>"]
}
```
#### Add like `PUT posts/like/crud`
Request
```json
{
    "post_id": "<PostID: ID of the post>",
    "user_id": "<UserID: User ID of the new like>"
}
```
Response
```json
{
    "post_id": "<PostID: ID of the post>",
    "user_id": "<UserID: User ID of the new like>"
}
```
#### Delete like `DELETE posts/like/crud`
Request
```json
{
    "post_id": "<PostID: ID of the post>",
    "user_id": "<UserID: ID of the user to be removed from likes>"
}
```
Response
```json
{
    "post_id": "<PostID: ID of the post>",
    "user_id": "<UserID>"
}
```
---

## Tags
#### Read tag `GET posts/tag/crud`
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
    "tags": ["<TagID>", "<TagID>", ...]
}
```
#### Add tag `POST posts/tag/crud`
Request
```json
{
    "post_id": "<PostID: ID of the post>",
    "tag_id": "<TagID>"
}
```
Response
```json
{
    "post_id": "<PostID: ID of the post>",
    "tag_id": "<TagID>"
}
```

#### Update tag `PUT posts/tag/crud`
Request
```json
{
    "post_id": "<PostID: ID of the post>",
    "old_tag_id": "<TagID: ID of the old tag>",
    "new_tag_id": "<TagID: ID of the new tag>"
}
```
Response
```json
{
    "post_id": "<PostID: ID of the post>",
    "tag_id": "<TagID: ID of the new tag>"
}
```
#### Delete tag `DELETE posts/tag/crud`
Request
```json
{
    "post_id": "<PostID: ID of the post>",
    "tag_id": "<TagID: ID of the tag to be deleted>"
}
```
Response
```json
{
    "post_id": "<PostID: ID of the post>",
    "tag_id": "<TagID: ID of the deleted tag>"
}
```



## Images
#### Read image `GET posts/like/crud`
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
    "image_id": "<ImageID: id of the image>"
}
```
#### Update image `PUT posts/like/crud`
Request
```json
{
    "post_id": "<PostID: ID of the post>",
    "image_id": "<ImageID: new image ID>"
}
```
Response
```json
{
    "post_id": "<PostID: ID of the post>",
    "image_id": "<ImageID: id of the image>"
}
```
#### Delete image `DELETE posts/like/crud`
Request
```json
{
    "post_id": "<PostID: ID of the post>",
    "image_id": "<ImageID: id of the image>"
}
```
Response
```json
{
    "post_id": "<PostID: ID of the post>",
    "image_id": "<ImageID: id of the image>"
}
```
---
