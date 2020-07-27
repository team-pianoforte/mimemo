<template>
  <div>
    <MemoBoard :board="board" :memos="memos" @change="updateMemo" @done="removeMemo" />
    {{ memos }}
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
  </div>
</template>

<script>
import { db } from '@/firebase'
import MemoBoard from '@/components/MemoBoard.vue'

const memos = db.collection('memos')

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
  }),
  watch: {
    id() { this.rebind() },
  },
  async created() {
    this.rebind()
  },
  computed: {
    boardRef() {
      return db.collection('boards').doc(this.id)
    },
  },
  methods: {
    async clickCreateButton() {
      await this.createMemo({ text: 'メモ: ' })
    },
    async rebind() {
      await this.$bind('board', this.boardRef)
      await this.$bind('memos', memos.where('board', '==', this.boardRef))
    },
    async createMemo({ text }) {
      const ref = await memos.add({ text, board: this.boardRef })
      await ref.update({ id: ref.id })
    },
    async removeMemo(memo) {
      await memos.doc(memo.id).delete()
    },
    async updateMemo({ id, text }) {
      await memos.doc(id).update({ text })
    },
  },
}
</script>
