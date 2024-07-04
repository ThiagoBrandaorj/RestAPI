package dados

import (
	"api/utils"
	"fmt"
	"math"
)

const VALOR_DELIVERY = 10.0

type Pedidos struct {
	Fila []Pedido
}

func (p *Pedidos) inicializar() {
	p.Fila = make([]Pedido, 0)
}

func (p *Pedidos) Adicionar(delivery bool, nomeProdutos []string) error {
	valorTotal := 0.0
	if delivery { valorTotal += VALOR_DELIVERY }

	for _, nome := range nomeProdutos {
		prod, err := ListaProdutos.Buscar(nome)
		if err != nil { return err }

		valorTotal += prod.Valor
	}
	valorTotal = math.Round(valorTotal*100)/100

	if valorTotal > 100{
		valorTotal = valorTotal * 0.9
		valorTotal = math.Round(valorTotal*100)/100
	}

	pedido := Pedido{Delivery: delivery, NomeProdutos: nomeProdutos, ValorTotal: valorTotal}
	pedido.RegistrarID()

	p.Fila = append(p.Fila, pedido)
	MetricasColetadas.PedidosEmAndamento++
	return nil
}

func (p *Pedidos) Expedir() error {
	if len(p.Fila) == 0 { return fmt.Errorf("fila de pedidos está vazia") }

	MetricasColetadas.PedidosEmAndamento--
	MetricasColetadas.PedidosEncerrados++
	MetricasColetadas.FaturamentoTotal += p.Fila[0].ValorTotal

	utils.LogMessage(fmt.Sprintf("Expedido pedido %d, valor R$ %.2f", p.Fila[0].Id, p.Fila[0].ValorTotal))

	p.Fila = p.Fila[1:]

	return nil
}

func (p *Pedidos) Listar() []Pedido {
	return p.Fila
}