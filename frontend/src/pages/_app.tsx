import '@/styles/globals.css'
import type { AppProps } from 'next/app'
import { useAsync } from 'react-use'
import axios from 'axios'
import { getCsrfToken } from './api/player'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'

export default function App({ Component, pageProps }: AppProps) {
  useAsync(async () => {
    axios.defaults.withCredentials = true
    await getCsrfToken()
  })
  const queryClient = new QueryClient({})
  return (
    <QueryClientProvider client={queryClient}>
      <Component {...pageProps} />
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  )
}
