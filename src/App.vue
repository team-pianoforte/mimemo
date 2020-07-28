<template>
  <router-view />
</template>

<script>
export default {
  name: 'App',
  created() {
    this.$store.dispatch('init')
    this.$store.dispatch('auth/signInAnonymously')
  },
}
</script>
