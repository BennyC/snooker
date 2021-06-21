package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fudge/snooker/internal/rest"
	"github.com/fudge/snooker/internal/storage/postgres"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// @TODO Allow only environment variables
	if err := viper.ReadInConfig(); err != nil {
		panic("cannot read environment")
	}

	db := postgres.New(viper.GetString("DB_CONNECTION_STRING"))
	users := &postgres.Users{Db: db}
	s := rest.New(users)

	srv := &http.Server{
		Addr:         viper.GetString("HOST_ADDRESS"),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      s.Routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
