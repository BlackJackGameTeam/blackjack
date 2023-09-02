import axios from 'axios'
import React from 'react'

export const fetcher = (url: string) => axios.get(url).then(res => res.data)