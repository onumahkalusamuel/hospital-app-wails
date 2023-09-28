<script lang="ts" setup>
import { RouterView } from 'vue-router';
</script>;

<template>
  <div class="container">
    <div class="inner-container">
      <router-view></router-view>
    </div>
  </div>
</template>

<style>
.container {
  display:flex;
  align-items: center;
  flex-direction: column;
  justify-content: center;
  height: 100vh;
}
.inner-container {
  max-width: 400px;
  min-width: 350px;
  text-align: center;
  border: 1px solid #888888;
  padding: 3rem 2rem;
  box-sizing: border-box;
}
</style>
