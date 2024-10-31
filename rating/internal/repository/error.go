package repository

import "errors"

//ErrNotFound is returnes when a requested record is not found.

var ErrNotFound = errors.New("not found")
