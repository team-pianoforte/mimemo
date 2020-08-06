export function mustFindIndexBy(prop, arr, obj) {
  const i = arr.findIndex((v) =>  v[prop] ===  obj[prop])
  if (i < 0) {
    throw new Error(`{ ${prop}: ${obj[prop]} } does not exist.`)
  }
  return i
}

export function handle404AsNull(err) {
  if (err.status === 404) {
    return null
  }
  throw err
}
