package dados

import "fmt"

type No struct{
	produto Produto
	esq *No
	dir *No
}

type Produtos struct {
	raiz *No
}

func (p *Produtos) inicializar() {
	p.raiz = nil
}
//funcionando...
func (p *Produtos) Adicionar(nome, descricao string, valor float64) error {
	if nome == "" || descricao == "" { return fmt.Errorf("nome ou descrição vazios") }
	if valor <= 0 { return fmt.Errorf("valor do produto inválido") }

	prod := Produto{
		Nome: nome,
		Descricao: descricao,
		Valor: valor,
	}

	prod.RegistrarID()
	p.raiz = p.adicionarRec(p.raiz, prod)
	MetricasColetadas.TotalProdutos++
	return nil
}

func (p *Produtos)adicionarRec(no *No, prod Produto) *No{
	if no == nil{
		return &No{produto: prod}
	}

	if prod.Nome < no.produto.Nome{
		no.esq = p.adicionarRec(no.esq, prod)
	} else{
		no.dir = p.adicionarRec(no.dir, prod)
	}
	return no
}

func (p *Produtos) Remover(nome string) error {
	var found bool
	p.raiz, found = p.removerRec(p.raiz, nome)
	if !found{
		return fmt.Errorf("nome não encontrado na lista de produtos")
	}
	MetricasColetadas.TotalProdutos--
	return nil
}

func (p *Produtos) removerRec(no *No, nome string)(*No, bool){
	if no == nil{
		return nil,false
	}

	if nome < no.produto.Nome{
		no.esq,_ = p.removerRec(no.esq, nome)
	} else if nome > no.produto.Nome{
		no.dir, _ = p.removerRec(no.dir, nome)
	} else{
		if no.esq == nil{
			return no.dir, true
		} else if no.dir == nil{
			return no.esq, true
		}

		no.produto = p.minValue(no.dir)
		no.dir, _ = p.removerRec(no.dir,  no.produto.Nome)
		return no , true
	}
	return no , true
}

func (p *Produtos) minValue(no *No) Produto{
	current := no
	for current.esq != nil{
		current = current.esq
	}
	return current.produto
}

func (p *Produtos) Buscar(nome string) (Produto, error) {
	no := p.buscarRec(p.raiz, nome)
	if no == nil{
		return Produto{}, fmt.Errorf("nome não encontrado na lista de produtos")
	}
	return no.produto , nil
}

func (p *Produtos) buscarRec(no *No , nome string) *No{
	if no == nil || no.produto.Nome == nome{
		return no
	}

	if nome < no.produto.Nome{
		return p.buscarRec(no.esq, nome)
	}
	return p.buscarRec(no.dir, nome)
}
//funcionando...
func (p *Produtos) Listar() []Produto {
	var produtos []Produto
	p.listarRec(p.raiz, &produtos)
	return produtos
}

func (p *Produtos) listarRec(no *No, produtos *[]Produto){
	if no != nil{
		p.listarRec(no.esq, produtos)
		*produtos = append(*produtos, no.produto)
		p.listarRec(no.dir,produtos)
	}
}