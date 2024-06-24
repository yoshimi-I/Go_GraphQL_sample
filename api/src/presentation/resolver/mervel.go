package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	graph_model "graphql_sample_directory/pkg/graphql/model"

	"github.com/google/uuid"
)

// CharacterCreate is the resolver for the characterCreate field.
func (r *mutationResolver) CharacterCreate(ctx context.Context, input graph_model.CharacterCreateInput) (*graph_model.CharacterCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CharacterCreate - characterCreate"))
}

// CharacterUpdate is the resolver for the characterUpdate field.
func (r *mutationResolver) CharacterUpdate(ctx context.Context, input graph_model.CharacterUpdateInput) (*graph_model.CharacterUpdatePayload, error) {
	panic(fmt.Errorf("not implemented: CharacterUpdate - characterUpdate"))
}

// Character is the resolver for the character field.
func (r *queryResolver) Character(ctx context.Context, id uuid.UUID) (*graph_model.Character, error) {
	panic(fmt.Errorf("not implemented: Character - character"))
}

// AllCharacters is the resolver for the allCharacters field.
func (r *queryResolver) AllCharacters(ctx context.Context) ([]*graph_model.Character, error) {
	panic(fmt.Errorf("not implemented: AllCharacters - allCharacters"))
}

// CharactersByType is the resolver for the charactersByType field.
func (r *queryResolver) CharactersByType(ctx context.Context, characterType graph_model.CharacterType) ([]*graph_model.Character, error) {
	panic(fmt.Errorf("not implemented: CharactersByType - charactersByType"))
}
