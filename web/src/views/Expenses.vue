<template>
  <v-container id="expenses" class="fill-height" justify="center" fluid>
    <v-row>
      <v-col>
        <ExpenseForm />
      </v-col>
      <v-col>
        <v-alert type="info">This is under construction.</v-alert>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
// @ is an alias to /src
import ExpenseForm from "@/components/ExpenseForm.vue";

export default {
  name: "Home",
  components: {
    ExpenseForm
  }
};
</script>