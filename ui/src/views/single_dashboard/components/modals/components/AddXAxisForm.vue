<script setup lang="ts">
import {funcMap} from "../../../../../modules/formatter";
const props = defineProps(['model']);

const addX = () => {
  props.model.options.x.push({
    id: props.model.options.x.length,
    title: '',
    fn: null,
  })
}

const deleteX = (rowId: number) => {
  props.model.options.x.splice(rowId, 1)
}

</script>

<template>
  <v-row>
    <v-col cols="12">
      <label>Ось X</label>
      <v-btn variant="text" icon="mdi-plus" @click="addX"></v-btn>
    </v-col>
  </v-row>

  <div class="data_row">
    <v-expansion-panels style="width: 50%">
      <v-expansion-panel v-for="x in props.model.options.x" :key="x.id" style="width: 100%">
        <v-expansion-panel-title>
          {{ x.title || 'Новая ось' }}
        </v-expansion-panel-title>
        <v-expansion-panel-text style="padding-top: 10px; width: 100%">
          <v-row>
            <v-col cols="12">
              <v-btn variant="text" color="red" @click="deleteX(x.id)" style="float: right">Удалить</v-btn>
            </v-col>
          </v-row>
        </v-expansion-panel-text>
      </v-expansion-panel>
    </v-expansion-panels>
  </div>
</template>

<style scoped>
.data_row {
  display: flex;
  justify-content: space-between;
  gap: 20px;
  align-items: flex-start;
}
</style>