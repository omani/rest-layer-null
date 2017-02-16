// Package null is an example REST backend storage that processes everything in memory and stores nothing.
package null

import (
	"context"

	"github.com/rs/rest-layer/resource"
)

// Handler is an example handler storing data nowhere.
type Handler struct{}

// NewHandler creates a new Handler.
func NewHandler() *Handler {
	return &Handler{}
}

// Insert doesn't insert new items
func (n *Handler) Insert(ctx context.Context, items []*resource.Item) (err error) {
	return nil
}

// Update doesn't replace anything
func (n *Handler) Update(ctx context.Context, item *resource.Item, original *resource.Item) (err error) {
	return nil
}

// Delete deletes nothing
func (n *Handler) Delete(ctx context.Context, item *resource.Item) (err error) {
	return nil
}

// Clear clears all items which didn't exist in the first place.
func (n *Handler) Clear(ctx context.Context, lookup *resource.Lookup) (total int, err error) {
	return 0, nil
}

// Find items out of nowhere matching the provided lookup
func (n *Handler) Find(ctx context.Context, lookup *resource.Lookup, offset, limit int) (list *resource.ItemList, err error) {
	return &resource.ItemList{}, nil
}
