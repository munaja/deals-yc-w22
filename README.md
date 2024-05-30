# Dating App
Dating Application

## Technical Design
- Database, used to store primary data with 4 main tables:
  - user
  - usertoken
  - profile
  - viewresult
  - subscription

- Memory Storage, used for anything that has expiration time. There are 2 flows that utilize memory-storage:
  - viewlist, will be expired at the end of the day
  - subscripotion, will be expired based on user input (in months)

## Technical Infrastructure
Used stacks:
- Programming Language: Go version 1.22.1 or above (https://go.dev/)
- Database: MariaDB version 10.10 (https://mariadb.org/) or MySQL with the evquivalent version (https://www.mysql.com)
- Memory Storage: Redis (https://redis.io/)
- Project Layout: Golang Standard Project Layout (https://github.com/golang-standards/project-layout) - this is just reference not reallly a stack

Based on the needs and according to the `go-project-standard-layout`, there are 3 directories that are being used:
- `internal`: consists any logic related to the business flow, which follows `Clean Architecture` (`CA`) approach
- `pkg`: consists any logic that are not related to the business flow
- `cmd`: consists any executabble parts that run the services (acctually only one service for exam purpose)
- `docs`: any documentation
- `test`: test script or request documentation
- `assets`: any assets used by the application

### More About `internal` Directory
To make it more organized, the files are grouped into 3 main parts or layers (according to the `CA` as mentioned previously) as follows:
- `entities` aka `domain`, containing files that define the data structure
- `use-cases` aka `application`, containing files that decide how the flows go
- `interfaces` aka `presentation`, containing files process user request and create responses in return.

Note:
- there are 2 models of `CA` approach: 3 layers and 4 layers. The 4 layers sparates interfaces into 2 more parts.
- the `cmd` directory also can be considered part of the `interfaces`

### More About `pkg` Directory
All of the codes within the directory are taken from several projects that are previously used by the author (the exam participant).


## Running The Project
### Preparation
- Make sure the main stack (`Go`, `MariaDB`, `Redis`) with the specified version above are already installed.
- Create database with the desired name on the `MariaDB` server.
- Clone the repository `https://github.com/munaja/deals-yc-w22` into the local PC
- Go to the project directory that has been cloned
- Run command `go mod tidy` to install dependency.

### Database Migration
- From the project directory, go to `cmd/all-db-migration`
- Create file named `.env` which can be copied from file `.env-example`
- Modify the file `.env` that has been created, adjust the database setting for the `dsn` key according to the environment, using format `{user}:{pass}@tcp({host}:{port})/dbname?charset=utf8mb4&parseTime=True&loc=Local`.
- Run command `go run .` to directly execute it, or
- Run command `go build .` to create the executabble file, then execute it.

### Running the Application
- From the project directory, go to `cmd/customer`
- Create file named `.env` which can be copied from file `.env-example`
- Modify the file `.env` that has been created, adjust the http server, database (which similar to the previous sample), and memory storage (example config uses default value) setting according to the environment.
- Run command `go run .` to directly execute it, or
- Run command `go build .` to create the executabble file, then execute it.

### API Documentation
Testing can be done using `postman` directly please import the collection which is available in the project repository under `test` directory: `https://github.com/munaja/deals-yc-w22/blob/main/test/Deals%20YC%20W22.postman_collection.json`. The file name is `Deals YC W22.postman_collection`, imported collection name should be `Deals YC W22`.

There is environment variable named `{{customer-host}}` which should be filled with host server created by the app, for example: `localhost:8100`

Some special flow of the request:
- Account registration needs an activation, which will be included in the result if the `env` in the `.env` is set to `development`, for example result will have a note similar to this  `For Dev Only: use this path '/account/confirm-by-email?email=santoso_03@gmail.com&token=fb8bbf42-c36d-4be2-a36a-3af4a471bf56' to activate account`. There for after registration please access the given path `/account/confirm-by-email?email=santoso_03@gmail.com&token=fb8bbf42-c36d-4be2-a36a-3af4a471bf56` from postman which available under `Account` directory request `Request Confirmation By Email`
- Login will return data with access token as one of it's field. Will be needed in most of the case as bearer token. Put in the token field, under the main direcotry > `Authorization` tab > type `Bearer Token`

There is also example data that can be imported by mysql: `https://github.com/munaja/deals-yc-w22/blob/main/test/deals-yc-w22.sql`