# Pokédex API

A Pokédex in the form of a REST API that returns Pokémon information.

The API has two main endpoints:
1. Return basic Pokémon information.
2. Return basic Pokémon information but with a 'fun' translation of the Pokémon Description.


### Run services locally

Clone the repository

```bash
git clone https://github.com/imogen-k/Pokedex.git
```

Run service on port 8080
```bash
go run cmd/main.go
```

### Visit in browser

localhost:8080/pokemon/{name}

e.g. http://localhost:8080/pokemon/metapod


localhost:8080/pokemon/translated/{name}

e.g. http://localhost:8080/pokemon/translated/metapod


### cURL command to call services

```bash 
curl http://localhost:8080/pokemon/charmander
```
```bash 
curl http://localhost:8080/pokemon/translated/charmander
```

### Build docker image & run locally

Build docker image

```bash
docker build . -t pokedex
```

Run service on port 8080

```bash
docker run -p 8080:8080 pokedex
```

Visit in browser or via cURL command

http://localhost:8080/pokemon/metapod

```bash 
curl http://localhost:8080/pokemon/charmander
```

### Tech stack
- Golang
- Chi
- Testify

### Resources
- https://pokeapi.co/
- https://funtranslations.com/api/yoda
- https://funtranslations.com/api/#shakespeare
