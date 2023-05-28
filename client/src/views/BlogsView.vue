<template>
  <div class="d-flex align-center flex-column">

    <div v-for='(product) in products' :key='product.id'>
      <v-card class="my-5" width="600" :title="product.id" subtitle="This is a subtitle" :text="product.name"></v-card>

    </div>
  </div>
</template>

<script setup>
import { ref } from "vue"

const products = ref([
  { id: 3, name: 'shirt' },
  { id: 1, name: 'jacket' },
  { id: 2, name: 'shoes' }
])

</script>
