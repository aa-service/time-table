package options

import (
	"errors"

	"gorm.io/gorm"
)

type Options struct {
	db *gorm.DB
}

func New(opts ...interface{}) (*Options, error) {
	options := &Options{}
	for _, opt := range opts {
		//
		if o, ok := opt.(*gorm.DB); ok {
			options.db = o
			continue
		}
	}
	if options.db == nil {
		return nil, errors.New("missing db")
	}
	return options, nil
}

func (o Options) DB() *gorm.DB {
	return o.db
}
