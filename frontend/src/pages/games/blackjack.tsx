import Head from 'next/head'
import { useState, useEffect } from 'react'
import dynamic from 'next/dynamic'

const DynamicComponentWithNoSSR = dynamic(
  () => import('../../games/blackjack/Preload'),
  {
    ssr: false,
  },
)

const blackjack = () => {
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    setLoading(true)
    console.log('set loading')
  }, [])

  return (
    <div>
      <Head>
        <title>BlackjackGame</title>
        <link rel="icon" href="/card.ico" />
      </Head>
      <div key={Math.random()} id="game"></div>
      {loading ? <DynamicComponentWithNoSSR /> : 'Loading...'}
    </div>
  )
}
export default blackjack
