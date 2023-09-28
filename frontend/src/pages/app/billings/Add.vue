<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import PageHeader from '@/components/PageHeader.vue';
import apiRequest from '@/services/http/api-requests';
import dayjs from 'dayjs'
import TextField from '@/components/form/TextField.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';
import { computed, onMounted, ref, watch } from 'vue';
import { PlusCircleIcon, TrashIcon } from '@heroicons/vue/24/outline';
import { Invoice, InvoiceDetails, Patient } from '@/interfaces';
import { popupStore, toasts } from '@/stores';
import SelectPatientPopup from '@/components/popups/SelectPatientPopup.vue';
import { v4 as uuid } from 'uuid';
import { useRoute, useRouter } from 'vue-router';
import ActionButton from '@/components/ActionButton.vue';
import { PencilIcon } from '@heroicons/vue/24/solid';

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Add Invoice", current: true },
] as BreadcrumbItem[]);

const router = useRouter();
const route = useRoute();
const selectPatientPopupId = ref('selectPatient');
const patient = ref({} as Patient);
const invoice = ref({} as Invoice);
const invoiceDetails = ref({ [uuid()]: {} } as { [key: string]: InvoiceDetails });
const subTotal = computed((): number => {
  let sum = 0;
  Object.values(invoiceDetails.value).forEach(element => {
    element.amount = (element.qty ?? 0) * (element.price ?? 0)
    sum += element.amount
  });
  return sum
})

const showSelectPatientPopup = () => { popupStore.id = selectPatientPopupId.value; popupStore.show = true; }

const addRow = () => {
  invoiceDetails.value[uuid()] = {} as InvoiceDetails;
}
const removeRow = (key: string) => {
  delete invoiceDetails.value[key]
  if (!Object.values(invoiceDetails.value).length) invoiceDetails.value[uuid()] = {} as InvoiceDetails
}

// const clearRows = () => {
//   invoiceDetails.value = {};
//   invoiceDetails.value[uuid()] = {} as InvoiceDetails
// }

const create = async () => {
  invoice.value.details = [];

  Object.values(invoiceDetails.value).forEach((item) => {
    if ((item.qty as number) > 0) {
      invoice.value.details.push({
        description: item.description,
        qty: item.qty || 0,
        price: item.price,
      });
    }
  });

  const create = await apiRequest.post('invoices', {
    ...invoice.value,
    invoice_date: dayjs(invoice.value.invoice_date as string).format('YYYY-MM-DDTHH:mm:ss[Z]'),
    due_date: dayjs(invoice.value.due_date as string).format('YYYY-MM-DDTHH:mm:ss[Z]'),
  });

  if (create?.id) {
    toasts.addToast({ message: "record created successfully.", type: 'success' });
    router.push({ name: 'view-invoice', params: { id: create.id } })
  }
}

onMounted(async () => {
  if (route.params.patient_id) {
    patient.value = await apiRequest.get(`patients/${route.params.patient_id}`);
  } else {
    const req = await apiRequest.get('patients?query=GUEST');
    patient.value = req.rows[0];
  }
})

