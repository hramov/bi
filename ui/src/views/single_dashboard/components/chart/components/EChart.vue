<script setup lang="ts">
import * as echarts from 'echarts'
import {onMounted, onUnmounted, ref, watch} from 'vue'
const props = defineProps(['id', 'title', 'data', 'options', 'styles', 'type'])

const chartData = ref({})
const chartOptions = {
  ...props.options,
}
const styles = {
  ...props.styles,
}

let myChart: any = null

function buildChart() {
  const labels = new Set()
  const datasets = new Map()

  for (const row of props.data) {
    if (props.options.labelFormatter) {
      labels.add(props.options.labelFormatter(row[props.options.xAxis.field]))
    } else {
      labels.add(row[props.options.axis.x.field])
    }

    for (let col in row) {
      const config = props.options.labels.find((c: any) => c.title === col)

      if (config) {
        if (!datasets.has(config.title)) {
          datasets.set(config.title, {
            type: config.type,
            name: config.display,
            label: config.display,
            data: config.formatter ? [config.formatter(row[config.title])] : [row[config.title]],
            stack: config.stack,
            yAxisIndex: config.yAxisID,
            itemStyle: {
              color: config.color || null,
              opacity: config.opacity || 1,
            },
            axisTicks: config.axisTicks,
          })
        } else {
          const dataset = datasets.get(col)
          dataset.data.push(config.formatter ? config.formatter(row[col]) : row[col])
        }
      }
    }
  }

  chartData.value = {
    labels: Array.from(labels),
    datasets: Array.from(datasets.values()),
    datasetsLabels: Array.from(datasets.keys()),
  }
}

const dialog = ref(false)
const name = ref('')
const seriesName = ref('')

function drawChart() {
  buildChart();
  const chartElement = document.getElementById(chartOptions.id);
  myChart = echarts.init(chartElement, 'dark')
  window.addEventListener('resize', () => {
    myChart.resize()
  })

  const options = {
    ...chartOptions,
    xAxis: {
      ...chartOptions.xAxis,
      data: (chartData.value as any).labels,
    },
    series: (chartData.value as any).datasets,
  }

  myChart.setOption(options)

  myChart.on('click', (ev: any) => {
    name.value = ev.name
    seriesName.value = ev.seriesName
    dialog.value = true
  })
}

onMounted(() => {
  drawChart()
})

onUnmounted(() => {
  window.removeEventListener('resize', myChart.resize)
})

watch(
    () => props.data,
    () => {
      window.removeEventListener('resize', myChart.resize)
      myChart.dispose()
      myChart = null
      drawChart()
    },
)
</script>

<template>
  <div :id="chartOptions.id" style="width: 100%; height: calc(100% - 36px);"></div>
</template>

<style scoped></style>
