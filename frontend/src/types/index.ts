export type History = {
  id: number
  win: number
  lose: number
  money: number
  created_at: Date
  updated_at: Date
}

export type CsrfToken = {
  csrf_token: string
}

export type Credential = {
  email: string
  password: string
}
