<script setup lang="ts">
import {reactive} from "vue";
import CreateModalFullScreen from "../../../layout/components/CreateModalFullScreen.vue";
import AddDataRowModal from "./AddDataRowModal.vue";
import ChartOptionsForm from "./components/ChartOptionsForm.vue";
const props = defineProps(['dialog']);
const emit = defineEmits(['close', 'save']);

const model = reactive({
  title: '',
  description: '',
  dashboard: '',
});

const onSave = () => {
  emit('save', model);
  clear();
}

const clear = () => {
  model.title = '';
  model.description = '';
  model.dashboard = '';
}

</script>

<template>
  <CreateModalFullScreen title="Добавить график" :dialog="props.dialog">
    <template v-slot:form>
      <v-form>
        <v-row>
          <v-col cols="10">
            <v-text-field label="Название" v-model="model.title" />
          </v-col>
          <v-col cols="2">
            <v-select label="Дашборд" v-model="model.dashboard" />
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="12">
            <v-textarea label="Описание" v-model="model.description" rows="2"></v-textarea>
          </v-col>
        </v-row>

<!--        <AddDataRowModal />-->
        <ChartOptionsForm />

      </v-form>
    </template>

    <template v-slot:actions>
      <v-btn color="green" @click="onSave">Создать</v-btn>
      <v-btn color="grey" @click="emit('close')">Закрыть</v-btn>
    </template>
  </CreateModalFullScreen>
</template>