<script setup lang="ts">
import {useAppStore} from "../../../modules/store/app.store.ts";

const emit = defineEmits(['drawer']);
const props = defineProps(['title'])
const store = useAppStore();

</script>

<template>
  <v-app-bar>
    <v-app-bar-nav-icon @click="emit('drawer')"></v-app-bar-nav-icon>

    <v-breadcrumbs v-if="props.title" :items="[...store.calcBreadcrumbs, props.title]">
      <template v-slot:divider>
        <v-icon icon="mdi-chevron-right"></v-icon>
      </template>
    </v-breadcrumbs>

    <v-breadcrumbs v-else :items="store.calcBreadcrumbs">
      <template v-slot:divider>
        <v-icon icon="mdi-chevron-right"></v-icon>
      </template>
    </v-breadcrumbs>

    <v-app-bar-title></v-app-bar-title>

    <v-toolbar-items>
      <slot name="toolbar"></slot>
    </v-toolbar-items>
  </v-app-bar>

  <v-main>
    <slot name="content"></slot>

    <slot name="modals"></slot>
  </v-main>
</template>