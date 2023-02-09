interface SubmitButtonProps {
  id: string;
  name: string;
  label: string;
}

export default function SubmitButton({ id, name, label }: SubmitButtonProps) {
  return (
    <button
      id={id}
      name={name}
      type="submit"
      className="w-full h-16 rounded-[4px] bg-indigo-600 text-neutral-50 font-bold hover:bg-indigo-800 transition duration-200"
    >
      {label}
    </button>
  );
}
