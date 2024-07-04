package dados

import (
	"fmt"
	"math"
	"time"
)

var ListaProdutos Produtos
var FilaPedidos Pedidos
var totalIdsProdutos, totalIdsPedidos = 1, 1
var lojaAberta = false
var MetricasColetadas Metricas

var StartTime = time.Now()

func AtualizarMetricas(){
	// Calcular Ticket Medio
	if MetricasColetadas.PedidosEncerrados > 0 {
		MetricasColetadas.TicketMedio = MetricasColetadas.FaturamentoTotal / float64(MetricasColetadas.PedidosEncerrados)
		MetricasColetadas.TicketMedio = math.Round(MetricasColetadas.TicketMedio * 100)/ 100
	} else {
		MetricasColetadas.TicketMedio = 0
	}
	// Calcular Tempo de Funcionamento
    MetricasColetadas.TempoFuncionamento = int64(time.Since(StartTime).Seconds())
}

func InicializarDados() {
	ListaProdutos = Produtos{}
	ListaProdutos.inicializar()

	FilaPedidos = Pedidos{}
	FilaPedidos.inicializar()

	MetricasColetadas = Metricas{
		TotalProdutos:      0,
		PedidosEncerrados:  0,
		PedidosEmAndamento: 0,
		FaturamentoTotal:   0.0,
		TicketMedio:		0.0,
		TempoFuncionamento: 0,
	}
}

func AbrirLoja() error {
	if lojaAberta {
		return fmt.Errorf("loja já está aberta")
	}

	lojaAberta = true
	return nil
}

func FecharLoja() error {
	if !lojaAberta {
		return fmt.Errorf("loja já está fechada")
	}

	lojaAberta = false
	return nil
}

func LojaEstaAberta() bool {
	return lojaAberta
}