<template>
  <v-dialog
    :value="value"
    @input="$emit('input', !dialog)"
    max-width="300"
  >
    <v-card>
      <v-card-title>ボードを追加</v-card-title>
      <v-card-text>
        <v-form>
          <v-text-field
            label="名前"
            autofocus
            v-model="name"
            required
          />
          <PasswordField
            label="共有用パスワード（任意）"
            v-model="password"
            counter
            hint="共有する用パスワードです"
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
import PasswordField from './PasswordField'

export default {
  name: 'CreateBoardDialog',
  components: { PasswordField },
  props: {
    value: {
      type: Boolean,
      required: true,
    },
  },
  data: () => ({
    name: '',
    password: '',
  }),
  methods: {
    create() {
      this.$emit('create', { name: this.name, password: this.password })
      this.$emit('input', false)
    },
  },
}
</script>
