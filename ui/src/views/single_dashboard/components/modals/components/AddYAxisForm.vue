<script setup lang="ts">
import {funcMap} from "../../../../../modules/formatter";
const props = defineProps(['model']);
</script>

<template>
  <v-row>
    <v-col cols="12">
      <label>Ось Y</label>
      <v-btn variant="text" icon="mdi-plus"></v-btn>
    </v-col>
  </v-row>

  <!--  <div class="data_row">-->
  <!--    <v-expansion-panels style="width: 50%">-->
  <!--      <v-expansion-panel v-for="y in props.model.options.y" :key="y.id" style="width: 100%">-->
  <!--        <v-expansion-panel-title>-->
  <!--          {{ y.title || 'Новый ряд' }}-->
  <!--        </v-expansion-panel-title>-->
  <!--        <v-expansion-panel-text style="padding-top: 10px; width: 100%">-->
  <!--          <v-row>-->
  <!--            <v-col cols="4">-->
  <!--              <v-select label="Источник" :items="store.sources" v-model="y.source" item-title="title" item-value="driver"/>-->
  <!--            </v-col>-->
  <!--            <v-col cols="4">-->
  <!--              <v-text-field label="Поле Х" v-model="y.xField" />-->
  <!--            </v-col>-->
  <!--            <v-col cols="4">-->
  <!--              <v-text-field label="Поле Y" v-model="y.yField" />-->
  <!--            </v-col>-->
  <!--          </v-row>-->

  <!--          <v-row>-->
  <!--            <v-col cols="4">-->
  <!--              <v-select label="Тип" :items="rowTypes" v-model="y.type" />-->
  <!--            </v-col>-->
  <!--            <v-col cols="4">-->
  <!--              <v-text-field label="Название ряда" v-model="y.title" />-->
  <!--            </v-col>-->
  <!--            <v-col cols="4">-->
  <!--              <v-select label="Форматирование" :items="Object.keys(funcMap)" v-model="y.fn" />-->
  <!--            </v-col>-->
  <!--          </v-row>-->

  <!--          <v-row>-->
  <!--            <v-col cols="12">-->
  <!--              <v-textarea label="Запрос" v-model="y.query"></v-textarea>-->
  <!--              <v-btn variant="text" @click="getData(y.id)">Выполнить</v-btn>-->
  <!--              <v-btn variant="text" color="yellow" @click="prepareChart">Отобразить</v-btn>-->
  <!--              <v-btn variant="text" color="red" @click="deleteY(y.id)" style="float: right">Удалить</v-btn>-->
  <!--            </v-col>-->
  <!--          </v-row>-->
  <!--        </v-expansion-panel-text>-->
  <!--      </v-expansion-panel>-->
  <!--    </v-expansion-panels>-->
  <!--  </div>-->
</template>

<style scoped>
.data_row {
  display: flex;
  justify-content: space-between;
  gap: 20px;
  align-items: flex-start;
}
</style>