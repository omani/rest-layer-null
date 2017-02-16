// Package null is an example REST backend storage that processes everything in memory and stores nothing.
package null

import (
	"context"

	"github.com/rs/rest-layer/resource"
)

// Handler is an example handler storing data nowhere.
type Handler struct{}

// NewHandler creates a new NullHandler.
func NewHandler() *Handler {
	return &Handler{}
}

// Insert inserts new items in memory
func (n *Handler) Insert(ctx context.Context, items []*resource.Item) (err error) {
	return nil
}

// Update replace an item by a new one in memory
func (n *Handler) Update(ctx context.Context, item *resource.Item, original *resource.Item) (err error) {
	return nil
}

// Delete deletes an item from memory
func (n *Handler) Delete(ctx context.Context, item *resource.Item) (err error) {
	return nil
}

// Clear clears all items from the memory store matching the lookup
func (n *Handler) Clear(ctx context.Context, lookup *resource.Lookup) (total int, err error) {
	return 0, nil
}

// Find items from memory matching the provided lookup
func (n *Handler) Find(ctx context.Context, lookup *resource.Lookup, offset, limit int) (list *resource.ItemList, err error) {
	return &resource.ItemList{}, nil
}
