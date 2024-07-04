package processamento

import (
	"api/dados"
	"api/utils"
	"time"
)

var intervaloExpedicao time.Duration = 30 * time.Second

func ProcessaPedidos(intervalo time.Duration) {
	intervaloExpedicao = intervalo
	for {
		if dados.LojaEstaAberta() {
			time.Sleep(intervaloExpedicao)
			dados.FilaPedidos.Expedir()
			utils.LogMessage("Pedidos expedidos")
		}
	}
}