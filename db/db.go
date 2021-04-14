package db

import (
	"context"
	"github.com/sobocinski/api-go-update/domain"
	"github.com/go-pg/pg/extra/pgdebug"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	log "github.com/sirupsen/logrus"
)

// NewDatabase create database connection
func NewDatabase(ctx context.Context) (*pg.DB, error) {
	log.Debug("Connecting database...")
	opt, err := pg.ParseURL("postgres://user:pass@localhost:5432/db_name?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	if err := db.Ping(ctx); err != nil {
		return db, err
	}

	err = createSchema(db)
	if err != nil {
		return db, err
	}

	//add logger
	db.AddQueryHook(pgdebug.DebugHook{
		// Print all queries.
		Verbose: true,
	})
	log.Debug("Database connected...")


	//adding dummy data to db
	user1 := &domain.User{
		Email: "user1@example.com",
		Lang: "en",
	}

	user2 := &domain.User{
		Email: "user2@example.com",
		Lang: "pl",
	}
	_, err = db.Model(user1).Insert()
	if err != nil {
		panic(err)
	}
	_, err = db.Model(user2).Insert()
	if err != nil {
		panic(err)
	}

	car1 := &domain.Car{
		UserId: user1.Id,
		Model: "BMW",
	}
	car2 := &domain.Car{
		UserId: user1.Id,
		Model: "Audi",
	}
	car3 := &domain.Car{
		UserId: user2.Id,
		Model: "Opel",
	}
	_, err = db.Model(car1).Insert()
	if err != nil {
		panic(err)
	}
	_, err = db.Model(car2).Insert()
	if err != nil {
		panic(err)
	}
	_, err = db.Model(car3).Insert()
	if err != nil {
		panic(err)
	}


	// Select user by primary key.
	car := &domain.Car{Base: domain.Base{Id: car1.Id}}
	err = db.Model(car).WherePK().Select()
	if err != nil {
		panic(err)
	}

	car.Model = "Subaru"
	_, err = db.Model(car).WherePK().Update()
	if err != nil {
		panic(err)
	}

	return db, nil
}


// createSchema creates database schema for models.
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*domain.User)(nil),
		(*domain.Car)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
			//IfNotExists: true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
