<script setup lang="ts">
import Layout from "./views/layout/Layout.vue";
import {useDatasourceStore} from "./modules/store/datasource.store.ts";
import {onMounted} from "vue";
import {useDashboardStore} from "./modules/store/dashboard.store.ts";

const store = useDatasourceStore();
const dashStore = useDashboardStore()

onMounted(() => {
  store.getSources()
  store.getDrivers()
  dashStore.getDashboards()
});
</script>

<template>
  <Layout />
</template>

<style scoped>
</style>
