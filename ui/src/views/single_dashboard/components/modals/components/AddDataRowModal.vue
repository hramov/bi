<script setup lang="ts">

</script>

<template>
  <v-row>
    <v-col cols="12">
      <label>Ряды данных</label>
      <v-btn variant="text" icon="mdi-plus"></v-btn>
    </v-col>
  </v-row>

  <v-expansion-panels>
    <v-expansion-panel>
      <v-expansion-panel-title>
        {{ 'Новый ряд' }}
      </v-expansion-panel-title>
      <v-expansion-panel-text style="padding-top: 10px">
        <v-row>
          <v-col cols="2">
            <v-select label="Источник" />
          </v-col>
          <v-col cols="2">
            <v-select label="Тип" />
          </v-col>
          <v-col cols="2">
            <v-text-field label="Поле в объекте данных" />
          </v-col>
          <v-col cols="3">
            <v-text-field label="Название ряда на графике" />
          </v-col>
          <v-col cols="3">
            <v-select label="Функция форматирования" />
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="6">
            <v-textarea label="Запрос"></v-textarea>
            <v-btn variant="text">Выполнить</v-btn>
          </v-col>
          <v-col cols="6">
            <v-textarea label="Результат"></v-textarea>
          </v-col>
        </v-row>
      </v-expansion-panel-text>
    </v-expansion-panel>
  </v-expansion-panels>
</template>