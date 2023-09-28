<script setup lang="ts">
import {ref} from "vue";

const userStore = ref({
  userAvatar: 'Х. С.',
  userFio: 'Храмов С. И.',
})
</script>

<template>
  <div style="background-color: rgb(33,33,33); height: 65px;">
    <v-card-item>
      <div class="avatar">
        <v-avatar color="warning" size="50"><span>{{ userStore.userAvatar}}</span></v-avatar>
        <div class="fio">
          Добро пожаловать
          <p>{{ userStore.userFio }}
          </p>
        </div>
      </div>
    </v-card-item>
  </div>
</template>

<style scoped>
.avatar {
  display: flex;
  justify-content: space-between;
}

.fio {
  width: 65%;
}
</style>
