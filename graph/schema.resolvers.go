package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/smoothbronco/blog_example/graph/generated"
	"github.com/smoothbronco/blog_example/graph/model"
)

// CreateArticle is the resolver for the createArticle field.
func (r *mutationResolver) CreateArticle(ctx context.Context, input model.CreateInput) (*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateArticle is the resolver for the updateArticle field.
func (r *mutationResolver) UpdateArticle(ctx context.Context, input model.UpdateInput) (*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteArticle is the resolver for the deleteArticle field.
func (r *mutationResolver) DeleteArticle(ctx context.Context, input int) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Article is the resolver for the article field.
func (r *queryResolver) Article(ctx context.Context, input int) (*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// Articles is the resolver for the articles field.
func (r *queryResolver) Articles(ctx context.Context) ([]*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
