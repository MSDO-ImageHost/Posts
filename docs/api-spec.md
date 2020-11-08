# API Specification




## Create a new post

Endpoint: `posts/createpost`

Request
```json
{
    "creator": "<user id>"
}
```

Response
```json
{
    "post_created": "<boolean>",
    "created_at": "<ISO8601 timestamp>",
    "post_id": "<postid>",
    
}
```