import { toast, ToastOptions, TypeOptions } from "react-toastify"

const toastConfig: ToastOptions = {
  autoClose: 5000,
  theme: 'dark',
  draggable: true,
  closeButton: true,
}

export function toaster(content: string, type: TypeOptions) {
  return toast(content, {
    ...toastConfig,
    type
  })
}
