<script lang="ts" setup>
import ActionButton from '@/components/ActionButton.vue';
import apiRequest from '@/services/http/api-requests';
import { Delivery } from '@/interfaces';
import { DocumentArrowDownIcon } from '@heroicons/vue/24/solid';
import { XCircleIcon } from '@heroicons/vue/24/outline';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import dayjs from 'dayjs';
import { GeneratePDF } from '@/utils/generate-pdf'
import { hospital } from '@/stores';

const route = useRoute();
const records = ref([] as Delivery[])
const printArea = ref(null);

const fetchDeliveries = async () => {
    records.value = await apiRequest.post(`deliveries/all`, route.query);
}

const downloadPdf = () => {
    const filename = `deliveries-${route.query.timeline}.pdf`;
    GeneratePDF(printArea.value as never as HTMLElement, filename);
}

onMounted(async () => {
    // get hospital
    const req = await apiRequest.get("hospital-details");
    if (req) hospital.setAll(req);
    await fetchDeliveries();
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
                    <div> {{ records.length }} deliveries</div>
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
            <div class="text-2xl text-center">Deliveries</div>
            <table class="w-full border-separate border text-left">
                <tr>
                    <th class="border-[1px] border-black p-1 w-[45px] text-center">S/N</th>
                    <th class="border-[1px] border-black p-1 w-[150px]">Delivery Time</th>
                    <th class="border-[1px] border-black p-1 w-[200px]">Patient Name</th>
                    <th class="border-[1px] border-black p-1 w-[100px]">Baby Sex</th>
                    <th class="border-[1px] border-black p-1 w-[100px]">Weight</th>
                    <th class="border-[1px] border-black p-1 w-[100px]">Condition</th>
                    <th class="border-[1px] border-black p-1 w-[150px]">Note</th>
                </tr>
                <tr v-if="records.length" v-for="delivery, i in records" :key="delivery.id">
                    <td class="border-[1px] border-black p-1 text-center">{{ i + 1 }}</td>
                    <td class="border-[1px] border-black p-1">
                        {{ dayjs(delivery.delivery_date_time).format('DD-MM-YYYY hh:mm A') }}
                    </td>
                    <td class="border-[1px] border-black p-1">
                        {{ `${delivery.patient?.lastname} ${delivery.patient?.firstname} ${delivery.patient?.middlename} - ${delivery.patient?.card_no}` }}
                    </td>
                    <td class="border-[1px] border-black p-1">{{ delivery.baby_sex }}</td>
                    <td class="border-[1px] border-black p-1">{{ delivery.baby_weight }} kg</td>
                    <td class="border-[1px] border-black p-1">{{ delivery.condition }}</td>
                    <td class="border-[1px] border-black p-1">{{ delivery.note }}</td>
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