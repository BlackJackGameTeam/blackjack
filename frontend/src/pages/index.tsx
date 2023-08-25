import { Inter } from 'next/font/google'
import { LoadingButton } from '../components/Button';

const inter = Inter({ subsets: ['latin'] })

export default function Home() {

  return (
    <>
      <div className="bg-white h-screen flex mx-auto w-full justify-center">
        <div className='my-96'>
          <LoadingButton></LoadingButton>
        </div>
      </div>
    </>
  )
}
