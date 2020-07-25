<template>
  <div class="home">
    HOME

    <div v-if="loggedIn">
      {{ user }}
    </div>
    <div v-else>
      logging in
    </div>
  </div>
</template>

<script>
import { mapActions, mapGetters, mapState } from 'vuex'
export default {
  name: 'Home',
  methods: {
    ...mapActions('auth', ['signInAnonymously']),
  },
  computed: {
    ...mapGetters('auth', ['loggedIn']),
    ...mapState('auth', ['user']),
  },
  async mounted() {
    await this.signInAnonymously()
  },
}
</script>
