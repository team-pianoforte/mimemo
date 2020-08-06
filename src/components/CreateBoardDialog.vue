<template>
  <v-dialog
    :value="value"
    max-width="300"
    @input="$emit('input', !value)"
  >
    <v-card>
      <v-card-title>ボードを追加</v-card-title>
      <v-card-text>
        <v-form>
          <v-text-field
            v-model="name"
            label="名前"
            autofocus
            required
          />
          <v-btn
            class="mt-8"
            block
            color="primary"
            @click="create"
          >
            作成
          </v-btn>
        </v-form>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: 'CreateBoardDialog',
  props: {
    value: {
      type: Boolean,
      required: true,
    },
  },
  data: () => ({
    name: '',
  }),
  methods: {
    create() {
      this.$emit('create', { name: this.name })
      this.$emit('input', false)
    },
  },
}
</script>
