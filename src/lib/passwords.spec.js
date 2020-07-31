import { hashPassword, verifyPassword } from './passwords'

describe('hashPassword', () => {
  it('returns password hash', () => {
    const hash = hashPassword('abc')
    expect(hash).toHaveLength(64)
    expect(typeof hash).toBe('string')
  })

  it('returns null when password is empty', () => {
    expect(hashPassword('')).toBe('')
    expect(hashPassword(undefined)).toBe('')
  })
})

describe('verifyPassword', () => {
  const password = 'abc'
  const hash = hashPassword(password)
  it('returns true when password is valid', () => {
    expect(verifyPassword(password, hash)).toBeTruthy()
  })

  it('returns true when password is invalid', () => {
    expect(verifyPassword('cdf', hash)).toBeFalsy()
  })

  it('returns hc when password and hash are empty', () => {
    expect(verifyPassword('', '')).toBeTruthy()
    expect(verifyPassword(undefined, '')).toBeTruthy()
    expect(verifyPassword('', undefined)).toBeTruthy()
  })
})
