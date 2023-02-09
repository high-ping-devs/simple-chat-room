import SubmitButton from "@/components/SubmitButton";
import { Form, Formik, useField, FieldHookConfig } from "formik";
import Link from "next/link";
import LoginInputs from "../components/LoginInputs";
import * as Yup from "yup";

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

export default function Login() {
  return (
    <div className="flex justify-center items-center w-screen h-screen bg-neutral-900">
      <div className="w-[408px] flex-row items-center px-4">
        <div className="flex flex-col gap-2 mb-6">
          <h1 className="font-bold text-neutral-50 text-2xl">Entrar</h1>
          <p className="text-neutral-300 text-base">
            Para acessar o chat, entre com sua conta
          </p>
        </div>

        <Formik
          initialValues={{
            email: "",
            password: "",
          }}
          validationSchema={Yup.object({
            firstName: Yup.string()
              .max(15, "Must be 15 characters or less")
              .required("Required"),
            lastName: Yup.string()
              .max(20, "Must be 20 characters or less")
              .required("Required"),
            email: Yup.string()
              .email("Invalid email address")
              .required("Required"),
          })}
          onSubmit={(values, { setSubmitting }) => {
            setTimeout(() => {
              alert(JSON.stringify(values, null, 2));
              setSubmitting(false);
            }, 400);
          }}
        >
          {(formik) => (
            <Form className="flex flex-col gap-4">
              <MyTextInput name="email" type="email" placeholder="E-mail" />
              <MyTextInput name="senha" type="password" placeholder="Senha" />
              <SubmitButton id={"button"} name={"button"} label={"Entrar"} />
            </Form>
          )}
        </Formik>

        <div className="w-full flex justify-center text-neutral-300 text-sm pt-4">
          <p>Não possui uma conta?</p>
          <Link href="/register" className="underline decoration-solid pl-1">
            Criar conta
          </Link>
        </div>
      </div>
    </div>
  );
}
