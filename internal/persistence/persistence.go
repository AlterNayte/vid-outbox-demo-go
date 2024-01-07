package persistence

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"vid-outbox-demo-go/internal"
)

type sqlConnection struct {
	*gorm.DB
}

type Connection interface {
	Select(query interface{}, args ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Preload(query string, args ...interface{}) (tx *gorm.DB)
}

func OpenSQLConnection() (Connection, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.Callback().Create().Before("gorm:create").Register("outbox:insert", insertOutbox)
	if err != nil {
		return nil, err
	}
	return &sqlConnection{
		DB: db,
	}, nil
}

func insertOutbox(db *gorm.DB) {
	if db.Statement.Schema != nil {
		if db.Error != nil {
			return
		}
		for _, field := range db.Statement.Schema.Fields {
			if field.Name == "Events" {
				val, isEmpty := field.ValueOf(db.Statement.Context, db.Statement.ReflectValue)
				if !isEmpty {
					if evts, ok := val.([]*internal.Event); ok {
						outboxSession := db.Session(&gorm.Session{SkipHooks: true})
						for _, evt := range evts {
							err := outboxSession.Exec(`INSERT INTO outbox 
    (id,aggregateid, aggregatetype, type, payload) VALUES (?,?,?,?,?)`,
								evt.ID, evt.AggregateId, evt.AggregateType, evt.EventType, evt.Payload).Error
							if err != nil {
								db.AddError(err)
							}
							err = outboxSession.Exec("DELETE FROM outbox WHERE id = ?", evt.ID).Error
							if err != nil {
								db.AddError(err)
							}
						}
					}
				}
				break
			}
		}
	}
}
