## SUMMARY ##
Sample go webapp using postgres as database, Bootstrap for CSS, jQuery for JS

## Based of ##
* https://github.com/EricLau1/Youtube/tree/master/go-webapp
* Related youtube playlist videos: https://www.youtube.com/playlist?list=PLFXr22TafQUs6JqgVOqst70iLtmGJ8nmc
* many thanks to : Eric Lau de Oliveira @ https://github.com/EricLau1

## Tech stack ##
go1.12.12, Bootstrap v4.2.1, JQuery v3.3.1, Postgres 12

## SETUP ## 

* Go setup:

Download go1.12.12.windows-amd64.msi from https://golang.org/dl/
Install under: C:\dev\go_installs\Go1.12.12\

Confirm installation:
go version
go version go1.12.12 windows/amd64
create directory :\projects\github\go-projects to hold all go projects
set GOPATH environment variable to C:\projects\github\go-projects

* DB installation and setup:

Download postgresql-12.0-1-windows-x64.exe from https://www.enterprisedb.com/downloads/postgres-postgresql-downloads
and install under: C:\dev\PostgreSQL\12
By default, PostgreSQL creates a user and a database called postgres. 
when prompted set super user: postgres password to postgres
It runs as a windows Service: postgresql-x64-12
After the installation you should be able to access the service using that username:postgres
Confirm installation:
add C:\dev\PostgreSQL\12\bin to path environment variable
psql --version
psql (PostgreSQL) 12.0
Now open new command prompt:
psql -U postgres
Password for user postgres: postgres

* setup table and add data:

Now open new command  prompt:
psql -U postgres
psql -U postgres postgres
psql (12.0)
WARNING: Console code page (437) differs from Windows code page (1252)
         8-bit characters might not work correctly. See psql reference
         page "Notes for Windows users" for details.
Type "help" for help.

postgres=#

copy paste the psql/create.sql and hit enter
go_psql_demo_webapp_db=#

copy paste the psql/data.sql and hit enter to setup data

create directory src under: C:\projects\github\go-projects
and Download the go-psql-demo-webapp project into it

##  Install Go packages ##
* cd C:\projects\github\go-projects\src\go-psql-demo-webapp

* go get -u github.com/gorilla/mux
* go get github.com/lib/pq
* go get github.com/gorilla/sessions
* go get github.com/badoux/checkmail
* go get -u golang.org/x/crypto/...

##  BUILDING ## 

* SET PATH=%PATH%;C:\dev\go_installs\Go1.12.12\bin
* SET GOROOT=C:\dev\go_installs\Go1.12.12
* SET GOPATH=C:\projects\github\go-projects
* go.exe build -o go-psql-demo-webapp.exe

##  RUNNING ## 
* cd to C:\projects\github\go-projects\src\go-psql-demo-webapp
* go run app.go
* open http://localhost:8080/

