import { z } from "zod";

const createDecimalValidator = (
  decimalPlaces: number,
  totalDigits: number,
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
        const num = parseFloat(val);
        if (isNaN(num)) return false;

        // Conta dígitos totais (antes e depois do ponto)
        const digitsOnly = val.replace(/[.,]/g, "");
        return digitsOnly.length <= totalDigits;
      },
      {
        message: `${fieldName} deve ter no máximo ${totalDigits} dígitos`,
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

const isValidDate = (dateString: string): boolean => {
  const [year, month, day] = dateString.split("-").map(Number);
  const date = new Date(year, month - 1, day);

  return (
    date.getFullYear() === year &&
    date.getMonth() === month - 1 &&
    date.getDate() === day
  );
};

export const offerSchema = z
  .object({
    price_per_mwh: createDecimalValidator(2, 10, 0.01, "Preço por MWh"),
    quantity_mwh: createDecimalValidator(3, 10, 0.001, "Quantidade"),
    period_start: z
      .string()
      .min(1, { message: "Data de início é obrigatória" })
      .superRefine((val, ctx) => {
        const date = new Date(val);
        if (isNaN(date.getTime()) || !isValidDate(val)) {
          ctx.addIssue({
            code: "custom",
            message: "Data de início inválida",
          });
          return z.NEVER;
        }

        const selectedDate = new Date(val + "T00:00:00");
        const today = new Date();
        today.setHours(0, 0, 0, 0);
        if (selectedDate < today) {
          ctx.addIssue({
            code: "custom",
            message: "Data de início não pode ser anterior a hoje",
          });
        }
      }),
    period_end: z
      .string()
      .min(1, { message: "Data de término é obrigatória" })
      .superRefine((val, ctx) => {
        const date = new Date(val);
        if (isNaN(date.getTime()) || !isValidDate(val)) {
          ctx.addIssue({
            code: "custom",
            message: "Data de término inválida",
          });
        }
      }),
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
      if (!isValidDate(data.period_start) || !isValidDate(data.period_end)) {
        return true;
      }
      const start = new Date(data.period_start);
      const end = new Date(data.period_end);
      return start < end;
    },
    {
      message: "Data de início deve ser anterior à data de término",
      path: ["period_end"],
    },
  );
