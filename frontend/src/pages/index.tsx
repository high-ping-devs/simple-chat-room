import Head from 'next/head'
import Image from 'next/image'
import { Inter } from '@next/font/google'
import styles from '@/styles/Home.module.css'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <div className="w-screen h-screen bg-white flex justify-center items-center">
      <h1 className='text-5xl text-slate-800 font-bold'>Setup frontend</h1>
    </div>
  )
}
