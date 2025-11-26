import type { ContractData } from "../types/responses/purchases";

export function generateContract(contractData: ContractData): void {
  const { supplier, offer, buyer } = contractData;

  const totalPrice = offer.price_per_mwh * offer.contracted_quantity_mwh;
  const currentDate = new Date();

  const ecoplyServiceFee = 0.10;
  const totalServiceFee = (offer.contracted_quantity_mwh * 1000 * ecoplyServiceFee).toFixed(2);

  const contractHTML = `
<!DOCTYPE html>
<html lang="pt-BR">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Contrato de Compra e Venda de Energia - Ecoply</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }

    body {
      font-family: 'Georgia', 'Times New Roman', serif;
      line-height: 1.8;
      color: #1a1a1a;
      background: #fff;
      padding: 40px;
      max-width: 1000px;
      margin: 0 auto;
    }

    .header {
      text-align: center;
      border-bottom: 4px double #10b981;
      padding-bottom: 25px;
      margin-bottom: 35px;
    }

    .header h1 {
      color: #10b981;
      font-size: 24px;
      font-weight: 700;
      text-transform: uppercase;
      letter-spacing: 1px;
      margin-bottom: 8px;
      line-height: 1.3;
    }

    .header .subtitle {
      color: #4b5563;
      font-size: 13px;
      font-style: italic;
      margin-top: 5px;
    }

    .contract-meta {
      background: #f9fafb;
      border: 1px solid #d1d5db;
      border-radius: 8px;
      padding: 20px;
      margin: 25px 0;
      font-size: 13px;
    }

    .contract-meta p {
      margin: 5px 0;
    }

    .contract-meta strong {
      color: #10b981;
      font-weight: 600;
    }

    .section {
      margin: 30px 0;
      page-break-inside: avoid;
    }

    .section h2 {
      color: #1f2937;
      font-size: 16px;
      font-weight: 700;
      margin-bottom: 18px;
      text-transform: uppercase;
      letter-spacing: 0.5px;
      border-bottom: 2px solid #10b981;
      padding-bottom: 8px;
    }

    .section h3 {
      color: #374151;
      font-size: 14px;
      font-weight: 600;
      margin: 20px 0 12px 0;
    }

    .party-info {
      background: #f9fafb;
      border-left: 4px solid #10b981;
      padding: 15px 20px;
      margin: 15px 0;
      font-size: 13px;
    }

    .party-info .party-title {
      font-weight: 700;
      color: #10b981;
      text-transform: uppercase;
      font-size: 12px;
      margin-bottom: 8px;
      letter-spacing: 0.5px;
    }

    .party-info p {
      margin: 5px 0;
      line-height: 1.6;
    }

    .data-table {
      width: 100%;
      border-collapse: collapse;
      margin: 20px 0;
      font-size: 13px;
      box-shadow: 0 1px 3px rgba(0,0,0,0.1);
    }

    .data-table th {
      background: #10b981;
      color: white;
      padding: 12px;
      text-align: left;
      font-weight: 600;
      text-transform: uppercase;
      font-size: 11px;
      letter-spacing: 0.5px;
    }

    .data-table td {
      padding: 12px;
      border-bottom: 1px solid #e5e7eb;
      background: #fff;
    }

    .data-table tr:last-child td {
      border-bottom: none;
    }

    .data-table tr:nth-child(even) td {
      background: #f9fafb;
    }

    .data-table td:first-child {
      font-weight: 600;
      color: #4b5563;
      width: 40%;
    }

    .highlight-value {
      color: #10b981;
      font-weight: 700;
      font-size: 15px;
    }

    .terms-text {
      font-size: 13px;
      color: #374151;
      line-height: 1.9;
      text-align: justify;
      margin: 15px 0;
    }

    .terms-text p {
      margin: 15px 0;
    }

    .terms-text strong {
      color: #1f2937;
      font-weight: 600;
    }

    .clause-number {
      font-weight: 700;
      color: #10b981;
    }

    .note-box {
      background: #fef3c7;
      border: 2px solid #f59e0b;
      border-radius: 8px;
      padding: 15px 20px;
      margin: 20px 0;
      font-size: 12px;
      line-height: 1.7;
    }

    .note-box strong {
      color: #b45309;
      display: block;
      margin-bottom: 8px;
      font-size: 13px;
    }

    .signature-section {
      margin-top: 60px;
      page-break-inside: avoid;
    }

    .signature-grid {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 50px;
      margin-top: 40px;
    }

    .signature-box {
      text-align: center;
    }

    .signature-line {
      border-top: 2px solid #1f2937;
      margin-bottom: 10px;
      padding-top: 80px;
    }

    .signature-box .name {
      font-weight: 700;
      color: #1f2937;
      font-size: 14px;
      margin-bottom: 5px;
    }

    .signature-box .role {
      font-size: 12px;
      color: #6b7280;
      font-style: italic;
    }

    .signature-box .details {
      font-size: 11px;
      color: #9ca3af;
      margin-top: 8px;
    }

    .footer {
      margin-top: 60px;
      padding-top: 25px;
      border-top: 2px solid #e5e7eb;
      text-align: center;
      font-size: 11px;
      color: #6b7280;
      line-height: 1.8;
    }

    .footer .disclaimer {
      background: #f3f4f6;
      padding: 15px;
      border-radius: 6px;
      margin-top: 15px;
      font-size: 10px;
      color: #4b5563;
    }

    @media print {
      body {
        padding: 20px;
      }
      .section {
        page-break-inside: avoid;
      }
      .signature-section {
        page-break-before: always;
      }
    }
  </style>
</head>
<body>
  <div class="header">
    <h1>Contrato de Compra e Venda de<br>Energia Elétrica Excedente</h1>
    <p class="subtitle">Versão Ecoply – Operação Bilateral no Ambiente de Contratação Livre (ACL)</p>
  </div>

  <div class="contract-meta">
    <p><strong>Número do Contrato:</strong> ${offer.uuid.substring(0, 8).toUpperCase()}-${currentDate.getFullYear()}</p>
    <p><strong>Data de Emissão:</strong> ${currentDate.toLocaleDateString("pt-BR", {
      day: "2-digit",
      month: "long",
      year: "numeric",
    })} às ${currentDate.toLocaleTimeString("pt-BR", { hour: "2-digit", minute: "2-digit" })}</p>
    <p><strong>Plataforma:</strong> Ecoply Plataforma de Energia Ltda.</p>
  </div>

  <div class="section">
    <h2>1. Qualificação das Partes</h2>

    <div class="party-info">
      <div class="party-title">Vendedor (Agente Fornecedor)</div>
      <p><strong>Razão Social:</strong> ${supplier.company_name}</p>
      <p><strong>CNPJ:</strong> ${formatCNPJ(supplier.cnpj)}</p>
      <p><strong>Código do Agente CCEE:</strong> ${supplier.ccee_code}</p>
      <p><strong>Submercado:</strong> ${supplier.submarket_name}</p>
    </div>

    <div class="party-info">
      <div class="party-title">Comprador (Agente Consumidor Livre ou Especial)</div>
      <p><strong>Razão Social:</strong> ${buyer.company_name}</p>
      <p><strong>CNPJ:</strong> ${formatCNPJ(buyer.cnpj)}</p>
      <p><strong>Código do Agente CCEE:</strong> ${buyer.ccee_code}</p>
      <p><strong>Submercado:</strong> ${buyer.submarket_name}</p>
    </div>

    <div class="party-info">
      <div class="party-title">Intermediadora (Plataforma Digital)</div>
      <p><strong>Razão Social:</strong> Ecoply Plataforma de Energia Ltda.</p>
      <p><strong>Função:</strong> Atuando exclusivamente como intermediadora tecnológica</p>
    </div>
  </div>

  <div class="section">
    <h2>2. Objeto do Contrato</h2>
    <div class="terms-text">
      <p>O presente contrato tem por objeto a <strong>compra e venda bilateral de energia elétrica excedente</strong>, medida em MWh, de origem incentivada, negociada diretamente entre as partes por intermédio da Plataforma Ecoply.</p>
      <p>A entrega da energia será <strong>virtual e sistêmica</strong>, conforme práticas operacionais do ONS e contabilização pela CCEE.</p>
    </div>
  </div>

  <div class="section">
    <h2>3. Base Regulatória e Condições de Agente</h2>
    <div class="terms-text">
      <p>As partes declaram-se agentes regulares e ativos da CCEE, comprometendo-se a manter sua habilitação vigente. Este contrato observa as <strong>Regras de Comercialização e os Procedimentos de Rede</strong>, inclusive quanto à apuração de medições, encargos e penalidades.</p>
      <p>O contrato não substitui o registro obrigatório de contrato bilateral (CCEAL) junto à CCEE, devendo as partes concluir o registro conforme o fluxo operacional da plataforma.</p>
    </div>
  </div>

  <div class="section">
    <h2>4. Preço e Condições Comerciais</h2>
    <table class="data-table">
      <tr>
        <td>Quantidade Contratada</td>
        <td class="highlight-value">${offer.contracted_quantity_mwh.toLocaleString("pt-BR", { minimumFractionDigits: 2, maximumFractionDigits: 2 })} MWh</td>
      </tr>
      <tr>
        <td>Preço Unitário</td>
        <td class="highlight-value">R$ ${offer.price_per_mwh.toLocaleString("pt-BR", { minimumFractionDigits: 2, maximumFractionDigits: 2 })}/MWh</td>
      </tr>
      <tr>
        <td>Valor Total da Energia</td>
        <td class="highlight-value">R$ ${totalPrice.toLocaleString("pt-BR", { minimumFractionDigits: 2, maximumFractionDigits: 2 })}</td>
      </tr>
      <tr>
        <td>Taxa de Serviço Ecoply</td>
        <td>R$ 0,10/kWh transacionado</td>
      </tr>
      <tr>
        <td>Valor Total da Taxa de Serviço</td>
        <td>R$ ${parseFloat(totalServiceFee).toLocaleString("pt-BR", { minimumFractionDigits: 2, maximumFractionDigits: 2 })}</td>
      </tr>
    </table>
    <div class="terms-text">
      <p>A Ecoply cobrará do comprador taxa de serviço de <strong>R$ 0,10/kWh transacionado</strong>, mediante emissão de nota fiscal.</p>
    </div>
  </div>

  <div class="section">
    <h2>5. Medição e Apuração</h2>
    <div class="terms-text">
      <p>A medição oficial será a registrada pela CCEE, com base nos pontos de medição homologados pela distribuidora ou transmissora local.</p>
      <p>Divergências serão resolvidas com base nos relatórios oficiais emitidos pela CCEE.</p>
    </div>
  </div>

  <div class="section">
    <h2>6. Liquidação Financeira</h2>
    <div class="terms-text">
      <p>A liquidação financeira será efetuada <strong>diretamente pela CCEE</strong>, conforme seu ciclo mensal de contabilização e liquidação.</p>
      <p>A Ecoply não transita, guarda ou liquida valores referentes à energia comercializada.</p>
    </div>
  </div>

  <div class="section">
    <h2>7. Garantias e Adimplemento</h2>
    <div class="terms-text">
      <p>As partes poderão estabelecer garantias de fiel cumprimento (caução, seguro-garantia ou fiança bancária), conforme negociação bilateral.</p>
    </div>
  </div>

  <div class="section">
    <h2>8. Rescisão</h2>
    <div class="terms-text">
      <p>O contrato poderá ser rescindido caso qualquer parte perca a condição de agente da CCEE ou em caso de inadimplemento das obrigações contratuais, observadas as penalidades previstas na regulamentação vigente.</p>
    </div>
  </div>

  <div class="section">
    <h2>9. Intermediação pela Ecoply</h2>
    <div class="terms-text">
      <p>A Ecoply atua como <strong>facilitadora digital</strong>, fornecendo infraestrutura segura, validação documental e emissão automatizada do termo de compra e venda.</p>
      <p>A Ecoply não é parte da relação comercial e não assume responsabilidade por lastro, adimplemento ou entrega física.</p>
    </div>
  </div>

  <div class="section">
    <h2>10. Resolução de Conflitos</h2>
    <div class="terms-text">
      <p>As partes envidarão esforços para resolver divergências de forma amigável. Persistindo o conflito, será submetido à Câmara de Arbitragem Especializada em Energia Elétrica ou ao foro da sede do comprador.</p>
    </div>
  </div>

  <div class="section">
    <h2>11. Disposições Finais</h2>
    <div class="terms-text">
      <p>Este contrato tem vigência durante o período de fornecimento definido no Anexo I.</p>
      <p>A assinatura digital na plataforma Ecoply equivale à assinatura física.</p>
      <p>As comunicações formais ocorrerão pelos canais cadastrados na plataforma.</p>
    </div>
  </div>

  <div class="section">
    <h2>Anexo I – Dados Técnicos da Transação</h2>
    <table class="data-table">
      <thead>
        <tr>
          <th>Especificação</th>
          <th>Valor</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>Quantidade Contratada (MWh)</td>
          <td>${offer.contracted_quantity_mwh.toLocaleString("pt-BR", { minimumFractionDigits: 2, maximumFractionDigits: 2 })}</td>
        </tr>
        <tr>
          <td>Submercado</td>
          <td>${offer.submarket}</td>
        </tr>
        <tr>
          <td>Tipo de Energia</td>
          <td>Incentivada – ${getEnergyTypeLabel(offer.energy_type)}</td>
        </tr>
        <tr>
          <td>Período de Fornecimento</td>
          <td>${new Date(offer.period_start).toLocaleDateString("pt-BR")} a ${new Date(offer.period_end).toLocaleDateString("pt-BR")}</td>
        </tr>
        <tr>
          <td>Preço Unitário (R$/MWh)</td>
          <td>R$ ${offer.price_per_mwh.toLocaleString("pt-BR", { minimumFractionDigits: 2, maximumFractionDigits: 2 })}</td>
        </tr>
        <tr>
          <td>Valor Total (R$)</td>
          <td>R$ ${totalPrice.toLocaleString("pt-BR", { minimumFractionDigits: 2, maximumFractionDigits: 2 })}</td>
        </tr>
        <tr>
          <td>Taxa de Serviço Ecoply (R$/kWh)</td>
          <td>0,10</td>
        </tr>
      </tbody>
    </table>
  </div>

  <div class="note-box">
    <strong>Observações Importantes:</strong>
    As partes declaram ter lido e compreendido todas as cláusulas deste contrato, concordando integralmente com seus termos. Este documento foi gerado eletronicamente pela Plataforma Ecoply e possui validade jurídica conforme a legislação brasileira vigente sobre documentos eletrônicos.
  </div>

  <div class="signature-section">
    <h2 style="text-align: center; margin-bottom: 40px;">Assinaturas</h2>
    <div class="signature-grid">
      <div class="signature-box">
        <div class="signature-line"></div>
        <div class="name">${supplier.company_name}</div>
        <div class="role">Vendedor (Agente Fornecedor)</div>
        <div class="details">
          CNPJ: ${formatCNPJ(supplier.cnpj)}<br>
          Código CCEE: ${supplier.ccee_code}
        </div>
      </div>
      <div class="signature-box">
        <div class="signature-line"></div>
        <div class="name">${buyer.company_name}</div>
        <div class="role">Comprador (Agente Consumidor)</div>
        <div class="details">
          CNPJ: ${formatCNPJ(buyer.cnpj)}<br>
          Código CCEE: ${buyer.ccee_code}
        </div>
      </div>
    </div>
  </div>

  <div class="footer">
    <p><strong>Ecoply Plataforma de Energia Ltda.</strong></p>
    <p>Plataforma Digital de Negociação de Energia no Mercado Livre</p>
    <p>www.ecoply.com.br | contato@ecoply.com.br</p>
    <div class="disclaimer">
      <strong>AVISO LEGAL:</strong> Este é um contrato legalmente vinculativo gerado eletronicamente. A assinatura digital possui a mesma validade que a assinatura física, conforme MP 2.200-2/2001 e Lei 14.063/2020. Este documento deve ser registrado junto à CCEE para validade operacional no Ambiente de Contratação Livre.
    </div>
  </div>
</body>
</html>
  `;

  const printWindow = window.open("", "_blank");
  if (printWindow) {
    printWindow.document.write(contractHTML);
    printWindow.document.close();
    printWindow.focus();

    setTimeout(() => {
      printWindow.print();
    }, 250);
  }
}

function getEnergyTypeLabel(type: string): string {
  const labels: Record<string, string> = {
    solar: "Energia Solar",
    eolic: "Energia Eólica",
    hydroelectric: "Energia Hidrelétrica",
    geothermal: "Energia Geotérmica",
  };
  return labels[type] || type;
}

function formatCNPJ(cnpj: string): string {
  const cleaned = cnpj.replace(/\D/g, "");

  if (cleaned.length === 14) {
    return cleaned.replace(
      /^(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})$/,
      "$1.$2.$3/$4-$5",
    );
  }

  return cnpj;
}
