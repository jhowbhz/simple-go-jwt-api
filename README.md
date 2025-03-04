# Simple API with Go, JWT, and SQLite

<img src="https://github.com/user-attachments/assets/451aef64-5bc9-4e01-8fb0-0e56d2337eea" width="100%" alt="Go lang"/>

Este projeto é um estudo simples utilizando a linguagem Go, SQLite como banco de dados, e Docker, com autenticação via tokens JWT.

# Como baixar o projeto?
```bash
git clone https://github.com/jhowbhz/simple-go-jwt-api.git simple-go-jwt-api && cd simple-go-jwt-api
```

# Como rodar?
```bash
docker compose up --build
```

# Tecnologias Utilizadas
Go: Linguagem de programação para o desenvolvimento da API.

SQLite: Banco de dados leve para armazenar as informações.

JWT (JSON Web Tokens): Tecnologia de autenticação e autorização utilizando tokens.

Docker: Containerização da aplicação.

# Endpoints

POST /register: Registra um novo usuário.

POST /login: Realiza o login do usuário e retorna um token JWT.

GET /profile: Retorna as informações do perfil do usuário autenticado.
