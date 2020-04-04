<template>
  <v-card max-width="500px" class="d-flex flex-column ma-auto justify-center">
    <v-toolbar color="indigo" dark>
      <v-app-bar-nav-icon></v-app-bar-nav-icon>

      <v-toolbar-title>Expenses</v-toolbar-title>

      <v-spacer></v-spacer>

      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-dots-vertical</v-icon>
      </v-btn>
    </v-toolbar>

    <v-list class="d-flex flex-column overflow-y-auto">
      <v-list-item
        v-for="item in items"
        :key="item.title"
        @click="onClick(item)"
      >
        <v-list-item-avatar>
          <v-icon :class="[item.iconClass]" v-text="item.icon"></v-icon>
        </v-list-item-avatar>

        <v-list-item-content>
          <v-list-item-title v-text="item.title"></v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>
  </v-card>
</template>

<script>
export default {
  data: () => ({
    items: [
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
        title: "S$15 on Jan 20, 2014",
        category: "Category 4",
        subtitle: "Jan 20, 2014"
      },
      {
        icon: "mdi-call_to_action",
        iconClass: "amber white--text",
        title: "S$25 for Restaurant",
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
        title: "S$15 on Jan 20, 2014",
        category: "Category 4",
        subtitle: "Jan 20, 2014"
      },
      {
        icon: "mdi-call_to_action",
        iconClass: "amber white--text",
        title: "S$25 for Restaurant",
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
        title: "S$15 on Jan 20, 2014",
        category: "Category 4",
        subtitle: "Jan 20, 2014"
      },
      {
        icon: "mdi-call_to_action",
        iconClass: "amber white--text",
        title: "S$25 for Restaurant",
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
        title: "S$15 on Jan 20, 2014",
        category: "Category 4",
        subtitle: "Jan 20, 2014"
      },
      {
        icon: "mdi-call_to_action",
        iconClass: "amber white--text",
        title: "S$25 for Restaurant",
        category: "Category 5",
        subtitle: "Jan 10, 2014"
      }
    ]
  }),
  methods: {
    onClick: function(item) {
      console.log("the button is clicked", item);
    }
  }
};
</script>
