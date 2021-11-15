package main

import (
	"html/template"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
	
)
type Webb struct {
	Id int `json "id"`
	Tovar_name string `json "tovar_name"`
	Tovar_buy_id int `json "tovar_buy_id"`
}
var tovars = []Webb{};
func Tovar(w http.ResponseWriter, r *http.Request){
	forPostGres := "user = postgres	password = Vinter1973 dbname = WebStore sslmode=disable"
	db,err := sql.Open("postgres",forPostGres)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	
	insert, err := db.Query("insert into web(tovar_name,tovar_buy_id) values($1,$2)", "Miban", 23588220)
	if err != nil {
		log.Print(err)
	}
	defer insert.Close()
	

}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		log.Print(err)
	}
	t.ExecuteTemplate(w, "index", nil)}

	
func buy_info(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/buy_info.html", "templates/header.html", "templates/footer.html")
	if err != nil{
		log.Print(err)
	}
	forPostGres := "user = postgres	password = Vinter1973 dbname = WebStore sslmode=disable"
	db,err := sql.Open("postgres",forPostGres)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	res,err := db.Query("select * from web")
	if err != nil {
		log.Print(err)
	}
	tovars = []Webb{}
	for res.Next(){
		var tovar Webb
		err = res.Scan(&tovar.Id,&tovar.Tovar_buy_id,&tovar.Tovar_name)
		if err != nil {
			log.Print(err)
		}
		tovars = append(tovars,tovar)

	}
	
	if err != nil {
		 log.Print(err)
	}
	t.ExecuteTemplate(w,"buy_info", tovars)
}
func handleFunc() {
	//Tovar();
	rtr:=mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/d", Tovar)
	rtr.HandleFunc("/buy",buy_info).Methods("GET")

	http.Handle("/", rtr)
	

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	
	http.ListenAndServe(":8080", nil)}
func main() {
	handleFunc()
}


