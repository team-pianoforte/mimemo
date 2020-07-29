import sha256 from 'sha256'

const stretchCount = 10

export function hashPassword(password) {
  if (password === '') {
    return ''
  }
  let hash = password
  for (let i = 0; i < stretchCount; i++) {
    hash = sha256(password)
  }
  return hash
}

export function verifyPassword(password, hash) {
  if (hash === '') {
    return true
  }
  return hashPassword(hash) === hash
}
