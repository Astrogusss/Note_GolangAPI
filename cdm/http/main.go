package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func NoteList (w http.ResponseWriter, r *http.Request){
	//se fizer alguma rota inesperada
	if r.URL.Path != "/"{
		//jeito mais facil do que http.error
		http.NotFound(w, r)
		return
	}
	//se quiser algo diferente de GET
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	//listando os templates para serem retornados
	files := []string{"views/templates/base.html", "views/templates/pages/NoteList.html"} 

	t, err := template.ParseFiles(files ...)
	if err != nil {
		http.Error(w, "Erro no servidor", http.StatusInternalServerError)
		return
	}

	t.ExecuteTemplate(w, "base", nil)
}

func NoteView (w http.ResponseWriter, r *http.Request){
	//temos que vizualizar qual o id da nota que queremos ver
	id := r.URL.Query().Get("id")
	nome := r.URL.Query().Get("Name")

	//caso nao for passado a id na url
	if id == ""{
		http.Error(w, "Nota não encontrada", http.StatusNotFound)
		return
	}
	if nome == ""{
		http.Error(w, "Nome não encontrada", http.StatusNotFound)
		return
	}

	//criacao de struct para mandar para o template
	type View struct {
		Id string
		Nome string
	}
	aux := View{Id: id, Nome: nome}

	//se não for o verbo http esperada
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	files := []string{"views/templates/base.html", "views/templates/pages/NoteView.html"}

	t, err := template.ParseFiles(files ...)
	if err != nil {
		http.Error(w, "Erro no servidor", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", aux)
}

func NoteNew (w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	
	files := []string{"views/templates/base.html", "views/templates/pages/NoteNew.html"}

	t, err := template.ParseFiles(files ...)
	if err != nil {
		http.Error(w, "Erro no servidor", http.StatusInternalServerError)
		return
	}

	t.ExecuteTemplate(w, "base", nil)
}

func NoteCreate (w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	files := []string{"views/templates/base.html", "views/templates/pages/NoteCreate.html"}

	t, err := template.ParseFiles(files ...)
	if err != nil {
		http.Error(w, "Erro no servidor", http.StatusInternalServerError)
		return
	}

	t.ExecuteTemplate(w, "base", nil)

}

func main(){
	//carregar as variaveis globais
	config := LoadConfig()

	mux := http.NewServeMux()
	//precisamos de um handler que sirva arquivos estáticos, assim como fazemos com html (pode ser feito com js, css, etc)
	// aqui estamos dizendo que todos os arquivos que estao em views/static serão servidos
	staticHandler := http.FileServer(http.Dir("views/static/"))

	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))
	//cuidado que tem que link o css tambem no html

	mux.HandleFunc("/", NoteList)
	mux.HandleFunc("/note/view", NoteView)
	mux.HandleFunc("/note/new", NoteNew)
	mux.HandleFunc("/note/create", NoteCreate)
	
	http.ListenAndServe(fmt.Sprintf(":%s", config.Server_Port), mux)
}