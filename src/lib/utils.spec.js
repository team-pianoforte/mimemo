import { mustFindIndexBy } from './utils'

describe('mustFindIndex', () => {
  const arr = [
    { a: 0 },
    { a: 1 },
  ]
  it('returns index if exists', () => {
    expect(mustFindIndexBy('a', arr, arr[0])).toEqual(0)
    expect(mustFindIndexBy('a', arr, arr[1])).toEqual(1)

    expect(mustFindIndexBy('b', arr, {})).toEqual(0)
  })

  it('throws error if not exists', () => {
    expect(() => { mustFindIndexBy('a', arr, { a: -1 }) }).toThrow()
    expect(() => { mustFindIndexBy('a', arr, {}) }).toThrow()
  })
})
