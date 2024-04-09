/*
Package main is the entry point of the application.

Important:
- "log" package is used for logging errors.
- "server/db" package is used to interact with the database.

The main function initializes the database and logs any errors that occur during the initialization.
*/
package main

import (
	"log"
	"server/db"
	"server/internal/users"
	"server/router"
)

/*
main function is the entry point of the application.

Parameters:
- None

Return values:
- (int, error): The first value is the number of files affected by the operation, and the second value is an error that occurred during the operation.

The function initializes the database using the NewDatabase function from the "server/db" package. If an error occurs during the initialization, it logs the error using the Fatal function from the "log" package.
*/
func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Error: %s", err)
	}

	userRep := users.NewRepository(dbConn.GetDB())
	userSvc := users.NewService(userRep)
	userHandler := users.NewHandler(userSvc)

	router.InitHandler(userHandler)
	router.Start("0.0.0.0:8080")

}
