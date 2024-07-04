package utils

import (
	"fmt"
	"time"

)

func LogMessage(mensagem string){
	tempoAtual := time.Now().Format("<02/01/2006 15:04:05>")
	fmt.Printf("%s %s\n", tempoAtual, mensagem)
}