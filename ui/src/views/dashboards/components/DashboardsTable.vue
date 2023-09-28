<script setup lang="ts">

import moment from "moment/moment";
import {DateFormat} from "../../../modules/constants";
import {useRouter} from "vue-router";

const props = defineProps(['dash']);
const router = useRouter();

const openSingleDashboard = (id: number) => {
  router.push('/dashboards/' + id)
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
        @click="openSingleDashboard(item.id)"
    >
      <td>{{ item.title }}</td>
      <td>{{ item.description }}</td>
      <td>{{ moment(item.last_updated).format(DateFormat) }}</td>
      <td>
        <div class="actions">
          <v-btn variant="text" icon="mdi-dots-horizontal"></v-btn>
          <v-menu location="bottom" offset="0" activator="parent">
            <v-list>
              <v-list-item
                  v-for="(item, index) in []"
                  :key="index"
                  :value="index"
              >
                <v-list-item-title>{{ item }}</v-list-item-title>
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