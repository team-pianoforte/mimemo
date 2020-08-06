export const messages = [
  'おつかれさまでした',
  'ご苦労さまでした',
  '一歩前進ですね',
  'この少しの一歩が、大きな成果への道かもしれませんよ',
  '焦らず無理せずやっていきましょう',
  'こつこつ頑張るあなたはえらい！',
]

export function pickMessage() {
  return messages[Math.floor(Math.random() * messages.length)]
}
