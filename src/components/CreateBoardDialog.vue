<template>
  <v-dialog
    :value="dialog"
    @input="input"
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
            :rules="[passwordRule]"
            :type="showPassword ? 'text' : 'password'"
            hint="共有する用パスワードです"
            @click:append="showPassword = !showPassword"
          />
          <v-btn
            class="mt-8"
            block
            color="primary"
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
    dialog: {
      type: Boolean,
      required: true,
    },
  },
  model: {
    prop: 'dialog',
    event: 'input',
  },
  data: () => ({
    name: '',
    password: '',
    showPassword: false,
  }),
  methods: {
    passwordRule(v){
      v.length >= 8 || '8文字以上にしてください'
    },
    input(v) {
      console.log('input', v)
    },
  },
}
</script>
