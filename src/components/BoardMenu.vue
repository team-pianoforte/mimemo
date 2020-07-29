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
    <v-list-item @click="createBoardDialog = true">
      <v-list-item-action>
        <v-icon>mdi-plus</v-icon>
      </v-list-item-action>
      <v-list-item-content>
        ボードを追加
      </v-list-item-content>
    </v-list-item>
    <CreateBoardDialog v-model="createBoardDialog" />
  </v-list>
</template>

<script>
import CreateBoardDialog from './CreateBoardDialog.vue'

export default {
  name: 'BoardMenu',
  components: { CreateBoardDialog },
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
    createBoardDialog: false,
  }),
}
</script>
