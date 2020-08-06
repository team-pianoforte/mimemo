<template>
  <v-app>
    <Toolbar @switch-nav="nav = !nav" />
    <v-navigation-drawer
      v-model="nav"
      app
      clipped
    >
      <BoardMenu
        :id="$route.params.id"
        :boards="boards"
        @create-board="createBoard"
      />
    </v-navigation-drawer>
    <v-main>
      <v-container class="content">
        <router-view />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import Toolbar from '@/components/Toolbar.vue'
import BoardMenu from '@/components/BoardMenu.vue'
import authMiddleware from '@/middleware/auth'

/* eslint-disable no-unused-vars, vue/no-unused-properties */

export default {
  components: { Toolbar, BoardMenu },
  middleware: [ authMiddleware ],
  data: () => ({
    nav: false,
    uid: null,
  }),
  computed: {
    ...mapState('boards', {
      boards: 'list',
    }),
  },
  methods: {
    ...mapActions('boards', {
      createBoard: 'create',
    }),
  },
}
</script>

<style scoped>
.content {
  margin-top: 64px;
}
</style>
