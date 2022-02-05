# otoklix-blog

instalation
1. Clone project repo
2. In you dont have postgresql server please make container with command below:
> docker run --name some-postgres -l 5431:5432 -e POSTGRES_PASSWORD=root -d postgres
3. And then run the project with
> go run main.go
4. In this project i have include postman collection to testing the project.
