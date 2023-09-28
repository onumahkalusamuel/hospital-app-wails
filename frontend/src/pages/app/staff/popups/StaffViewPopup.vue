<script setup lang="ts">
import dayjs from 'dayjs';
import PopUp from '@/components/PopUp.vue';
import { Staff } from '@/interfaces';
import { UserIcon } from '@heroicons/vue/24/solid';
defineProps<{ popupId: string, staff: Staff }>()
</script>

<template>
    <pop-up :width="600" title="Staff Details" :id="popupId">
        <div class="mb-5">
            <div class="flex">
                <div class="w-[80px] h-[80px] mr-2">
                    <component :is="UserIcon" class="w-[80px] h-[80px] text-[#0078d4]" />
                </div>
                <div class="flex flex-col justify-center">
                    <div class=" text-2xl text-[#333]">{{ staff.lastname }} {{ staff.firstname }} {{ staff.middlename }}</div>
                    <div class="page-header-subtitle">{{ staff.sex }} - {{ staff.role == 2 ? 'Doctor' : 'Nurse' }} </div>
                </div>
            </div>
            <div class="grid grid-cols-2 gap-x-2 mt-4">

                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Phone:</div>
                    <div class="">{{ staff.phone || "N/A" }}</div>
                </div>
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Email:</div>
                    <div class="">{{ staff.email || "N/A" }}</div>
                </div>
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Username:</div>
                    <div class="">{{ staff.username || "N/A" }}</div>
                </div>
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Created At:</div>
                    <div class="">{{ dayjs(staff.created_at).format('DD-MM-YYYY hh:mm A') }}</div>
                </div>
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