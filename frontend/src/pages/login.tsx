import Head from 'next/head'
import React, { FormEvent } from 'react'
import Image from 'next/image'
import backgroundImage from '../../public/assets/backgroudPicture.png'
import { useState } from 'react'
import { playerMutateAuth } from '@/hooks/playerMutateAuth'
import { ArrowPathIcon } from '@heroicons/react/20/solid'
import { useRouter } from 'next/router'

export const login = () => {
  const [email, setEmail] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const [isLogin, setIsLogin] = useState<boolean>(true)
  const { loginMutation, registerMutation } = playerMutateAuth()
  const router = useRouter()

  const submitAuthHandler = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (isLogin) {
      loginMutation.mutate({
        email: email,
        password: password,
      })
      setTimeout(() => {
        router.push('/main')
      }, 1000)
    } else {
      await registerMutation
        .mutateAsync({
          email: email,
          password: password,
        })
        .then(() => {
          loginMutation.mutate({
            email: email,
            password: password,
          })
          setTimeout(() => {
            router.push('/main')
          }, 2000)
        })
    }
  }
  return (
    <>
      <div className="hero min-h-screen bg-base-200">
        <Image src={backgroundImage} alt="card_game" fill />
        <div className="hero-content flex-col lg:flex-row-reverse">
          <div className="text-center lg:text-left">
            <h1 className="text-5xl font-bold">
              {isLogin ? 'Login now!' : 'Create account now!'}
            </h1>
            <p className="py-6">
              Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda
              excepturi exercitationem quasi. In deleniti eaque aut repudiandae
              et a id nisi.
            </p>
          </div>

          <div className="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100">
            <div className="card-body">
              <form onSubmit={submitAuthHandler}>
                <div className="form-control">
                  <label className="label">
                    <span className="label-text">Email</span>
                  </label>
                  <input
                    name="email"
                    type="email"
                    autoFocus
                    placeholder="email"
                    onChange={(e) => setEmail(e.target.value)}
                    value={email}
                    className="input input-bordered"
                  />
                </div>
                <div className="form-control">
                  <label className="label">
                    <span className="label-text">Password</span>
                  </label>
                  <input
                    name="password"
                    type="password"
                    placeholder="password"
                    onChange={(e) => setPassword(e.target.value)}
                    value={password}
                    className="input input-bordered"
                  />
                </div>
                <div className="form-control mt-6">
                  <button
                    className="btn btn-primary"
                    disabled={!email || !password}
                    type="submit"
                  >
                    {isLogin ? 'Login' : 'Create Account'}
                  </button>
                  <ArrowPathIcon
                    onClick={() => setIsLogin(!isLogin)}
                    className="h-6 w-6 my-2 text-blue-500 cursor-pointer"
                  />
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </>
  )
}

export default login
