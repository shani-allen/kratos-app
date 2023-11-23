package data

import (
	"basic-kratos-app/internal/conf"
	"database/sql"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
	sqlDB *sql.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	l:=log.NewHelper(logger)
	var(
		db *gorm.DB
		sqlDB *sql.DB
		err error
	)

	if c.Database.Driver=="mysql"{
		dsn := fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", c.Database.Source)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
		if err != nil {
			l.Errorf("Unable to open db connection: %v", err)
			return nil, nil, fmt.Errorf("unable to open db connection: %w", err)
		}

		sqlDB, err = db.DB()
		if err != nil {
			l.Errorf("Unable to get generic db interface: %v", err)
			return nil, nil, fmt.Errorf("unable to get generic db interface: %w", err)
		}
				 // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
				// sqlDB.SetMaxIdleConns(int(c.Database.MaxIdleConns))

				// // SetMaxOpenConns sets the maximum number of open connections to the database.
				// sqlDB.SetMaxOpenConns(int(c.Database.MaxOpenConns))
		
				// // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
				// sqlDB.SetConnMaxLifetime(time.Duration(c.Database.GetMaxConnLifetimeInMins()) * time.Minute)
	 }

	d := &Data{db, sqlDB}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	
	return d, cleanup, nil
}
