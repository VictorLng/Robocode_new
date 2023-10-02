package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func envPage(w http.ResponseWriter, r *http.Request) {

	envPage := "src/pagina de envios.html"

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

	index := "src/Index.html"
	user := "user 1"
	utcTimeLoc := time.FixedZone("UTC", -3)
	t := time.Now()
	s := t.In(utcTimeLoc).Format(http.TimeFormat)
	tpl, err := template.ParseFiles(index)

	if err != nil {
		fmt.Println("ERRO 404, NÃO FOI ENCONTRADO PAGINA")

		fmt.Fprintf(w, "404 SEM ROTA")

		return
	}
	strVar := map[string]any{
		"user": user,
		"data": s,
	}

	w.WriteHeader(http.StatusOK)

	tpl.Execute(w, strVar)
}

func main() {

	http.HandleFunc("/env", envPage)
	http.HandleFunc("/", index)

	folderSystem := http.FileServer(http.Dir("./src/"))

	pathRequest := http.StripPrefix("/src/", folderSystem)

	http.Handle("/src/", pathRequest)
	http.ListenAndServe(":8080", nil)

}