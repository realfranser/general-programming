package main

import (
	"log"

	"microservice_example/controllers"
	"microservice_example/db/orm"
	"microservice_example/db/plain_sql"
	"microservice_example/services"
	"microservice_example/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Flag parse

	// Database connections
	sql, err := plain_sql.NewPostgreSQLsql()
	if err != nil {
		log.Fatalf("Could not initialize sql Database connection %s", err)
	}
	defer sql.Close()

	sqlx, err := plain_sql.NewPostgreSQLsqlx()
	if err != nil {
		log.Fatalf("Could not initialize sqlx Database connection %s", err)
	}
	defer sqlx.Close()

	pgx, err := plain_sql.NewPostgreSQLpgx()
	if err != nil {
		log.Fatalf("Could not initialize pgx Database connection %s", err)
	}
	defer pgx.Close()

	_, err = orm.NewPostgreSQLgorm()
	if err != nil {
		log.Fatalf("Could not initialize gorm Database connection %s", err)
	}

	// Memcache connection
	memcacheClient, err := services.NewMemcached()
	if err != nil {
		log.Fatalf("Could not initialize Memcached client %s", err)
	}
	defer memcacheClient.CloseMemcacheClient()

	// Start app with fiber
	app := fiber.New()

	// Middlewares
	app.Use(recover.New())
	app.Use(logger.New())

	// Routing
	namesApi := app.Group("/names") // /names

	// Memcached endpoints
	memcachedApi := namesApi.Group("/memcached")          // /names/memcached
	memcachedApi.Get("/:id", controllers.GetNameMemcache) // GET /names/memcached/:id

	// DB endpoints
	dbApi := namesApi.Group("/db")                  // /names/db
	dbApi.Get("/gorm/:id", controllers.GetNameGorm) // GET /names/db/gorm/:id
	dbApi.Get("/sql/:id", controllers.GetNameSQL)   // GET /names/db/sql/:id
	dbApi.Get("/sqlx/:id", controllers.GetNameSQLX) // GET /names/db/sqlx/:id
	dbApi.Get("/pgx/:id", controllers.GetNamePGX)   // GET /names/db/pgx/:id

	// Handle not founds
	app.Use(utils.NotFound)

	log.Fatal(app.Listen(":8080"))
}
