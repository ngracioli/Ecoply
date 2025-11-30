package resources

type CceeAgent struct {
	Label     string `json:"label"`
	Value     string `json:"value"`
	Submarket string `json:"submarket"`
}

type CceeApiResponse struct {
	Success bool          `json:"success"`
	Result  CceeApiResult `json:"result"`
}

type CceeApiResult struct {
	Records []CceeRecord `json:"records"`
}

type CceeRecord struct {
	CodAgente        int    `json:"COD_AGENTE"`
	SiglaAgente      string `json:"SIGLA_AGENTE"`
	NomeEmpresarial  string `json:"NOME_EMPRESARIAL"`
	Cnpj             string `json:"CNPJ"`
	CodPerfAgente    int    `json:"COD_PERF_AGENTE"`
	SiglaPerfAgente  string `json:"SIGLA_PERFIL_AGENTE"`
	ClassePerfAgente string `json:"CLASSE_PERFIL_AGENTE"`
	StatusPerfil     string `json:"STATUS_PERFIL"`
	CategoriaAgente  string `json:"CATEGORIA_AGENTE"`
	Submercado       string `json:"SUBMERCADO"`
	Varejista        string `json:"VAREJISTA"`
}
