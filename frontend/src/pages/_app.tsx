import '@/styles/globals.css'
import type { AppProps } from 'next/app'
import { useAsync } from 'react-use'
import axios from 'axios'
import { getCsrfToken } from './api/player'
import { useRouter } from 'next/router'

export default function App({ Component, pageProps }: AppProps) {
  useAsync(async () => {
    axios.defaults.withCredentials = true
    await getCsrfToken()
  })
  return <Component {...pageProps} />
}
