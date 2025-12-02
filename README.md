# 🌱 Ecoply – Marketplace de Energia Excedente Renovável

<p align="center">
  <img src="https://img.shields.io/badge/Vue-3.0-42b883?style=for-the-badge&logo=vue.js&logoColor=white"/>
  <img src="https://img.shields.io/badge/Vite-5.0-646CFF?style=for-the-badge&logo=vite&logoColor=white"/>
  <img src="https://img.shields.io/badge/TailwindCSS-3.0-38BDF8?style=for-the-badge&logo=tailwindcss&logoColor=white"/>
  <img src="https://img.shields.io/badge/PrimeVue-Latest-4CAF50?style=for-the-badge"/>
  <img src="https://img.shields.io/badge/TypeScript-5.0-3178C6?style=for-the-badge&logo=typescript&logoColor=white"/>
  <img src="https://img.shields.io/badge/Go-1.22-00ADD8?style=for-the-badge&logo=go&logoColor=white"/>
  <img src="https://img.shields.io/badge/Postman-Usado%20nos%20Testes-orange?style=for-the-badge&logo=postman"/>
  <img src="https://img.shields.io/badge/Docker-Latest-2496ED?style=for-the-badge&logo=docker&logoColor=white"/>
  <img src="https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white"/>
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
* [📸 Artefatos e Anexos](#-artefatos-e-anexos)
* [🔗 Links Importantes](#-links-importantes)

---

## 🚀 Visão Geral

No Mercado Livre de Energia (ACL), negociar excedentes energéticos ainda é um processo:

* **Manual, Lento e Burocrático:** Exige negociações bilaterais fragmentadas.
* **Pouco Transparente:** Falta liquidez imediata para o ativo energético.

A Ecoply transforma esse processo, criando um marketplace centralizado onde:

1.  **Vendedores** publicam ofertas de excedente com preço fixo.
2.  **Compradores** adquirem a energia de forma instantânea.

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

### 📝 Contrato Bilateral Assistido

A compra gera automaticamente a documentação necessária:

* **Termo de Acordo Digital:** Comprador, vendedor, quantidade e preço.
* **Resumo Formatado:** Dados prontos para o registro na CCEE.

---

## 🛠️ Tecnologias Utilizadas

| Categoria | Tecnologia | Versão Principal |
| :--- | :--- | :--- |
| **Frontend** | Vue | 3.x |
| | Vite, TailwindCSS, PrimeVue | Latest |
| | TypeScript | 5.x |
| **Backend** | Golang (Go) | 1.22+ |
| | **Arquitetura:** API com autenticação JWT e camadas (Handlers, Services, Repositórios). | |
| **Banco de Dados** | **PostgreSQL** | Relacional e robusto. |
| **Infraestrutura** | **Docker** | Containerização para desenvolvimento e testes. |
|  | Postman | Utilizado para requisições e validação das APIs. |

## 📦 Funcionalidades e Requisitos

### Requisitos Principais

* ✔ Cadastro e autenticação de usuários com **validação documental**.
* ✔ Criação e gestão de ofertas de venda de energia.
* ✔ Filtros avançados: tipo de energia, submercado e data inicio/fim.
* ✔ **Recomendação por Geolocalização** (ofertas próximas em destaque).
* ✔ Emissão automática de **contratos e notas fiscais** após a transação.

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

### 🗂️ Banco de Dados

O banco de dados relacional utiliza do **PostgreSQL** e usa as seguintes entidades principais para o controle de usuários, transações e localizações:

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

## 🔗 Links Importantes

* 🌐 **Link do Website**: https://ecoply.app/

* 📘 **Documentação interna**: https://docs.google.com/document/d/1UwBPm3Txfcy7cO2q9TO8K7J3G8PPCgLKKpbg-3Oc5T8/edit?usp=sharing

* 📄 **Link do Trello**: https://trello.com/b/QihAB3w8/pi2s2025
