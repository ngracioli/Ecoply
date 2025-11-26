# Ecoply: Negociação de Energia Excedente ⚡️

Um marketplace digital focado na compra e venda direta de **excedentes energéticos renováveis** no Brasil para Agentes da CCEE (Câmara de Comercialização de Energia Elétrica). A Ecoply simplifica a contratação bilateral, reduz o desperdício de energia e promove a eficiência no Mercado Livre de Energia (ACL).



## ✨ Visão Geral do Projeto

O mercado de energia livre no Brasil frequentemente lida com **excedentes contratuais** que se tornam ativos valiosos, mas de difícil liquidez. O processo atual de negociação bilateral é manual, fragmentado e ineficiente.

A **Ecoply** resolve isso, atuando como uma plataforma centralizada e transparente:

* **Vendedor (Geradores/Comercializadores):** Publica seu excedente de energia renovável a um preço competitivo.
* **Comprador (Consumidores Livres/Comercializadores):** Adquire a energia a um custo significativamente menor que o mercado regulado.
* **Ecoply:** Gera receita por meio de uma **taxa de serviço (fee)** fixa por kWh transacionado, criando um ciclo virtuoso para todos.

O projeto está alinhado com o **ODS 7 – Energia Limpa e Acessível**.

---

## 💻 Tecnologias Utilizadas

Este projeto é dividido em uma arquitetura de frontend e backend, utilizando um conjunto de tecnologias modernas para garantir **escalabilidade** e **rápido tempo de resposta**.

| Categoria | Tecnologia | Descrição |
| :--- | :--- | :--- |
| **Frontend** | **Vue.js** (com **Vite**) | Framework reativo e rápido para a interface do usuário. |
| | **TypeScript** | Garante código mais robusto e menos propenso a erros. |
| | **PrimeVue** | Biblioteca de componentes UI para uma interface intuitiva. |
| | **Tailwind CSS** | Framework utility-first para estilização rápida e responsiva. |
| **Backend** | **Golang (Go)** | Linguagem de alto desempenho para o desenvolvimento da API. |
| **Infraestrutura** | **Docker** | Containerização para ambientes de desenvolvimento e produção consistentes. |
| | **Postman** | Utilizado para testes e documentação da API. |

---

## ⚙️ Funcionalidades Principais

A plataforma Ecoply oferece as seguintes funcionalidades para agentes da CCEE:

* **Autenticação Segura:** Cadastro e login de usuários com **validação documental** e **autenticação por token**.
* **Marketplace de Venda Direta:**
    * Criação e gestão de ofertas de energia excedente (preço fixo por MWh).
    * Filtros avançados (por submercado, preço, tipo de energia e vendedor).
    * **Recomendação por Geolocalização** (ofertas próximas aparecem primeiro na listagem).
* **Formalização Simplificada:**
    * Geração de um **"termo de acordo"** digital após a compra.
    * Emissão automática de **contratos** e **notas fiscais**.
    * Geração de um **"resumo de registro"** com dados formatados para registro assistido na CCEE.
* **Notificações:** Alertas sobre atividades (compras, fim de estoque) e atualizações de status.

---

## 🚀 Como Executar o Projeto Localmente

Para clonar e executar o Ecoply em sua máquina local, siga os passos abaixo.

### Pré-requisitos

Certifique-se de ter as seguintes ferramentas instaladas:

* **Git**
* **Docker**
* **Docker Compose**

### 1. Clonar o Repositório

```bash
git clone [https://github.com/ngracioli/Ecoply.git](https://github.com/ngracioli/Ecoply.git)
cd Ecoply
