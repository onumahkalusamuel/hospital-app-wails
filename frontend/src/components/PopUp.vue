<script setup lang="ts">
import { XMarkIcon } from '@heroicons/vue/24/solid'
import { popupStore } from '@/stores';

defineProps<{ id: string; title: string, width?: number, popClass?: string }>();

const close = () => {
  closeChild()
  popupStore.id = '';
  popupStore.show = false;
}
const closeChild = () => {
  popupStore.child_id = '';
  popupStore.child_show = false;
}

</script>
<template>
  <div class="bg-[#00000043] h-screen fixed top-0 left-0 w-full flex justify-center items-center"
    v-if="popupStore.show && popupStore.id == id" @click.self="close">
    <div :class="width ? `w-[${width}px] ${popClass}` : `w-[500px] ${popClass}`"
      class="bg-white lg:min-w-[500px] min-h-[50px] mb-2 p-2 border-[1px] rounded">
      <div class="header flex justify-between items-center w-full px-2">
        <div class="font-bold pt-2 pb-1">{{ title }}</div>
        <XMarkIcon class="w-6 h-6 cursor-pointer hover:bg-stone-300 rounded" @click="close" />
      </div>
      <div class="my-2 border-t-[1px] border-gray-300"></div>
      <div class="max-h-[85vh] pb-5 popup-scroll-area">
        <slot></slot>
      </div>
    </div>
  </div>

  <div class="bg-[#000000cc] h-screen fixed top-0 left-0 w-full flex justify-center items-center"
    v-if="popupStore.child_show && popupStore.child_id == id" @click.self="closeChild">
    <div :class="width ? `w-[${width}px] ${popClass}` : `w-[1020px] ${popClass}`"
      class="bg-white lg:min-w-[500px] min-h-[500px] mb-2 p-2 border-[1px] rounded">
      <div class="header flex justify-between items-center w-full px-2">
        <div class="font-bold pt-2 pb-1">{{ title }}</div>
        <XMarkIcon class="w-6 h-6 cursor-pointer hover:bg-stone-300 rounded" @click="closeChild" />
      </div>
      <div class="my-2 border-t-[1px] border-gray-300"></div>
      <div class="max-h-[85vh] pb-5 popup-scroll-area">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<style scoped>
.popup-scroll-area {
  overflow-y: scroll;
  scrollbar-color: white;
  scrollbar-width: none;
}

.popup-scroll-area::-webkit-scrollbar {
  display: none;
}
</style>