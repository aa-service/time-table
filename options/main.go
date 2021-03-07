package options

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Options struct {
	authorizator gin.HandlerFunc
	db           *gorm.DB
}

func New(opts ...interface{}) (*Options, error) {
	options := &Options{}
	for _, opt := range opts {
		if o, ok := opt.(gin.HandlerFunc); ok {
			options.authorizator = o
			continue
		}
		//
		if o, ok := opt.(*gorm.DB); ok {
			options.db = o
			continue
		}
	}
	if options.db == nil {
		return nil, errors.New("missing db")
	}
	if options.authorizator == nil {
		return nil, errors.New("missing authorization middleware")
	}
	return options, nil
}

func (o Options) Authorizator() gin.HandlerFunc {
	return o.authorizator
}

func (o Options) DB() *gorm.DB {
	return o.db
}
