package produtos

import (
	"api/dados"
	"api/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func Adicionar(w http.ResponseWriter, r *http.Request) {
	var novoProduto dados.Produto

	err := json.NewDecoder(r.Body).Decode(&novoProduto)
	if err != nil {
		http.Error(w, "Erro ao decodificar dados", http.StatusBadRequest)
		return
	}

	err = dados.ListaProdutos.Adicionar(novoProduto.Nome, novoProduto.Descricao, novoProduto.Valor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Sucesso"))
	utils.LogMessage(fmt.Sprintf("Produto %s adicionado com sucesso", novoProduto.Nome))
}

func Remover(w http.ResponseWriter, r *http.Request) {
	nome := r.FormValue("nome")

	err := dados.ListaProdutos.Remover(nome)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Sucesso"))
	utils.LogMessage(fmt.Sprintf("Produto com nome %s removido com sucesso", nome))
}

func Buscar(w http.ResponseWriter, r *http.Request) {
	nome := r.FormValue("nome")

	prod, err := dados.ListaProdutos.Buscar(nome)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prodJSON, err := json.Marshal(prod)
	if err != nil {
		http.Error(w, "Erro ao converter mensagens para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(prodJSON)
	utils.LogMessage(fmt.Sprintf("Produto com nome %s localizado", nome))
}

func Listar(w http.ResponseWriter, r *http.Request) {
	listaJSON, err := json.Marshal(dados.ListaProdutos.Listar())
	if err != nil {
		http.Error(w, "Erro ao converter mensagens para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(listaJSON)
	utils.LogMessage("Lista de produtos obtida")
}