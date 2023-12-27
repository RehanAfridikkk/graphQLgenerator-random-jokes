// graph/schema.resolvers.go
package graph

import (
	"context"
	"math/rand"
	"time"

	"graphQLGEN/graph/model"
)

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Joke(ctx context.Context, args struct{ ID int }) (*model.Joke, error) {
	// Your implementation here

	id := args.ID

	if id == 0 {
		// If id is 0, return a random joke
		return r.GetRandomJoke(), nil
	}

	// Fetch the joke with the specified ID
	return r.GetJokeByID(id), nil
}

// GetJokeByID retrieves a joke from the database by its ID
func (r *Resolver) GetJokeByID(id int) *model.Joke {
	var joke model.Joke
	if err := r.DB.First(&joke, id).Error; err != nil {
		return &model.Joke{
			Content: "Joke not found",
			ID:      0,
		}
	}
	return &joke
}

// GetRandomJoke returns a random joke
func (r *Resolver) GetRandomJoke() *model.Joke {
	var jokes []model.Joke
	r.DB.Find(&jokes)

	if len(jokes) == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(jokes))
	return &jokes[index]
}

// Todos is required to satisfy the QueryResolver interface
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	// Implement the Todos method if needed
	return nil, nil
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

// CreateTodo is a mutation operation to create a new Todo
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	newTodo := model.Todo{
		Text: input.Text,
		Done: false,
	}

	// Save the Todo to the database
	if err := r.DB.Create(&newTodo).Error; err != nil {
		return nil, err
	}

	// Return the created Todo
	return &newTodo, nil
}
