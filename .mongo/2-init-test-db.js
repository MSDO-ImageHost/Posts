// https://docs.mongodb.com/manual/reference/operator/query/jsonSchema/#op._S_jsonSchema

db = db.getSiblingDB('posts-test-db');

creator_id_policy = {
   bsonType: 'string',
   description: 'must be a string and is required'
}

post_policy = {
   validator: {
      $jsonSchema: {
         bsonType: 'object',
         required: ['creator_id'],
         properties: {
            creator: creator_id_policy
         }
      }
   }
}

content_policy = {
   validator: {
      $jsonSchema: {
         bsonType: 'object',
         required: ['creator_id', 'data'],
         properties: {
            creator: creator_id_policy,
            data: {
               bsonType: 'string',
               description: 'must be a string and is required'
            }
         }
      }
   }
}

db.createCollection('posts', post_policy)
db.createCollection('postbodies', content_policy)
db.createCollection('postheaders', content_policy)


db.posts.insert