<script setup lang="ts">
import EChart from './components/EChart.vue';

const props = defineProps(['type', 'title', 'options', 'styles', 'loading', 'data']);

const title = props.title;
</script>

<template>
  <v-card style="height: 100%; margin: 0; width: 100%">
    <v-card-title>{{ title }}</v-card-title>
    <v-skeleton-loader v-if="props.loading"></v-skeleton-loader>
    <EChart v-else :type="props.type" :title="props.title" :styles="styles" :options="options" :data="props.data" />
  </v-card>
</template>

<style scoped></style>