watch(() => patient.value.id, () => {
  invoice.value.name = `${patient.value.lastname} ${patient.value.firstname}`
  invoice.value.patient_id = patient.value.id;
  if (patient.value.address) invoice.value.billing_address = `${patient.value.address}`
});
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="New Invoice" :icon-src="PlusCircleIcon"></PageHeader>
    <hr class="ml-4 mr-4" />
    <div class="">
      <form method="POST" v-on:submit.prevent="create" class="p-[15px]">
        <div class="flex justify-between">
          <div class="flex">
            <div class="font-bold pr-3">
              <div class="py-2">Bill To:</div>
              <div class="w-[120px]">
                <ActionButton dark :icon-src="PencilIcon" @click="showSelectPatientPopup">Pick</ActionButton>
              </div>
            </div>
            <div>
              <div class="w-[300px]">
                <TextField label="Patient Name" v-model="invoice.name" required />
              </div>
              <div>
                <TextField label="Billing Address" v-model="invoice.billing_address" required />
              </div>
            </div>
          </div>
          <div class="">
            <div class="space-y-2 flex flex-col">
              <div class="flex items-center">
                <div class="font-bold w-[150px] text-right">Invoice Number:</div>
                <div class="w-[250px] ml-3">
                  <TextField placeholder="Invoice number (optional)" v-model="invoice.invoice_number"></TextField>
                </div>
              </div>
              <div class="flex items-center justify-between">
                <div class="font-bold w-[150px] text-right">Invoice Date:</div>
                <div class="w-[250px] ml-3">
                  <TextField v-model="invoice.invoice_date" type="date" required></TextField>
                </div>
              </div>
              <div class="flex items-center justify-between">
                <div class="font-bold w-[150px] text-right">Due Date:</div>
                <div class="w-[250px] ml-3">
                  <TextField v-model="invoice.due_date" type="date" required></TextField>
                </div>
              </div>
            </div>
          </div>
        </div>
        <table class="w-full my-5">
          <thead class="border-t-[1px] border-stone-400 border-b-[1px] h-[40px]">
            <tr>
              <th class="text-left">Description</th>
              <th class="text-left w-[120px]">Quantity</th>
              <th class="text-left w-[120px]">Price</th>
              <th class="text-right w-[120px]">Amount</th>
              <th class="text-left w-[50px]"></th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="d, key in invoiceDetails" :key="key">
              <td>
                <TextField class="w-full" v-model="d.description" />
              </td>
              <td>
                <TextField class="w-[120px]" v-model.number="d.qty" type="number" />
              </td>
              <td>
                <TextField class="w-[120px]" v-model.number="d.price" type="number" />
              </td>
              <td class="text-right px-2">{{ ((d.qty || 0) * (d.price || 0)).toLocaleString("en-US") }}</td>
              <td class=" w-[50px]">
                <button type="button" title="Delete item" @click="() => removeRow(`${key}`)"
                  class="rounded border-[1px] w-[50px] border-blue-600 py-1 hover:bg-blue-600 hover:text-white active:bg-blue-700 flex items-center justify-center text-center">
                  <TrashIcon class="w-5 h-5" />
                </button>
              </td>
            </tr>
          </tbody>
          <tfoot class="border-t-[1px] border-stone-400 border-b-[1px]">
            <tr>
              <th class="py-2 text-right" colspan="2">SUB TOTAL</th>
              <th colspan="3" class="py-2 text-xl font-bold text-right">{{ subTotal.toLocaleString("en-US") }}</th>
            </tr>
            <tr>
              <th class="py-2 text-right" colspan="2">DISCOUNT</th>
              <th colspan="3" class="py-2 text-xl font-bold text-right pl-3">
                <TextField v-model.number="invoice.discount" type="number" class="text-right appearance-none" />
              </th>
            </tr>
            <tr>
              <th class="py-2 text-right" colspan="2">GRAND TOTAL</th>
              <th colspan="3" class="py-2 text-xl font-bold text-right">{{ (subTotal - (invoice.discount ||
                0)).toLocaleString("en-US") }}</th>
            </tr>
          </tfoot>
        </table>
        <div class="flex space-x-2">
          <div class="flex">
            <PrimaryButton type="submit">Submit</PrimaryButton>
          </div>
          <div class="flex">
            <SecondaryButton type="button" v-on:click.prevent="() => $router.go(-1)">Cancel</SecondaryButton>
          </div>
          <div class="flex">
            <PrimaryButton type="button" @click="addRow">Add Row</PrimaryButton>
          </div>
        </div>
      </form>
    </div>

    <SelectPatientPopup :popup-id="selectPatientPopupId" v-model="patient" />
  </div>
</template>