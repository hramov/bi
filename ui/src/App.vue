<script setup lang="ts">
import Layout from "./views/layout/Layout.vue";
import {useDatasourceStore} from "./modules/store/datasource.store";
import {onMounted} from "vue";
import {useDashboardStore} from "./modules/store/dashboard.store";
import {useNotificationsStore, NotificationService, EButtonsType} from "./modules/store/notifications.store";

import notifications from "./views/components/notifications/NotificationsComponent.vue";

const store = useDatasourceStore();
const dashStore = useDashboardStore();
const notificationService: NotificationService = useNotificationsStore();

function showNotification() {
  notificationService.showNotification.error({
    text: 'произошел троллинг',
    duration: 200000,
    actions: [
      {
        callback:() => {
            console.log(123123)
        },
        type: EButtonsType.retry,
      }
    ]
  });
}

onMounted(() => {
  store.getSources();
  store.getDrivers();
  dashStore.getDashboards();
});

</script>

<template>
  <v-app>
    <Layout />
    <v-btn @click="showNotification">ok</v-btn>
    <notifications v-if="notificationService.notifications.length"/>
  </v-app>
</template>

<style scoped>
</style>
