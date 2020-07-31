<template>
  <v-app-bar
    app
    color="primary"
    dark
    clipped-left
  >
    <v-app-bar-nav-icon @click="$emit('switch-nav')" />

    <v-spacer />
    <v-toolbar-title>みめも</v-toolbar-title>
    <v-spacer />

  </v-app-bar>
</template>

<script>
export default {
  name: 'Toolbar',
}
</script>
