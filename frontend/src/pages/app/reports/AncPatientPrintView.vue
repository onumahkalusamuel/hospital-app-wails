<script lang="ts" setup>
import ActionButton from '@/components/ActionButton.vue';
import apiRequest from '@/services/http/api-requests';
import { Patient } from '@/interfaces';
import { DocumentArrowDownIcon } from '@heroicons/vue/24/solid';
import { XCircleIcon } from '@heroicons/vue/24/outline';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import dayjs from 'dayjs';
import { hospital } from '@/stores';
import { GeneratePDF } from '@/utils/generate-pdf';

const route = useRoute();
const records = ref([] as Patient[])
const printArea = ref(null);

const fetchPatients = async () => {
    records.value = await apiRequest.post(`patients/anc-patient`, route.query);
}


const downloadPdf = () => {
    const filename = `anc-patients-${route.query.timeline}.pdf`;
    GeneratePDF(printArea.value as never as HTMLElement, filename);
}

onMounted(async () => {
    // get hospital
    const req = await apiRequest.get("hospital-details");
    if (req) hospital.setAll(req);
    await fetchPatients();
})
</script>;

<template>
    <div class="page-wrapper w-[920px] my-5 m-auto">
        <div class="flex justify-between">
            <div class="px-[15px] flex justify-center py-2 gap-2">
                <ActionButton dark v-on:click="downloadPdf" :icon-src="DocumentArrowDownIcon">Download as PDF</ActionButton>
            </div>
            <ActionButton v-on:click="() => $router.go(-1)" :icon-src="XCircleIcon">Exit</ActionButton>
        </div>
        <div class="border-[1px] p-2 border-stone-950 m-2 gap-5 flex flex-col" ref="printArea">
            <div class="flex border-b-[1px] p-2 border-stone-600">
                <div class="m-auto flex gap-3 items-center">
                    <img :src="`${hospital.asset_base_url}/files/logo/${hospital.hospital_logo}`" class="max-w-[100px]"
                        alt="logo">
                    <div>
                        <div class="text-2xl font-bold" style="text-transform:uppercase">{{ hospital.hospital_name }}</div>
                        <div>{{ hospital.hospital_address }}</div>
                        <div>{{ hospital.hospital_phone }}</div>
                    </div>
                </div>
            </div>

            <div class="flex border-b-[1px] p-2 border-stone-600 justify-between">
                <div class="m-auto flex gap-3 items-center">
                    <div class="font-bold">Total:</div>
                    <div> {{ records.length }} patients</div>
                </div>
                <div class="m-auto flex gap-3 items-center">
                    <div class="font-bold">Timeline:</div>
                    <div v-if="route.query.timeline == 'thisyear'">This Year</div>
                    <div v-else-if="route.query.timeline == 'alltime'"> All Time</div>
                    <div v-else> {{ dayjs(route.query.date_from as string).format('DD-MM-YYYY') }} - {{
                        dayjs(route.query.date_to as string).format('DD-MM-YYYY') }}</div>
                </div>
            </div>
            <!-- Record Proper -->
            <div class="text-2xl text-center">Ante Natal Patients</div>
            <table class="w-full border-separate border text-left">
                <tr>
                    <th class="border-[1px] border-black p-1 w-[45px] text-center">S/N</th>
                    <th class="border-[1px] border-black p-1 w-[100px]">Card No</th>
                    <th class="border-[1px] border-black p-1 w-[200px]">Name</th>
                    <th class="border-[1px] border-black p-1 w-[100px]">Gender</th>
                    <th class="border-[1px] border-black p-1 w-[100px]">Age</th>
                    <th class="border-[1px] border-black p-1 w-[100px]">Phone</th>
                    <th class="border-[1px] border-black p-1 w-[100px]">Marital Status</th>
                </tr>
                <tr v-if="records.length" v-for="patient, i in records" :key="patient.id">
                    <td class="border-[1px] border-black p-1 text-center">{{ i + 1 }}</td>
                    <td class="border-[1px] border-black p-1">{{ patient.card_no }}</td>
                    <td class="border-[1px] border-black p-1"> {{ `${patient.lastname} ${patient.firstname}
                                            ${patient.middlename}` }}</td>
                    <td class="border-[1px] border-black p-1">{{ patient.sex }}</td>
                    <td class="border-[1px] border-black p-1">{{ patient.age }} yrs</td>
                    <td class="border-[1px] border-black p-1">{{ patient.phone }}</td>
                    <td class="border-[1px] border-black p-1">{{ patient.marital_status }}</td>
                </tr>
                <tr v-else>
                    <td colspan="7">
                        <div class="text-center mt-4 mb-4 pt-4">No records found</div>
                    </td>
                </tr>
            </table>
        </div>
    </div>
</template>

<style scoped></style>