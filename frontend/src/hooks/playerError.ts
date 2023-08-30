import React from 'react'
import { getCsrfToken } from '@/pages/api/player'

export const playerError = () => {
    const switchErrorHandling = (msg: string) => {
        switch (msg) {
          case 'invalid csrf token':
            getCsrfToken()
            alert('CSRF token is invalid, please try again')
            break
          case 'invalid or expired jwt':
            alert('access token expired, please login')
            break
          case 'missing or malformed jwt':
            alert('access token is not valid, please login')
            break
          case 'duplicated key not allowed':
            alert('email already exist, please use another one')
            break
          case 'crypto/bcrypt: hashedPassword is not the hash of the given password':
            alert('password is not correct')
            break
          case 'record not found':
            alert('email is not correct')
            break
          default:
            alert(msg)
        }
      }
      return { switchErrorHandling }
}
