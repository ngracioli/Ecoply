import type {
  RegisterFirstStepProps,
  RegisterRequest,
  RegisterSecondStepProps,
} from "./intefaces";

export function transformFormData(
  firstStep: RegisterFirstStepProps | null,
  secondStep: RegisterSecondStepProps | null,
): RegisterRequest {
  return {
    name: firstStep?.name ?? "",
    email: firstStep?.email ?? "",
    password: firstStep?.password ?? "",
    address: secondStep?.address ?? "",
    city: secondStep?.city ?? "",
    zip: secondStep?.zip ?? "",
  };
}
