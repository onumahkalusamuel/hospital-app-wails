<script lang="ts" setup>
import ActionButton from '@/components/ActionButton.vue';
import apiRequest from '@/services/http/api-requests';
import { Invoice, Patient } from '@/interfaces';
import { DocumentArrowDownIcon } from '@heroicons/vue/24/solid';
import { XCircleIcon } from '@heroicons/vue/24/outline';
import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import { hospital, user } from '@/stores';
import { GeneratePDF } from '@/utils/generate-pdf';

const route = useRoute();
const patient = ref({} as Patient);
const invoice = ref({} as Invoice);
const router = useRouter();
const printArea = ref(null);

const downloadPdf = () => {
  const filename = `${patient.value.lastname}-${patient.value.firstname}-receipt.pdf`;
  
  GeneratePDF(printArea.value as never as HTMLElement, filename, 'a7');
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

    <div class="" ref="printArea">
      <div class="w-[220px] bg-white text-stone-950 text-[9px] p-[10px] pt-1">
        <div class="flex gap-2 items-center justify-center border-[1px] border-stone-800 p-1">
          <img :src="`${hospital.asset_base_url}/files/logo/${hospital.hospital_logo}`" class="max-w-[40px]" alt="logo">
          <div>
            <div class="text-[12px]">{{ hospital.hospital_name }}</div>
            <div>{{ hospital.hospital_address }}</div>
            <div>{{ hospital.hospital_phone }}</div>
          </div>
        </div>
        <div class="text-[12px] text-center py-1">INVOICE {{ invoice.invoice_number }}</div>
        <table>
          <thead class="border-y-[1px] border-black">
            <tr>
              <th class="w-[30px]">QTY</th>
              <th class="w-[100px]">DESCRIPTION</th>
              <th class="w-[40px]">PRICE</th>
              <th class="w-[40px]">AMOUNT</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item, key in invoice.details" :key="key">
              <td class="text-center">{{ item.qty.toLocaleString("en-US") }}</td>
              <td class="w-[1px]">{{ item.description }}</td>
              <td class="w-[1px] text-right">{{ item.price.toLocaleString("en-US") }}</td>
              <td class="w-[1px] text-right">{{ item.amount?.toLocaleString("en-US") }}</td>
            </tr>
          </tbody>
          <tfoot class="border-y-[1px] border-black">
            <tr>
              <td class=""></td>
              <td class="w-[1px]" colspan="2">SUBTOTAL</td>
              <td class="text-right font-bold">{{ (invoice.sub_total || 0).toLocaleString("en-US") }}</td>
            </tr>
            <tr>
              <td class=""></td>
              <td class="w-[1px]" colspan="2">DISCOUNT</td>
              <td class="text-right font-bold">-{{ (invoice.discount || 0).toLocaleString("en-US") }}</td>
            </tr>
            <tr>
              <td class=""></td>
              <td class="w-[1px]" colspan="2">GRAND TOTAL</td>
              <td class="text-right font-bold">{{ (invoice.amount || 0).toLocaleString("en-US") }}</td>
            </tr>
            <tr>
              <td class=""></td>
              <td class="w-[1px]" colspan="2">AMOUNT PAID</td>
              <td class="text-right font-bold">{{ ((invoice.amount || 0) - (invoice.balance || 0)).toLocaleString("en-US")
              }}</td>
            </tr>
            <tr>
              <td class=""></td>
              <td class="w-[1px]" colspan="2">BALANCE</td>
              <td class="text-right font-bold">{{ (invoice.balance || 0).toLocaleString("en-US") }}</td>
            </tr>
          </tfoot>
        </table>
        <div class="mt-[10px]"> Thank you for your patronage</div>
        <div> Trnx total is VAT inclusive</div>
        <div> NO REFUNDS. NO RETURNS. NO EXCHANGES.</div>
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