<template>
  <v-app-bar flat>
    <v-app-bar-title>

      <router-link to="/" @click="setActive(0, '/')">
        <v-icon icon="mdi-circle-slice-4" />

        Blogging System

      </router-link>
    </v-app-bar-title>

    <v-list >
      <v-list-item>
        <v-list-item-title>
          <router-link to="/"> Home
          </router-link>
        </v-list-item-title>
      </v-list-item>
      <v-list-item>
        <v-list-item-title>
          <router-link to="/">            <v-icon icon="mdi-account"></v-icon>


          </router-link>
        </v-list-item-title>
      </v-list-item>
    </v-list>
  </v-app-bar>
</template>

<script setup>
  //
</script>
