<script setup lang="ts">
import Page from "../layout/components/Page.vue";
import { nextTick, onMounted, ref, shallowRef } from "vue";
import CreateChartModal from "./components/modals/CreateChartModal.vue";
import { useDashboardStore } from "../../modules/store/dashboard.store";
import { GridItem, GridLayout } from 'vue3-grid-layout-next';
import { useRoute } from "vue-router";
import { replaceFunctions } from "../../modules/formatter";
import ChartBlueprint from "./components/chart/ChartBlueprint.vue";
import ApiManager from "../../modules/api";

const store = useDashboardStore();

const route = useRoute();

const filterDrawer = ref(false);

const layout = ref([] as any);

const getDataForRow = async (row: any) => {

  const result = [];

  const res = await ApiManager.performQuery(row.data_queries[row.options.yAxis[0].yAxisID]);
  if (Array.isArray(res)) {
    for (const r of res) {
      if (r.length > 2) {
        result.push(JSON.parse(r))
      }
    }
  }

  return result;
}

const loadData = async () => {
  await store.getDashboard(route.params.dashboard_id as string);
  await store.getAvailableTypes();

  let counter = 0;

  if (store.dashboard.items) {
    for (const item of store.dashboard.items) {
      if (item.item_type === 1) {
        layout.value.push({
          x: (counter % 2) * (item.width || 6),
          y: (Math.abs((counter - 1)) % 2) * (item.width || 11),
          w: item.width || 6,
          h: item.height || 11,
          i: item.id,
          component: shallowRef(ChartBlueprint),
          title: item.title,
          data: await getDataForRow(item),
          options: item.options,
          styles: item.styles
        })
      }
      counter++;
    }
  }

  await nextTick(() => onContainerResized());
}
onMounted(async () => {
  await loadData();
});

const createChartDialog = ref(false);

const onChartDialogClose = () => {
  createChartDialog.value = false;
}

const onChartDialogSave = async () => {
  createChartDialog.value = false;
  await loadData();
}

const onMenuClick = (el: string) => {
  switch (el) {
    case 'chart':
      createChartDialog.value = true;
      return;
  }
}

const onContainerResized = () => {
  window.dispatchEvent(new Event('resize'))
}

const applyFilters = () => {

}

</script>

<template>
  <Page :title="{ title: store.dashboard.title, path: store.dashboard.title + '/' + store.dashboard.dash_id}">
    <template v-slot:toolbar>
      <v-btn color="green" style="margin-right: 20px">
        Добавить
        <v-menu location="bottom" offset="0" activator="parent">
          <v-list>
            <v-list-item v-for="item in store.availableTypes" :key="item.id" :value="item.id"
              @click="onMenuClick(item.name)">
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
          :use-css-transforms="true">
        <grid-item
            style="touch-action: none; min-width: 500px; min-height: 350px"
            v-for="item in layout"
            :key="item.i"
            :x="item.x"
            :y="item.y"
            :w="item.w"
            :h="item.h"
            :i="item.i"
            :min-w="3"
            :min-h="8"
            :max-w="6"
            :max-h="11"
            @resized="onContainerResized"
            @container-resized="onContainerResized">
          <component
              :is="item.component"
              :title="item.title"
              :data="item.data"
              :options="replaceFunctions(item.options)"
              :styles="item.styles"
          ></component>
        </grid-item>
      </grid-layout>

      <v-navigation-drawer v-model="filterDrawer" location="right" temporary>
        <v-list style="height: 100%">
          <v-list-item>
            <v-select label="Зам" />
          </v-list-item>
          <v-list-item>
            <v-select label="ЦТС" />
          </v-list-item>

          <v-btn variant="text" color="green" style="position: absolute; bottom: 10px; width: 100%"
            @click="applyFilters">Применить</v-btn>
        </v-list>
      </v-navigation-drawer>
    </template>

    <template v-slot:modals>
      <CreateChartModal v-if="store.dashboard.dash_id" :dialog="createChartDialog" @close="onChartDialogClose" @save="onChartDialogSave" />
    </template>
  </Page>
</template>

<style scoped></style>