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
          <router-link v-else :to="{ name: 'board', params: { id: board.id }}">
            {{ board.name }}
          </router-link>
        </v-list-item-content>
      </v-list-item>
    </v-list-item-group>
  </v-list>
</template>

<script>
export default {
  name: 'BoardList',
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
}
</script>
