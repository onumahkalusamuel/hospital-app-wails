<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import PageHeader from '@/components/PageHeader.vue';
import apiRequest from '@/services/http/api-requests';
import dayjs from 'dayjs'
import TextField from '@/components/form/TextField.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';
import { onMounted, ref } from 'vue';
import { router } from '@/router';
import { PlusCircleIcon } from '@heroicons/vue/24/outline';
import { Patient } from '@/interfaces';
import { toasts, popupStore } from '@/stores';
import SelectPatientPopup from './popups/SelectPatientPopup.vue';
import { useRoute } from 'vue-router';
import { PencilIcon } from '@heroicons/vue/24/solid';
import CheckboxField from '@/components/form/CheckboxField.vue';

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Deliveries", link: { name: "deliveries" } },
  { title: "Add Delivery", current: true },
] as BreadcrumbItem[]);

const createForm = ref(null);
const selectPatientPopupId = ref('selectPatient');
const patient = ref({} as Patient);
const route = useRoute();
const form = ref({
  baby_sex: '',
  delivery_mode: '',
  condition: ''
});

const showSelectPatientPopup = () => { popupStore.id = selectPatientPopupId.value; popupStore.show = true; }

const create = async () => {
  if (!patient.value.id) {
    showSelectPatientPopup();
    return;
  } 
  const formData = Object.fromEntries(new FormData(createForm.value as never as HTMLFormElement).entries())
  if (formData.delivery_date_time) {
    formData.delivery_date_time = dayjs(formData.delivery_date_time as string).format('YYYY-MM-DDTHH:mm:ss[Z]');
  }

  const create = await apiRequest.post('deliveries', formData);
  if (create?.id) {
    toasts.addToast({ message: "record created successfully.", type: 'success' });
    router.push({ name: 'deliveries' })
  }
}

onMounted(async () => {
  if (route.params.patient_id) {
    patient.value = await apiRequest.get(`patients/${route.params.patient_id}`);
  }
})

</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="New Delivery" :icon-src="PlusCircleIcon">
      <template #append>
        <div v-if="patient.id">{{ `${patient.lastname} ${patient.firstname} - ${patient.card_no}` }}</div>
        <div v-else>Choose Patient ==> </div>
        <PencilIcon title="Select Patient"
          class="w-6 w-6 cursor-pointer rounded p-1 hover:bg-blue-800 bg-[#0078d4] text-white"
          @click="showSelectPatientPopup" />
      </template>
    </PageHeader>
    <hr class="ml-4 mr-4" />
    <div class="page-scroll-area">
      <form method="POST" class="p-5" v-on:submit.prevent="create" ref="createForm">
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 mb-4 gap-3">
          <input type="hidden" name="patient_id" v-model="patient.id" />
          <CheckboxField :options="['Vaginal', 'C-section']" name="delivery_mode" label="Delivery Mode" required />
          <CheckboxField :options="['Male', 'Female']" name="baby_sex" label="Baby Sex" required />
          <CheckboxField :options="['Baby cried', 'Baby did not cry']" name="condition" label="Condition" required />
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 mt-5">
          <div>
            <TextField label="Baby Weight (kg)" placeholder="5.43" name="baby_weight" required type="number" step="any">
            </TextField>
          </div>
          <div>
            <TextField label="Delivery Date/Time" name="delivery_date_time" type="datetime-local" required></TextField>
          </div>
          <div>
            <TextField label="Note" placeholder="Safe delivery" name="note"></TextField>
          </div>
        </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
          <div class="flex gap-2 mt-5">
            <PrimaryButton type="submit">Submit</PrimaryButton>
            <SecondaryButton type="button" v-on:click.prevent="() => $router.go(-1)">Cancel</SecondaryButton>
          </div>
        </div>
      </form>
    </div>

    <SelectPatientPopup :popup-id="selectPatientPopupId" v-model="patient" />
  </div>
</template>