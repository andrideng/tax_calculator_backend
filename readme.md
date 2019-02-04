# GO RESTful [tax-calculator]

## Features

* Get List of bills complete with the total price, total tax price and grand total.
* Add New Bill.

## Requirement

* Golang v1.5 or above.
* Mysql v5.7 or above

## Step to create database

```
1. Create database with name [tax_calculator_db]
2. Import query from testdata/db.sql
```

## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. Requires Go 1.5 or above.

After installing Go, run the following commands to download and install this project:

```shell
# install the project
go get github.com/andrideng/tax-calculator

# install dep
$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# fetch the dependent packages
cd $GOPATH/andrideng/tax-calculator
dep ensure
```

Start Application

```shell
go run server.go
```

The application runs as an HTTP server at port 8080. It provides the following RESTful endpoints:

* `GET /api/`: welcoming text api
* `GET /api/ping`: a ping service mainly provided for health check purpose
* `GET /api/bills`: list all bills with the price total, tax total and grand total.
* `POST /api/bills`: create a bill.

API Documentation [https://documenter.getpostman.com/view/528724/RztmsovD]


For example, if you access the URL `http://localhost:8080/api/ping` in a browser, you should see the browser
displays something like `PONG!`.

## Project Structure

This project divides the whole project into four main packages:

* `models`: contains the data structures used for communication between different layers.
* `services`: contains the main business logic of the application.
* `daos`: contains the DAO (Data Access Object) layer that interacts with persistent storage.
* `apis`: contains the API layer that wires up the HTTP routes with the corresponding service APIs.

[Dependency inversion principle](https://en.wikipedia.org/wiki/Dependency_inversion_principle)
is followed to make these packages independent of each other and thus easier to test and maintain.

The rest of the packages in the project are used globally:
 
* `app`: contains routing middlewares and application-level configurations
* `errors`: contains error representation and handling
* `util`: contains utility code

The main entry of the application is in the `server.go` file. It does the following work:

* load external configuration
* establish database connection
* instantiate components and inject dependencies
* start the HTTP server