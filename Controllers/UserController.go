package Controllers

import (
	"Go_Web_Project/DBProvider"
	"Go_Web_Project/TemplateParser"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func GetUsers(htmlCevap http.ResponseWriter, htmlIstek *http.Request, p httprouter.Params) {
	_, users := DBProvider.Mgr.GetAllUsers()
	sablon, _ := template.ParseFiles("Views/User/index.html")
	sablon.Execute(htmlCevap, users)
}

func GetUserEdit(htmlCevap http.ResponseWriter, htmlIstek *http.Request, p httprouter.Params) {

	postedVal := p.ByName("userid")
	myUser := DBProvider.Mgr.GetUserByID(postedVal)
	TemplateParser.Display(htmlCevap, "useredit", myUser)
}

func GetUserProfile(htmlCevap http.ResponseWriter, htmlIstek *http.Request, p httprouter.Params) {

	postedVal := p.ByName("userid")
	myUserProfile := DBProvider.Mgr.GetUserProfileByUserID(postedVal)
	TemplateParser.Display(htmlCevap, "userprofileindex", myUserProfile)
}

func SaveUser(htmlCevap http.ResponseWriter, htmlIstek *http.Request, p httprouter.Params){

  		id :=	htmlIstek.FormValue("id")
  		name := htmlIstek.FormValue("name")
		surname := htmlIstek.FormValue("surname")
		myUser := DBProvider.Mgr.GetUserByID(id)
		myUser.Name = name
		myUser.SurName = surname
		DBProvider.MyDB.Save(myUser)

		_, users := DBProvider.Mgr.GetAllUsers()
		sablon, _ := template.ParseFiles("Views/User/index.html")
		sablon.Execute(htmlCevap, users)
}