package resources

type BrazilState struct {
	ID     int    `json:"id"`
	Sigla  string `json:"sigla"`
	Nome   string `json:"nome"`
	Regiao struct {
		ID    int    `json:"id"`
		Sigla string `json:"sigla"`
		Nome  string `json:"nome"`
	} `json:"regiao"`
}

type BrasilApiCnpj struct {
	Cnpj                           string  `json:"cnpj"`
	RazaoSocial                    string  `json:"razao_social"`
	NomeFantasia                   string  `json:"nome_fantasia"`
	DataInicioAtividade            string  `json:"data_inicio_atividade"`
	SituacaoCadastral              int     `json:"situacao_cadastral"`
	DescricaoSituacaoCadastral     string  `json:"descricao_situacao_cadastral"`
	DataSituacaoCadastral          string  `json:"data_situacao_cadastral"`
	MotivoSituacaoCadastral        int     `json:"motivo_situacao_cadastral"`
	DescricaoMotivoSituacaoCadastral string `json:"descricao_motivo_situacao_cadastral"`
	NaturezaJuridica               string  `json:"natureza_juridica"`
	CapitalSocial                  float64 `json:"capital_social"`
	Porte                          string  `json:"porte"`
	Logradouro                     string  `json:"logradouro"`
	Numero                         string  `json:"numero"`
	Complemento                    string  `json:"complemento"`
	Bairro                         string  `json:"bairro"`
	Cep                            string  `json:"cep"`
	Municipio                      string  `json:"municipio"`
	Uf                             string  `json:"uf"`
	Email                          string  `json:"email"`
	DddTelefone1                   string  `json:"ddd_telefone_1"`
	DddTelefone2                   string  `json:"ddd_telefone_2"`
	CnaeFiscal                     int     `json:"cnae_fiscal"`
	CnaeFiscalDescricao            string  `json:"cnae_fiscal_descricao"`
}

type BrasilApiCep struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}
