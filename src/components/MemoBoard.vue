<template>
  <v-row
    class="px-16"
    justify="center"
    align="center"
  >
    <v-card
      max-width="800"
      class="flex-grow-1"
    >
      <template v-for="(memo, i) in memos">
        <v-touch
          :key="memo.id"
          @swipeleft="done(memo)"
          @swiperight="done(memo)"
        >
          <v-list-item>
            <v-list-item-content>
              <v-text-field
                :value="memo.text"
                hide-details
                solo
                flat
                @input="change(memo, $event)"
                @focus="focusedId = memo.id"
                @blur="focusedId = null"
              />
            </v-list-item-content>
            <v-list-item-icon>
              <v-btn
                icon
                color="green"
                @click="done(memo)"
              >
                <v-icon>mdi-check-bold</v-icon>
              </v-btn>
            </v-list-item-icon>
          </v-list-item>
          <v-divider v-if="i + 1 < memos.length" />
        </v-touch>
      </template>
      <v-snackbar v-model="snackbar">
        <p class="text-center ma-0">
          {{ message }}
        </p>
      </v-snackbar>
    </v-card>
  </v-row>
</template>

<script>
import { pickMessage } from '@/messages'

export default {
  name: 'MemoBoard',
  props: {
    memos: {
      type: Array,
      required: true,
    },
  },
  data: () => ({
    focusedId: null,
    snackbar: false,
    message: '',
  }),
  methods: {
    done(memo) {
      this.$emit('done', memo)
      this.showMessage()
    },
    change(memo, text) {
      this.$emit('change', { ...memo, text })
    },
    showMessage() {
      this.message = pickMessage()
      this.snackbar = true
    },
  },
}
</script>
