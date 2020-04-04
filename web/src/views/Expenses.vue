<template>
  <MainLayout>
    <v-container id="expenses" class="fill-height" fluid>
      <v-row class="fill-height" no-gutters>
        <v-col>
          <ExpenseForm class="expense-form" style="width: 90%" />
        </v-col>
        <v-col v-if="$vuetify.breakpoint.lgAndUp" class="fill-height">
          <TableExpenses class="fill-height" />
        </v-col>
      </v-row>
    </v-container>
  </MainLayout>
</template>

<script>
// @ is an alias to /src
import MainLayout from "@/layouts/MainLayout.vue";
import ExpenseForm from "@/components/ExpenseForm.vue";
// import ListTransactions from "@/components/ListTransactions.vue";
import TableExpenses from "@/components/TableExpenses.vue";

export default {
  name: "Home",
  components: {
    MainLayout,
    ExpenseForm,
    TableExpenses
  }
};
</script>
