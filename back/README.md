# Version

Go 1.25.0

# How to run for development

### 1. Create Environment File
Copy the example environment file and configure it:
```bash
cp .env.example .env
```

### 2. Run the API with Docker
```bash
make docker-run
```

# Dependencies

- [Gin](https://gin-gonic.com/) Web framework
- [Gorm](https://gorm.io/) Database ORM
- [Env](https://github.com/caarlos0/env) Handling environment variables with structs
- [DotEnv](https://github.com/joho/godotenv) Reading env from dot env files
- [JWT](https://github.com/golang-jwt/jwt) Parse and management of JWT tokens
- [UUID](https://github.com/google/uuid) uuid generation