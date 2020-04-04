---
to: "src/layouts/<%= h.changeCase.pascalCase(name) %>.vue"
---
<template>
  <v-app>
    <v-content>
      <slot />
    </v-content>
  </v-app>
</template>
