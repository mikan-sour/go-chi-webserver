# go-chi-webserver

# Getting used to using Chi framework

- https://github.com/go-chi/chi

## Running Docker Postgres

- command to be used:

```docker run --rm --name pg-testing -e POSTGRES_PASSWORD=password -d -p 5433:5432 postgres```
* must have a table called `pokemon` in it.

# To Do

- [ ] (5/31) pokemon service should return the response struct and not the pokemon struct
- [ ] (5/31) figure out how to get an error from the GORM and return w. provider
- [ ] (5/31) pokemon unit tests