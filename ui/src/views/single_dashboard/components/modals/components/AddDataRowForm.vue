<script setup lang="ts">
import {useDatasourceStore} from "../../../../../modules/store/datasource.store.ts";
import ApiManager from "../../../../../modules/api";
import {funcMap, replaceFunctions} from "../../../../../modules/formatter";
import ChartBlueprint from "../../chart/ChartBlueprint.vue";
import {ref} from "vue";
import ShowDataTable from "./ShowDataTable.vue";
import {chartMapper} from "../../../../../modules/mapper";

const props = defineProps(['model']);
const store = useDatasourceStore();

const rowTypes = ['bar', 'line'];

const currentRow = ref(-1);

const showType = ref('')

const getData = async (rowId: number) => {
  showType.value = '';
  currentRow.value = rowId;
  const res = await ApiManager.performQuery(props.model.options.y[rowId]);
  if (Array.isArray(res)) {
    props.model.options.y[rowId].sampleData = [];
    for (const r of res) {
      if (r.length > 2) {
        props.model.options.y[rowId].sampleData.push(JSON.parse(r))
      }
    }
  }
  showType.value = 'table';
}

const getDataForAllRows = async (options: any) => {
  const result = [];

  for (const opt of options) {
    const res = await ApiManager.performQuery(opt);
    if (Array.isArray(res)) {
      for (const r of res) {
        if (r.length > 2) {
          result.push(JSON.parse(r))
        }
      }
    }
  }
  return result;
}

const addY = () => {
  props.model.options.y.push({
    id: props.model.options.y.length,
    source: null,
    query: `SELECT * FROM`,
    xField: '',
    yField: '',
    type: null,
    title: '',
    fn: null,
    sampleData: [],
  })
}

const deleteY = (rowId: number) => {
  props.model.options.y.splice(rowId, 1)
}

const prepareChart = async () => {
  showType.value = '';
  const mapped = chartMapper({ ...props.model.options, title: props.model.title, id: props.model.container_id});
  props.model.raw_options = replaceFunctions(mapped.rawOptions);
  props.model.data_queries = mapped.dataQueries;

  props.model.data = await getDataForAllRows(Object.values(props.model.data_queries));
  showType.value = 'chart';
}
</script>

<template>
  <v-row>
    <v-col cols="12">
      <label>Ряды данных</label>
      <v-btn variant="text" icon="mdi-plus" @click="addY"></v-btn>
    </v-col>
  </v-row>

  <div class="data_row">
    <v-expansion-panels style="width: 50%">
      <v-expansion-panel v-for="y in props.model.options.y" :key="y.id" style="width: 100%">
        <v-expansion-panel-title>
          {{ y.title || 'Новый ряд' }}
        </v-expansion-panel-title>
        <v-expansion-panel-text style="padding-top: 10px; width: 100%">
          <v-row>
            <v-col cols="4">
              <v-select label="Источник" :items="store.sources" v-model="y.source" item-title="title" item-value="driver"/>
            </v-col>
            <v-col cols="4">
              <v-text-field label="Поле Х" v-model="y.xField" />
            </v-col>
            <v-col cols="4">
              <v-text-field label="Поле Y" v-model="y.yField" />
            </v-col>
          </v-row>

          <v-row>
            <v-col cols="4">
              <v-select label="Тип" :items="rowTypes" v-model="y.type" />
            </v-col>
            <v-col cols="4">
              <v-text-field label="Название ряда" v-model="y.title" />
            </v-col>
            <v-col cols="4">
              <v-select label="Форматирование" :items="Object.keys(funcMap)" v-model="y.fn" />
            </v-col>
          </v-row>

          <v-row>
            <v-col cols="12">
              <v-textarea label="Запрос" v-model="y.query"></v-textarea>
              <v-btn variant="text" @click="getData(y.id)">Выполнить</v-btn>
              <v-btn variant="text" color="yellow" @click="prepareChart">Отобразить</v-btn>
              <v-btn variant="text" color="red" @click="deleteY(y.id)" style="float: right">Удалить</v-btn>
            </v-col>
          </v-row>
        </v-expansion-panel-text>
      </v-expansion-panel>
    </v-expansion-panels>

    <div style="width: 50%; height: 500px;">
      <ShowDataTable v-if="showType === 'table' && currentRow >= 0 && model.options.y[currentRow]" :y="model.options.y[currentRow]" />
    </div>
  </div>
</template>

<style scoped>
.data_row {
  display: flex;
  justify-content: space-between;
  gap: 20px;
  align-items: flex-start;
}
</style>