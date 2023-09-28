<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import ActionButton from '@/components/ActionButton.vue';
import PageHeader from '@/components/PageHeader.vue';
import apiRequest from '@/services/http/api-requests';
import { Patient, Pagination } from '@/interfaces'
import TextField from '@/components/form/TextField.vue';
import { UsersIcon, MagnifyingGlassIcon, TrashIcon, UserPlusIcon } from '@heroicons/vue/24/solid';
import { onMounted, ref, watch } from 'vue';
import { toasts } from '@/stores/toasts';
import Paging from '@/components/Paging.vue';
import { useDebounce } from '@/utils/debounce';
const debounce = useDebounce()

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Patients", current: true },
] as BreadcrumbItem[]);

const pagination = ref({
  limit: 10,
  page: 1,
  rows: [] as Patient[],
  sort_key: '',
  sort_order: 'desc',
  total_pages: 0,
  total_rows: 0,
  query: ''
} as Pagination);

const deletItem = async (id: string)  => {
  const del = await apiRequest.deleteRecord(`patients/${id}`);
  if(del && del.message) toasts.addToast({message: del.message, type: 'success'});
  fetchPatients();
}

const fetchPatients = async () => {
  debounce(async () => {
    const filter = {...pagination.value, rows: null };
    pagination.value = await apiRequest.get(`patients?${new URLSearchParams(filter as never as Record<string, string>)}`);
  });
}

onMounted(() => {fetchPatients()})
watch(() => pagination.value.limit, fetchPatients);
watch(() => pagination.value.page, fetchPatients);
watch(() => pagination.value.query, fetchPatients);
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="Patients" subtitle="Manage patients" :icon-src="UsersIcon">
    </PageHeader>
    <div class="px-[15px] flex justify-between border-t-[1px] border-[#333] py-2">
      <div class="pr-2">
        <ActionButton dark @click="() => $router.push({name: 'add-patient'})" :icon-src="UserPlusIcon">Add patient</ActionButton>
      </div>
      <div class="flex">
        <TextField placeholder="Search" v-model="pagination.query">
          <template #prepend>
            <MagnifyingGlassIcon class="h-5 w-5" />
          </template>
        </TextField>
      </div>
    </div>

    <hr class="ml-4 mr-4" />
    <div class="page-scroll-area">
      <div v-if="!pagination.rows.length">
        <div class="text-center mt-4 mb-4 pt-4">No records found</div>
        <hr/>
      </div>
      <div v-else class="table">
        <div class="table-row table-header">
          <div class="cell-data cell-sn">S/N</div>
          <div class="cell-data cell-size-1">Name</div>
          <div class="cell-data">Card No</div>
          <div class="cell-data">Sex</div>
          <div class="cell-data cell-size-1">Phone</div>
          <div class="cell-data cell-size-1">Current Appnt.</div>
          <div class="cell-data">Actions</div>
        </div>
        <div class="table-body">
          <div class="table-row" v-for="patient, i in pagination.rows" :key="patient.id">
            <div class="cell-data cell-sn">{{ i + 1 }}.</div>
            <div class="cell-data cell-size-1">
              <router-link class="text-blue-600 hover:underline" :to="{name: 'view-patient', params: { id: patient.id }}">
                {{ `${patient.firstname} ${patient.lastname}` }}
              </router-link>
            </div>
            <div class="cell-data">{{ patient.card_no }}</div>
            <div class="cell-data">{{ patient.sex }}</div>
            <div class="cell-data cell-size-1">{{ patient.phone }}</div>
            <div class="cell-data cell-size-1">{{ patient.current_appointment }}</div>
            <div class="cell-data gap-x-2">
              <ActionButton outline v-on:click="deletItem(patient.id)" :icon-src="TrashIcon" >Delete</ActionButton>
            </div>
          </div>
        </div>
      </div>
    </div>
    <Paging v-model="pagination" />
  </div>
</template>