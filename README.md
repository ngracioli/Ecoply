# 🌱 Ecoply – Marketplace de Energia Excedente Renovável

<p align="center">
  <img src="https://img.shields.io/badge/Vue-3.0-42b883?style=for-the-badge&logo=vue.js&logoColor=white"/>
  <img src="https://img.shields.io/badge/Vite-5.0-646CFF?style=for-the-badge&logo=vite&logoColor=white"/>
  <img src="https://img.shields.io/badge/TailwindCSS-3.0-38BDF8?style=for-the-badge&logo=tailwindcss&logoColor=white"/>
  <img src="https://img.shields.io/badge/PrimeVue-Latest-4CAF50?style=for-the-badge"/>
  <img src="https://img.shields.io/badge/TypeScript-5.0-3178C6?style=for-the-badge&logo=typescript&logoColor=white"/>
  <img src="https://img.shields.io/badge/Go-1.22-00ADD8?style=for-the-badge&logo=go&logoColor=white"/>
  <img src="https://img.shields.io/badge/Postman-Usado%20nos%20Testes-orange?style=for-the-badge&logo=postman"/>
</p>

A Ecoply é uma plataforma digital que conecta agentes do Mercado Livre de Energia (ACL) para comprar e vender **energia excedente renovável** de forma rápida, transparente e segura.

O projeto cria um marketplace totalmente automatizado, onde geradores, comercializadores e consumidores livres podem negociar excedentes energéticos sem a burocracia dos contratos tradicionais.

---

## 📌 Sumário

