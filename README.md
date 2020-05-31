# api-transaction
Criar e consultar contas. Transacionar com as contas criadas e apresentar as transações efetuadas em um determinado período

## Pré-requisitos:
* Clone ou download [api-transaction](xxxxxx).
* Go instalado.
* Code Editor.
* Postman para realizar requisições na API.
* Criar o banco de dados, o nome o banco deve ser: bd_transactions.
* Executar o script sql que consta na pasta do projeto em api-transaction/dbConnection/scripts.



## Por onde eu começo?
Crie uma conta! O /accounts  é o endpoint responsável por criar uma conta.

- Informe a URL da requisição: http://localhost:8080/accounts
- Informe o método da requisição: POST
- Deverá ser enviado no body o conteúdo a seguir, respeitando o formato e o tipo Json:
	{
		"document_number": {número da conta}
    }

Exemplo requisição:

```
http://localhost:8080/accounts
```

## Resultado esperado:
* Caso o retorno da requisição seja 200, será apresentada uma mensagem informando que houve sucesso na requisição e será sinalizado o ID da conta cadastrada.



## Como consulto as informação de uma conta?
Após criar a conta, você poderá consultá-la informando o seu ID.
O /accounts/{accountId}  é o endpoint responsável por apresentar as informações de uma conta..

- Informe a URL da requisição: http://localhost:8080/accounts/{accountId}
- Informe o método da requisição: GET
- Deverá ser preenchido no “{accountId}” o ID da conta que se deseja obter informações.

Exemplo requisição:

```
http://localhost:8080/accounts/1
```

## Resultado esperado:
* Caso o retorno da requisição seja 200, serão apresentadas todas as informações referente ao ID da conta informada na URL. 



## Como incluo uma transação?
Após criar a conta, você poderá transacionar.
O /transactions é o endpoint responsável por cadastrar no banco de dados uma transação

- Informe o método da requisição: POST
- Informe a URL da requisição: http://localhost:8080/transactions
- Deverá ser enviado no body o conteúdo a seguir, respeitando o formato e o tipo Json:
	{
		"account_id": {id atrelado a uma conta},
		"operation_type_id": {tipo da operação},
		"amount": {valor transacionado}
    }

Exemplo requisição:

```
http://localhost:8080/transactions
```

## Resultado esperado:
* Caso o retorno da requisição seja 200, será apresentada uma mensagem informando que houve sucesso na requisição e será sinalizado o ID da transação cadastrada.



## O que é operation_type_id?
É o tipo da transação efetuada, a saber:

* 1 = (COMPRA A VISTA)
* 2 = (COMPRA PARCELADA)
* 3 = (SAQUE)
* 4 = (PAGAMENTO)



## Como consulto as informação de uma transação?
O /transactions é o endpoint responsável por apresentar as transações que uma conta efetuou em um determinado período.

- Informe a URL da requisição: http://localhost:8080/transactions?accountid={accountId}&initialdate={dataDeReferênciaInicial}&enddate={dataDeReferênciaFinal}
- Informe o método da requisição: GET 
- Deverá ser preenchido em “{accountId}” na URL o ID da conta que se deseja obter as informações sobre o que foi transacionado.
- Deverá ser preenchido em “{dataDeReferênciaInicial}” e “{dataDeReferênciaFinal}” o período em que houve transação para a conta “{accountId}”.

Exemplo requisição:

```
http://localhost:8080/transactions?accountid=1&initialdate=2020-05-30&enddate=2020-05-31
```

## Resultado esperado:
* Caso o retorno da requisição seja 200, serão apresentadas todas as transações relacionada à conta ({accountId}) no período (“{dataDeReferênciaInicial}” e “{dataDeReferênciaFinal}”). Mais detalhes:

- Conta que possua somente uma transação no período informado, então somente um registro será apresentado;
- Conta que possua somente uma transação no período informado, então somente um registro será apresentado;
- Conta que possua mais de uma transação no período informado, então uma lista de registros serão apresentados;
- O campo transaction_id: indica o ID da transação
- O campo account_id: indica o ID da conta
- O campo operation_type_id: indica o tipo da operação, a saber
    caso o operation_type_id seja 1, 2, 3: o amount será negativo 
	caso o operation_type_id seja 4: o amount será positivo
- O campo amount: indica o valor transacionado
- O campo event_date: é a data em que houve transação para a conta no período indicado.
