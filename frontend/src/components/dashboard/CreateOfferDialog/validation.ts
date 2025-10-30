import { z } from "zod";

const createDecimalValidator = (
  decimalPlaces: number,
  maxTotalChars: number,
  minValue: number,
  fieldName: string,
) => {
  return z
    .union([z.string(), z.number()])
    .transform((val) => String(val))
    .refine((val) => val.length > 0, {
      message: `${fieldName} é obrigatório`,
    })
    .refine(
      (val) => {
        const num = parseFloat(val);
        if (isNaN(num)) return false;
        const parts = val.split(".");
        if (parts.length === 2) {
          return parts[1].length <= decimalPlaces;
        }
        return true;
      },
      {
        message: `${fieldName} deve ter no máximo ${decimalPlaces} casas decimais`,
      },
    )
    .refine(
      (val) => {
        const digitsOnly = val.replace(/[.,]/g, "");
        return digitsOnly.length <= maxTotalChars;
      },
      {
        message: `${fieldName} deve ter no máximo ${maxTotalChars} caracteres no total`,
      },
    )
    .refine(
      (val) => {
        const num = parseFloat(val);
        return !isNaN(num) && num >= minValue;
      },
      { message: `${fieldName} deve ser maior ou igual a ${minValue}` },
    );
};

export const offerSchema = z
  .object({
    price_per_mwh: createDecimalValidator(2, 10, 0.01, "Preço por MWh"),
    quantity_mwh: createDecimalValidator(3, 10, 0.001, "Quantidade"),
    period_start: z
      .string()
      .min(1, { message: "Data de início é obrigatória" })
      .refine(
        (val) => {
          const date = new Date(val);
          return !isNaN(date.getTime());
        },
        { message: "Data de início inválida" },
      )
      .refine(
        (val) => {
          const selectedDate = new Date(val + "T00:00:00");
          const today = new Date();
          today.setHours(0, 0, 0, 0);
          return selectedDate >= today;
        },
        { message: "Data de início não pode ser anterior a hoje" },
      ),
    period_end: z
      .string()
      .min(1, { message: "Data de término é obrigatória" })
      .refine(
        (val) => {
          const date = new Date(val);
          return !isNaN(date.getTime());
        },
        { message: "Data de término inválida" },
      ),
    description: z
      .string()
      .min(5, { message: "Descrição deve ter no mínimo 5 caracteres" })
      .max(500, { message: "Descrição deve ter no máximo 500 caracteres" }),
    energy_type: z.enum(["solar", "eolic", "hydroelectric", "geothermal"], {
      message: "Tipo de energia inválido",
    }),
  })
  .refine(
    (data) => {
      const start = new Date(data.period_start);
      const end = new Date(data.period_end);
      return start < end;
    },
    {
      message: "Data de início deve ser anterior à data de término",
      path: ["period_end"],
    },
  );
