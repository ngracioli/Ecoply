🌱 Ecoply – Marketplace de Energia Excedente Renovável
<p align="center"> <img src="https://img.shields.io/badge/Vue-3.0-42b883?style=for-the-badge&logo=vue.js&logoColor=white"/> <img src="https://img.shields.io/badge/Vite-5.0-646CFF?style=for-the-badge&logo=vite&logoColor=white"/> <img src="https://img.shields.io/badge/TailwindCSS-3.0-38BDF8?style=for-the-badge&logo=tailwindcss&logoColor=white"/> <img src="https://img.shields.io/badge/PrimeVue-Latest-4CAF50?style=for-the-badge"/> <img src="https://img.shields.io/badge/TypeScript-5.0-3178C6?style=for-the-badge&logo=typescript&logoColor=white"/> <img src="https://img.shields.io/badge/Go-1.22-00ADD8?style=for-the-badge&logo=go&logoColor=white"/> <img src="https://img.shields.io/badge/Postman-Usado%20nos%20Testes-orange?style=for-the-badge&logo=postman"/> </p>

A Ecoply é uma plataforma digital que conecta agentes do Mercado Livre de Energia (ACL) para comprar e vender energia excedente renovável de forma rápida, transparente e segura.
O projeto cria um marketplace totalmente automatizado, onde geradores, comercializadores e consumidores livres podem negociar excedentes energéticos sem a burocracia dos contratos tradicionais.

📌 Sumário

🚀 Visão Geral

⚡ Objetivo da Plataforma

💡 Como Funciona

📸 Artefatos e Anexos

🛠️ Tecnologias Utilizadas

📂 Arquitetura do Projeto

📦 Funcionalidades

🔐 Requisitos Não Funcionais

🧠 Requisitos Inovadores

🌍 Agentes da CCEE

🗂️ Banco de Dados

🏗️ Como Rodar o Projeto

🚀 Visão Geral

No Mercado Livre de Energia (ACL), empresas podem negociar energia diretamente — mas negociar excedentes energéticos ainda é um processo:

✔ manual
✔ lento
✔ burocrático
✔ pouco transparente

A Ecoply transforma esse processo criando um marketplace digital onde:

vendedores publicam ofertas de excedente

compradores adquirem energia com preço fixo

a plataforma emite termos, documentos e guias para registro na CCEE

Tudo isso com agilidade, segurança e transparência.

⚡ Objetivo da Plataforma

Criar um ecossistema digital que:

gere liquidez para excedentes energéticos

reduza custos do comprador

simplifique o processo de contratação bilateral

incentive o uso de energia renovável (ODS 7)

minimize desperdícios e aumente a eficiência do setor elétrico brasileiro

💡 Como Funciona
🔎 Validação dos Agentes

Somente agentes ativos e validados da CCEE podem operar.

💰 Marketplace de Ofertas

Vendedores criam ofertas com preço fixo por MWh.

Compradores filtram, visualizam e compram energia disponível.

Submercado do comprador deve coincidir com o da oferta.

📝 Contrato Bilateral Assistido

A compra gera automaticamente:

Termo de acordo digital

Resumo formatado para registro na CCEE

Passo a passo do registro do CCEAL

💵 Liquidação Financeira

Realizada pela própria CCEE no ciclo mensal.
A Ecoply cobra apenas um fee de R$0,10/kWh negociado.

📸 Artefatos e Anexos
🧭 Landing Page

📎 https://github.com/ngracioli/Ecoply/tree/main/anexos/frontend/landingPage

📊 Dashboard

📎 https://github.com/ngracioli/Ecoply/tree/main/anexos/dashboard

🗺️ Diagrama Físico do Banco

📎 https://github.com/ngracioli/Ecoply/blob/main/anexos/banco_diagrama_fisico.png

🛠️ Tecnologias Utilizadas
Frontend

Vue 3

Vite

TailwindCSS

PrimeVue

TypeScript

Backend

Golang

API com autenticação JWT

Arquitetura com camadas (handlers, services, repositories)

Integração com banco SQL

Testes

Postman para requisições, coleções e validação das APIs

📂 Arquitetura do Projeto
Ecoply/
├─ frontend/
│  ├─ src/
│  ├─ components/
│  ├─ pages/
│  └─ ...
├─ backend/
│  ├─ cmd/
│  ├─ internal/
│  ├─ models/
│  ├─ handlers/
│  └─ ...
└─ anexos/

📦 Funcionalidades

✔ Cadastro e autenticação com validação documental
✔ Criação e gestão de ofertas
✔ Filtros avançados (nome, preço, submercado, vendedor)
✔ Recomendação por geolocalização
✔ Compra de energia com atualização em tempo real
✔ Perfil com edição de dados
✔ Emissão automática de contratos e documentos
✔ Tutoriais para registro no CCEE

🔐 Requisitos Não Funcionais

Autenticação por token segura

Interface moderna, clara e responsiva

Baixa latência

Escalabilidade e modularidade

🧠 Requisitos Inovadores

Transparência total

Automação documental

Redução de burocracia no ACL

Gestão descentralizada de excedentes

🌍 Agentes da CCEE

A plataforma atende:

Geradores

Comercializadores

Consumidores Livres

Consumidores Especiais

🗂️ Banco de Dados

Principais entidades:

users

agents

offers

purchases

submarkets

user_types

addresses

Diagrama disponível no repositório.

🏗️ Como Rodar o Projeto
🔧 Pré-requisitos

Node.js 18+

Go 1.22+

npm ou yarn

Banco relacional (PostgreSQL recomendado)

Postman (opcional, porém recomendado)

🖥️ Frontend (Vue + Vite)
cd frontend
npm install
npm run dev


Acesse:

👉 http://localhost:5173/

⚙️ Backend (Golang)
cd backend
go mod tidy
go run cmd/main.go


A API subirá normalmente em:

👉 http://localhost:8080/