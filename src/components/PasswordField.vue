<template>
  <v-text-field
    :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
    :type="showPassword ? 'text' : 'password'"
    :value="value"
    @input="$emit('input', value)"
    @click:append="showPassword = !showPassword"
    v-bind="$attrs"
  />
</template>

<script>
export default {
  name: 'PasswordField',
  props: {
    value: {
      type: String,
      default: '',
    },
  },
  data: () => ({
    showPassword: false,
  }),
}
</script>
