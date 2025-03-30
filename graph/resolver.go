package graph

import (
	"github.com/Jeanga7/graphql-go-api.git/graph/resolvers"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
    DB           *gorm.DB
    UserResolver *resolvers.UserResolver
}

func NewResolver(db *gorm.DB) *Resolver {
    return &Resolver{
        DB: db,
        UserResolver: &resolvers.UserResolver{
            DB: db,
        },
    }
}