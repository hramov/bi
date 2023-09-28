<script setup lang="ts">
import EChart from './components/EChart.vue';

const props = defineProps(['type', 'title', 'options', 'styles', 'loading', 'data']);

const title = props.title;
</script>

<template>
  <v-card style="height: 100%; margin: 0; width: 100%">
    <v-card-title>
      <div style="display:flex; justify-content: space-between; align-items: center">
        <span>{{ title }}</span>
        <div style="width: 20%; text-align: right">
          <v-btn variant="text" icon="mdi-dots-horizontal" size="40"></v-btn>
          <v-menu location="bottom" offset="0" activator="parent">
            <v-list>
              <v-list-item
                  v-for="(item, index) in ['Опция 1', 'Опция 2']"
                  :key="index"
                  :value="index"
              >
                <v-list-item-title>{{ item }}</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>
        </div>
      </div>
    </v-card-title>
    <v-skeleton-loader v-if="props.loading"></v-skeleton-loader>
    <EChart v-else :type="props.type" :title="props.title" :styles="styles" :options="options" :data="props.data" />
  </v-card>
</template>

<style scoped></style>
