package metricas

import (
	"api/dados"
	"api/utils"
	"encoding/json"
	"net/http"
)

func Metricas(w http.ResponseWriter, r *http.Request) {
	 // Atualiza as métricas antes de convertê-las para JSON
	 dados.AtualizarMetricas()

	metricasJSON, err := json.Marshal(dados.MetricasColetadas)
	if err != nil {
		http.Error(w, "Erro ao converter mensagens para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(metricasJSON)
	utils.LogMessage("Métricas coletadas com sucesso")
}