package main

import (
	"context"
	"database/sql"
	db "example-crud/db/sqlc"
	"example-crud/util"
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/lib/pq"
)

const (
	DriverName     = "postgres"
	DataSourceName = "postgresql://root:root@localhost:5432/example_crud?sslmode=disable"
)

func main() {
	// notify interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		ossignal := <-c
		log.Printf("OS Signal %v", ossignal)
		cancel()
	}()

	sqldb, err := sql.Open(DriverName, DataSourceName)
	if err != nil {
		panic(err)
	}

	q := db.New(sqldb)

	// create employee
	emp, err := q.CreateEmployee(ctx, db.CreateEmployeeParams{
		Code:        util.RandomCode(),
		Name:        util.RandomName(),
		Email:       util.RandomEmail(),
		PhoneNumber: util.RandomPhoneNumber(),
	})

	if err != nil {
		log.Fatalf("Error create employee, %v", err)
	}

	log.Printf("Create Employee : %+v, %p", emp, &emp)

	// sleep
	log.Println("SLEEP")
	time.Sleep(5 * time.Second)

	// update employee
	emp, err = q.UpdateEmployee(ctx, db.UpdateEmployeeParams{
		ID:          emp.ID,
		Name:        util.RandomName(),
		Email:       util.RandomEmail(),
		PhoneNumber: util.RandomPhoneNumber(),
	})
	if err != nil {
		log.Printf("Error update employee, %v", err)
	}

	log.Printf("Update Employee : %+v, %p", emp, &emp)

	// sleep
	log.Println("SLEEP")
	time.Sleep(5 * time.Second)

	// get employee
	emp, err = q.GetEmployee(ctx, emp.ID)
	if err != nil {
		log.Printf("Error get employee, %v", err)
	}

	log.Printf("Get Employee : %+v, %p", emp, &emp)

	// sleep
	log.Println("SLEEP")
	time.Sleep(5 * time.Second)

	// list employee
	emps, err := q.ListEmployees(ctx, db.ListEmployeesParams{
		Limit:  100,
		Offset: 0,
	})
	if err != nil {
		log.Printf("Error List employee, %v", err)
	}

	time.Sleep(5 * time.Second)
	for _, em := range emps {
		log.Printf("List Employee : %+v, %p", em, &em)
	}

	// gracefully shutdown
	<-ctx.Done()
	log.Println("STOPPED")
	if sqldb != nil {
		sqldb.Close()
	}
}
