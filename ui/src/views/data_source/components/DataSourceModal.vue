<script setup lang="ts">

import CreateModal from "../../layout/components/CreateModal.vue";
import {computed, ref, watch} from "vue";
import ApiManager from "../../../modules/api";
import {useDatasourceStore} from "../../../modules/store/datasource.store";
import { useNotificationsStore, NotificationService } from "../../../modules/store/notifications.store";

const props = defineProps(['dialog', 'data']);
const emit = defineEmits(['close', 'save']);

const store = useDatasourceStore();
const notificationService: NotificationService = useNotificationsStore();

const onClose = () => {
  status.value = {
    color: 'yellow',
    text: 'не проверено',
  }
  reset();
  emit('close')
}

const reset = () => {
  model.value = {
    driver: null,
    title: '',
    host: '',
    port: '5432',
    user: '',
    password: '',
    database: '',
    sslmode: 'disable',
    checked: false,
  };
}

const loading = ref(false);
const status = ref({
  color: 'yellow',
  text: 'не проверено',
});

const model = ref({
  driver: null,
  title: '',
  host: '',
  port: '5432',
  user: '',
  password: '',
  database: '',
  sslmode: 'disable',
  checked: false,
});

watch(() => props.dialog, () => {
  model.value = { ...props.data, ...credentials(props.data.dsn), sslmode: 'disable'};
})

const canCheckConnection = computed(() => {
  return model.value.driver && model.value.title && model.value.host && model.value.port && model.value.user && model.value.password;
})

const checkConnection = async () => {
  loading.value = true;
  const result = await ApiManager.checkConnection(model.value);
  loading.value = false;

  if (result.status === 'accepted') {
    model.value.checked = true;
    status.value.color = 'green';
    status.value.text = 'Успешно';
    notificationService.showNotification.success({
      text: "Проверка подключения прошла успешно",
      duration: 2000,
    })
  } else {
    model.value.checked = false;
    status.value.color = 'red';
    status.value.text = 'Ошибка';
    notificationService.showNotification.error({
      text: result.error,
      duration: 10000
    })
  }
}

const dsn = computed(() => {
  return `postgresql://${model.value.user || '[user]'}:${model.value.password || '[password]'}@${model.value.host || '[host]'}:${model.value.port || '[port]'}/${model.value.database || '[database]'}?sslmode=${model.value.sslmode || '[sslmode]'}`
});

const credentials = (dsn: string) => {
  if (!dsn) return {};

  const driverPrefix = dsn.split('://')[0];
  const user = dsn.split('://')[1].split(':')[0];
  const password = dsn.split('://')[1].split(':')[1].split('@')[0];
  const host = dsn.split('://')[1].split(':')[1].split('@')[1].split(':')[0];
  const port = dsn.split('://')[1].split(':')[2].split('/')[0]
  const database = dsn.split('://')[1].split(':')[2].split('/')[1].split('?')[0]
  const sslmode = dsn.split('://')[1].split(':')[2].split('/')[1].split('?')[1].split('sslmode=')[1]

  return {
    driverPrefix, user,password, host, port, database, sslmode
  }
};

const onSave = async () => {
  model.value.dsn = dsn;
  try {
    await store.saveSource(model.value);
    notificationService.showNotification.success({
      text: "Источник данных был успешно сохранен",
      duration: 2000,
    })
    emit('save');
  } catch (e) {
    notificationService.showNotification.error({
      text: e.message,
      duration: 10000,
    })
  }
  
}
</script>

<template>
  <CreateModal title="Создать источник данных" :dialog="props.dialog" :width="700">
    <template v-slot:form>
      <v-form>
        <v-row><v-col>{{ dsn }}</v-col></v-row>

        <v-row>
          <v-col>
            <v-select label="Драйвер *" v-model="model.driver" :items="store.drivers" item-title="title" item-value="code"/>
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
            <v-text-field label="Название БД" v-model="model.database" />
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
      <v-btn color="green" @click="onSave" :disabled="!model.checked">Сохранить</v-btn>
      <v-btn color="grey" @click="onClose">Закрыть</v-btn>
    </template>
  </CreateModal>
</template>