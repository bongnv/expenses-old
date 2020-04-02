<template>
  <v-autocomplete
    :value="value"
    :items="items"
    :rules="rules"
    :label="name"
    required
    @input="$emit('input', $event)"
  ></v-autocomplete>
</template>

<script>
export default {
  data: vm => ({
    rules: [v => !!v || vm.name + " is required"]
  }),

  props: {
    name: String,
    items: Array,
    value: String
  }
};
</script>
