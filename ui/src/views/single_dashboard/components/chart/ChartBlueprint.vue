<script setup lang="ts">
import {ref} from 'vue'
import Chart from './Chart.vue';

const props = defineProps(['filter', 'options', 'styles', 'title', 'data', 'minValueFunc']);

const chartData = ref(props.data);

const loading = ref(false);

</script>

<template>
  <Chart :title="props.title" :data="chartData" :options="props.options" :styles="props.styles" :loading="loading"/>
</template>

<style scoped></style>
