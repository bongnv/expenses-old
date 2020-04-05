<template>
  <v-data-table
    :headers="headers"
    :items="computedItems"
    class="elevation-1 overflow-y-hidden"
    disable-sort
    disable-pagination
    hide-default-footer
    fill-height
  >
    <template v-slot:item.category="{ item }">
      <v-avatar>
        <v-icon :class="[item.iconClass]" v-text="item.icon"></v-icon>
      </v-avatar>
    </template>
    <template v-slot:item.actions="{ item }">
      <v-icon small class="mr-2" @click="editItem(item)">mdi-pencil</v-icon>
      <v-icon small @click="deleteItem(item)">mdi-delete</v-icon>
    </template>
  </v-data-table>
</template>

<script>
import categories from "@/data/categories";
import { mapActions } from "vuex";

export default {
  data: () => ({
    headers: [
      {
        text: "Category",
        value: "category"
      },
      {
        text: "Amount",
        align: "start",
        sortable: false,
        value: "amount"
      },
      { text: "Date", value: "date" },
      { text: "Actions", value: "actions" }
    ]
  }),

  computed: {
    computedItems() {
      const iconMap = categories.reduce(function(map, obj) {
        map[obj.value] = obj.icon;
        return map;
      }, {});

      const months = [
        "Jan",
        "Feb",
        "Mar",
        "Apr",
        "May",
        "Jun",
        "Jul",
        "Aug",
        "Sep",
        "Oct",
        "Nov",
        "Dec"
      ];
      const getDate = date => {
        const obj = new Date(date);
        return (
          obj.getDate() + "-" + months[obj.getMonth()] + "-" + obj.getFullYear()
        );
      };
      const formatAmount = item => item.currency + " " + item.amount;

      return this.items.filter(Boolean).map(item => ({
        ...item,
        amount: formatAmount(item),
        icon: iconMap[item.category],
        date: getDate(item.date)
      }));
    }
  },

  methods: {
    editItem: function() {},
    deleteItem(item) {
      this.deleteExpense(item.id).then(() => this.listExpenses());
    },
    ...mapActions("expenses", ["deleteExpense", "listExpenses"])
  },

  props: {
    items: {
      type: Array,
      default: () => []
    }
  }
};
</script>
