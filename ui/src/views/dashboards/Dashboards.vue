<script setup lang="ts">
  import Page from "../layout/components/Page.vue";
  import {ref} from "vue";
  import DashboardsTable from "./components/DashboardsTable.vue";
  import CreateDashboardModal from "./components/CreateDashboardModal.vue";
  import {useAppStore} from "../../modules/store/app.store.ts";
  import {useDashboardStore} from "../../modules/store/dashboard.store.ts";

  const appStore = useAppStore();

  const store = useDashboardStore();

  const createDashboardDialog = ref(false);

  const onCreateDialogClose = () => {
    createDashboardDialog.value = false;
  }

  const onCreateDialogSave = (model) => {}
</script>

<template>
  <Page>
    <template v-slot:toolbar>
      <v-text-field clearable="" density="compact" variant="outlined" label="Название" style="width: 400px; margin-top: 10px"/>
      <v-select density="compact" variant="outlined" label="Сортировка" :items="appStore.sortOptions" style="width: 200px; margin-left: 10px; margin-top: 10px" item-title="label" item-value="value"/>
      <v-btn icon="mdi-folder-plus" color="yellow" @click="createDashboardDialog = true" title="Создать"></v-btn>
    </template>

    <template v-slot:content>
      <DashboardsTable :dash="store.dashboards" />
    </template>

    <template v-slot:modals>
      <CreateDashboardModal :dialog="createDashboardDialog" @close="onCreateDialogClose" @save="onCreateDialogSave" />
    </template>
  </Page>
</template>