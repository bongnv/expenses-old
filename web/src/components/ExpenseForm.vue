<template>
  <v-container fill-width>
    <v-form ref="form" v-model="valid" :lazy-validation="lazy" class="justify-center">
      <DatePicker v-model="expense.date" />

      <SimpleDropdown v-model="expense.currency" name="Currency" :items="currencies"></SimpleDropdown>
      <SimpleDropdown v-model="expense.category" name="Category" :items="categories"></SimpleDropdown>
      <v-text-field
        v-model.number="expense.amount"
        :counter="amountLength"
        :rules="amountRules"
        label="Amount"
        required
        placeholder="0"
        type="number"
        @keyup.enter="submit(expense)"
      ></v-text-field>

      <v-text-field
        v-model="expense.note"
        :counter="noteLength"
        :rules="noteRules"
        label="Note"
        placeholder="A note on the expense."
      ></v-text-field>

      <v-row>
        <v-spacer />
        <v-btn :disabled="!valid" class="ma-4" color="success" @click="submit(expense)">Submit</v-btn>
        <v-spacer />
        <v-btn color="error" class="ma-4" @click="reset">Reset Form</v-btn>
        <v-spacer />
      </v-row>
    </v-form>
  </v-container>
</template>

<script>
import SimpleDropdown from "@/components/SimpleDropdown";
import DatePicker from "@/components/DatePicker";
import categories from "@/data/categories.json";
import currencies from "@/data/currencies.json";
import { mapActions } from "vuex";

// data, amount, category, account, note, currency, category group
export default {
  data: vm => ({
    expense: {
      date: new Date(),
      currency: "SGD"
    },
    valid: true,
    noteLength: 64,
    noteRules: [
      v =>
        !v ||
        v.length <= vm.noteLength ||
        "Note must not be more than " + vm.noteLength + " characters"
    ],
    amountLength: 20,
    amountRules: [
      v => !!v || "Amount is required",
      v => v > 0 || "Amount must be valid"
    ],
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
    submit(expense) {
      this.createExpense(expense).then(() => {
        this.listExpenses();
        this.reset();
      });
    },
    ...mapActions("expenses", ["createExpense", "listExpenses"])
  }
};
</script>
