import Link from "next/link";
import LoginInputs from "../components/LoginInputs";
import * as Yup from "yup";
import * as React from "react";
import * as ReactDOM from "react-dom";
import { Formik, Field, Form, FormikHelpers, useField } from "formik";

const MyTextInput = (props: any) => {
  // useField() returns [formik.getFieldProps(), formik.getFieldMeta()]
  // which we can spread on <input>. We can use field meta to show an error
  // message if the field is invalid and it has been touched (i.e. visited)
  const [field, meta] = useField(props);
  return (
    <>
      <input
        className="w-full bg-neutral-800 h-16 rounded-[4px] text-neutral-300 text-base pl-4 placeholder:text-neutral-300 placeholder:text-base"
        {...field}
        {...props}
      />
      {meta.touched && meta.error ? (
        <span className="text-red-600">{meta.error}</span>
      ) : null}
    </>
  );
};

interface Values {
  email: string;
  password: string;
  confirmPassword: string;
}

const ValidateInputs = Yup.object({
  email: Yup.string().email("Invalid email address").required("Required"),
  password: Yup.string()
    .min(4, "Must be 4 characters or more")
    .required("Required"),
  confirmPassword: Yup.string()
    .min(4, "Must be 4 characters or more")
    .required("Required"),
});

export default function Login() {
  return (
    <div className="flex justify-center items-center w-screen h-screen bg-neutral-900">
      <div className="w-[408px] flex-row items-center px-4">
        <div className="flex flex-col gap-2 mb-6">
          <h1 className="font-bold text-neutral-50 text-2xl">Criar conta</h1>
          <p className="text-neutral-300 text-base">
            Crie uma conta para acessar o chat
          </p>
        </div>

        <Formik
          initialValues={{
            email: "",
            password: "",
            confirmPassword: "",
          }}
          validationSchema={ValidateInputs}
          onSubmit={(
            values: Values,
            { setSubmitting }: FormikHelpers<Values>
          ) => {
            setTimeout(() => {
              alert(JSON.stringify(values, null, 2));
              setSubmitting(false);
            }, 500);
          }}
        >
          {({ errors, touched }) => (
            <Form className="flex flex-col gap-4">
              <MyTextInput id="email" name="email" placeholder="E-mail" />

              <MyTextInput
                id="password"
                name="password"
                placeholder="Senha"
                type="password"
              />
              <MyTextInput
                id="confirmPassword"
                name="confirmPassword"
                placeholder="Confirmar senha"
                type="password"
              />
              <button
                className="w-full h-16 rounded-[4px] bg-indigo-600 text-neutral-50 font-bold hover:bg-indigo-800 transition duration-200"
                type="submit"
              >
                Criar conta
              </button>
            </Form>
          )}
        </Formik>

        <div className="w-full flex justify-center text-neutral-300 text-sm pt-4">
          <p>Já possui uma uma conta?</p>
          <Link href="/" className="underline decoration-solid pl-1">
            Entrar agora
          </Link>
        </div>
      </div>
    </div>

    // <div>
    //   <h1>Signup</h1>

    // </div>
  );
}
