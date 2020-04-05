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
        v-model="date"
        label="Date"
        prepend-icon="mdi-calendar"
        v-on="on"
        @keyup.enter="onEnter"
        :rules="rules"
      ></v-text-field>
    </template>
    <v-date-picker v-model="date" no-title scrollable />
  </v-menu>
</template>

<script>
const validateDate = v => v && v.length == 10 && !isNaN(Date.parse(v));
export default {
  data: () => ({
    menu: false,
    valid: false,
    rules: [v => validateDate(v) || "Date is invalid."]
  }),

  computed: {
    date: {
      get() {
        return this.value.toISOString().substr(0, 10);
      },

      set(newDate) {
        if (validateDate(newDate)) {
          this.$emit("input", new Date(newDate));
        }
      }
    }
  },

  props: {
    value: Date,
    default: new Date()
  },

  methods: {
    onEnter() {
      this.menu = false;
    }
  }
};
</script>
