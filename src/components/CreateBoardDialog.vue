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
          <v-text-field
            label="共有用パスワード（任意）"
            v-model="password"
            counter
            :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
            :type="showPassword ? 'text' : 'password'"
            hint="共有する用パスワードです"
            @click:append="showPassword = !showPassword"
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
    password: '',
    showPassword: false,
  }),
  methods: {
    create() {
      this.$emit('create', { name: this.name, password: this.password })
      this.$emit('input', false)
    },
  },
}
</script>
