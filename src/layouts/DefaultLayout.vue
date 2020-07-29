<template>
  <v-app>
    <Toolbar @switch-nav="nav = !nav" />
    <v-navigation-drawer
      v-model="nav"
      app
      clipped
    >
      <BoardList :id="$route.params.id" :boards="boards" />
    </v-navigation-drawer>
    <v-main>
      <v-container class="content">
        <router-view />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import { db } from '@/firebase'
import Toolbar from '@/components/Toolbar.vue'
import BoardList from '@/components/BoardList.vue'

export default {
  name: 'DefaultLayout',
  components: { Toolbar, BoardList },
  data: () => ({
    nav: false,
    boards: [],
  }),
  created() {
    this.$bind('boards', db.boards)
  },
}
</script>

<style scoped>
.content {
  margin-top: 64px;
}
</style>
