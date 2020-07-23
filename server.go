package main

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"soutube/graph"
	"soutube/graph/generated"
	"soutube/postgres"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")

	pgDB := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "131112awt",
		Database: "Soutube",
	})

	//opt, err := pg.ParseURL("postgres://oadxreyqmffgnh:1bd75bcd83cf9e171d63dc62dc19be625150e7541b0dc2f2b4ff8fe56706fa69@ec2-54-247-118-139.eu-west-1.compute.amazonaws.com:5432/d7sbv5fgdtfjdr?sslmode=require")
	//if err != nil {
	//	panic(err)
	//}
	//
	//pgDB := pg.Connect(opt)
	pgDB.AddQueryHook(postgres.DBLogger{})

	defer pgDB.Close()
	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: pgDB}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}



