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
import { db, auth } from '@/firebase'
import { hashPassword } from '@/lib/passwords'
import Toolbar from '@/components/Toolbar.vue'
import BoardMenu from '@/components/BoardMenu.vue'

export default {
  name: 'DefaultLayout',
  components: { Toolbar, BoardMenu },
  data: () => ({
    nav: false,
    boards: [],
    uid: null,
  }),
  async created() {
    auth.onAuthStateChanged(async (user) => {
      this.userChanged(user)
    })

    console.log(auth.currentUser)
    await auth.signInAnonymously()
  },
  methods: {
    async userChanged(user) {
      this.uid = user.uid
      this.$bind('boards', await db.userBoards(user.uid))
    },
    async createBoard({ name, password }) {
      const passwordHash = hashPassword(password)
      const ref = await db.boards.add({ name, passwordHash })
      const id = (await ref.get()).id
      await db.boards.doc(id).update({ id })
      await db.addBoardToUser(this.uid, id)
      await this.userChanged(auth.currentUser)
    },
  },
}
</script>

<style scoped>
.content {
  margin-top: 64px;
}
</style>
