## Iniciar modulo

Para iniciar o modulo go no projeto, rode o comando:

```bash
go mod tidy
```

## Subir Postgres

```bash
docker run --name pg -e POSTGRES_PASSWORD=pass -e POSTGRES_USER=user \
  -e POSTGRES_DB=myapp -p 5432:5432 -d postgres:18
```

## Criar tabela

```sql
CREATE TABLE users (
    id         text PRIMARY KEY,
    name       text NOT NULL,
    email      text NOT NULL,
    created_at timestamptz NOT NULL
);
```

## Exportar `DATABASE_URL`

```bash
export DATABASE_URL="postgres://user:pass@localhost:5432/postgres?sslmode=disable"
```

## Rodar o projeto

```bash
go run ./cmd/api
```