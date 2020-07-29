import * as firebase from 'firebase/app'

import 'firebase/analytics'
import 'firebase/auth'
import 'firebase/firestore'

const env = Object.fromEntries(
  Object.entries(process.env)
    .filter(([k]) => k.includes('FIREBASE'))
    .map(([k, v]) => [k.replace('VUE_APP_FIREBASE_', ''), v]),
)
const config = {
  apiKey: env.API_KEY,
  authDomain: env.AUTH_ADMIN,
  databaseURL: env.DATABASE_URL,
  projectId: env.PROJECT_ID,
  storageBucket: env.STORAGE_BUCKET,
  messagingSenderId: env.MESSAGING_SENDER_ID,
  appId: env.APP_ID,
  measurementId: env.MEASUREMENT_ID,
}

firebase.initializeApp(config)
firebase.analytics()

export const firestore = firebase.firestore()
export const auth = firebase.auth()
export const Timestamp = firebase.firestore.Timestamp


const boards = firestore.collection('boards')
const memos = (boardRef) => boardRef.collection('memos')
const users = firestore.collection('users')
const user = (uid) => users.doc(uid)

export const db = {
  boards,
  memos,
  users,
  user,
}
