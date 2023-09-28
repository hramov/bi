<script setup lang="ts">
import Page from "../layout/components/Page.vue";
import DataSourceModal from "./components/DataSourceModal.vue";
import {ref} from "vue";

const dialog = ref(false);

const showBanner = ref(true);

const dataSourceData = ref({} as any);

const editDataSource = (row: any) => {
  dataSourceData.value = row;
  dialog.value = true;
}

const deleteDataSource = () => {

}
</script>

<template>
  <Page>
    <template v-slot:toolbar>
    </template>

    <template v-slot:content>
      <div v-if="showBanner" style="padding-left: 20px; padding-right: 20px">
        <v-banner
            lines="one"
            icon="$warning"
            color="warning"
            class="mt-4"
        >
          <v-banner-text style="font-size: 18px">
            Учетная запись, используемая в источниках, должна иметь права ТОЛЬКО на чтение данных!
          </v-banner-text>

          <v-banner-actions>
            <v-btn @click="showBanner = false">Понятно</v-btn>
          </v-banner-actions>
        </v-banner>
      </div>

      <div class="data_sources" style="padding-top: 20px">
        <v-card
            variant="elevated"
        >
          <v-card-item>
            <div>
              <div class="text-overline mb-1">
              </div>
              <div class="text-h6 mb-1">
                Postgres
              </div>
              <div class="text-caption" style="max-width: 200px">Host: localhost</div>
              <div class="text-caption" style="max-width: 200px">Port: 5432</div>
              <div class="text-caption" style="max-width: 200px">User: postgres</div>
              <div class="text-caption" style="max-width: 200px">Database: bi</div>
            </div>
          </v-card-item>

          <v-card-actions>
            <v-btn icon="mdi-pencil" @click="editDataSource"></v-btn>
            <v-btn icon="mdi-delete" @click="deleteDataSource"></v-btn>
          </v-card-actions>
        </v-card>

        <v-card
            variant="elevated"
        >
          <v-card-item style="height: 100%">
            <v-btn icon="mdi-plus" @click="dialog = true"></v-btn>
          </v-card-item>
        </v-card>
      </div>
    </template>

    <template v-slot:modals>
      <DataSourceModal :dialog="dialog" @close="dialog = false" :data="dataSourceData"/>
    </template>
  </Page>
</template>

<style scoped>
.data_sources {
  width: 100%;
  display: flex;
  justify-content: left;
  padding: 0 20px;
  gap: 20px;
}
</style>