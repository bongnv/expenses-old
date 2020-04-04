<template>
  <v-container fill-width>
    <v-form ref="form" v-model="valid" :lazy-validation="lazy" class="justify-center">
      <DatePicker v-model="date" />

      <SimpleDropdown v-model="currency" name="Currency" :items="currencies"></SimpleDropdown>
      <SimpleDropdown v-model="category" name="Category" :items="categories"></SimpleDropdown>
      <v-text-field
        v-model="amount"
        :counter="amountLength"
        :rules="amountRules"
        label="Amount"
        required
        placeholder="0"
        type="number"
      ></v-text-field>

      <v-text-field
        v-model="note"
        :counter="noteLength"
        :rules="noteRules"
        label="Note"
        placeholder="A note on the expense."
      ></v-text-field>

      <v-row>
        <v-spacer />
        <v-btn :disabled="!valid" class="ma-4" color="success" @click="submit">Submit</v-btn>
        <v-spacer />
        <v-btn color="error" class="ma-4" @click="reset">Reset Form</v-btn>
        <v-spacer />
      </v-row>
    </v-form>
  </v-container>
</template>

<script>
import SimpleDropdown from "@/components/SimpleDropdown.vue";
import DatePicker from "@/components/DatePicker.vue";
import categories from "@/data/categories.json";
import currencies from "@/data/currencies.json";

// data, amount, category, account, note, currency, category group
export default {
  data: vm => ({
    date: new Date().toISOString().substr(0, 10),
    valid: true,
    noteLength: 64,
    note: "",
    noteRules: [
      v =>
        !v ||
        v.length <= vm.noteLength ||
        "Note must not be more than " + vm.noteLength + " characters"
    ],
    amount: undefined,
    amountLength: 20,
    amountRules: [
      v => !!v || "Amount is required",
      v => v > 0 || "Amount must be valid"
    ],
    category: "",
    currency: "SGD",
    categories,
    currencies,
    lazy: false
  }),

  components: {
    SimpleDropdown,
    DatePicker
  },

  methods: {
    reset() {
      this.$refs.form.reset();
    },
    submit() {
      this.$refs.form.resetValidation();
    }
  }
};
</script>
