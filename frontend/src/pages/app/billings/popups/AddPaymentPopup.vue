<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import { useRoute } from 'vue-router';
import apiRequest from '@/services/http/api-requests';
import PopUp from '@/components/PopUp.vue';
import { toasts, popupStore } from '@/stores';
import { Payment, Invoice } from '@/interfaces'
import TextField from '@/components/form/TextField.vue';
import SelectField from '@/components/form/SelectField.vue';
import TextArea from '@/components/form/TextArea.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';

defineProps<{ popupId: string }>()
const route = useRoute();
const payment = ref({} as Payment)
const invoice = ref({} as Invoice)

const addPayment = async () => {
  payment.value.patient_id = invoice.value.patient?.id || '';
  payment.value.balance = (invoice.value.balance || 0) - payment.value.amount_paid;

  const add = await apiRequest.post(`invoices/${route.params.id}/payments`, payment.value);
  console.log(add)

  if(add.id) {
    toasts.addToast({message: "payment added successfully.", type: 'success'});
    popupStore.show = false;
    await fetchInvoice();
  }
}

const fetchInvoice = async () => {
  invoice.value = await apiRequest.get(`invoices/${route.params.id}`);
}

onMounted(async() => { await fetchInvoice() })
watch(() => popupStore.show, async (newValue) => { if (newValue === true) await fetchInvoice(); })
</script>
<template>
  <pop-up title="Add payment to order" :id="popupId">
    <form method="POST" v-on:submit.prevent="addPayment">
      <div class="text-center font-bold py-3">Balance: {{ ((invoice.balance || 0) - (payment.amount_paid || 0)).toLocaleString('en-US') }}</div>
      <div class="flex gap-x-3">
        <div class="w-full"><TextField label="Payment Amount" placeholder="100000" v-model.number="payment.amount_paid" type="number" name="amount_paid" required :max="invoice.balance"></TextField></div>
        <div class="w-full"><SelectField label="Payment Status" v-model="payment.status" name="payment.status" required :options="[['paid'], ['waived']]"></SelectField></div>
      </div>
      <div class="flex">
        <div class="w-full"><TextField label="Payment Reference" placeholder="XFyg3sstr" v-model="payment.payment_reference" name="payment_reference"></TextField></div>
      </div>
      <div class="flex">
        <div class="w-full"><TextArea label="Note" placeholder="Note" v-model="payment.note" name="note"></TextArea></div>
      </div>
      <div class="flex gap-3 mt-5">
        <div><PrimaryButton type="submit">Submit</PrimaryButton></div>
        <div><SecondaryButton type="button" @click.prevent="()=>popupStore.show=false">Cancel</SecondaryButton></div>
      </div>
    </form>
  </pop-up>
</template>