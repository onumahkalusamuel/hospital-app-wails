<script setup lang="ts">
import { FunctionalComponent, HTMLAttributes, VNodeProps } from 'vue';

defineProps<{
  iconSrc?: FunctionalComponent<HTMLAttributes & VNodeProps>,
  title: string,
  subtitle?: string;
  class?: string;
}>()
</script>
<template>
  <div class="flex py-[10px] px-[15px] justify-between items-center flex-wrap gap-y-5" :class="class">
    <div class="flex">
      <div class="w-[80px] h-[80px] mr-2" v-if="iconSrc">
        <component :is="iconSrc" class="w-[80px] h-[80px] text-[#0078d4]" />
      </div>
      <div class="flex flex-col justify-center">
        <div class=" text-2xl text-[#333]">{{ title }}</div>
        <div v-if="subtitle || $slots.append" class="color-[#555] mt-[8px] flex gap-2">
          {{ subtitle }}
          <slot name="append"></slot>
        </div>
      </div>
    </div>
    <slot></slot>
  </div>
</template>