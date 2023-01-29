import Link from "next/link";
import Inputs from "../components/Inputs";

export default function Login() {
  return (
    <div className="flex justify-center items-center w-screen h-screen bg-neutral-900">
      <div className="w-[408px] flex-row items-center px-4">
        <div className="flex flex-col gap-2 mb-6">
          <h1 className="font-bold text-neutral-50 text-2xl">Entrar</h1>
          <p className="text-neutral-300 text-base">
            Para acessar o chat entre com sua conta
          </p>
        </div>
        <Inputs />
        <div className="w-full flex justify-center text-neutral-300 text-sm pt-4">
          <p>Não possui uma conta?</p>
          <Link href="/cadastro" className="underline decoration-solid pl-1">
            Criar conta
          </Link>
        </div>
      </div>
    </div>
  );
}
