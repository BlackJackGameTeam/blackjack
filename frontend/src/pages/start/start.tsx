import Head from 'next/head'
import React, { useEffect } from 'react'
import Image from 'next/image'
import { useRouter } from 'next/router'
import axios from 'axios'
import { CsrfToken } from '@/types'

const startPage = () => {
  useEffect(() => {
    // backendからcsrfTokenを取得する処理
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(
        `${process.env.REACT_APP_API_URL}/csrf`
      )
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    getCsrfToken()
  }, [])
  const router = useRouter()
  const handlePage = () => {
    router.push('/login/login', 'login')
  }

  return (
    <>
      <div
        className="hero min-h-screen"
        style={{
          backgroundImage:
            'url(https://daisyui.com/images/stock/photo-1507358522600-9f71e620c44e.jpg)',
        }}
      >
        <div className="hero-overlay bg-opacity-60"></div>
        <div className="hero-content text-center text-neutral-content">
          <div className="max-w-md">
            <h1 className="mb-5 text-5xl font-bold">Hello World</h1>
            <p className="mb-5">
              Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda
              excepturi exercitationem quasi. In deleniti eaque aut repudiandae
              et a id nisi.
            </p>
            <button className="btn btn-primary" onClick={handlePage}>
              Get Started
            </button>
          </div>
        </div>
      </div>
    </>
  )
}

export default startPage