* [🚀 Visão Geral](#-visão-geral)
* [⚡ Objetivo da Plataforma](#-objetivo-da-plataforma)
* [💡 Como Funciona](#-como-funciona)
* [🛠️ Tecnologias Utilizadas](#-tecnologias-utilizadas)
* [📦 Funcionalidades e Requisitos](#-funcionalidades-e-requisitos)
* [📂 Arquitetura do Projeto](#-arquitetura-do-projeto)
* [📸 Artefatos e Anexos](#-artefatos-e-anexos)
* [🏗️ Como Rodar o Projeto](#-como-rodar-o-projeto)

---

## 🚀 Visão Geral

No Mercado Livre de Energia (ACL), negociar excedentes energéticos ainda é um processo:

* **Manual, Lento e Burocrático:** Exige negociações bilaterais fragmentadas.
* **Pouco Transparente:** Falta liquidez imediata para o ativo energético.

A Ecoply transforma esse processo, criando um marketplace centralizado onde:

1.  **Vendedores** publicam ofertas de excedente com preço fixo.
2.  **Compradores** adquirem a energia de forma instantânea.
3.  A plataforma automatiza a **emissão de termos, documentos e guias** para registro oficial na CCEE (Contratos de Comercialização de Energia no Ambiente Livre - CCEAL).

## ⚡ Objetivo da Plataforma

O projeto está alinhado com o **ODS 7 – Energia Limpa e Acessível** e visa:

* **Gerar Liquidez:** Monetizar excedentes energéticos de forma rápida.
* **Reduzir Custos:** O comprador adquire energia a um valor menor que o mercado regulado.
* **Simplificar:** Reduzir drasticamente o tempo e a burocracia da contratação bilateral.
* **Incentivar:** Promover a eficiência e o uso de energia renovável no setor elétrico brasileiro.

## 💡 Como Funciona

### 🔎 Validação e Segurança

Somente **Agentes da CCEE**, ativos e validados, podem operar na plataforma, garantindo a legitimidade de todas as transações.

### 💰 Marketplace de Ofertas

* **Criação:** Vendedores criam ofertas com preço fixo por MWh e quantidade disponível.
* **Compra:** Compradores filtram, visualizam e compram a energia.
* **Regra de Submercado:** O submercado do comprador deve **coincidir** com o submercado onde a oferta foi criada.

### 📝 Contrato Bilateral Assistido

A compra gera automaticamente a documentação necessária:

* **Termo de Acordo Digital:** Comprador, vendedor, quantidade e preço.
* **Resumo Formatado:** Dados prontos para o registro na CCEE.
* **Tutorial:** Guia passo a passo para finalizar o registro do CCEAL.

### 💵 Liquidação Financeira

A **liquidação financeira da energia** é realizada pela própria **CCEE** em seu ciclo mensal. A Ecoply fatura apenas a sua taxa de serviço (*fee*), atualmente R$0,10/kWh negociado, separadamente, através de uma fatura de prestação de serviços.

---

## 🛠️ Tecnologias Utilizadas

| Categoria | Tecnologia | Versão Principal |
| :--- | :--- | :--- |
| **Frontend** | Vue | 3.x |
| | Vite, TailwindCSS, PrimeVue | Latest |
| | TypeScript | 5.x |
| **Backend** | Golang (Go) | 1.22+ |
| | **Arquitetura:** API com autenticação JWT e camadas (Handlers, Services, Repositories). | |
| **Testes** | Postman | Utilizado para requisições e validação das APIs. |

## 📦 Funcionalidades e Requisitos

### Requisitos Funcionais Principais

* ✔ Cadastro e autenticação de usuários com **validação documental**.
* ✔ Criação e gestão de ofertas de venda de energia.
* ✔ Filtros avançados: nome, preço, submercado e vendedor.
* ✔ **Recomendação por Geolocalização** (ofertas próximas em destaque).
* ✔ Emissão automática de **contratos e notas fiscais** após a transação.
* ✔ Notificações sobre atividades e status de compras.

### 🔐 Requisitos Não Funcionais

* **Segurança:** Autenticação por token (JWT) e proteção de rotas.
* **Usabilidade:** Interface moderna, clara e responsiva (Vue + PrimeVue).
* **Performance:** Baixa latência e rápido tempo de resposta.
* **Arquitetura:** Alta escalabilidade e modularidade.

### 🧠 Requisitos Inovadores

* **Transparência Total:** Redução de burocracia no ACL através da automação documental.
* **Gestão Descentralizada:** Facilidade na comercialização de excedentes, mitigando riscos.

### 🌍 Agentes da CCEE Atendidos

A plataforma é projetada para atender:

* Geradores
* Comercializadores
* Consumidores Livres
* Consumidores Especiais

---

## 📂 Arquitetura do Projeto

O projeto segue uma arquitetura separada entre Frontend e Backend:

Ecoply/ ├─ frontend/ # Aplicação Vue/Vite/TypeScript │ ├─ src/ │ ├─ components/ │ └─ pages/ ├─ backend/ # API em Golang │ ├─ cmd/ # Ponto de entrada (main.go) │ ├─ internal/ # Lógica de negócio, serviços, repositórios │ ├─ models/ │ └─ handlers/ └─ anexos/ # Documentação, diagramas e prints


### 🗂️ Banco de Dados

O banco de dados relacional utiliza as seguintes entidades principais para o controle de usuários, transações e localizações:

* `users`, `agents`, `user_types`
* `offers`, `purchases`
* `submarkets`, `addresses`

O **Diagrama Físico** completo está disponível na seção de anexos.

---

## 📸 Artefatos e Anexos

| Artefato | Descrição | Link |
| :--- | :--- | :--- |
| 🧭 **Landing Page** | Vitrine inicial da plataforma e guia de *onboarding*. | [anexos/frontend/landingPage](https://github.com/ngracioli/Ecoply/tree/main/anexos/frontend/landingPage) |
| 📊 **Dashboard** | Ponto central de interação e visualização de ofertas. | [anexos/dashboard](https://github.com/ngracioli/Ecoply/tree/main/anexos/dashboard) |
| 🗺️ **Diagrama Físico do Banco** | Definição das tabelas e relações do banco de dados. | [anexos/banco\_diagrama\_fisico.png](https://github.com/ngracioli/Ecoply/blob/main/anexos/banco_diagrama_fisico.png) |

---

## 🏗️ Como Rodar o Projeto

### 🔧 Pré-requisitos

* **Node.js** (18+) e **npm** ou **yarn**
* **Go** (1.22+)
* Banco de dados relacional (PostgreSQL recomendado)
* **Postman** (Opcional, mas recomendado para testes de API)

### 1. 🖥️ Rodando o Frontend (Vue + Vite)

```bash
cd frontend
npm install
npm run dev
```
* Acesse a aplicação em: 👉 http://localhost:5173/

### 2. ⚙️ Rodando o Backend (Golang)

```Bash

cd backend
go mod tidy
go run cmd/main.go
```

A API estará rodando em: 👉 http://localhost:8080/

🧪 Testes de API com Postman

Uma collection de Postman está disponível para facilitar o envio de requisições REST e a validação de todas as rotas da aplicação.
