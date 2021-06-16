package main

import (
	"bot/config"
	"bot/constants"
	"bot/daos"
	"bot/handlers"
	m "bot/middlewares"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"bot/redis"
	"bot/routes"
	"bot/services"
	"bot/shared"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

var port string = ""

func init() {
	conf := config.ConfigReader
	port = conf.GetString(os.Getenv(constants.SOILPROTECTIONENV) + ".port")
}
func main() {
	argsWithoutProg := os.Args[1:]
	sh := shared.NewShared(SplitCmdArguments(argsWithoutProg))
	fmt.Println(sh)
	redis := redis.Connect()
	fmt.Println("redis cunt", redis.Pool.Stats().ActiveCount)
	db := daos.GetDaos(sh, redis)
	fmt.Println(db)
	route := routes.GetRoute(handlers.GetHandler(services.GetService(db, sh, redis), sh, redis), sh, redis)
	r := mux.NewRouter()
	r.Use(m.Log)
	r.Use(m.AllowCors)
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})
	route.Category(r)

	log.Println("listening in port : ", sh.GetCmdArg(constants.PORT))
	log.Fatal(http.ListenAndServe("0.0.0.0:"+sh.GetCmdArg(constants.PORT), r))
}

//SplitCmdArguments : ""
func SplitCmdArguments(args []string) map[string]string {
	m := make(map[string]string)
	for _, v := range args {
		strs := strings.Split(v, "=")
		if len(strs) == 2 {
			m[strs[0]] = strs[1]
		} else {
			log.Println("not proper arguments", strs)
		}
	}
	return m
}
