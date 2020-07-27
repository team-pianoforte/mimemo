<template>
  <v-row class="px-16" justify="center" align="center">
    <v-card max-width="800" class="flex-grow-1">
      <template v-for="(memo, i) in memos">
        <v-touch :key="memo.id" @swipeleft="done(memo)" @swiperight="done(memo)">
          <v-list-item>
            <v-list-item-content>
              <v-text-field
                hide-details solo flat
                v-model="memo.text"
                @focus="focusedId = memo.id"
                @blur="focusedId = null"
              />
            </v-list-item-content>
            <v-list-item-icon>
              <v-btn icon color="green" @click="done(memo)">
                <v-icon>mdi-check-bold</v-icon>
              </v-btn>
            </v-list-item-icon>
          </v-list-item>
          <v-divider v-if="i + 1 < memos.length" />
        </v-touch>
      </template>
    </v-card>
  </v-row>
</template>

<script>
export default {
  name: 'MemoBoard',
  data: () => ({
    memos: [{ id: 'a', text: 'a' }, { id: 'b', text: 'b' }],
    focusedId: null,
  }),
  methods: {
    done(memo) {
      this.memos = this.memos.filter((v) => v.id !== memo.id)
    },
  },
}
</script>
