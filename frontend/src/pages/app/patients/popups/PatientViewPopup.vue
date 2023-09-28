<script setup lang="ts">
import PopUp from '@/components/PopUp.vue';
import { Patient } from '@/interfaces';
import { UserIcon } from '@heroicons/vue/24/solid';
import dayjs from 'dayjs';

defineProps<{ popupId: string, patient: Patient }>()

</script>
<template>
    <pop-up :width="600" title="Patient Details" :id="popupId">
        <div class="mb-5">
            <div class="flex">
                <div class="w-[80px] h-[80px] mr-2">
                    <component :is="UserIcon" class="w-[80px] h-[80px] text-[#0078d4]" />
                </div>
                <div class="flex flex-col justify-center">
                    <div class=" text-2xl text-[#333]"> {{ patient.title }} {{ patient.lastname }} {{ patient.firstname }}
                        {{ patient.middlename }}</div>
                    <div class="page-header-subtitle">{{ patient.sex }} - {{ patient.card_no }}</div>
                </div>
            </div>
            <div class="grid grid-cols-2 gap-x-2 mt-4">
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Marital Status:</div>
                    <div class="">{{ patient.marital_status || "N/A" }}</div>
                </div>
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Age:</div>
                    <div class="">{{ patient.age || "N/A" }}</div>
                </div>
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Occupation:</div>
                    <div class="">{{ patient.occupation || "N/A" }}</div>
                </div>
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Phone:</div>
                    <div class="">{{ patient.phone || "N/A" }}</div>
                </div>
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Email:</div>
                    <div class="">{{ patient.email || "N/A" }}</div>
                </div>
                <div class="flex mb-2 border-[1px] p-2 rounded"
                    :class="patient.current_status == 'Admitted' ? 'bg-red-100' : 'bg-blue-100'">
                    <div class="title">Status:</div>
                    <div class="">{{ patient.current_status || "N/A" }}</div>
                </div>
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">Registered:</div>
                    <div class="">{{ dayjs(patient.created_at).format('DD-MM-YYYY hh:mm A') || "N/A" }}</div>
                </div>
                <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                    <div class="title">ANC Patient?</div>
                    <div class="">{{ patient.anc || "N/A" }}</div>
                </div>
            </div>
            <div class="flex mb-2 border-[1px] bg-blue-100 p-2 rounded">
                <div class="title">Address:</div>
                <div class="">{{ patient.address || "N/A" }}</div>
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