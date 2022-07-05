# Go api with hexagonal architecture
A simple api to show how to configure an application following a hexagonal architecture. 

**Disclaimer**🚨: This is my first application in go from scratch, and it may have a lot of bugs 🐛

## 🏛 Architecture 
This project is composed for the following modules:

* ### application
The main go file. It contains the configuration to launch the application such as the endpoints, db dependency injection...

* ### domain
The pure logic of the api. Contains the model, the services with user case logic and port interfaces to implement the dependency inversion principle (DIP)

* ### adapters
External resources of the application like mysql db

* ### rest-api
It is another adapter, and is the entrypoint for the application. That's why I prefer to have it in a separate module for clarity

## 🚀 Usage
### 1. Run dependencies
All the external resources needed to up and run the api they are configured in a docker-compose. To run it just execute:
```
make containers 
```

### 2. Run api
To run the application locally just execute:
```
make run 
```

## 🚧 TODO list
* [ ] Unit and integrated test. (I need to get more knowledge in this point)
* [ ] Explore ways to avoid the manual injection of dependencies
* [ ] MongoDB integration
* [ ] Integration to other services with gRPC
* [ ] Integration to other services with http
* [ ] kafka integration