<script setup lang="ts">
import {onMounted, ref, watch} from 'vue'
import moment from 'moment'
import {defaultTooltipFormatter} from '../../../../modules/utils';
import Chart from './Chart.vue';
import {useDashboardStore} from "../../../../modules/store/dashboard.store";

const props = defineProps(['filter']);

const chartData = ref([] as any[])
const loading = ref(false)

const store = useDashboardStore();

const loadData = async (filter) => {
  loading.value = true
  chartData.value = await store.getAvailability(filter.period)
  loading.value = false
}

onMounted(async () => {
  await loadData(props.filter);
  calcMinValue('Значение');
});

watch(() => props.filter.period, async (val) => {
  loading.value = true;
  chartData.value = await store.getAvailability(val);
  calcMinValue('Значение');
  loading.value = false;
});

const calcMinValue = (field: string) => {
  let min = Number.MAX_VALUE, max = Number.MIN_VALUE;
  for (const d of chartData.value) {
    min = Math.min(d[field], min)
    max = Math.max(d[field], max)
  }
  options.value.yAxis[0].min = ((min - ((max - min) * 0.1)) * 100).toFixed(2);
}

const title = 'Доступность2';

const styles = {};

const options = ref({
  id: 'availability2',
  grid: { containLabel: true, x: '7%', y: '7%', x2: '5%', y2: '7%' },
  labels: [
    {
      type: 'line',
      title: 'Значение',
      display: 'Доступность',
      formatter: (val: number) => {
        if (Number(val)) {
          return (val * 100).toFixed(2)
        }
        return val
      },
    },
    {
      yAxisID: 1,
      type: 'bar',
      title: 'Всего',
      display: 'Всего',
      color: '#c0bdbd',
      opacity: 0.5,
    },
  ],
  labelFormatter: (val: string) => {
    return moment(val).format('DD.MM.YYYY')
  },
  yAxis: [
    {
      type: 'value',
      position: 'left',
      name: 'Доступность',
      nameLocation: 'middle',
      nameGap: 70,
      axisLabel: {
        formatter: (val: number) => {
          return val + '%'
        },
      },
      min: 'dataMin',
    },
    {
      type: 'value',
      position: 'right',
      splitLine: {
        show: false,
      },
      axisLabel: {
        show: false,
      },
    },
  ],
  xAxis: {
    field: 'Дата',
    type: 'category',
  },
  tooltip: {
    trigger: 'axis',
    cross: true,
    axisPointer: {
      type: 'shadow',
    },
    formatter: defaultTooltipFormatter(['default', 'percent-sign']),
  },
  legend: {},
})

</script>

<template>
  <Chart :key="options.yAxis[0].min" type="bar" :title="title" :data="chartData" :options="options" :styles="styles" :loading="loading" style="position: relative" />
</template>

<style scoped></style>
