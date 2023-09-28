<script lang="ts" setup>
import ActionButton from '@/components/ActionButton.vue';
import apiRequest from '@/services/http/api-requests';
import { Pagination, Patient, PatientHistory } from '@/interfaces'
import { ArrowDownIcon, PencilIcon, TrashIcon, BookOpenIcon } from '@heroicons/vue/24/solid';
import { onMounted, ref, watch } from 'vue';
import { popupStore, toasts, user } from '@/stores';
import { useRoute } from 'vue-router';
import dayjs from 'dayjs';
import Paging from '@/components/Paging.vue';
import { useDebounce } from '@/utils/debounce';
import DownloadPatientHistoryPopup from './popups/DownloadPatientHistoryPopup.vue';
import HistoryViewPopup from './popups/HistoryViewPopup.vue';
import AddPatientHistoryPopup from './popups/AddPatientHistoryPopup.vue';
import HistoryEditPopup from './popups/HistoryEditPopup.vue';

const downloadHistoryPopupId = ref('downloadhistory');
const historyPopupId = ref('history');
const addHistoryPopupId = ref('addHistory');
const editHistoryPopupId = ref('editHistory');

const pagination = ref({
  limit: 10,
  page: 1,
  rows: [] as PatientHistory[],
  sort_key: '',
  sort_order: 'desc',
  total_pages: 0,
  total_rows: 0,
  query: ''
} as Pagination);


const route = useRoute();
const debounce = useDebounce();
const patient = ref({} as Patient);
const currentHistory = ref({} as PatientHistory);

const showDownloadHistoryPopup = () => { popupStore.id = downloadHistoryPopupId.value; popupStore.show = true; }
const showAddHistoryPopup = () => { popupStore.id = addHistoryPopupId.value; popupStore.show = true; }
const viewHistory = (history: PatientHistory) => {
  currentHistory.value = history;
  popupStore.id = historyPopupId.value;
  popupStore.show = true;
}
const editHistory = (history: PatientHistory) => {
  currentHistory.value = history;
  popupStore.id = editHistoryPopupId.value;
  popupStore.show = true;
}

const deletItem = async (id: string) => {
  const del = await apiRequest.deleteRecord(`patients/${patient.value.id}/patient-history/${id}`);
  if (del && del.message) toasts.addToast({ message: del.message, type: 'success' });
  await fetchPatientHistory();
}

const fetchPatientHistory = async () => {
  debounce(async () => {
    const filter = { ...pagination.value, rows: null };
    pagination.value = await apiRequest.get(`patients/${route.params.id}/patient-history?${new URLSearchParams(filter as never as Record<string, string>)}`);
  })
}

const fetchPatient = async () => {
  patient.value = await apiRequest.get(`patients/${route.params.id}`);
  await fetchPatientHistory();
}

onMounted(async () => {
  await fetchPatient()
})

watch(() => pagination.value.limit, fetchPatientHistory);
watch(() => pagination.value.page, fetchPatientHistory);
watch(() => pagination.value.query, fetchPatientHistory);
</script>;

<template>
  <div class="px-[15px] border-t-[1px] border-[#333] py-2 flex justify-between items-center">
    <div class=" text-xl">History Details</div>
    <div class="flex gap-2">
      <ActionButton dark @click="showAddHistoryPopup" :icon-src="BookOpenIcon">Add history</ActionButton>
      <ActionButton dark @click="showDownloadHistoryPopup" :icon-src="ArrowDownIcon">Download</ActionButton>
    </div>
  </div>
  <hr />
  <div class="page-scroll-area">
    <div v-if="!pagination.rows.length">
      <div class="text-center mt-4 mb-4 pt-4">No history records found.</div>
      <hr />
    </div>
    <div v-else class="table">
      <div class="table-row table-header">
        <div class="cell-data cell-sn">S/N</div>
        <div class="cell-data cell-size-3">Subject</div>
        <div class="cell-data cell-size-1">Date</div>
        <div class="cell-data">Actions</div>
      </div>
      <div class="table-body">
        <div class="table-row" v-for="history, i in pagination.rows" :key="patient.id">
          <div class="cell-data cell-sn">{{ i + 1 }}.</div>
          <div class="cell-data cell-size-3">
            <router-link class="text-blue-600 hover:underline" to="#" @click="viewHistory(history)">
              {{ history.subject }}
            </router-link>
          </div>
          <div class="cell-data cell-size-1">
            {{ dayjs(history.date_time).format('DD-MM-YYYY hh:mm A') }}
          </div>
          <div class="cell-data gap-x-2">
            <ActionButton v-on:click="editHistory(history)" :icon-src="PencilIcon" outline>Edit
            </ActionButton>
            <ActionButton outline v-if="user.role == '1'" @click="deletItem(history.id)" :icon-src="TrashIcon">Delete
            </ActionButton>
          </div>
        </div>
      </div>
    </div>
  </div>
  <Paging v-model="pagination" />
  <AddPatientHistoryPopup :popup-id="addHistoryPopupId" :patient="patient" @update:closed="fetchPatientHistory" />
  <DownloadPatientHistoryPopup :popup-id="downloadHistoryPopupId" />
  <HistoryViewPopup :popup-id="historyPopupId" :history="currentHistory" :patient="patient" />
  <HistoryEditPopup :popup-id="editHistoryPopupId" :history="currentHistory" :patient="patient"
    @update:closed="fetchPatientHistory" />
</template>

<style scoped></style>