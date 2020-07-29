<template>
  <v-list>
    <v-subheader>Boards</v-subheader>
    <v-list-item-group>
      <v-list-item
        v-for="board in boards"
        :key="board.id"
        nav
      >
        <v-list-item-content>
          <span v-if="board.id === id">
            {{ board.name }}
          </span>
          <router-link
            v-else
            :to="{ name: 'board', params: { id: board.id }}"
          >
            {{ board.name }}
          </router-link>
        </v-list-item-content>
      </v-list-item>
    </v-list-item-group>
    <v-divider />
    <v-list-item @click="dialog = true">
      <v-list-item-action>
        <v-icon>mdi-plus</v-icon>
      </v-list-item-action>
      <v-list-item-content>
        ボードを追加
      </v-list-item-content>
    </v-list-item>
    <v-dialog
      v-model="dialog"
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
  </v-list>
</template>

<script>
export default {
  name: 'BoardMenu',
  props: {
    id: {
      type: String,
      default: '',
    },
    boards: {
      type: Array,
      required: true,
    },
  },
  data: () => ({
    dialog: false,
    name: '',
    password: '',
    showPassword: false,
  }),
  methods: {
    passwordRule(v){
      v.length >= 8 || '8文字以上にしてください'
    },
  },
}
</script>
