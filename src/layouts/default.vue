<template>
  <div>
    <Toolbar />
    <router-view />
  </div>
</template>

<script>
import Toolbar from '@/components/Tollbar.vue'
export default {
  name: 'DefaultLayout',
  components: { Toolbar },
}
</script>
