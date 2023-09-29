<script setup lang="ts">
import Page from "../layout/components/Page.vue";
import {nextTick, onMounted, ref, shallowRef} from "vue";
import CreateChartModal from "./components/modals/CreateChartModal.vue";
import {useDashboardStore} from "../../modules/store/dashboard.store.ts";
import {GridItem, GridLayout} from 'vue3-grid-layout-next';
import {useRoute} from "vue-router";
import {data, optionsStr, replaceFunctions} from "../../modules/utils";
import ChartBlueprint from "./components/chart/ChartBlueprint.vue";

const dashboardStore = useDashboardStore();

const route = useRoute();

const filterDrawer = ref(false);

const layout = ref([] as any);

const dashboardData = ref({
  id: 1,
  dash_id: '123',
  title: 'Дашборд 1',
  items: [
    {
      id: 2,
      type: 'chart',
      title: '123456',
      data: data,
      options: optionsStr,
      styles: {},
    }
  ]
});

onMounted(async () => {
  await dashboardStore.getAvailableTypes();
  let counter = 0;
  for (const item of dashboardData.value.items) {
    if (item.type === 'chart') {
      layout.value.push({
        x: counter * (item.width || 6), y: 0, w: item.width || 6, h: item.height || 11, i: item.id,
        component: shallowRef(ChartBlueprint),
        title: item.title,
        data: item.data,
        options: item.options,
        styles: item.styles
      })
    }
    counter++;
  }
  await nextTick(() => onContainerResized());
});

const createChartDialog = ref(false);

const onChartDialogClose = () => {
  createChartDialog.value = false;
}

const onChartDialogSave = () => {

}

const onMenuClick = (el: string) => {
  switch(el) {
    case 'chart':
      createChartDialog.value = true;
      return;
  }
}

const onContainerResized = () => {
  window.dispatchEvent(new Event('resize'))
}

const retrieveFilters = () => {

}

const applyFilters = () => {

}

</script>

<template>
  <Page title="Дашборд 1">
    <template v-slot:toolbar>
      <v-btn
          color="green"
          style="margin-right: 20px"
      >
        Добавить
        <v-menu location="bottom" offset="0" activator="parent">
          <v-list>
            <v-list-item
                v-for="item in dashboardStore.availableTypes"
                :key="item.id"
                :value="item.id"
                @click="onMenuClick(item.name)"
            >
              <v-list-item-title>{{ item.title }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </v-btn>

      <v-btn color="yellow" @click="filterDrawer = !filterDrawer">Фильтр</v-btn>
    </template>

    <template v-slot:content>
      <grid-layout
          :layout="layout"
          :col-num="12"
          :row-height="30"
          :is-draggable="true"
          :is-resizable="true"
          :is-mirrored="false"
          :vertical-compact="true"
          :margin="[10, 10]"
          :use-css-transforms="true"
      >
        <grid-item
            style="touch-action: none"
            v-for="item in layout"
            :key="item.i"
            :x="item.x"
            :y="item.y"
            :w="item.w"
            :h="item.h"
            :i="item.i"
            @resized="onContainerResized"
        >
          <component :is="item.component" :title="item.title" :data="item.data" :options="replaceFunctions(item.options)" :styles="item.styles"></component>
        </grid-item>
      </grid-layout>

      <v-navigation-drawer
          v-model="filterDrawer"
          location="right"
          temporary
      >
        <v-list style="height: 100%">
          <v-list-item>
            <v-select label="Зам" />
          </v-list-item>
          <v-list-item>
              <v-select label="ЦТС" />
          </v-list-item>

          <v-btn variant="text" color="green" style="position: absolute; bottom: 10px; width: 100%" @click="applyFilters">Применить</v-btn>
        </v-list>
      </v-navigation-drawer>
    </template>

    <template v-slot:modals>
      <CreateChartModal :dialog="createChartDialog" @close="onChartDialogClose" @save="onChartDialogSave"/>
    </template>
  </Page>
</template>

<style scoped>
</style>