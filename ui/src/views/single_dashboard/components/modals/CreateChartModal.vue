<script setup lang="ts">
import {onMounted, ref} from "vue";
import CreateModalFullScreen from "../../../layout/components/CreateModalFullScreen.vue";
import AddDataRowForm from "./components/AddDataRowForm.vue";
import {chartMapper} from "../../../../modules/mapper";
import {useDashboardStore} from "../../../../modules/store/dashboard.store";
import AddXAxisForm from "./components/AddXAxisForm.vue";
import AddYAxisForm from "./components/AddYAxisForm.vue";
import ChartBlueprint from "../chart/ChartBlueprint.vue";
const props = defineProps(['dialog', 'dash_id']);
const emit = defineEmits(['close', 'save']);

const store = useDashboardStore();

const model = ref({} as any);

onMounted(() => {
  model.value = {
    container_id: '',
    item_type: 1,
    title: '',
    description: '',
    dash_id: store.dashboard.dash_id,
    data: [],
    options: {
      x: [] as any[],
      y: [] as any[],
    },
    raw_options: '',
    data_queries: {},
  }
})

const onSave = async () => {
  const mapped = chartMapper(model.value.options);
  model.value.options = mapped.rawOptions;
  model.value.data_queries = mapped.dataQueries;

  const result = await store.saveChart(model.value);
  console.log(result);
  emit('save', model);
  clear();
}

const clear = () => {
  model.value = {} as any;
}

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
            <v-text-field label="Идентификатор" v-model="model.container_id" />
          </v-col>
          <v-col cols="2">
            <v-select :items="store.dashboards" label="Дашборд" v-model="model.dash_id" item-title="title" item-value="dash_id" />
          </v-col>
        </v-row>

        <div class="settings">
          <AddXAxisForm :model="model" />
          <AddYAxisForm :model="model" />
          <AddDataRowForm :model="model"/>
        </div>

        <div class="chart">
          <ChartBlueprint
              v-if="model.raw_options"
              :title="model.title"
              :data="model.data"
              :options="model.raw_options"
              :styles="{}"
          />
        </div>

      </v-form>
    </template>

    <template v-slot:actions>
      <v-btn color="green" @click="onSave">Создать</v-btn>
      <v-btn color="grey" @click="emit('close')">Закрыть</v-btn>
    </template>
  </CreateModalFullScreen>
</template>