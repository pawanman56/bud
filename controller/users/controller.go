package users

import (
	context "context"
)

// Controller for users
type Controller struct {
}

// User struct
type User struct {
	ID int `json:"id"`
}

// Index of users
// GET /users
func (c *Controller) Index(ctx context.Context) (users []*User, err error) {
	return []*User{}, nil
}

// New returns a view for creating a new user
// GET /users/new
func (c *Controller) New(ctx context.Context) {
}

// Create user
// POST /users
func (c *Controller) Create(ctx context.Context) (user *User, err error) {
	return &User{
		ID: 0,
	}, nil
}

// Show user
// GET /users/:id
func (c *Controller) Show(ctx context.Context, id int) (user *User, err error) {
	return &User{
		ID: id,
	}, nil
}

// Edit returns a view for editing a user
// GET /users/:id/edit
func (c *Controller) Edit(ctx context.Context, id int) (user *User, err error) {
	return &User{
		ID: id,
	}, nil
}

// Update user
// PATCH /users/:id
func (c *Controller) Update(ctx context.Context, id int) (user *User, err error) {
	return &User{
		ID: id,
	}, nil
}

// Delete user
// DELETE /users/:id
func (c *Controller) Delete(ctx context.Context, id int) error {
	return nil
}
