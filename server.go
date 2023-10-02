package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func envPage(w http.ResponseWriter, r *http.Request) {

	envPage := "pages/pagina de envios.html"

	tpl, err := template.ParseFiles(envPage)

	if err != nil {
		fmt.Println(`ocorreu um erro fatal em sua pagina, o arquivo não foi encontrado.`)

		fmt.Fprintf(w, `[ERRO], Não foi encontrado uma rota para a pagina selecionada, va para o 
		 youtube e se divirtar`)

		return
	}
	data := map[string]any{
		"title": "minha pagina",
		"subT":  "minha pagina toda em go",
	}
	w.WriteHeader(http.StatusOK)

	tpl.Execute(w, data)
}

func index(w http.ResponseWriter, r *http.Request) {

	index := "pages/Index.html"
	user := "user 1"
	utcTimeLoc := time.FixedZone("UTC", -3)
	t := time.Now()
	seet := t.In(utcTimeLoc).Format(http.TimeFormat)
	tpl, err := template.ParseFiles(index)

	if err != nil {
		fmt.Println("ERRO 404, NÃO FOI ENCONTRADO PAGINA")

		fmt.Fprintf(w, "404 SEM ROTA")

		return
	}
	strVar := map[string]any{
		"user": user,
		"data": seet,
	}

	w.WriteHeader(http.StatusOK)

	tpl.Execute(w, strVar)
}
func styles(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/css/styles.css")
}
func main() {

	http.HandleFunc("/env", envPage)
	http.HandleFunc("/", index)
	http.HandleFunc("/style", styles)
	folderSystem := http.FileServer(http.Dir("./pages/"))

	pathRequest := http.StripPrefix("/pages/", folderSystem)

	http.Handle("/pages/", pathRequest)
	http.Handle("/src/backend/", http.StripPrefix("/src/backend/", http.FileServer(http.Dir("./src/backend/"))))
	http.ListenAndServe(":8080", nil)

}
