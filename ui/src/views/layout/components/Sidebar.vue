<script setup lang="ts">
</script>

<template>
  <v-list density="compact" nav style="margin-top: 10px;">
  <v-row style="height: calc(100vh - 70px);">
    <v-col class="align-self-start justify-start" cols="12">
      <v-list-item to="/" prepend-icon="mdi-home" title="Главная" value="home"></v-list-item>
      <v-list-item to="/dashboards" prepend-icon="mdi-view-dashboard" title="Дашборды" value="dashboards"></v-list-item>
    </v-col>
    <v-col dense class="align-self-end justify-end" cols="12">
      <v-divider class="mb-4"></v-divider>
      <v-list-item to="/data_sources" prepend-icon="mdi-database-arrow-up-outline" title="Источники данных" value="data_sources"></v-list-item>
      <v-list-item to="/settings" prepend-icon="mdi-cog" title="Настройки" value="settings"></v-list-item>
      <v-list-item to="/help" prepend-icon="mdi-help-circle-outline" title="Помощь" value="help"></v-list-item>
    </v-col>
  </v-row>
  </v-list>
</template>

<style scoped>
* {
  font-size: 20px;
}
</style>