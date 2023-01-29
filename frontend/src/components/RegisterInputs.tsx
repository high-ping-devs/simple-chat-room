import { toaster } from "@/utils/toaster";
import { useFormik } from "formik";

const resolveAfter3Sec = new Promise(resolve => {
  setTimeout(resolve, 3000)
});

const SignupForm = () => {
  const formik = useFormik({
    initialValues: {
      email: "",
      password: "",
      confirmPassword: "",
    },
    onSubmit: () => {
      toaster("Sucesso", 'success')
    },
  });
  return (
    <div>
      <form onSubmit={formik.handleSubmit}>
        <input
          id="email"
          name="email"
          type="email"
          placeholder="E-mail"
          onChange={formik.handleChange}
          value={formik.values.email}
          className="mb-4 w-full bg-neutral-800 h-16 rounded-[4px] text-neutral-300 text-base pl-4 placeholder:text-neutral-300 placeholder:text-base"
        />

        <input
          id="password"
          name="password"
          type="password"
          placeholder="Senha"
          onChange={formik.handleChange}
          value={formik.values.password}
          className="mb-4 w-full bg-neutral-800 h-16 rounded-[4px] text-neutral-300 text-base pl-4 placeholder:text-neutral-300 placeholder:text-base"
        />

        <input
          id="confirmPassword"
          name="confirmPassword"
          type="password"
          placeholder="Confirmar senha"
          onChange={formik.handleChange}
          value={formik.values.confirmPassword}
          className="mb-4 w-full bg-neutral-800 h-16 rounded-[4px] text-neutral-300 text-base pl-4 placeholder:text-neutral-300 placeholder:text-base"
        />
        <button
          type="submit"
          className="w-full h-16 rounded-[4px] bg-indigo-600 text-neutral-50 font-bold hover:bg-indigo-800 transition duration-200"
        >
          Criar conta
        </button>
      </form>
    </div>
  );
};

export default function CadastroInput() {
  return SignupForm();
}
