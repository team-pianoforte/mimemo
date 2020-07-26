<template>
  <v-app>
    <Toolbar />
    <v-main>
      <v-container class="content">
        <router-view />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import Toolbar from '@/components/Toolbar.vue'
export default {
  name: 'DefaultLayout',
  components: { Toolbar },
}
</script>

<style scoped>
.content {
  margin-top: 64px;
}
</style>
