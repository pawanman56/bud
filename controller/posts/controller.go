package posts

import (
	context "context"
)

// Controller for posts
type Controller struct {
}

// Post struct
type Post struct {
	ID int `json:"id"`
}

// Index of posts
// GET /posts
func (c *Controller) Index(ctx context.Context) (posts []*Post, err error) {
	return []*Post{}, nil
}

// New returns a view for creating a new post
// GET /posts/new
func (c *Controller) New(ctx context.Context) {
}

// Create post
// POST /posts
func (c *Controller) Create(ctx context.Context) (post *Post, err error) {
	return &Post{
		ID: 0,
	}, nil
}

// Show post
// GET /posts/:id
func (c *Controller) Show(ctx context.Context, id int) (post *Post, err error) {
	return &Post{
		ID: id,
	}, nil
}

// Edit returns a view for editing a post
// GET /posts/:id/edit
func (c *Controller) Edit(ctx context.Context, id int) (post *Post, err error) {
	return &Post{
		ID: id,
	}, nil
}

// Update post
// PATCH /posts/:id
func (c *Controller) Update(ctx context.Context, id int) (post *Post, err error) {
	return &Post{
		ID: id,
	}, nil
}

// Delete post
// DELETE /posts/:id
func (c *Controller) Delete(ctx context.Context, id int) error {
	return nil
}
