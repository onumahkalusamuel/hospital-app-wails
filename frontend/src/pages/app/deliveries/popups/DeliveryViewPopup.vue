<script setup lang="ts">
import dayjs from 'dayjs';
import PopUp from '@/components/PopUp.vue';
import { Delivery } from '@/interfaces';
import { UserIcon } from '@heroicons/vue/24/solid';
defineProps<{ popupId: string, delivery: Delivery }>()
</script>

<template>
    <pop-up :width="600" title="Delivery Details" :id="popupId">
        <div class="mb-5">
            <div class="flex">
                <div class="w-[80px] h-[80px] mr-2">
                    <component :is="UserIcon" class="w-[80px] h-[80px] text-[#0078d4]" />
                </div>
                <div class="flex flex-col justify-center">
                    <div class=" text-2xl text-[#333]">{{ delivery.patient?.lastname }} {{ delivery.patient?.firstname }} {{
                        delivery.patient?.middlename }}</div>
                    <div class="page-header-subtitle">{{ delivery.patient?.sex }} </div>
                </div>
            </div>
            <div class="grid grid-cols-2 gap-x-2 mt-4">
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Delivery Mode:</div>
                    <div class="">{{ delivery.delivery_mode || "N/A" }}</div>
                </div>
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Baby Sex:</div>
                    <div class="">{{ delivery.baby_sex || "N/A" }}</div>
                </div>
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Baby Weight:</div>
                    <div class="">{{ delivery.baby_weight || "N/A" }} kgs</div>
                </div>
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Condition:</div>
                    <div class="">{{ delivery.condition || "N/A" }}</div>
                </div>
            </div>
            <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                <div class="title">Delivered At:</div>
                <div class="">{{ dayjs(delivery.delivery_date_time).format('DD-MM-YYYY hh:mm A') }}</div>
            </div>
            <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                <div class="title">Note:</div>
                <div class="">{{ delivery.note || "N/A" }}</div>
            </div>
        </div>
    </pop-up>
</template>

<style scoped>
.title {
    color: #0078d4;
    border-radius: 5px;
    font-weight: bold;
    margin-right: 5px;
}
</style>