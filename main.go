package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type config struct{
	Port string `json:"port"`
	ConnectionString string `json:"ConnectionString"`
}

type test interface{}

func ReadConfig() (*config, error){
	var C config
	// var T test
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		return nil, err
	}
	fmt.Println(C)
	// fmt.Println(T["port"])

	return &C, nil

}

func Process(w http.ResponseWriter, r *http.Request){

}

func main(){

	conf, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conf)
	http.HandleFunc("/", Process)
	http.ListenAndServe(":8082", nil)
}