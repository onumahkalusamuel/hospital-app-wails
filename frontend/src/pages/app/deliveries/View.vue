<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import PageHeader from '@/components/PageHeader.vue';
import apiRequest from '@/services/http/api-requests';
import dayjs from 'dayjs'
import TextField from '@/components/form/TextField.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';
import { onMounted, ref } from 'vue';
import { Delivery } from '@/interfaces';
import { UserIcon } from '@heroicons/vue/24/solid'
import { useRoute, useRouter } from 'vue-router';
import { toasts } from '@/stores/toasts';
import CheckboxField from '@/components/form/CheckboxField.vue';

const route = useRoute();
const router = useRouter();

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Deliveries", link: { name: "deliveries" } },
  { title: 'Delivery details', current: true }
] as BreadcrumbItem[]);

const delivery = ref({} as Delivery);
const form = ref(null);

const fetchDelivery = async () => {
  delivery.value = await apiRequest.get(`deliveries/${route.params.id}`);
}
const update = async () => {
  const formData = Object.fromEntries(new FormData(form.value as never as HTMLFormElement).entries())
  if (formData.delivery_date_time) {
    formData.delivery_date_time = dayjs(formData.delivery_date_time as string).format('YYYY-MM-DDTHH:mm:ss[Z]');
  }

  const update = await apiRequest.put(`deliveries/${delivery.value.id}`, formData);
  if (update.message) {
    toasts.addToast({ message: update.message, type: 'success' });
    router.push({ name: 'view-delivery', params: { id: update.id } })
  }
}

onMounted(async () => { await fetchDelivery() });
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="Delivery details" :icon-src="UserIcon">
      <template #append> Patient
        <router-link :to="{ name: 'view-patient', params: { id: delivery.patient_id } }">
          {{ `${delivery.patient?.lastname} ${delivery.patient?.firstname} - ${delivery.patient?.card_no}` }}
        </router-link>
      </template>
    </PageHeader>

    <div style="padding: 0 15px; display: flex; justify-content: space-between; border-bottom:1px solid #333">
    </div>
    <div class="page-scroll-area">
      <form method="POST" v-on:submit.prevent="update" class="p-5" ref="form">
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 mb-4 gap-3">
          <CheckboxField :options="['Vaginal', 'C-section']" name="delivery_mode" label="Delivery Mode" required
            :model-value="delivery.delivery_mode" />
          <CheckboxField :options="['Male', 'Female']" name="baby_sex" label="Baby Sex" required
            :model-value="delivery.baby_sex" />
          <CheckboxField :options="['Baby cried', 'Baby did not cry']" name="condition" label="Condition" required
            :model-value="delivery.condition" />
        </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-3">
          <div>
            <TextField label="Baby Weight (kg)" placeholder="5.43" name="baby_weight" required type="number" step="any"
              :value="`${delivery.baby_weight}`"></TextField>
          </div>
          <div>
            <TextField label="Delivery Date/Time" type="datetime-local"
              :value="dayjs(delivery.delivery_date_time).format('YYYY-MM-DD HH:mm')" readonly></TextField>
          </div>
          <div>
            <TextField label="Note" placeholder="Safe delivery" name="note" :value="delivery.note"></TextField>
          </div>
        </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
          <div class="flex gap-2 mt-5">
            <PrimaryButton type="submit" class="w-full">Update</PrimaryButton>
            <SecondaryButton type="button" class="w-full" v-on:click.prevent="() => $router.go(-1)">Cancel
            </SecondaryButton>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>