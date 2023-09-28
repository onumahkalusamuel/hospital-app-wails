<script lang="ts" setup>
import ActionButton from '@/components/ActionButton.vue';
import apiRequest from '@/services/http/api-requests';
import { Invoice, Patient } from '@/interfaces';
import { DocumentArrowDownIcon } from '@heroicons/vue/24/solid';
import { XCircleIcon } from '@heroicons/vue/24/outline';
import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import dayjs from 'dayjs';
import { hospital, user } from '@/stores';
import { GeneratePDF } from '@/utils/generate-pdf'

const route = useRoute();
const patient = ref({} as Patient);
const invoice = ref({} as Invoice);
const router = useRouter();
const printArea = ref(null as never as HTMLElement);

const downloadPdf = () => {
  const filename = `${patient.value.lastname}-${patient.value.firstname}-invoice.pdf`;
  GeneratePDF(printArea.value as never as HTMLElement, filename);
}

const fetchInvoice = async () => {
  invoice.value = await apiRequest.get(`invoices/${route.params.id}`);
  patient.value = invoice.value.patient as Patient

  setTimeout(() => { downloadPdf(); router.go(-1) }, 200);
}

onMounted(async () => {
  // get hospital
  const req = await apiRequest.get("hospital-details");
  if (req) hospital.setAll(req);
  await fetchInvoice()
})
</script>;

<template>
  <div class="page-wrapper w-[840px] my-5 m-auto">
    <div class="flex justify-between">
      <div class="px-[15px] flex justify-center py-2 gap-2">
        <ActionButton dark v-on:click="downloadPdf" :icon-src="DocumentArrowDownIcon">Download as PDF</ActionButton>
      </div>
      <ActionButton v-on:click="() => $router.go(-1)" :icon-src="XCircleIcon">Exit</ActionButton>
    </div>

    <div class="border-[1px] p-2 border-stone-950 m-2" ref="printArea">
      <div class="p-[15px]">
        <div class="flex gap-2">
          <div class="flex gap-3 items-center w-[50%] border-[1px] p-2 border-stone-600">
            <img :src="`${hospital.asset_base_url}/files/logo/${hospital.hospital_logo}`" class="max-w-[50px]" alt="logo">
            <div>
              <div class="text-[12px] font-bold" style="text-transform:uppercase">{{ hospital.hospital_name }}</div>
              <div>{{ hospital.hospital_address }}</div>
              <div>{{ hospital.hospital_phone }}</div>
            </div>
          </div>
          <div class="border-[1px] p-2 border-stone-600 w-[50%]">
            <div class="mb-2 font-bold">Bill To:</div>
            <div>{{ invoice.name }}</div>
            <div>{{ invoice.billing_address }}</div>
          </div>
        </div>
        <div class="font-bold flex justify-between pt-4">
          <div>Invoice No: {{ invoice.invoice_number }}</div>
          <div>Invoice Date: {{ dayjs(invoice.invoice_date).format('DD-MM-YYYY') }}</div>
          <div>Due Date: {{ dayjs(invoice.due_date).format('DD-MM-YYYY') }}</div>
        </div>
        <table class="w-full my-5">
          <thead class="border-t-[1px] border-stone-400 border-b-[1px] h-[40px]">
            <tr>
              <th class="text-center w-[50px]">S/N</th>
              <th class="text-left">Description</th>
              <th class="text-center w-[100px]">Quantity</th>
              <th class="text-right w-[100px]">Price</th>
              <th class="text-right w-[100px]">Amount</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="d, key in invoice.details" :key="key" class="border-t-[1px] border-stone-300 h-[30px]">
              <td class="text-center">{{ key + 1 }}.</td>
              <td>
                <div class="w-full">{{ d.description }}</div>
              </td>
              <td>
                <div class="w-[120px] text-center">{{ d.qty.toLocaleString("en-US") }}</div>
              </td>
              <td>
                <div class="text-right w-[120px]">{{ d.price.toLocaleString("en-US") }}</div>
              </td>
              <td class="text-right">{{ d.amount?.toLocaleString("en-US") }}</td>
            </tr>
          </tbody>
          <tfoot class="border-t-[1px] border-stone-400 border-b-[1px]">
            <tr>
              <th class="text-right" colspan="5">SUB TOTAL</th>
              <th colspan="2" class="font-bold text-right">{{ invoice.sub_total?.toLocaleString("en-US") }}</th>
            </tr>
            <tr>
              <th class="text-right" colspan="5">DISCOUNT</th>
              <th colspan="2" class="font-bold text-right">{{ (invoice.discount || 0).toLocaleString("en-US") }}</th>
            </tr>
            <tr>
              <th class="text-right" colspan="5">GRAND TOTAL</th>
              <th colspan="2" class="font-bold text-right">{{ invoice.amount?.toLocaleString("en-US") }}</th>
            </tr>
          </tfoot>
        </table>
        <br />
        <div class="text-[8px]">Printed by: {{ user.username }};<br />Date: {{ new Date().toLocaleDateString('en-us', {
          weekday: "short", year: "numeric", month: "short", day: "numeric", hour: "2-digit", minute: "2-digit", second:
            "2-digit"
        }) }}</div>
      </div>
    </div>

  </div>
</template>

<style scoped></style>