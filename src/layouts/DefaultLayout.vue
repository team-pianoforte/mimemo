<template>
  <v-app>
    <Toolbar @switch-nav="nav = !nav" />
    <v-navigation-drawer
      v-model="nav"
      app
      clipped
    >
      <BoardMenu :id="$route.params.id" :boards="boards" @create-board="createBoard" />
    </v-navigation-drawer>
    <v-main>
      <v-container class="content">
        <router-view />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import { mapState } from 'vuex'
import { firestore, db } from '@/firebase'
import { hashPassword } from '@/lib/passwords'
import Toolbar from '@/components/Toolbar.vue'
import BoardMenu from '@/components/BoardMenu.vue'

export default {
  name: 'DefaultLayout',
  components: { Toolbar, BoardMenu },
  data: () => ({
    nav: false,
    boards: [],
  }),
  computed: {
    ...mapState('auth', ['user']),
  },
  created() {
    this.$bind('boards', db.boards)
  },
  methods: {
    async createBoard({ name, password }) {
      const passwordHash = hashPassword(password)
      const { boardId } = await db.boards.add({ name, passwordHash })
      db.user(this.user.uid).update(firestore.FieldValue.arrayUnion(boardId))
    },
  },
}
</script>

<style scoped>
.content {
  margin-top: 64px;
}
</style>
