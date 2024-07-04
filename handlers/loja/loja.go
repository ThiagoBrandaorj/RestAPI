package loja

import (
	"api/dados"
	"api/processamento"
	"api/utils"
	"net/http"
	"strconv"
	"time"
)

func Abrir(w http.ResponseWriter, r *http.Request) {
	err := dados.AbrirLoja()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar se o parâmetro 'intervalo' está presente
	intervaloStr := r.URL.Query().Get("intervalo")
	if intervaloStr != "" {
		intervalo, err := strconv.Atoi(intervaloStr)
		if err != nil || intervalo <= 0 {
			http.Error(w, "Intervalo inválido", http.StatusBadRequest)
			return
		}
		go processamento.ProcessaPedidos(time.Duration(intervalo) * time.Second)
	} else {
		go processamento.ProcessaPedidos(30 * time.Second)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Loja aberta"))
	utils.LogMessage("Loja Aberta")
}

func Fechar(w http.ResponseWriter, r *http.Request) {
	err := dados.FecharLoja()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Loja fechada"))
	utils.LogMessage("Loja Fechada")
}