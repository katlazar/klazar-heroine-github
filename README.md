# INTRODUCTION 
## Autor: Katarzyna Łazar k.lazar@dtpoland.com
## Cel: Przetarcie szlaków dla przyszłych programistek w DTP :-) 

# HEROES ACADEMY

1. Git
2. Angular
3. C# .net
4. Golang
5. Docker
6. Azure
7. OpenID

# 1. Git
Klonowanie repozytorium zdalnego na lokalne:

`$ git clone https://dtpdevops@dev.azure.com/dtpdevops/HEROES/_git/klazar-heroine`

Utworzenie repozytorium na github.com:

https://github.com/katlazar/klazar-heroine-github


# 2. Angular
## Run dev
`cd Heroes-Angular/angular-tour-of-heroes`

`ng serve`

# 3. C# .net
## Run dev
`cd Heroes-Dotnet/HeroesApi`

`dotnet run`

## Run static Angular app
`cd Heroes-Angular/angular-tour-of-heroes`

`ng build`

`cd ..`

`cd ..`

`cp -R Heroes-Angular/angular-tour-of-heroes/dist/angular-tour-of-heroes Heroes-Dotnet/HeroesApi/AngularApp`

`cd Heroes-Dotnet/HeroesApi`

`dotnet run`

### Angular app is served on port 5000

# 4. Go
## Run dev
`cd Heroes-Golang/Herosi`

`go run main.go`

## Run static Angular app
`cd Heroes-Angular/angular-tour-of-heroes`

`ng build`

`cd ..`

`cd ..`

`cp -R Heroes-Angular/angular-tour-of-heroes/dist/angular-tour-of-heroes Heroes-Golang/Herosi/AngularApp`

`cd Heroes-Golang/Herosi`

`go run main.go`


## Run test
`cd Heroes-Golang/Herosi/controllers`

`go test`

# 5. Docker

## .Net

### DockerHub
`docker pull katlaz/klazar-heroine:netapi`
### Run Docker Container
`docker run -d -p 80:80 katlaz/klazar-heroine:netapi`

## Go

### DockerHub
`docker pull katlaz/klazar-heroine:goapi`
### Run Docker Container
`docker run -d -p 8080:8080 katlaz/klazar-heroine:goapi`