// graph/resolver.go
package graph

import "gorm.io/gorm"

type Resolver struct {
	DB *gorm.DB // Assuming you pass the database connection to the resolver
}

func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{DB: db}
}

// Add resolver methods as needed
