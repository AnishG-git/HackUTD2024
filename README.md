# HackUTD2024

### Steps to run project
 - Ensure docker daemon is running
 - Open a terminal
 - Make sure you are in hackutd2024 directory
 - Run `make build` (equivalent to docker compose build)
 - Run `make run`   (equivalent to docker compose up)





# IGNORE FOR NOW
For WSL docker setup:

`sudo docker ps`

Get the db container id and run:

`sudo docker exec -it <containerid> psql -h localhost -p 5432 -U postgres -d postgres`

`\c postgres`

`CREATE DATABASE hackutd2024;`

