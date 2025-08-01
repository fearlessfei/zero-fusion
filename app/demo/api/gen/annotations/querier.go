package annotations

import "gorm.io/gen"

type Querier interface {
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id int64) (*gen.T, error)

	// SELECT * FROM @@table
	GetAll() ([]*gen.T, error)
}
