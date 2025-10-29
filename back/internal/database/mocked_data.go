package database

import (
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/services"
	"fmt"

	"gorm.io/gorm"
)

func insertMockedData(db *gorm.DB) {
	var mockedData models.MockedData = models.MockedData{
		Inserted: 1,
	}
	var mockedInserted int64

	// Check if mocked data has already been inserted
	db.Model(models.MockedData{}).Count(&mockedInserted)
	if mockedInserted > 1 {
		return
	}

	// States
	states := []models.AddressState{
		{State: "São Paulo", Initials: "SP"},
		{State: "Rio de Janeiro", Initials: "RJ"},
		{State: "Minas Gerais", Initials: "MG"},
		{State: "Bahia", Initials: "BA"},
	}
	db.Create(&states)

	// Cities
	cities := []models.AddressCity{
		{City: "São Paulo", StateId: 1},
		{City: "Rio de Janeiro", StateId: 2},
		{City: "Belo Horizonte", StateId: 3},
		{City: "Salvador", StateId: 4},
	}
	db.Create(&cities)

	// Neighborhoods
	neighborhoods := []models.AddressNeighborhood{
		{Neighborhood: "Centro", CityId: 1},
		{Neighborhood: "Zona Norte", CityId: 2},
		{Neighborhood: "Zona Sul", CityId: 3},
		{Neighborhood: "Litoral", CityId: 4},
	}
	db.Create(&neighborhoods)

	// Streets
	streets := []models.AddressStreet{
		{Street: "Rua A", NeighborhoodId: 1},
		{Street: "Rua B", NeighborhoodId: 2},
		{Street: "Rua C", NeighborhoodId: 3},
		{Street: "Rua D", NeighborhoodId: 4},
	}
	db.Create(&streets)

	// Addresses
	addresses := []models.Address{
		{Cep: "01001000", Complement: "Apt 101", Number: "123", StreetId: 1},
		{Cep: "02001000", Complement: "Apt 102", Number: "456", StreetId: 2},
		{Cep: "03001000", Complement: "Apt 103", Number: "789", StreetId: 3},
		{Cep: "04001000", Complement: "Apt 104", Number: "101", StreetId: 4},
	}
	db.Create(&addresses)

	// --- Create 10 unique agents ---
	agents := []models.Agent{
		{Cnpj: "00000000000001", CompanyName: "Agent SE_CO", CceeCode: "00001", SubmarketId: 1, AddressId: 1},
		{Cnpj: "00000000000002", CompanyName: "Agent S", CceeCode: "00002", SubmarketId: 2, AddressId: 2},
		{Cnpj: "00000000000003", CompanyName: "Agent NE", CceeCode: "00003", SubmarketId: 3, AddressId: 3},
		{Cnpj: "00000000000004", CompanyName: "Agent N", CceeCode: "00004", SubmarketId: 4, AddressId: 4},
		{Cnpj: "00000000000005", CompanyName: "Agent SE_CO 2", CceeCode: "00005", SubmarketId: 1, AddressId: 2},
		{Cnpj: "00000000000006", CompanyName: "Agent S 2", CceeCode: "00006", SubmarketId: 2, AddressId: 3},
		{Cnpj: "00000000000007", CompanyName: "Agent NE 2", CceeCode: "00007", SubmarketId: 3, AddressId: 4},
		{Cnpj: "00000000000008", CompanyName: "Agent N 2", CceeCode: "00008", SubmarketId: 4, AddressId: 1},
		{Cnpj: "00000000000009", CompanyName: "Agent SE_CO 3", CceeCode: "00009", SubmarketId: 1, AddressId: 3},
		{Cnpj: "00000000000010", CompanyName: "Agent S 3", CceeCode: "00010", SubmarketId: 2, AddressId: 4},
	}
	db.Create(&agents)

	// --- Create users (10 users for each submarket) ---
	var users []models.User
	password := "Teste@123"
	submarkets := []string{"SE_CO", "S", "NE", "N"}

	for i, submarket := range submarkets {
		{
			users = append(users, models.User{
				Uuid:       services.NewUuidV7String(),
				Name:       fmt.Sprintf("Fornecedor %s", submarket),
				Email:      fmt.Sprintf("%s.supplier@mail.com", submarket),
				Password:   services.Hash256String(password),
				UserTypeId: 1,
				AgentId:    uint(i * 5),
			})
		}

		{
			users = append(users, models.User{
				Uuid:       services.NewUuidV7String(),
				Name:       fmt.Sprintf("Comprador %s", submarket),
				Email:      fmt.Sprintf("%s.buyer@mail.com", submarket),
				Password:   services.Hash256String(password),
				UserTypeId: 2,
				AgentId:    uint(i * 5),
			})
		}
	}

	// Create users in batch
	db.Create(&users)

	// Mark the mocked data as inserted
	db.Create(&mockedData)
}
