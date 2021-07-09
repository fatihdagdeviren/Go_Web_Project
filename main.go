package main

import (
	"Go_Web_Project/Controllers"
	"Go_Web_Project/DBProvider"
	"Go_Web_Project/Models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)




func TodoIndex(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//param1s := r.URL.Query()["param1"]
	//http://localhost:8000/todos?data=hi&data2=hi2
	todos := Models.Todos{
		Models.Todo{Name: "Write presentation"},
		Models.Todo{Name: "Host meetup"},
		Models.Todo{Name:  r.URL.Query().Get("data") },
	}

	json.NewEncoder(w).Encode(todos)
}


func TodoMultiple(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//http://localhost:8000/todosMultiple?param1=1&param2=2
	param1s := r.URL.Query()["param1"]
	param2s := r.URL.Query()["param2"]
	todos := Models.Todos{
		Models.Todo{Name: "Write presentation"},
		Models.Todo{Name: "Host meetup"},
		Models.Todo{Name:  param1s[0] },
		Models.Todo{Name:  param2s[0] },
	}

	json.NewEncoder(w).Encode(todos)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func PostmanPost(htmlCevap http.ResponseWriter, htmlIstek *http.Request, p httprouter.Params) {
	body, _ := ioutil.ReadAll(htmlIstek.Body) // check for errors
	//[{"Cumle":"hellos"},{"Cumle":"hellos3"},{"Cumle":"hellos4"}]
	keyVal := make([]Models.Soz, 1)
	json.Unmarshal(body, &keyVal) // check for errors

	name := keyVal[0].Cumle
	fmt.Println(name)
}


func main() {

	DBProvider.Mgr.Migrate()
	//DBProvider.Mgr.AddProduct(&Models.Product{Code:"100", Price:100})
	//myUser := Entities.User{
	//	Name:    "Omer",
	//	SurName: "Dagdeviren",
	//}
	//DBProvider.Mgr.AddUser(&myUser)
	//DBProvider.Mgr.AddUserProfile(&Entities.UserProfile{
	//	User:        myUser,
	//	Email:       "omerdgdvrn@gmail.com",
	//	PhoneNumber: "05465675080",
	//	Description: "Ömer Profili",
	//})



	r := httprouter.New()
	r.GET("/",Controllers.SozListele)
	r.GET("/gizli",Controllers.GizliGet)
	r.POST("/gizli", Controllers.GizliPost)
	r.GET("/todos/*data",TodoIndex)
	r.GET("/todosMultiple",TodoMultiple)
	r.GET("/hello/:name", Hello)
	r.POST("/postmanPost", PostmanPost)
	r.GET("/user",Controllers.GetUsers)
	r.GET("/userProfile/:userid",Controllers.GetUserProfile)
	r.GET("/userEdit/:userid",Controllers.GetUserEdit)
	r.POST("/userSave", Controllers.SaveUser)
	//r.GET("/gizli",GizemliSeylerYap)
	http.ListenAndServe(":8000", r)
}

