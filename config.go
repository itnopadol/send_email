package main
//
//import (
//	//"log"
//	//
//	"github.com/BurntSushi/toml"
//	"log"
//)
//
//type Config struct {
//	Server   string
//	Port     int
//	Email    string
//	Password string
//}
//
//func (c *Config) Read() {
//
//	//c.Server = "smtp.gmail.com"
//	//c.Port = 587
//	//c.Email= "satit@nopadol.com"
//	//c.Password="815309917"
//
//	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
//		log.Fatal(err)
//	}
//
//}