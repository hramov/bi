<script setup lang="ts">
import {onMounted, ref} from "vue";
import CreateModalFullScreen from "../../../layout/components/CreateModalFullScreen.vue";
import AddDataRowModal from "./components/AddDataRowModal.vue";
import {chartMapper} from "../../../../modules/mapper";
import {useDashboardStore} from "../../../../modules/store/dashboard.store.ts";
const props = defineProps(['dialog']);
const emit = defineEmits(['close', 'save']);

const store = useDashboardStore();

const model = ref({} as any);

onMounted(() => {
  model.value = {
    id: '',
    type: 1,
    title: '',
    description: '',
    dashboard: 1,
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
  }
})

const onSave = async () => {
  const mapped = chartMapper(model.value.options);
  model.value.rawOptions = mapped.rawOptions;
  model.value.dataQueries = mapped.dataQueries;

  const result = await store.saveChart(model.value);
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
            <v-select :items="store.dashboards" label="Дашборд" v-model="model.dashboard" item-title="title" item-value="id" />
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="8">
            <v-text-field label="Описание" v-model="model.description" />
          </v-col>
          <v-col cols="2">
            <v-text-field label="Название оси Х" v-model="model.options.x.title" />
          </v-col>
          <v-col cols="2">
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