package character

import "context"

// Character Repository Interface
// Character manipulation requires that Character.Id is set.
type Repository interface {
	Insert(ctx context.Context, inCharacter Character) (Character, error)
	Get(ctx context.Context, id string) (Character, error)
	GetWithName(ctx context.Context, name string) (Character, error)
	Delete(ctx context.Context, id string) (Character, error)
	List(ctx context.Context) ([]Character, error)
}
