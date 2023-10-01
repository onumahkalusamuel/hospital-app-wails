<script setup lang="ts">
import PopUp from '@/components/PopUp.vue';
import VueQrcode from '@chenfengyuan/vue-qrcode';

defineProps<{ popupId: string; remoteAddress: string }>()

</script>
<template>
    <pop-up title="Login with another device" :id="popupId" class="min-h-[80vh]">
        <div class="flex">
            <div class="m-auto text-xl text-center font-bold text-white bg-[#0078d4] border-[5px] border-white p-[15px]"
                v-if="remoteAddress">
                <vue-qrcode :value="remoteAddress" :options="{ width: 200 }"></vue-qrcode>
                <div class="pt-[10px]">SCAN WITH PHONE</div>
            </div>
        </div>
    </pop-up>
</template>