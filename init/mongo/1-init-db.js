// https://docs.mongodb.com/manual/reference/operator/query/jsonSchema/#op._S_jsonSchema

db = db.getSiblingDB('posts-db');

creator_id_policy = {
   bsonType: 'string',
   description: 'must be a string and is required'
}

scaffold_policy = {
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

db.createCollection('scaffolds', scaffold_policy)
db.createCollection('bodies', content_policy)
db.createCollection('headers', content_policy)
