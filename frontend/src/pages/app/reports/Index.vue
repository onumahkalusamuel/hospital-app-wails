<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import PageHeader from '@/components/PageHeader.vue';
import { BookOpenIcon, QueueListIcon, TableCellsIcon, UserGroupIcon } from '@heroicons/vue/24/solid';
import { onMounted, ref } from 'vue';
import DownloadPatientHistoryPopup from './popups/DownloadPatientHistoryPopup.vue';
import { popupStore } from '@/stores';
import DownloadAncPatientPopup from './popups/DownloadAncPatientPopup.vue';
import DownloadDeliveriesPopup from './popups/DownloadDeliveriesPopup.vue';

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Reports", current: true },
] as BreadcrumbItem[]);

const showPatientHistoryPopup = () => { popupStore.id = 'patient-history'; popupStore.show = true }
const showAncPatientPopup = () => { popupStore.id = 'anc-patient'; popupStore.show = true }
const showDeliveriesPopup = () => { popupStore.id = 'deliveries'; popupStore.show = true }

onMounted(() => { })
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="Reports" subtitle="Generate reports" :icon-src="QueueListIcon">
    </PageHeader>
    <div class="px-[15px] flex justify-between border-t-[1px] border-[#333] py-2">
    </div>

    <div class="flex flex-wrap justify-around gap-2 pt-5 pb-10">
      <router-link @click="showPatientHistoryPopup"
        class="flex flex-col justify-center items-center text-white rounded bg-[#0078d4] hover:bg-blue-700 dashboard-body-item"
        to="#">
        <div class="h-[156px] w-[156px]">
          <component :is="BookOpenIcon" />
        </div>
        <div class="mt-5 text-2xl font-bold">Patient History</div>
      </router-link>
      <router-link @click="showAncPatientPopup"
        class="flex flex-col justify-center items-center text-white rounded bg-[#0078d4] hover:bg-blue-700 dashboard-body-item"
        to="#">
        <div class="h-[156px] w-[156px]">
          <component :is="TableCellsIcon" />
        </div>
        <div class="mt-5 text-2xl font-bold">ANC Patients</div>
      </router-link>
      <router-link @click="showDeliveriesPopup"
        class="flex flex-col justify-center items-center text-white rounded bg-[#0078d4] hover:bg-blue-700 dashboard-body-item"
        to="#">
        <div class="h-[156px] w-[156px]">
          <component :is="UserGroupIcon" />
        </div>
        <div class="mt-5 text-2xl font-bold">Deliveries</div>
      </router-link>
    </div>

  </div>
  <DownloadPatientHistoryPopup popup-id="patient-history" />
  <DownloadAncPatientPopup popup-id="anc-patient" />
  <DownloadDeliveriesPopup popup-id="deliveries" />
</template>

<style scoped>
.dashboard-body-item {
  width: 212px;
  height: 259px;
  margin-top: 47px;
  cursor: pointer;
  text-decoration: none;
}
</style>