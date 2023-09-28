<script setup lang="ts">

import CreateModal from "../../layout/components/CreateModal.vue";
import {computed, reactive, ref} from "vue";

const props = defineProps(['dialog', 'data']);
const emit = defineEmits(['close', 'save']);

const onSave = () => {
  status.value = {
    color: 'yellow',
    text: 'не проверено',
  }
  reset();
  emit('save');
}

const onClose = () => {
  status.value = {
    color: 'yellow',
    text: 'не проверено',
  }
  reset();
  emit('close')
}

const reset = () => {
  model.driver = null;
  model.title = '';
  model.host = '';
  model.port = 5432;
  model.user = '';
  model.password = '';
  model.checked = false;
}

const loading = ref(false);
const status = ref({
  color: 'yellow',
  text: 'не проверено',
});

const model = reactive(props.data || {
  driver: null,
  title: '',
  host: '',
  port: 5432,
  user: '',
  password: '',
  checked: false,
});

const canCheckConnection = computed(() => {
  return model.driver && model.title && model.host && model.port && model.user && model.password;
})

const checkConnection = () => {
  loading.value = true;
  setTimeout(() => {
    loading.value = false;
    model.checked = true;
    status.value.color = 'green';
    status.value.text = 'Успешно';
  }, 2000);
}
</script>

<template>
  <CreateModal title="Создать источник данных" :dialog="props.dialog" :width="700">
    <template v-slot:form>
      <v-form>
        <v-row>
          <v-col>
            <v-select label="Драйвер *" v-model="model.driver" />
          </v-col>
          <v-col>
            <v-text-field label="Название *" v-model="model.title" />
          </v-col>
        </v-row>

        <v-row>
          <v-col>
            <v-text-field label="Хост *" v-model="model.host" />
          </v-col>
          <v-col>
            <v-text-field label="Порт *" placeholder="5432" v-model="model.port" />
          </v-col>
        </v-row>

        <v-row>
          <v-col>
            <v-text-field label="Пользователь *" v-model="model.user" />
          </v-col>
          <v-col>
            <v-text-field label="Пароль *" v-model="model.password" />
          </v-col>
        </v-row>

        <v-row>
          <v-col>
            <v-text-field label="Название БД" v-model="model.user" />
          </v-col>
        </v-row>

        <v-row>
          <v-col>
            <v-btn variant="text" :loading="loading" :disabled="!canCheckConnection" @click="checkConnection">Проверить подключение</v-btn>
          </v-col>
          <v-col>
            Статус подключения: <span :style="{color: status.color}">{{ status.text }}</span>
          </v-col>
        </v-row>
      </v-form>
    </template>

    <template v-slot:actions>
      <v-btn color="green" @click="onSave" :disabled="!model.checked">Создать</v-btn>
      <v-btn color="grey" @click="onClose">Закрыть</v-btn>
    </template>
  </CreateModal>
</template>