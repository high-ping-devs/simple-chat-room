import RegisterInput from "@/components/RegisterInputs";
import Link from "next/link";

export default function cadastro() {
  return (
    <div className="flex justify-center items-center w-screen h-screen bg-neutral-900">
      <div className="w-[408px] flex-row items-center px-4">
        <div className="flex flex-col gap-2 mb-6">
          <h1 className="font-bold text-neutral-50 text-2xl">Criar conta</h1>
          <p className="text-neutral-300 text-base">
            Crie uma conta para acessar o chat
          </p>
        </div>
        <RegisterInput></RegisterInput>
        <div className="w-full flex justify-center text-neutral-300 text-sm pt-4">
          <p>Já possui uma conta?</p>
          <Link href="/" className="underline decoration-solid pl-1">
            Entrar agora
          </Link>
        </div>
      </div>
    </div>
  );
}
