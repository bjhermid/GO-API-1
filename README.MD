## GO-API-1 
### USING JWT, MYSQL and LEARN to TESTING

### FEATURE
1. GET /user/{id}
2. POST / register
3. POST / login
4. GET / PRODUCT
5. POST / cart/cheackout

### to ------>
API SERVER | DOCKER

### to ------>
MYSQL Database


## FOLDER STRUCTURE in Go
cmd/ --> handler : Contains application entry points. Each subdirectory within cmd represents a distinct executable, typically with a main.go file.

-bin (the build)
-cmd (the machine)
    |-api
    |-migrate
-config 
    |-server
    |-database
-db
-services 
    |-auth
    |-cart
    |-order
    |-product
-types
-utils

## FRAMEWORK
- JWT,mux, godotenv, mysql (go-sql-driver), crypto
check it at go.mod file at root folder

## RUN
1. clone the repo
2. 
```bash
## run build app
make run 
## migrate table to db
make migrate-up

```
3. check api in http://localhost:addr/api/v1 as base url and
    - /login 
    - /register 
    - /products
    - /cart/checkout

