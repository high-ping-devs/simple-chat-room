import { useFormik } from "formik";

const SignupForm = () => {
  // Pass the useFormik() hook initial form values and a submit function that will
  // be called when the form is submitted
  const formik = useFormik({
    initialValues: {
      email: "",
      password: "",
    },
    onSubmit: (values) => {
      alert(JSON.stringify(values, null, 2));
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
      </form>

      <button
        id="button"
        name="button"
        type="submit"
        className="w-full h-16 rounded-[4px] bg-indigo-600 text-neutral-50 font-bold hover:bg-indigo-800 transition duration-200"
      >
        Entrar
      </button>
    </div>
  );
};

export default function LoginInputs() {
  return SignupForm();
}
