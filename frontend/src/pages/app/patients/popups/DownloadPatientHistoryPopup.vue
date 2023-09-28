<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import apiRequest from '@/services/http/api-requests';
import PopUp from '@/components/PopUp.vue';
import { popupStore } from '@/stores';
import TextField from '@/components/form/TextField.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';
import dayjs from 'dayjs';

defineProps<{ popupId: string }>()
const historyRef = ref(null);
const route = useRoute();
const router = useRouter();
const form = ref({ timeline: 'thisyear', date_from: '', date_to: '' })

const downloadHistory = async () => {
  // if(form.value.timeline !== 'alltime') {
    form.value.date_from = dayjs(form.value.date_from).format('YYYY-MM-DDTHH:mm:ss[Z]');
    form.value.date_to = dayjs(form.value.date_to).format('YYYY-MM-DDTHH:mm:ss[Z]');
  // }

  router.push({ name: 'patient-history-print-view', query: form.value })
  const getRecords = await apiRequest.post(`patients/${route.params.id}/patient-history/all`, form.value);
  if(getRecords.message) {
    popupStore.show = false;
  }
}

onMounted(()  => {
  form.value.date_to = dayjs(new Date()).format('YYYY-MM-DD')
  form.value.date_from = dayjs(new Date()).subtract(1, 'month').format('YYYY-MM-DD')
})

</script>
<template>
  <pop-up title="Download Patient History" :id="popupId">
    <form method="POST" v-on:submit.prevent="downloadHistory" ref="historyRef">
      <div class="font-bold">Timeline:</div>
      <div class="flex gap-x-3 mb-2">
        <label><input type="radio" v-model="form.timeline" value="thisyear"/>&nbsp;This year</label>
        <label><input type="radio" v-model="form.timeline" value="alltime"/>&nbsp;All Time</label>
        <label><input type="radio" v-model="form.timeline" value="daterange"/>&nbsp;Date Range</label>
      </div>
      <div class="flex gap-x-3 my-3" v-if="form.timeline == 'daterange'">
        <div class="w-full"><TextField label="From" v-model="form.date_from" required type="date"/></div>
        <div class="w-full"><TextField label="To" v-model="form.date_to" required type="date" /></div>
      </div>
      <!-- <div class=" my-3 font-bold">File Type:</div>
      <div class="flex gap-x-3 mb-2">
        <label><input type="radio" v-model="form.format" value="pdf"/>&nbsp;PDF</label>
        <label><input type="radio" v-model="form.format" value="csv"/>&nbsp;CSV</label>
        <label><input type="radio" v-model="form.format" value="html"/>&nbsp;HTML</label>
      </div> -->
      <div class="flex gap-3 mt-5">
        <div><PrimaryButton type="submit">Download</PrimaryButton></div>
        <div><SecondaryButton type="button" v-on:click.prevent="()=>popupStore.show=false">Cancel</SecondaryButton></div>
      </div>
    </form>
  </pop-up>
</template>