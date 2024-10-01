# Pomodoro services

## Tasks

- [x] Projects management  
- [x] Tasks management  
- [x] Payment  
- [ ] Authentication

## Running the project

Copy the `.env.example` file to `.env` and fill the environment variables.

```bash
docker-compose up --build -d
```

## Running migrations

```bash
go run . migrate up
```
