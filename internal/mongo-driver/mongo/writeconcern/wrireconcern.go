// Package writeconcern defines write concerns for MongoDB operations.
package writeconcern // import "go.mongodb.org/mongo-driver/mongo/writeconcern"

import "errors"

// ErrInconsistent indicates that an inconsistent write concern was specified.
var ErrInconsistent = errors.New("a write concern cannot have both w=0 and j=true")
