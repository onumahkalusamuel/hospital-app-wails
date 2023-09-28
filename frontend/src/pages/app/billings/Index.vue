<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import ActionButton from '@/components/ActionButton.vue';
import PageHeader from '@/components/PageHeader.vue';
import apiRequest from '@/services/http/api-requests';
import { Pagination, Invoice } from '@/interfaces'
import TextField from '@/components/form/TextField.vue';
import { PlusIcon, MagnifyingGlassIcon, TrashIcon, BanknotesIcon } from '@heroicons/vue/24/solid';
import { onMounted, ref, watch } from 'vue';
import dayjs from 'dayjs';
import { user, toasts } from '@/stores';
import Paging from '@/components/Paging.vue';
import { useDebounce } from '@/utils/debounce';
const debounce = useDebounce()

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Billings", current: true },
] as BreadcrumbItem[]);

const pagination = ref({
  limit: 10,
  page: 1,
  rows: [] as Invoice[],
  sort_key: '',
  sort_order: 'desc',
  total_pages: 0,
  total_rows: 0,
  query: ''
} as Pagination);


const deletItem = async (id: string)  => {
  const del = await apiRequest.deleteRecord(`invoices/${id}`);
  if(del && del.message) toasts.addToast({message: del.message, type: 'success'});
  fetchInvoices();
}

const fetchInvoices =async () => {
  debounce(async () => {
    const filter = {...pagination.value, rows: null };
    pagination.value = await apiRequest.get(`invoices?${new URLSearchParams(filter as never as Record<string, string>)}`);
  });
}

onMounted(async() => {fetchInvoices()})
watch(() => pagination.value.limit, fetchInvoices);
watch(() => pagination.value.page, fetchInvoices);
watch(() => pagination.value.query, fetchInvoices);
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="Billings" subtitle="Manage billings" :icon-src="BanknotesIcon">
    </PageHeader>
    <div class="px-[15px] flex justify-between border-t-[1px] border-[#333] py-2">
      <div class="pr-2">
        <ActionButton dark v-on:click="() => $router.push({name: 'add-invoice'})" :icon-src="PlusIcon">Add invoice</ActionButton>
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
          <div class="cell-data">Amount</div>
          <div class="cell-data">Status</div>
          <div class="cell-data cell-size-1">Due Date</div>
          <div class="cell-data">Actions</div>
        </div>
        <div class="table-body">
          <div class="table-row" v-for="invoice, i in pagination.rows" :key="invoice.id">
            <div class="cell-data cell-sn">{{ i + 1 }}</div>
            <div class="cell-data cell-size-1">
              <router-link class="text-blue-600 hover:underline" :to="{name: 'view-invoice', params: { id: invoice.id }}">
                {{ `${invoice.patient.firstname} ${invoice.patient.lastname}` }}
              </router-link>
            </div>
            <div class="cell-data">{{ invoice.amount.toLocaleString("en-US") }}</div>
            <div class="cell-data">{{ invoice.completed ? 'paid' : 'pending' }}</div>
            <div class="cell-data cell-size-1">{{ dayjs(invoice.due_date).format('DD-MM-YYYY hh:mm A') }}</div>
            <div class="cell-data gap-x-2">
              <ActionButton outline v-if="user.role == '1'" v-on:click="deletItem(invoice.id)" :icon-src="TrashIcon">Delete</ActionButton>
            </div>
          </div>
        </div>
      </div>
    </div>
    <Paging v-model="pagination" />
  </div>
</template>