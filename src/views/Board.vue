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
    memosRef() {
      return this.boardRef.collection('memos')
    },
  },
  methods: {
    async clickCreateButton() {
      await this.createMemo({ text: 'メモ: ' })
    },
    async rebind() {
      await this.$bind('board', this.boardRef)
      await this.$bind('memos', this.memosRef)
    },
    async createMemo({ text }) {
      const ref = await this.memosRef.add({ text })
      await ref.update({ id: ref.id })
    },
    async removeMemo(memo) {
      await this.memosRef.doc(memo.id).delete()
    },
    async updateMemo({ id, text }) {
      await this.memosRef.doc(id).update({ text })
    },
  },
}
</script>
