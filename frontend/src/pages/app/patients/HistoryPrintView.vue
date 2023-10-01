<script lang="ts" setup>
import ActionButton from '@/components/ActionButton.vue';
import apiRequest from '@/services/http/api-requests';
import { Patient, PatientHistory } from '@/interfaces';
import { DocumentArrowDownIcon, EyeIcon } from '@heroicons/vue/24/solid';
import { XCircleIcon } from '@heroicons/vue/24/outline';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import dayjs from 'dayjs';
import { hospital, popupStore } from '@/stores';
import { GeneratePDF } from '@/utils/generate-pdf';
import DocumentViewPopup from '@/components/popups/DocumentViewPopup.vue';

const route = useRoute();
const patient = ref({} as Patient);
const records = ref([] as PatientHistory[])
const printArea = ref(null);
const printing = ref(false);

const fetchPatientHistory = async () => {
  records.value = await apiRequest.post(`patients/${route.params.id}/patient-history/all`, route.query);
}


const fetchPatient = async () => {
  patient.value = await apiRequest.get(`patients/${route.params.id}`);
}

const downloadPdf = () => {
  printing.value = true;
  const filename = `${patient.value.lastname}-${patient.value.firstname}-${route.query.date_from}-${route.query.date_to}.pdf`;
  GeneratePDF(printArea.value as never as HTMLElement, filename);
  setTimeout(() => { printing.value = false; }, 2000)
}

const currentDocumentUrl = ref('');
const viewDocument = (documentUrl: string) => {
  currentDocumentUrl.value = documentUrl;
  popupStore.child_id = 'document';
  popupStore.child_show = true
}


onMounted(async () => {
  // get hospital
  const req = await apiRequest.get("hospital-details");
  if (req) hospital.setAll(req);
  await fetchPatient();
  await fetchPatientHistory();
})
</script>;

<template>
  <div class="page-wrapper w-[840px] my-5 m-auto">
    <div class="px-[10px] flex items-center justify-between lg:justify-center">
      <ActionButton dark v-on:click="downloadPdf" :icon-src="DocumentArrowDownIcon">Download as PDF</ActionButton>
      <ActionButton class="inline lg:fixed lg:top-5 lg:right-5" outline v-on:click="() => $router.go(-1)"
        :icon-src="XCircleIcon">Exit</ActionButton>
    </div>

    <div class="border-[1px] p-2 border-stone-950 m-2 gap-5 flex flex-col" ref="printArea">

      <div class="flex border-b-[1px] p-2 border-stone-600">
        <div class="m-auto flex gap-3 items-center">
          <img :src="`${hospital.asset_base_url}/files/logo/${hospital.hospital_logo}`" class="max-w-[100px]" alt="logo">
          <div>
            <div class="text-2xl font-bold" style="text-transform:uppercase">{{ hospital.hospital_name }}</div>
            <div>{{ hospital.hospital_address }}</div>
            <div>{{ hospital.hospital_phone }}</div>
          </div>
        </div>
      </div>

      <div class="text-2xl text-center">Patient Details</div>
      <table class="w-full border-separate border text-left">
        <tr>
          <th class="border-[1px] border-black p-1 w-[100px]">Name</th>
          <td class="border-[1px] border-black p-1 w-[170px]">{{ `${patient.title} ${patient.lastname},
                      ${patient.firstname} ${patient.middlename}` }}</td>
          <th class="border-[1px] border-black p-1 w-[100px]">Card No.</th>
          <td class="border-[1px] border-black p-1 w-[120px]">{{ patient.card_no }}</td>
          <th class="border-[1px] border-black p-1 w-[120px]">Phone No.</th>
          <td class="border-[1px] border-black p-1 w-[170px]">{{ patient.phone }}</td>
        </tr>
        <tr>
          <th class="border-[1px] border-black p-1">Email</th>
          <td class="border-[1px] border-black p-1">{{ patient.email }}</td>
          <th class="border-[1px] border-black p-1">Sex</th>
          <td class="border-[1px] border-black p-1">{{ patient.sex }}</td>
          <th class="border-[1px] border-black p-1">Marital Status</th>
          <td class="border-[1px] border-black p-1">{{ patient.marital_status }}</td>
        </tr>
        <tr>
          <th class="border-[1px] border-black p-1">Age</th>
          <td class="border-[1px] border-black p-1">
            <div v-if="!!patient.age">{{ patient.age ?? 0 }} yrs</div>
          </td>
          <th class="border-[1px] border-black p-1">Appointment</th>
          <td class="border-[1px] border-black p-1">{{ patient.current_appointment }}</td>
          <th class="border-[1px] border-black p-1">Registered</th>
          <td class="border-[1px] border-black p-1">{{ dayjs(patient.created_at).format('DD-MM-YYYY hh:mm A') }}</td>
        </tr>
      </table>
      <!-- Record Proper -->
      <div class="text-2xl text-center">History Details</div>
      <table class="w-full border-separate border text-left">
        <tr>
          <th class="border-[1px] border-black p-1 w-[45px] text-center">S/N</th>
          <th class="border-[1px] border-black p-1 w-[150px]">Date</th>
          <th class="border-[1px] border-black p-1">Details</th>
        </tr>
        <tr v-if="records.length" v-for="history, i in records" :key="patient.id">
          <td class="border-[1px] border-black text-center align-top p-2">{{ i + 1 }}</td>
          <td class="border-[1px] border-black align-top p-2">
            <div>{{ dayjs(history.date_time).format('DD-MM-YYYY') }}</div>
            <div>{{ dayjs(history.date_time).format('hh:mm A') }}</div>
          </td>
          <td class="border-[1px] border-black p-1">
            <div class="text-xl font-bold uppercase text-center"> {{ history.subject }}</div>
            <div v-for="details, j in history.details" :key="i">
              <div v-if="details.note || details.document" class="border-[1px] border-stone-900 p-2 mt-2">
                <div class="underline uppercase font-bold pb-2">{{ j }}</div>
                <div>{{ details.note }}</div>
                <div v-if="j == 'admission'"> Ward: {{ details.ward_number }}; Room: {{
                  details.room_number }}</div>
                <div v-if="j == 'appointment' && details.appointment_type"> Appt: {{ details.appointment_type }}</div>
                <div v-if="details.document" class="flex justify-right">
                  <div class="py-1 rounded justify-center" v-if="!printing">
                    <div
                      @click="viewDocument(`${hospital.asset_base_url}/files/images_${patient.id}/${details.document}`)"
                      class="cursor-pointer text-white bg-blue-700 hover:bg-blue-800 flex gap-x-2 items-center rounded px-3 py-2">
                      <EyeIcon class="w-5 h-5" />
                      View document
                    </div>
                  </div>
                  <div v-else class="underline">*File attachment*</div>
                </div>
              </div>
            </div>
          </td>
        </tr>
        <tr v-else>
          <td colspan="4">
            <div class="text-center mt-4 mb-4 pt-4">No records found</div>
          </td>
        </tr>
      </table>
    </div>
    <DocumentViewPopup :document-url="currentDocumentUrl" popup-id="document" />
  </div>
</template>

<style scoped></style>