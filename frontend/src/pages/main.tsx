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

  return (
    <div className="flex justify-center">
      <div className="mockup-phone border-primary mt-20">
        <ArrowRightOnRectangleIcon
          onClick={logout}
          className="h-6 w-6 my-6 text-blue-500 cursor-pointer"
        />
        <div className="camera"></div>
        <div className="display">
          <div className="artboard artboard-demo phone-1">Hi.</div>
        </div>
      </div>
    </div>
  )
}

export default main
