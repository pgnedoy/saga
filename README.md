# saga

DB_URL:

```postgres://{user}:{password}@{host | "localhost"}:5432/{database}?sslmode=disable```

To create migration: 

```migrate create -ext sql -dir db/migrations -seq create_users_table```