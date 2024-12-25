# DOT Golang Test

## Architecture
- Arsitektur yang saya gunakan menggunakan “clean architecture” yang diusulkan oleh Robert C.
Martin (lebih dikenal sebagai Paman Bob). Dimana dalam “clean architecture” tersebut terdapat 4
layer. yaitu Entity, Use Case,Interface Adapters, Frameworks And Drivers. Dalam implementasinya
saya menggunakan repository https://github.com/herryg91/go-clean-architecture sebagai contoh. Lalu
saya menggunakan MySQL dan Redis sebagai database karena penggunaannya lebih mudah dan cepat untuk
proses developing 
- The clean architecture that I made was based on [here](https://github.com/herryg91/go-clean-architecture)

## Requirement
- MySQL and Redis for Database

## Library
- [Gorilla mux](https://github.com/gorilla/mux) for routing
- [GORM](https://gorm.io/) for database ORM
- [Validator](github.com/go-playground/validator) for input validation
- [JWT](github.com/golang-jwt/jwt) for authentication login
- [godotenv](github.com/joho/godotenv) for create .ENV
- [Redis](github.com/redis/go-redis/v9) for database Redis


## How To Setup on Local Environment
- Git clone this repository to your local environment
- Copy env file to .env
- Change configuration with your own configuration
- Open file main.go
- If you want to run database migration you can remove comment on this line

- go to terminal and run 

```go
go run main.go
```

- Install MySQL

```bash
sudo apt-get install -y mysql-server
```

- Setup MySQL

```bash
sudo mysql_secure_installation
```

- Install Redis

```bash
curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg

echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list

sudo apt-get update
sudo apt-get install redis
```

- Start Redis

```bash
sudo service redis-server start
```

- Connect to Redis

```bash
redis-cli
```

- Create directory for golang app

```bash
mkdir /home/ubuntu/go
```

- Set golang environtment variable

```bash
# GOROOT is the location where Go package is installed on your system
export GOROOT=/usr/lib/go
# GOPATH is the location of your work directory
export GOPATH=$HOME/go
```

- Move to go directory and clone repository from git

```bash
cd /home/ubuntu/go
git clone https://github.com/edwinjordan/golang_test_dot.git
cd golang_test_dot
```

- Copy env to .env and change configuration with your own configuration

```bash
cp env .env
```

- Check if the app is running normally

```bash
go run main.go
```

- Build golang app

```bash
go build
```
