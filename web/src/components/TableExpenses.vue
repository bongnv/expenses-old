<template>
  <v-data-table
    :headers="headers"
    :items="items"
    class="elevation-1 overflow-y-hidden"
    disable-sort
    disable-pagination
    hide-default-footer
    fill-height
  >
    <template v-slot:item.icon="{ item }">
      <v-avatar>
        <v-icon :class="[item.iconClass]" v-text="item.icon"></v-icon>
      </v-avatar>
    </template>
    >
  </v-data-table>
</template>

<script>
export default {
  data() {
    return {
      headers: [
        {
          text: "Category",
          value: "icon"
        },
        {
          text: "Amount",
          align: "start",
          sortable: false,
          value: "title"
        },
        { text: "Date", value: "subtitle" }
      ],
      items: [
        {
          icon: "mdi-folder",
          iconClass: "grey lighten-1 white--text",
          title: "S$1.25",
          subtitle: "Jan 9, 2014"
        },
        {
          icon: "mdi-folder",
          iconClass: "grey lighten-1 white--text",
          title: "S$2.25",
          category: "Category 2",
          subtitle: "Jan 17, 2014"
        },
        {
          icon: "mdi-folder",
          iconClass: "grey lighten-1 white--text",
          title: "S$5",
          category: "Category 3",
          subtitle: "Jan 28, 2014"
        },
        {
          icon: "mdi-food-fork-drink",
          iconClass: "blue white--text",
          title: "S$15",
          category: "Category 4",
          subtitle: "Jan 20, 2014"
        },
        {
          icon: "mdi-call_to_action",
          iconClass: "amber white--text",
          title: "S$25",
          category: "Category 5",
          subtitle: "Jan 10, 2014"
        },
        {
          icon: "mdi-folder",
          iconClass: "grey lighten-1 white--text",
          title: "S$1.25",
          subtitle: "For Restaurant on Jan 9, 2014"
        },
        {
          icon: "mdi-folder",
          iconClass: "grey lighten-1 white--text",
          title: "S$2.25",
          category: "Category 2",
          subtitle: "For Restaurant on Jan 17, 2014"
        },
        {
          icon: "mdi-folder",
          iconClass: "grey lighten-1 white--text",
          title: "S$5",
          category: "Category 3",
          subtitle: "For Restaurant on  Jan 28, 2014"
        },
        {
          icon: "mdi-food-fork-drink",
          iconClass: "blue white--text",
          title: "S$15",
          category: "Category 4",
          subtitle: "Jan 20, 2014"
        },
        {
          icon: "mdi-call_to_action",
          iconClass: "amber white--text",
          title: "S$25",
          category: "Category 5",
          subtitle: "Jan 10, 2014"
        },
        {
          icon: "mdi-folder",
          iconClass: "grey lighten-1 white--text",
          title: "S$2.25",
          category: "Category 2",
          subtitle: "For Restaurant on Jan 17, 2014"
        },
        {
          icon: "mdi-folder",
          iconClass: "grey lighten-1 white--text",
          title: "S$5",
          category: "Category 3",
          subtitle: "For Restaurant on  Jan 28, 2014"
        },
        {
          icon: "mdi-food-fork-drink",
          iconClass: "blue white--text",
          title: "S$15",
          category: "Category 4",
          subtitle: "Jan 20, 2014"
        },
        {
          icon: "mdi-call_to_action",
          iconClass: "amber white--text",
          title: "S$25",
          category: "Category 5",
          subtitle: "Jan 10, 2014"
        },
        {
          icon: "mdi-folder",
          iconClass: "grey lighten-1 white--text",
          title: "S$5",
          category: "Category 3",
          subtitle: "For Restaurant on  Jan 28, 2014"
        },
        {
          icon: "mdi-food-fork-drink",
          iconClass: "blue white--text",
          title: "S$15",
          category: "Category 4",
          subtitle: "Jan 20, 2014"
        },
        {
          icon: "mdi-call_to_action",
          iconClass: "amber white--text",
          title: "S$25",
          category: "Category 5",
          subtitle: "Jan 10, 2014"
        }
      ]
    };
  }
};
</script>
