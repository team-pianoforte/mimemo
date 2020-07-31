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
  </div>
</template>

<script>
import { Timestamp, db, auth } from '@/firebase'
import MemoBoard from '@/components/MemoBoard.vue'

export default {
  name: 'Board',
  components: { MemoBoard },
  props: {
    id: {
      type: String,
      required: true,
    },
  },
  data: () => ({
    board: null,
    memos: [],
    authenticated: false,
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
    auth.onAuthStateChanged(this.authenticate)
    this.bind()
  },
  methods: {
    async authenticate(user) {
      if ((await db.userBoardIds(user.uid)).includes(this.id)) {
        this.authenticated = true
      } else {
        this.authenticated = false
      }
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
