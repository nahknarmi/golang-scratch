package main

import (
	"fmt"
	"net/http"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Person struct {
	Name string
	Phone string
}


func main() {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "123"}, &Person{"Foo", "456"})

	if err != nil {
		panic(err)
	}

	result := Person{}

	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Phone: " + result.Phone)

//	m := martini.Classic()
//	m.Get("/", func() string {
//			return "Hello world"
//		})
//	m.Get("/hello/:name", Auth, func(params martini.Params) string {
//			return "Hello " + params["name"]
//		})
//	m.Run()
}

func Auth(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("API-KEY") != "123" {
		http.Error(res, "Nope", 401)
	}
}

