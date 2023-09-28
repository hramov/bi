<script setup lang="ts">

import CreateModal from "../../../layout/components/CreateModal.vue";
import {reactive} from "vue";
const props = defineProps(['dialog']);
const emit = defineEmits(['close', 'save']);

const model = reactive({
  title: '',
  description: '',
});

const onSave = () => {
  emit('save', model);
  clear();
}

const clear = () => {
  model.title = '';
  model.description = '';
}

</script>

<template>
  <CreateModal title="Добавить график" :dialog="props.dialog">
    <template v-slot:form>
      <v-form>
        <v-text-field label="Название" v-model="model.title" />
      </v-form>
    </template>

    <template v-slot:actions>
      <v-btn color="green" @click="onSave">Создать</v-btn>
      <v-btn color="grey" @click="emit('close')">Закрыть</v-btn>
    </template>
  </CreateModal>
</template>