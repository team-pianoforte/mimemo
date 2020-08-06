<template>
  <v-text-field
    :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
    :type="showPassword ? 'text' : 'password'"
    :value="value"
    v-bind="$attrs"
    @input="$emit('input', $event)"
    @click:append="showPassword = !showPassword"
  />
</template>

<script>
export default {
  name: 'PasswordField',
  props: {
    value: {
      type: String,
      required: true,
    },
  },
  data: () => ({
    showPassword: false,
  }),
}
</script>
