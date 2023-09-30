<script setup lang="ts">
import {reactive} from "vue";
import CreateModalFullScreen from "../../../layout/components/CreateModalFullScreen.vue";
import AddDataRowModal from "./components/AddDataRowModal.vue";
import {chartMapper} from "../../../../modules/mapper";
import {useDashboardStore} from "../../../../modules/store/dashboard.store.ts";
const props = defineProps(['dialog']);
const emit = defineEmits(['close', 'save']);

const store = useDashboardStore();

let model = reactive({
  id: '',
  type: 1,
  title: '',
  description: '',
  dashboard: store.dashboard,
  data: [],
  options: {
    x: {
      title: '',
      fn: null,
    },
    y: [] as any[],
  },
  rawOptions: '',
  dataQueries: {},
});

const onSave = async () => {
  const mapped = chartMapper(model.options);
  model.rawOptions = mapped.rawOptions;
  model.dataQueries = mapped.dataQueries;

  const result = await store.saveChart(model);
  console.log(result);
  // emit('save', model);
  // clear();
}

// const clear = () => {
//   model = {} as any;
// }

</script>

<template>
  <CreateModalFullScreen title="Добавить график" :dialog="props.dialog">
    <template v-slot:form>
      <v-form>
        <v-row>
          <v-col cols="8">
            <v-text-field label="Название" v-model="model.title" />
          </v-col>
          <v-col cols="2">
            <v-text-field label="Идентификатор" v-model="model.id" />
          </v-col>
          <v-col cols="2">
            <v-select :items="store.dashboards" label="Дашборд" v-model="model.dashboard" />
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="12">
            <v-textarea label="Описание" v-model="model.description" rows="2"></v-textarea>
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="3">
            <v-text-field label="Название оси Х" v-model="model.options.x.title" />
          </v-col>
          <v-col cols="3">
            <v-select label="Форматирование оси Х" v-model="model.options.x.fn" />
          </v-col>
        </v-row>

        <AddDataRowModal :model="model"/>

      </v-form>
    </template>

    <template v-slot:actions>
      <v-btn color="green" @click="onSave">Создать</v-btn>
      <v-btn color="grey" @click="emit('close')">Закрыть</v-btn>
    </template>
  </CreateModalFullScreen>
</template>