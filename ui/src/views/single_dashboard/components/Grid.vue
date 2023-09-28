<script setup lang="ts">

</script>

<template>
  <div class="grid">
    <v-row>
      <v-col cols="6" class="lt">
        <slot name="lt"></slot>
      </v-col>
      <v-col cols="6" class="rt">
        <slot name="rt"></slot>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="6" class="lb">
        <slot name="lb"></slot>
      </v-col>
      <v-col cols="6" class="rb">
        <slot name="rb"></slot>
      </v-col>
    </v-row>
  </div>
</template>

<style scoped>
  .lt {
    margin-top: 5px;
    height: calc(50vh - 30px);
    border-right: 2px dashed grey;
    border-bottom: 2px dashed grey;
    padding: 0;
  }

  .rt {
    margin-top: 5px;
    height: calc(50vh - 30px);
    border-bottom: 2px dashed grey;
    padding: 0;
  }

  .lb {
    height: calc(50vh - 30px);
    border-right: 2px dashed grey;
    padding: 0;
  }

  .rb {
    height: calc(50vh - 30px);
    padding: 0;
  }
</style>