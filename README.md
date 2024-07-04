# RESTAPI

## Funcionalidades do Sistema
### Adicionar um Novo Produto:
Um usuário pode adicionar um novo produto na árvore de produtos enviando uma requisição HTTP POST utilizando a rota /produto com o nome, descrição e valor do produto no corpo da requisição em formato JSON.
* POST /produto\
  Por exemplo:
{
    "nome": "Camiseta",
    "descricao": "Camiseta branca de algodão",
    "valor": 25.99
}

### Remover um Produto Existente:
Para remover um produto existente da árvore de produtos, o usuário pode enviar uma requisição HTTP DELETE utilizando a rota /produto, selecionando-o pelo nome do produto na forma chave e valor na requisição.
* DELETE /produto\
Por exemplo: http://localhost:8080/produto?nome=camiseta

### Buscar um Produto Existente:
Um usuário pode buscar um produto específico da árvore de produtos enviando uma requisição HTTP GET na rota /produto, pelo nome do produto na forma chave e valor na requisição.
* GET /produto\
Por exemplo: http://localhost:8080/produto?nome=camiseta

### Listar Todos os Produtos:
Para listar todos os produtos disponíveis da árvore de produtos, o usuário pode enviar uma requisição HTTP GET utilizando a rota /produtos , a árvore de produtos irá retornar da maneira em que está ordenada. Ou seja, os produtos mais a esquerda da árvore vem primeiro. 
* GET /produtos

### Adicionar um Pedido :
Para adicionar um pedido na fila de pedidos, o usuário pode enviar uma requisiçao HTTP
POST para pedido/ com o valor booleano de delivery e nome_produtos(um array de strings com o nome dos produtos) no corpo da requisição no formato JSON.
* POST /pedido\
Por exemplo:
{
    "delivery": true,
    "nome_produtos": ["camiseta"]
    "valor_total": 35.99
}

### Exibir Pedidos em Aberto:
Um usuário pode obter uma lista de todos os pedidos em aberto enviando uma requisição HTTP GET para a rota /pedidos. O usuário pode ainda fornecer um parâmetro que irá definir a ordenação do retorno da lista de pedidos pelo valor, sendo 3 possíveis : bubblesort, quicksort e mergesort. Caso o usuário não passe nenhum parâmetro, o retorno da lista será no formato original em que os pedidos foram adicionados, ou seja, em formato de fila.
* GET /pedidos\
  Por exemplo: http://localhost:8080/pedidos?sort=bubblesort

### Obter Métricas do Sistema:
Para obter as métricas do sistema que são: total de produtos cadastrados, número de pedidos encerrados, número de pedidos em andamento, faturamento total, ticket médio e tempo de funcionamento, o usuário pode enviar uma requisição HTTP GET utilizando a rota /metricas. 
* GET /metricas

### Abrir ou Fechar a Loja:
O usuário pode abrir ou fechar a loja enviando uma requisição HTTP POST para /abrir ou /fechar, respectivamente. O usuário pode ainda quando abir a loja , definir o tempo em que cada pedido levará para ser expedido por meio de um parametro, caso o usuario não passe nenhum parametro o tempo de expedição de cada pedido será o de 30 segundos.
* POST /abrir\
  Por exemplo: http://localhost:8080/abrir?intervalo=10

* POST /fechar

Esses são exemplos de como um usuário pode interagir com a API REST desenvolvida em golang, por meio de requisições HTTP fornecidas pelo sistema. Cada requisição deve incluir as informações necessárias no corpo da requisição ou na URL, dependendo da operação desejada.
