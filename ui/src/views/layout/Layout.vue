<script setup lang="ts">
import { ref } from 'vue'
import Sidebar from "./components/Sidebar.vue";
import Avatar from "./components/Avatar.vue";

const drawer = ref(null);

</script>

<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer">
      <Avatar />
      <v-divider></v-divider>
      <Sidebar />
    </v-navigation-drawer>

   <router-view @drawer="drawer = !drawer"></router-view>

  </v-app>
</template>

<style scoped>
</style>