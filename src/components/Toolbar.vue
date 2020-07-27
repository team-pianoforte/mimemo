<template>
  <v-app-bar
    app
    color="primary"
    dark
  >
    <v-spacer />
    <v-toolbar-title>みめも</v-toolbar-title>
    <v-spacer />

    <v-btn
      v-if="loggedIn"
      icon
    >
      <v-icon>mdi-account</v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'Toolbar',
  computed: {
    ...mapGetters('auth', ['loggedIn']),
  },
}
</script>
