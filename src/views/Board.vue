<template>
  <div>
    <template v-if="authenticated">
      <MemoBoard
        :board="board"
        :memos="memos"
        @change="updateMemo"
        @done="removeMemo"
      />
      <v-btn
        fixed
        fab
        right
        bottom
        color="primary"
        class="mb-8 mr-8"
        @click="clickCreateButton"
      >
        <v-icon>mdi-clipboard-plus-outline</v-icon>
      </v-btn>
    </template>
    <v-dialog
      v-model="passwordDialog"
      max-width="300"
    >
      <v-card>
        <v-card-title>
          ボードに参加
        </v-card-title>
        <v-card-text>
          <PasswordField
            label="パスワード"
            hint="ボードのパスワードを入力してください"
            v-model="password"
          />
          <v-btn
            block
            color="primary"
            @click="joinToBoard"
          >
            参加する
          </v-btn>
        </v-card-text>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import { Timestamp, db, auth } from '@/firebase'
import PasswordField from '@/components/PasswordField.vue'
import MemoBoard from '@/components/MemoBoard.vue'

/* eslint-disable */

export default {
  name: 'Board',
  components: { MemoBoard, PasswordField },
  props: {
    id: {
      type: String,
      required: true,
    },
  },
  data: () => ({
    user: null,
    board: null,
    memos: [],
    authenticated: undefined,
    passwordDialog: false,
    password: '',
  }),
  computed: {
    boardRef() {
      return db.boards.doc(this.id)
    },
    memosRef() {
      return db.memos(this.boardRef)
    },
  },
  async created() {
    this.bind()
    auth.onAuthStateChanged(this.authenticate)
  },
  methods: {
    async authenticate(user) {
      await this.$bind('user', db.user(user.uid))
      if ((await db.userBoardIds(user.uid)).includes(this.id)) {
        this.authenticated = true
        this.passwordDialog = false
      } else {
        this.authenticated = false
        this.passwordDialog = true
      }
    },
    async joinToBoard(password) {
      this.user.update(firestore.FieldValue.arrayUnion(boardId))
    },
    async clickCreateButton() {
      await this.createMemo({ text: 'メモ: ' })
    },
    async bind() {
      await this.$bind('board', this.boardRef)
      await this.$bind('memos', this.memosRef)
    },
    async createMemo({ text }) {
      const ref = await this.memosRef.add({
        text,
        createdAt: Timestamp.now(),
        updatedAt: Timestamp.now(),
      })
      await ref.update({ id: ref.id })
    },
    async removeMemo(memo) {
      await this.memosRef.doc(memo.id).delete()
    },
    async updateMemo({ id, text }) {
      await this.memosRef.doc(id).update({
        text,
        updatedAt: Timestamp.now(),
      })
    },
  },
}
</script>
