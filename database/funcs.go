package database

import "context"

// IsQueryableContext checks if the given context is a QueryableContext.
//
// Parameters:
// - ctx: The context to check.
//
// Returns:
// - bool: True if the context is a QueryableContext, false otherwise.
func IsQueryableContext(ctx context.Context) bool {
	if _, ok := ctx.(QueryableContext); ok {
		return true
	}

	return false
}

// Context returns a new context with the given QueryableInterface.
// It is a shortcut for NewQueryableContext.
//
// Example:
// 	ctx := database.Context(context.Background(), tx)
//
// Parameters:
// - ctx: The parent context.
// - queryable: The QueryableInterface to be associated with the context.
//
// Returns:
// - QueryableContext: A new context with the given QueryableInterface.
func Context(ctx context.Context, queryable QueryableInterface) QueryableContext {
	return NewQueryableContext(ctx, queryable)
}
