<script setup lang="ts">
const props = defineProps(['y'])
</script>

<template>
  <v-table
      v-if="props.y.sampleData.length"
      height="500px"
      fixed-header
  >
    <thead>
    <tr>
      <th class="text-center" v-for="h in Object.keys(props.y.sampleData[0])" :key="h">
        {{ h }}
      </th>
    </tr>
    </thead>
    <tbody>
    <tr
        v-for="item in props.y.sampleData"
        :key="item[Object.keys(item)[0]]"
    >
      <td class="text-center" v-for="col in Object.values(item)" :key="col as string">{{ col }}</td>
    </tr>
    </tbody>
  </v-table>
</template>