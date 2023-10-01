<script setup lang="ts">

import moment from "moment/moment";
import {DateFormat} from "../../../modules/constants";
import {useRouter} from "vue-router";

const props = defineProps(['dash']);
const router = useRouter();

const openSingleDashboard = (dash_id: string) => {
  router.push('/dashboards/' + dash_id)
}
</script>

<template>
  <v-table
      fixed-header
  >
    <thead>
    <tr>
      <th class="text-center">
        Название
      </th>
      <th class="text-center">
        Описание
      </th>
      <th class="text-center">
        Создан
      </th>
      <th class="text-center">
        Последнее изменение
      </th>
      <th class="text-center">

      </th>
    </tr>
    </thead>
    <tbody>
    <tr
        v-for="item in props.dash"
        :key="item.id"
        class="text-center clickable"
        @click="openSingleDashboard(item.dash_id)"
    >
      <td>{{ item.title }}</td>
      <td>{{ item.description }}</td>
      <td>{{ moment(item.created_at).format(DateFormat) }}</td>
      <td>{{ moment(item.updated_at).year() > 2000 ? moment(item.updated_at).format(DateFormat) : '' }}</td>
      <td>
        <div class="actions">
          <v-btn variant="text" icon="mdi-dots-horizontal"></v-btn>
          <v-menu location="bottom" offset="0" activator="parent">
            <v-list>
              <v-list-item>
                <v-list-item-title>Доступ</v-list-item-title>
              </v-list-item>
              <v-list-item>
                <v-list-item-title>Удалить</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>
        </div>
      </td>
    </tr>
    </tbody>
  </v-table>
</template>

<style scoped>
.clickable {
  cursor: pointer;
}

.clickable:hover {
  background-color: #2c353a
}

.actions {

}
</style>