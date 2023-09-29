<script setup lang="ts">
import {ref} from "vue";
import ChartBlueprint from "../../chart/ChartBlueprint.vue";
import {useDatasourceStore} from "../../../../../modules/store/datasource.store.ts";
import ApiManager from "../../../../../modules/api/api.ts";

const store = useDatasourceStore();

const model = ref({
  data: `[]`,
  options: `{}`,
  query: `SELECT * FROM table`,
  source: null,
});

const result = ref([] as any)

const replaceFunctions = (options: string) => {
  const optObject = JSON.parse(options);
  replaceFunctionsImpl(optObject);
}

const replaceFunctionsImpl = (root: any) => {
  if (typeof root === 'function') {
    throw new Error('Functions are not allowed');
  }
  if (typeof root !== 'object') {
    if (root.startsWith('customfn_')) {
      const handler = root.replace('customfn_', '');
      console.log(handler);
    }
  }
}

const getData = async () => {
  const res = await ApiManager.performQuery(model.value);
  if (Array.isArray(res)) {
    result.value = [];
    for (const r of res) {
      if (r.length > 2) {
        result.value.push(JSON.parse(r))
      }
    }
  }
}
</script>

<template>
  <v-row>
    <v-col cols="12">
      <label>Данные и описание</label>
    </v-col>
  </v-row>

  <v-row>
    <v-col cols="6">
      <v-select label="Источник" :items="store.sources" v-model="model.source" item-title="title" item-value="driver"/>
      <v-textarea label="Запрос" rows="5" v-model="model.query"></v-textarea>
      <v-btn @click="getData" class="mb-4">Выполнить</v-btn>
      <v-textarea label="Параметры (JSON)" rows="8" v-model="model.options"></v-textarea>
    </v-col>

    <v-col cols="6">
<!--      <ChartBlueprint v-if="JSON.parse(model.options).id" title="Проверка" :data="JSON.parse(model.data)" :options="replaceFunctions(model.options)" />-->
      <v-table
          v-if="result[0]"
          height="250px"
          fixed-header
      >
        <thead>
        <tr>
          <th class="text-center" v-for="h in Object.keys(result[0])" :key="h">
            {{ h }}
          </th>
        </tr>
        </thead>
        <tbody>
        <tr
            v-for="item in result"
            :key="item[Object.keys(item)[0]]"
        >
          <td class="text-center" v-for="col in Object.values(item)" :key="col as string">{{ col }}</td>
        </tr>
        </tbody>
      </v-table>
    </v-col>
  </v-row>

</template>