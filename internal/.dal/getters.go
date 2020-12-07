package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PostData getters
func (pd *PostData) GetID() string            { return pd.ID.(string) }
func (pd *PostData) GetAuthorID() string      { return pd.AuthorID }
func (pd *PostData) GetCreatedAt() time.Time  { return pd.CreatedAt }
func (pd *PostData) GetUpdatedAt() *time.Time { return pd.UpdatedAt }
func (pd *PostData) GetMarkedDeleted() bool   { return *pd.MarkedDeleted }
func (pd *PostData) GetHeader() interface{}   { return pd.Headers }
func (pd *PostData) GetBody() interface{}     { return pd.Bodies }

// PostContent getters
func (pc *PostContent) GetID() string           { return pc.ID.(primitive.ObjectID).Hex() }
func (pc *PostContent) GetAuthorID() string     { return pc.AuthorID }
func (pc *PostContent) GetCreatedAt() time.Time { return pc.CreatedAt }
func (pc *PostContent) GetData() string         { return pc.Data }
func (pc *PostContent) GetMarkedDeleted() bool  { return *pc.MarkedDeleted }
