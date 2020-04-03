<template>
  <v-menu
    ref="menu"
    v-model="menu"
    close-on-content-click
    transition="scale-transition"
    offset-y
    max-width="290px"
    min-width="290px"
  >
    <template v-slot:activator="{ on }">
      <v-text-field
        :value="value"
        @input="update"
        label="Month"
        prepend-icon="mdi-calendar"
        v-on="on"
        @keyup.enter="onEnter"
      ></v-text-field>
    </template>
    <v-date-picker :value="value" @input="update" no-title scrollable />
  </v-menu>
</template>

<script>
export default {
  data: () => ({
    menu: false,
    modal: false
  }),

  props: {
    value: String,
    default: new Date().toISOString().substr(0, 10)
  },

  methods: {
    onEnter() {
      this.menu = false;
    },
    update(newValue) {
      this.$emit("input", newValue);
    }
  }
};
</script>