package Controllers

import (
	"Go_Web_Project/Models"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)


func GizliGet(htmlCevap http.ResponseWriter, htmlIstek *http.Request, p httprouter.Params){
	
	liste := make([]Models.Soz, 1) //var liste [3]Soz
	liste[0] = Models.Soz{Cumle:"Gizli ekran acildi", Soyleyen:"Abert Einstein"}
	sablon, _ := template.ParseFiles("Views/Gizli/index.html")
	sablon.Execute(htmlCevap, liste)
	//log.Println(ornekSoz1.Soyleyen," : ", ornekSoz1.Cumle)
}


func GizliPost(htmlCevap http.ResponseWriter, htmlIstek *http.Request, p httprouter.Params){

	liste := make([]Models.Soz, 1) //var liste [3]Soz
	myVal := htmlIstek.FormValue("postEdilecekData1")
	//
	//err := r.ParseForm()
	//if err != nil {
	//	// handle err
	//}
	//
	//name := r.FormValue("name")

	liste[0] = Models.Soz{Cumle:myVal, Soyleyen:"Abert Einstein"}
	sablon, _ := template.ParseFiles("Views/Gizli/index.html")
	sablon.Execute(htmlCevap, liste)
	//log.Println(ornekSoz1.Soyleyen," : ", ornekSoz1.Cumle)
}
