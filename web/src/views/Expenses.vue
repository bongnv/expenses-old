<template>
  <MainLayout>
    <v-container id="expenses" class="fill-height" fluid>
      <v-row class="fill-height" no-gutters>
        <v-col>
          <ExpenseForm style="width: 90%" />
        </v-col>
        <v-col v-if="$vuetify.breakpoint.lgAndUp" class="fill-height">
          <TableExpenses class="fill-height" :items="expenses" />
        </v-col>
      </v-row>
    </v-container>
  </MainLayout>
</template>

<script>
// @ is an alias to /src
import MainLayout from "@/layouts/MainLayout";
import ExpenseForm from "@/components/ExpenseForm";
// import ListTransactions from "@/components/ListTransactions.vue";
import TableExpenses from "@/components/TableExpenses";
import { mapActions } from "vuex";

export default {
  name: "Home",
  computed: {
    expenses() {
      return this.$store.state.expenses.expenses;
    }
  },

  components: {
    MainLayout,
    ExpenseForm,
    TableExpenses
  },
  methods: {
    ...mapActions("expenses", ["listExpenses"])
  },

  mounted: function() {
    this.listExpenses();
  }
};
</script>
