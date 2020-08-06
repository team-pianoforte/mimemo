<template>
  <div>
    <MemoBoard
      :board="board"
      :memos="memos"
      @change="updateMemo"
      @complete="completeMemo"
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
  </div>
</template>

<script>
import { mapState } from 'vuex'
import MemoBoard from '@/components/MemoBoard.vue'

export default {
  components: { MemoBoard },
  async fetch({ store, params, error }) {
    await store.dispatch('boards/fetchAll')
    await store.dispatch('boards/setCurrent', params.id)
    if (!store.state.boards.current) {
      error({ errorCode: 404 })
    }
  },
  data: () => ({
    memos: [],
  }),
  computed: {
    ...mapState('boards', {
      board: 'current',
      boards: 'list',
    }),
  },
  methods: {
    async completeMemo(memo) {
      await this.deleteMemo(memo)
    },
    async clickCreateButton() {
      await this.createMemo({ text: 'メモ: ' })
    },
    async createMemo() {
    },
    async updateMemo() {
    },
  },
}
</script>
