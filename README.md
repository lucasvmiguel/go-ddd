# Go-DDD

- Entities: Mutable Identifiable Structs.
- Repository: A implementation of storing aggregates or other information
- Factory: A constructor to create complex objects and make creating new instance easier for the developers of other domains
- Service: A collection of repositories and sub-services that builds together the business flow

Reference: https://programmingpercy.tech/blog/how-to-domain-driven-design-ddd-golang/

### Migrations

#### Create migration

```
 migrate create -ext sql -dir db/migrations -seq create_users_table
```

#### Run migration

```
go run cmd/migrate-up/migrate-up.go
```

#### Rollback migration

```
go run cmd/migrate-down/migrate-down.go
```

Reference: https://github.com/golang-migrate/migrate
