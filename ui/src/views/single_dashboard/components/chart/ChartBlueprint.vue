<script setup lang="ts">
import {ref} from 'vue'
import Chart from './Chart.vue';

const props = defineProps(['filter', 'options', 'styles', 'title', 'data', 'minValueFunc']);

const options = ref(props.options);

const chartData = ref(props.data);

const loading = ref(false);

</script>

<template>
  <Chart :key="options.yAxis[0].min" :title="props.title" :data="chartData" :options="options" :styles="props.styles" :loading="loading" />
</template>

<style scoped></style>
