<script setup lang="ts">

import CreateModal from "../../layout/components/CreateModal.vue";
import {ref} from "vue";
import ApiManager from "../../../modules/api";
const props = defineProps(['dialog']);
const emit = defineEmits(['close', 'save']);

const model = ref({} as any);

const onSave = async () => {
  const res = await ApiManager.post('/dashboards/', model.value);
  if (res.status < 400) {
    clear();
  }
}

const clear = () => {
  model.value = {};
}

</script>

<template>
  <CreateModal title="Создать дашборд" :dialog="props.dialog">
    <template v-slot:form>
      <v-form>
        <v-text-field label="Название" v-model="model.title" />

        <v-textarea rows="3" type="textarea" label="Описание" v-model="model.description"></v-textarea>

        <v-select label="Уровень доступа" v-model="model.access" />
      </v-form>
    </template>

    <template v-slot:actions>
      <v-btn color="green" @click="onSave">Создать</v-btn>
      <v-btn color="grey" @click="emit('close')">Закрыть</v-btn>
    </template>
  </CreateModal>
</template>