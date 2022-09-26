# Falemais Telzir API#

É uma API que calcula e compara os valores de uma ligação telefônica entre regiões diferentes 
com base em planos de tarifas pré-definidos.

## Table of contents

- [Visão Geral do Projeto](#general-view)
    - [O Desafio](#o-desafio)
- [O Desenvolvimento](#the-development-process)
    - [Skills Usadas](#tools-used)
    - [Lessons learned](#lessons-learned)
- [Uso](#usage)
- [Author](#author)

## General view

### O desafio

É uma API que calcula e compara os valores de uma ligação telefônica entre regiões diferentes 
com base em planos de tarifas pré-definidos.

**O usuário será capaz de**

  - endpoint /api/v1/dddcost/swagger/index.html
    * visualizar a documentação da API

  - endpoint GET /api/v1/dddcost
    * Lista todos os códigos de DDDs cadastrados com DDD de origem e DDD de destino
    
   
  Exemplo de retorno:
   ```json
   [
        {
            "origin": 11,
            "destination": 16
        },
        {
            "origin": 16,
            "destination": 11
        },
        {
            "origin": 11,
            "destination": 17
        },
        {
            "origin": 17,
            "destination": 11
        },
        {
            "origin": 11,
            "destination": 18
        },
        {
            "origin": 18,
            "destination": 11
        }
   ]
```
  - endpoint POST /api/v1/dddcost
   * Retorna o calculo do valor da ligação com base nos planos de tarifas cadastrados no sistema, 
     ddd de origem, ddd de destino e tempo de duração da ligação
   * O endpoint retorna o Status 200 com o os custos da ligação com e sem o plano de tarifa

  Exemplo de requisição:

    ```json
    {
        "origin": 11,
        "destination": 18,
        "minutes": 200,
    }
    ```
  - Validações do endpoint
  * O campo origin refere ao DDD de origem, é obrigatório e deve ser maior que 9
  * O campo destination refere ao DDD de destino,é obrigatório e deve ser maior que 9
  * O campo minutes refere ao tempo de duração da ligação em minutos, é obrigatório e deve ser maior que 0
  
- Caso a validação do campo origin falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
	 "message": "origin must be greater than 10"
   }
  ```
- Caso a validação do campo destination falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
	  "message": "destination must be greater than 10"
   }
  ```
- Caso a validação do campo minutes falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
	 "message": "minutes must be greater than 0"
   }
  ```

## O Desenvolvimento

Foi utilizado o método TDD, para o desenvolvimento. Foi escolhido chi como router por
este possuir o melhor benchmark de performance, e o banco de dados escolhido foi o PostgreSQL. Todo o backend está dockerizado.
Para documentação foi utilizado o Swagger.

### Ferramentas usadas

- [Golang](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Swagger](https://swagger.io/)
- [Chi](https://github.com/go-chi/chi)
- [Docker](https://www.docker.com/)

## Usage


- Utilize o comando abaixo para subir os containers do Projeto em modo de desenvolvimento
```bash
docker-compose -f docker-compose.dev.yml up -d --build
```
- exitem dois bancos, um para uso continuo e outro somente para testes,
- para que não haja conflito entre os dados de teste e os dados de uso continuo.
- Por padrão ele vem setado para o banco de teste, para mudar deve renomear o .env para .env.test e renomear o .env.dev para .env

- Utilize o comando abaixo para subir os containers do Projeto em modo de testes
```bash
docker-compose -f docker-compose.test.yml up -d --build
```

- Para rodar o backend de modo local use o comando abaixo:

```bash
make run
```
- Para rodar o backend no container:

- Para acessar o container do backend use o comando abaixo:

```bash
docker exec -it telzir_backend bash
```

- Para rodar o backend de modo
- 
```bash
make run
```
- Para rodar o backend de modo  use o comando abaixo:

## Test

* O Banco de testes deve estar ativo e as variaveis de ambientes de testes também, senão nao rodará todos os testes somente os mockados.
    - Nos testes dos handlers, o o banco de dados é populado no inico de cada teste e depois é feito um drop.

- Para conseguir rodar todos os testes deve ser acessado o container do backend:

- Para acessar o container de tests do backend use o comando abaixo:

```bash
docker exec -it telzir_backend_test bash
```
- Para rodar os testes use o comando abaixo:

```bash
     make test
```

## Author

- LinkedIn - [Alberto Ouverney Paz](https://www.linkedin.com/in/beto-ouverney-paz/)
