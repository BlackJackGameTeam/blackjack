import { Inter } from 'next/font/google'
import { LoadingButton } from '../components/Button'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  const router = useRouter()
  const [shouldRedirect, setShouldRedirect] = useState<boolean>(true)

  useEffect(() => {
    const redirectTimer = setTimeout(() => {
      setShouldRedirect(false)
      router.push('start/start', 'start')
    }, 5000)

    return () => {
      clearTimeout(redirectTimer) // コンポーネントがアンマウントされる際にタイマーをクリアする
    }
  }, [])

  return (
    <>
      <div className="bg-white h-screen flex mx-auto w-full justify-center">
        <div className="my-96">
          {shouldRedirect == true ? (
            <LoadingButton></LoadingButton>
          ) : (
            <div>Redirecting...</div>
          )}
        </div>
      </div>
    </>
  )
}
