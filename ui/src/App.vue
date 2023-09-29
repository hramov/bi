<script setup lang="ts">
import Layout from "./views/layout/Layout.vue";
import {useDatasourceStore} from "./modules/store/datasource.store.ts";
import {onMounted} from "vue";

const store = useDatasourceStore();

onMounted(() => {
  store.getSources()
  store.getDrivers()
});
</script>

<template>
  <Layout />
</template>

<style scoped>
</style>
