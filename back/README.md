# The Project

REST API backend built with Go for Brazil's surplus energy marketplace. Enables energy producers, suppliers, and resellers to trade excess energy capacity through a streamlined platform.

## Offer System

The system implements a **stock market-like trading mechanism** for energy offers:

- **Offer Management**: Suppliers create energy offers with price per MWh, quantity (in MWh), validity period, energy type, and submarket classification
- **Purchase System**: Buyers place purchase orders against available offers, similar to stock transactions. Each purchase decrements the offer's remaining quantity
- **Order Lifecycle**: Offers transition through statuses (`fresh` → `open` → `fulfilled`/`expired`) based on quantity availability and time constraints

## Others

- **Persistence Layer**: PostgreSQL with GORM for relational data modeling
- **Authentication**: JWT-based authentication with role-based access control for different market participants (producers, suppliers)
- **Business Validation**: CNPJ validation for Brazilian company registration, energy type classification, and submarket segmentation

## How to run for development

### 1. Create Environment File

Copy the example environment file and configure it:

```bash
cp .env.example .env
```

### 2. Run the API with Docker

```bash
make docker-run
```

## Tools

- [Air](https://github.com/air-verse/air) For go live updating, helped a lot during development

## Dependencies

- [Gin](https://gin-gonic.com/) Web framework, to make http routing more easy
- [Gorm](https://gorm.io/) Database ORM
- [Env](https://github.com/caarlos0/env) Handling environment variables with structs
- [DotEnv](https://github.com/joho/godotenv) Reading env from dot env files
- [JWT](https://github.com/golang-jwt/jwt) Parse and management of JWT tokens
- [UUID](https://github.com/google/uuid) uuid generation
