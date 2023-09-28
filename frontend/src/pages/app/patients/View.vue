<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import ActionButton from '@/components/ActionButton.vue';
import PageHeader from '@/components/PageHeader.vue';
import apiRequest from '@/services/http/api-requests';
import { Patient } from '@/interfaces'
import HistoryIndex from './HistoryIndex.vue';
import { useRoute } from 'vue-router';
import { onMounted, ref } from 'vue';
import { UserIcon, PencilIcon, EyeIcon, 
  // FolderPlusIcon, FolderMinusIcon,
   UserGroupIcon, BanknotesIcon,
    // PhoneArrowUpRightIcon
   } from '@heroicons/vue/24/solid'
import { popupStore } from '@/stores';
import AdmitPatientPopup from './popups/AdmitPatientPopup.vue';
import DischargePatientPopup from './popups/DischargePatientPopup.vue';
import AppointmentPopup from './popups/AppointmentPopup.vue';
import PatientViewPopup from './popups/PatientViewPopup.vue';

const route = useRoute();

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Patients", link: { name: "patients" } },
] as BreadcrumbItem[]);

const patient = ref({} as Patient);
const admitPatientPopupId = ref('admitpatient');
const dischargePatientPopupId = ref('dischargepatient');
const appointmentPopupId = ref('appointment');
const viewPatientPopupId = ref('viewpatient');
const displayKey = ref(1);

const fetchPatient = async () => {
  patient.value = await apiRequest.get(`patients/${route.params.id}`);
  if (!breadcrumbs.value[breadcrumbs.value.length - 1]?.current) {
    breadcrumbs.value.push({ title: `${patient.value.lastname} ${patient.value.firstname}`, current: true });
  }
  displayKey.value++;
}

// const showAdmitPatientPopup = () => { popupStore.id = admitPatientPopupId.value; popupStore.show = true; }
// const showDischargePatientPopup = () => { popupStore.id = dischargePatientPopupId.value; popupStore.show = true; }
// const showAppointmentPopup = () => { popupStore.id = appointmentPopupId.value; popupStore.show = true; }
const showPatientViewPopup = () => { popupStore.id = viewPatientPopupId.value; popupStore.show = true; }

onMounted(async () => { await fetchPatient() });

</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader :title="`${patient.lastname} ${patient.firstname}`" :subtitle="`${patient.sex} - ${patient.card_no}`"
      :icon-src="UserIcon">

      <div class="">
        <div class="grid grid-cols-2 gap-2">
          <ActionButton dark @click="showPatientViewPopup" :icon-src="EyeIcon">View</ActionButton>
          <ActionButton dark v-on:click="() => $router.push({ name: 'edit-patient' })" :icon-src="PencilIcon">Edit
          </ActionButton>
          <!-- <ActionButton outline v-if="'Admitted' !== patient.current_status" :icon-src="FolderPlusIcon"
            @click="showAdmitPatientPopup">Admit</ActionButton>
          <ActionButton outline v-if="'Admitted' === patient.current_status" :icon-src="FolderMinusIcon"
            @click="showDischargePatientPopup">Discharge</ActionButton> -->
          <!-- <ActionButton dark :icon-src="PhoneArrowUpRightIcon" @click="showAppointmentPopup">Initiate appt.</ActionButton> -->
          <ActionButton outline v-on:click="() => $router.push({ name: 'patient-deliveries' })" :icon-src="UserGroupIcon">Deliveries
          </ActionButton>
          <ActionButton outline v-on:click="() => $router.push({ name: 'patient-billings' })" :icon-src="BanknotesIcon">Billings
          </ActionButton>
        </div>
      </div>

      <!-- <div class="flex flex-col space-y-2 font-bold text-center" style="justify-content: end;">
        <div class="flex items-center">
          <div class="border-[1px] border-gray-500 px-2 py-1 rounded ml-5 min-w-[150px]" :class="patient.current_appointment == 'Emergency' ?
            'bg-red-500 text-white border-red-500' :
            patient.current_appointment == 'Regular' ?
              'bg-yellow-300' :
              'bg-gray-200'
            ">
            {{ patient.current_appointment || 'No appointment' }}
          </div>
        </div>
        <div class="flex items-center jusitify-end">
          <div class="border-[1px] border-gray-500 px-2 py-1 rounded ml-5 min-w-[150px]" :class="patient.current_status == 'Admitted' ?
            'bg-red-500 text-white border-red-500' :
            patient.current_status == 'Discharged' ?
              'bg-green-500 text-white border-green-500' :
              'bg-gray-200'
            ">
            {{ patient.current_status || 'Not admitted' }}
          </div>
        </div>
      </div> -->
    </PageHeader>
    <HistoryIndex :key="displayKey" />
    <AdmitPatientPopup :popup-id="admitPatientPopupId" @update:closed="fetchPatient" />
    <DischargePatientPopup :popup-id="dischargePatientPopupId" @update:closed="fetchPatient" />
    <AppointmentPopup :popup-id="appointmentPopupId" @update:closed="fetchPatient" />
    <PatientViewPopup :popup-id="viewPatientPopupId" :patient="patient" />

  </div>
</template>