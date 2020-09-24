package services

import "errors"

// ErrUserNotFound ...
var ErrUserNotFound = errors.New("User not found")

// ErrBadPassword ...
var ErrBadPassword = errors.New("Bad password")
