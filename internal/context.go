package internal

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

var max_retry = 10

func NewContext(db_dsn string) (*Context, error) {
	var ctx Context
	var db *gorm.DB
	var err error

	for i := 0; i < max_retry; i++ {
		db, err = gorm.Open("postgres", db_dsn)
		if err == nil {
			break
		}
		if i == max_retry - 1 {
			return nil, fmt.Errorf("Too much retry to connect to %q, last error: %v",
				db_dsn,
				err,
			)
		}
		w := (1 << uint(i)) * time.Second
		log.Printf("Error %q while connecting to %q, waiting %v", err, db_dsn, w)
		time.Sleep(w)
	}
	ctx.Db = db
	if err = InitDatabase(&ctx); err != nil {
		return nil, fmt.Errorf("NewContext: %f", err)
	}
	return &ctx, nil
}
