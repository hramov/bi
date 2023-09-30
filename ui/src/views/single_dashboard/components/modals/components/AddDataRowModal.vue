<script setup lang="ts">
import {useDatasourceStore} from "../../../../../modules/store/datasource.store.ts";
import ApiManager from "../../../../../modules/api";
import { funcMap} from "../../../../../modules/formatter";
const props = defineProps(['model']);
const store = useDatasourceStore();

const getData = async (rowId: number) => {
  const res = await ApiManager.performQuery(props.model.options.y[rowId]);
  if (Array.isArray(res)) {
    props.model.options.y[rowId].sampleData = [];
    for (const r of res) {
      if (r.length > 2) {
        props.model.options.y[rowId].sampleData.push(JSON.parse(r))
      }
      if (props.model.options.y[rowId].sampleData.length > 10) {
        return;
      }
    }
  }
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

const rowTypes = ['bar', 'line'];
</script>

<template>
  <v-row>
    <v-col cols="12">
      <label>Ряды данных</label>
      <v-btn variant="text" icon="mdi-plus" @click="addY"></v-btn>
    </v-col>
  </v-row>

  <v-expansion-panels>
    <v-expansion-panel v-for="y in props.model.options.y" :key="y.id">
      <v-expansion-panel-title>
        {{ y.title || 'Новый ряд' }}
      </v-expansion-panel-title>
      <v-expansion-panel-text style="padding-top: 10px">
        <v-row>
          <v-col cols="2">
            <v-select label="Источник" :items="store.sources" v-model="y.source" item-title="title" item-value="driver"/>
          </v-col>
          <v-col cols="2">
            <v-text-field label="Поле Х" v-model="y.xField" />
          </v-col>
          <v-col cols="2">
            <v-text-field label="Поле Y" v-model="y.yField" />
          </v-col>
          <v-col cols="2">
            <v-select label="Тип" :items="rowTypes" v-model="y.type" />
          </v-col>
          <v-col cols="2">
            <v-text-field label="Название ряда" v-model="y.title" />
          </v-col>
          <v-col cols="2">
            <v-select label="Форматирование" :items="Object.keys(funcMap)" v-model="y.fn" />
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="6">
            <v-textarea label="Запрос" v-model="y.query"></v-textarea>
            <v-btn variant="text" @click="getData(y.id)">Выполнить</v-btn>
            <v-btn variant="text" color="red" @click="deleteY(y.id)">Удалить</v-btn>
          </v-col>
          <v-col cols="6">
            <v-table
                v-if="y.sampleData.length"
                height="200px"
                fixed-header
            >
              <thead>
              <tr>
                <th class="text-center" v-for="h in Object.keys(y.sampleData[0])" :key="h">
                  {{ h }}
                </th>
              </tr>
              </thead>
              <tbody>
              <tr
                  v-for="item in y.sampleData"
                  :key="item[Object.keys(item)[0]]"
              >
                <td class="text-center" v-for="col in Object.values(item)" :key="col as string">{{ col }}</td>
              </tr>
              </tbody>
            </v-table>
            <div v-else>Нет данных для отображения</div>
          </v-col>
        </v-row>
      </v-expansion-panel-text>
    </v-expansion-panel>
  </v-expansion-panels>
</template>