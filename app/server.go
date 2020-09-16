package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/satioO/users/config"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	_userDelivery "github.com/satioO/users/app/users/delivery/http"
	_userRepository "github.com/satioO/users/app/users/repository/mongo"
	_userUsecase "github.com/satioO/users/app/users/usecase"

	_bookDelivery "github.com/satioO/users/app/books/delivery/http"
	_bookRepository "github.com/satioO/users/app/books/repository/mongo"
	_bookUsecase "github.com/satioO/users/app/books/usecase"

	_authDelivery "github.com/satioO/users/app/auth/delivery/http"
	_authRepository "github.com/satioO/users/app/auth/repository"
	_authUsecase "github.com/satioO/users/app/auth/usecase"
)

// Server defines the routing configuration
type Server struct {
	HTTPServer *http.Server
	Router     *mux.Router
	DB         *mongo.Database
}

// Initialize the storage and router
func (a *Server) Initialize() {
	a.DB = ConnectDB("users_ms")
	a.Router = mux.NewRouter().StrictSlash(true)

	// Auth
	authRepo := _authRepository.NewAuthRepository(a.DB)
	authUsecase := _authUsecase.NewAuthUsecase(authRepo)
	_authDelivery.MakeAuthHandlers(a.Router, authUsecase)

	// Users
	usersRepo := _userRepository.NewUsersRepository(a.DB)
	usersUsecase := _userUsecase.NewUsersUsecase(usersRepo)
	_userDelivery.MakeUserHandlers(a.Router, usersUsecase)

	// Books
	booksRepo := _bookRepository.NewBookRepository(a.DB)
	booksUsecase := _bookUsecase.NewBookUsecase(booksRepo)
	_bookDelivery.MakeBookHandlers(a.Router, booksUsecase)

}

// Run the application
func (a *Server) Run(port string) error {
	// Setup configuration
	configuration := SetupConfig()

	// HTTP Server
	a.HTTPServer = &http.Server{
		Addr:           ":" + configuration.Server.Port,
		Handler:        a.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		fmt.Println("User Management Service started listening on port", configuration.Server.Port)
		if err := a.HTTPServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.HTTPServer.Shutdown(ctx)
}

// ConnectDB estabilishes the connection with the database
func ConnectDB(dbName string) *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//To close the connection at the end
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	return client.Database(dbName)
}

// SetupConfig manages configuration across application
func SetupConfig() *config.Configurations {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	var configuration config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return &configuration
}
