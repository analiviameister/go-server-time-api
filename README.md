# Go Server Time API

API simples em Go que retorna o horário do servidor com cache de 1 minuto utilizando Redis.

---

## Funcionalidades

- Retorna o horário atual do servidor
- Utiliza Redis para cache (TTL de 1 minuto)
- Indica se a resposta veio do cache
- Exibe resposta diretamente no navegador (texto simples)

---

## Endpoints

| Endpoint | Descrição |
|--------|----------|
| `/` | Retorna o horário do servidor |

---

## Como funciona

1. A API verifica se existe um valor no Redis (`server_time`)
2. Se existir:
   - Retorna o valor em cache
3. Se não existir:
   - Gera o horário atual
   - Armazena no Redis por 1 minuto
   - Retorna o novo valor

---

## Variáveis de ambiente

| Variável | Descrição | Default |
|--------|----------|--------|
| `PORT` | Porta da aplicação | `8080` |
| `REDIS_ADDR` | Endereço do Redis | `localhost:6379` |
| `REDIS_PASSWORD` | Senha do Redis | vazio |

---

## Rodando localmente

```bash
go mod tidy
go run main.go
```

## Rodando com docker

```bash
docker build -t go-time-api .