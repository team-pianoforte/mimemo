<template>
  <div>
    <MemoBoard />
    <v-btn fixed fab right bottom color="primary" class="mb-8 mr-8">
      <v-icon>mdi-clipboard-plus-outline</v-icon>
    </v-btn>
  </div>
</template>

<script>
import MemoBoard from '@/components/MemoBoard.vue'

export default {
  name: 'Home',
  components: { MemoBoard },
}
</script>
