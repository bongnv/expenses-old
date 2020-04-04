<template>
  <MainLayout>
    <div class="home">
      <img alt="Vue logo" src="../assets/logo.png" />
      <HelloWorld msg="Welcome to Your Vue.js App" />
    </div>
  </MainLayout>
</template>

<script>
// @ is an alias to /src
import MainLayout from "@/layouts/MainLayout.vue";
import HelloWorld from "@/components/HelloWorld.vue";

export default {
  name: "Home",
  components: {
    HelloWorld,
    MainLayout
  }
};
</script>
