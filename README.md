# Go User Microservice Tutorial 

Go user microservice with postgresql and docker

## File Structure

```
|- bin
|- cmd
  |- core
|- pkg
  |- config
  |- controllers
  |- models
  |- routers
  |- util

```

### bin
This folder contains the built-in go server.

### cmd
This folder contains the core application(s). In this service, main.go in cmd/core is the main application.

### pkg
The codes written for the service to work are located in this folder. 
- config : The settings codes of the application are found here.
- controllers : Codes with all operations and logic are found here.
- models : The codes where database operations are made are coded here.
- routers : Here are the codes that define the routing of incoming requests.
- util : Tools coded for the application are stored in this folder.

## Installation
To install, you must first authorize the service.sh file. Then type the install command.
```bash
sudo chmod +x ./service.sh
sudo ./service.sh install
```

## Building Services
The services must be built before running the application. Service building codes can be done easily with the service.sh file commands.
#### Create the postgresql database
```bash
sudo ./service.sh build_db
```
#### Build the go service
```bash
sudo ./service.sh build
```

## Staring Go Microservice App
Starting database service
```bash
sudo ./service.sh run_db
```
Starting go microservice 
```bash
sudo ./service.sh run
```

## Usage
### Endpoints
Below are all of the endpoints provided by the server, displayed by their
relative URL, and the HTTP method with which you access them.

#### Get All Users
##### GET /api/user
Example Request:

```http
GET /api/user/ HTTP/1.1
Host: 192.168.1.2:8000
```

#### Get User By Id
##### GET /api/user/{id}
Example Request:

```http
GET /api/user/232 HTTP/1.1
Host: 192.168.1.2:8000
```

#### Create User
##### POST /api/user/{id}
Example Request:

```http
POST /api/user/232 HTTP/1.1
Host: 192.168.1.2:8000
Content-Type: application/json
{
    "Name": "Kadir Umut",
    "Email": "sanelkadir@gmail.com",
    "Password": "secret"
}
```

#### Delete User
##### DELETE /api/user/{id}
Example Request:

```http
DELETE /api/user/232 HTTP/1.1
Host: 192.168.1.2:8000
```

#### Update User
##### PUT /api/user/{id}
Example Request:

```http
PUT /api/user/232 HTTP/1.1
Host: 192.168.1.2:8000
Content-Type: application/json
{
    "Name": "Kadir Umut",
    "Email": "sanelkadir@gmail.com",
    "Password": "secret"
}
```

## Stoping Postgresql Database Service
```bash
sudo ./service.sh stop_db
```


## License
[MIT](https://choosealicense.com/licenses/mit/)
