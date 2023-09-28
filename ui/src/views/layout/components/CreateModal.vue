<script setup lang="ts">
const props = defineProps(['dialog', 'title', 'width']);
</script>

<template>
  <div class="text-center">
    <v-dialog
        :model-value="props.dialog"
        :width="props.width || 500"
        persistent
    >
      <v-card>
        <v-card-text>{{ props.title }}</v-card-text>
        <v-card-text>
          <slot name="form"></slot>
        </v-card-text>
        <v-card-actions>
          <div class="actions_container">
            <slot name="actions"></slot>
          </div>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<style scoped>
  .actions_container {
    display: flex;
    justify-content: flex-end;
    width: 100%;
    padding-left: 15px;
    padding-right: 15px;
  }
</style>