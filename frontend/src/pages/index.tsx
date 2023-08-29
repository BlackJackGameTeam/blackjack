import { Inter } from 'next/font/google'
import { LoadingButton } from '../components/Button'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { getCsrfToken } from './api/player'
import { useAsync } from 'react-use'
import axios from 'axios'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  const router = useRouter()
  setTimeout(() => {
    router.push('/start')
  }, 5000)

  return (
    <>
      <div className="bg-white h-screen flex mx-auto w-full justify-center">
        <div className="my-96">
          <LoadingButton></LoadingButton>
        </div>
      </div>
    </>
  )
}
