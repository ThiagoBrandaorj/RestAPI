package pedidos

import (
	"api/dados"
	"api/utils"
	"encoding/json"
	"net/http"
)

func Adicionar(w http.ResponseWriter, r *http.Request) {
	var novoPedido dados.Pedido

	err := json.NewDecoder(r.Body).Decode(&novoPedido)
	if err != nil {
		http.Error(w, "Erro ao decodificar dados", http.StatusBadRequest)
		return
	}

	err = dados.FilaPedidos.Adicionar(novoPedido.Delivery, novoPedido.NomeProdutos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Sucesso"))
	utils.LogMessage("Pedido adicionado com sucesso")
}

func Listar(w http.ResponseWriter, r *http.Request) {
	sortParam := r.URL.Query().Get("sort")

	pedidos := dados.FilaPedidos.Listar()

	switch sortParam {
	case "bubblesort":
		BubbleSort(pedidos)
	case "quicksort":
		QuickSort(pedidos, 0, len(pedidos)-1)
	
	case "mergesort":
		MergeSort(pedidos)
	}
	
	filaJSON, err := json.Marshal(pedidos)
	if err != nil {
		http.Error(w, "Erro ao converter mensagens para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(filaJSON)
	utils.LogMessage("Fila de pedidos obtida")
}

func BubbleSort(pedidos []dados.Pedido) {
	n := len(pedidos)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if pedidos[j].ValorTotal > pedidos[j+1].ValorTotal {
				pedidos[j], pedidos[j+1] = pedidos[j+1], pedidos[j]
			}
		}
	}
}

func QuickSort(pedidos []dados.Pedido, inf, sup int) {
	if inf < sup {
		p := partirPorValor(pedidos, inf, sup)
		QuickSort(pedidos, inf, p-1)
		QuickSort(pedidos, p+1, sup)
	}
}

func partirPorValor(pedidos []dados.Pedido, inf, sup int) int {
	pivo := pedidos[sup].ValorTotal
	i := inf - 1
	for j := inf; j < sup; j++ {
		if pedidos[j].ValorTotal < pivo {
			i++
			pedidos[i], pedidos[j] = pedidos[j], pedidos[i]
		}
	}
	pedidos[i+1], pedidos[sup] = pedidos[sup], pedidos[i+1]
	return i + 1
}

func MergeSort(pedidos []dados.Pedido) []dados.Pedido {
	if len(pedidos) <= 1 {
		return pedidos
	}
	metade := len(pedidos) / 2
	esq := MergeSort(pedidos[:metade])
	dir := MergeSort(pedidos[metade:])
	return merge(esq, dir)
}

func merge(esq, dir []dados.Pedido) []dados.Pedido {
	res := make([]dados.Pedido, 0, len(esq)+len(dir))
	i, j := 0, 0
	for i < len(esq) && j < len(dir) {
		if esq[i].ValorTotal < dir[j].ValorTotal {
			res = append(res, esq[i])
			i++
		} else {
			res = append(res, dir[j])
			j++
		}
	}
	res = append(res, esq[i:]...)
	res = append(res, dir[j:]...)
	return res
}
