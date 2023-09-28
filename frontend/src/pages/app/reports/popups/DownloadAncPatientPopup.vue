<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import PopUp from '@/components/PopUp.vue';
import { popupStore } from '@/stores';
import TextField from '@/components/form/TextField.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';
import dayjs from 'dayjs';

defineProps<{ popupId: string }>()
const historyRef = ref(null);
const router = useRouter();
const form = ref({ timeline: 'thisyear', date_from: '', date_to: '' })

const downloadHistory = async () => {
  form.value.date_from = dayjs(form.value.date_from).format('YYYY-MM-DDTHH:mm:ss[Z]');
  form.value.date_to = dayjs(form.value.date_to).format('YYYY-MM-DDTHH:mm:ss[Z]');

  router.push({ name: 'anc-patient-print-view', query: form.value })
}

onMounted(async () => {
  form.value.date_to = dayjs(new Date()).format('YYYY-MM-DD')
  form.value.date_from = dayjs(new Date()).subtract(1, 'month').format('YYYY-MM-DD')
})
</script>
<template>
  <pop-up title="Ante Natal Patients" :id="popupId">
    <form method="POST" v-on:submit.prevent="downloadHistory" ref="historyRef">
      <div class="font-bold text-left py-2">Timeline of registration:</div>
      <div class="flex gap-x-3 mb-2">
        <label><input type="radio" v-model="form.timeline" value="thisyear" />&nbsp;This year</label>
        <label><input type="radio" v-model="form.timeline" value="alltime" />&nbsp;All Time</label>
        <label><input type="radio" v-model="form.timeline" value="daterange" />&nbsp;Date Range</label>
      </div>
      <div class="flex gap-x-3 my-3" v-if="form.timeline == 'daterange'">
        <div class="w-full">
          <TextField label="From" v-model="form.date_from" required type="date" />
        </div>
        <div class="w-full">
          <TextField label="To" v-model="form.date_to" required type="date" />
        </div>
      </div>
      <div class="flex gap-3 mt-5">
        <div>
          <PrimaryButton type="submit">View</PrimaryButton>
        </div>
        <div>
          <SecondaryButton type="button" v-on:click.prevent="() => popupStore.show = false">Cancel</SecondaryButton>
        </div>
      </div>
    </form>
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