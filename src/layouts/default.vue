<template>
  <div>
    <Toolbar />
    <div class="content">
      h
      <router-view />
    </div>
  </div>
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
