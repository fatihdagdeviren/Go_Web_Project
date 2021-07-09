package Controllers

import (
	"Go_Web_Project/Models"
	"Go_Web_Project/TemplateParser"
	"github.com/julienschmidt/httprouter"
	//"html/template"
	"net/http"
)



func SozListele(htmlCevap http.ResponseWriter, htmlIstek *http.Request, p httprouter.Params){
	//ornekSoz1 := Soz{Cumle:"Herseyi olabildigince sade yapin, ama basit degil", Soyleyen:"Abert Einstein"}
	//var liste [3]Soz
	liste := make([]Models.Soz, 3) //var liste [3]Soz
	liste[0] = Models.Soz{Cumle:"Herseyi olabildigince sade yapin, ama basit degil", Soyleyen:"Abert Einstein"}
	liste[1] = Models.Soz{Cumle:"Herseyi olabildigince1 sade yapin, ama basit degil", Soyleyen:"Abert1 Einstein"}
	liste[2] = Models.Soz{Cumle:"Herseyi olabildigince2 sade yapin, ama basit degil", Soyleyen:"Abert2 Einstein"}
	TemplateParser.Display(htmlCevap, "homeindex", liste)
	//sablon, _ := template.ParseFiles("Views/Home/index.html")
	//sablon.Execute(htmlCevap, liste)
	//log.Println(ornekSoz1.Soyleyen," : ", ornekSoz1.Cumle)
}
