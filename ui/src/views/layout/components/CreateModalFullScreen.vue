<script setup lang="ts">
const props = defineProps(['dialog', 'title']);
</script>

<template>
  <div>
    <v-dialog
        :model-value="props.dialog"
        fullscreen
        :scrim="false"
        transition="dialog-bottom-transition"
    >
      <v-card>
        <v-card-title>{{ props.title }}</v-card-title>
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