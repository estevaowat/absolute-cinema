# absolute-cinema

this project is my answer to this challenge:

- [populate database with thousands of movies](https://app.devgym.com.br/challenges/ec36e7e2-6a2d-4406-98e1-3029f843b5c3)


## challenge
Nesse desafio você deve criar um programa de linha de comando (cli) que lê um arquivo csv de filmes e popula um banco de dados pensando em performance e esperando que o arquivo pode crescer muito.

Cada linha do csv contém colunas que devem ser salvas em colunas/campos separados no banco de dados:
ID - número inteiro que identifica o filme encontrado na 1a coluna do csv
title - título do filme encontrado na 2a coluna do csv. Com o valor "Jumanji (1995)" o title é Jumanji
Year - ano do filme encontrado na 2a coluna do csv. Com o valor "Jumanji (1995)" o year é 1995
Genres - múltiplos valores com os gêneros do filme separados por |. Encontrado na 3a coluna.

O script deve pensar em performance e tirar proveito de concorrência/paralelismo para popular o banco de dados.


## What I am learning with this project?
- Golang (http, json, tests and dabatase packages)
- Concurrency/paralellism
- Postgres


## How to use this cli?

**Pre-requisites**
- Golang
- Kotlin
- podman (I used podman because it's free!)



1- First you should need this microservice running into your machine to create thousands of movies
(ADD_KOTLIN_MICROSERVICE_HERE)

2- `git clone` this repo, go to repo folder and run `go mod tidy` to download dependencies packages

3- Run `go run main.go` to check if the project is running smooth


## Here is my todo list

### Feature 1: Generate a csv with movies (sequential)
- [x] Connect to an API get movies infos
- [x] Parse movies to csv format (123,spiderman(2002),comedy|action|superhero)
- [x] Create csv file with a random name
- [x] Add movies to csv file
- [x] when finish adding movies, print filename created with movies
- [x] study how to try catch http.Get
- [x] measure time and write down Here

1000 movies ~= 760ms <br/>
10_000 movies ~= 1.67s  <br/>
100_000 movies ~= 15.63s <br/>
1_000_000 movies ~= 4m35s <br/>



#### Feature 1.1 (Bonus): Generate a csv file with movies using go routines
- [x] add movies to csv file using go routines
- [x] measue time and write down here


1000 movies ~= 35.383ms <br/>
10_000 movies ~= 344.383ms  <br/>
100_000 movies ~= 5.42s <br/>
1_000_000 movies ~= 37.517s <br/>


### Prerequisites to Feature 2, 3 and 4
- [x] learn how golang works with databases
- [x] learn how to save in database one by one
- [x] measure time to save one by one
1_000_000 movies took more than 20min to save
- [x] learn how to sabe in database using chunks
1_000_000 movies took aprox. 25s


### Feature 2: Save csv with movies in a database (sequential)
- [x] Read all csv file
- [x] Loop all movies in the csv file
- [x] Save one by one in database
- [x] Measure time elapsed

### Feature 2.1: Save csv with movies in a database using chunks
- [x] read csv file
- [x] loop through all mvoies in csv file
- [x] save in chunks of 10_000 movies in database
- [x] measure time elapsed

1_000_000 movies took aprox. 25s


### Feature 3: Save csv with movies in database (goroutines, one by one)
- [ ] Read csv file
- [ ] Loop through movies in csv file
- [ ] Using goroutines parse line and save in database one by one
- [ ] Measure time elapsed

### Feature 4: Save csv with movies in database(goroutines, chunks by 100 of size)
- [ ] Read csv file
- [ ] Loop through movies in csv file
- [ ] Using goroutines get chunks by 100 of size and save in database
- [ ] Measure time elapsed

#### Premises
- All uuids to save in database are different

