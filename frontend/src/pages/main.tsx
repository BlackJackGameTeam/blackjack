import { playerMutateAuth } from '@/hooks/playerMutateAuth'
import React, { useRef, useEffect } from 'react'
import {
  ArrowRightOnRectangleIcon,
  ShieldCheckIcon,
} from '@heroicons/react/20/solid'
import { useRouter } from 'next/router'

export const main = () => {
  const { logoutMutation } = playerMutateAuth()
  const router = useRouter()

  const logout = async () => {
    await logoutMutation.mutateAsync()
    setTimeout(() => {
      router.push('/login')
    }, 1000)
  }

  const handlePageInBlackjack = () => {
    router.push('/games/blackjack')
  }

  return (
    <div className="bg-blue-400 py-48">
      <div className="card w-96 bg-base-100 shadow-xl mx-10">
        <figure>
          <img src="/assets/background.jpg" alt="backGround" />
        </figure>
        <div className="card-body">
          <h2 className="card-title">BlackjackGame</h2>
          <p>Let`s Play together?</p>
          <div className="card-actions justify-end">
            <button className="btn btn-primary" onClick={handlePageInBlackjack}>
              Start Now
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default main
