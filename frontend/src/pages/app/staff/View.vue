<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import PageHeader from '@/components/PageHeader.vue';
import apiRequest from '@/services/http/api-requests';
import { Staff } from '@/interfaces'
import SelectField from '@/components/form/SelectField.vue';
import TextField from '@/components/form/TextField.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';
import { useRoute, useRouter } from 'vue-router';
import { onMounted, ref } from 'vue';
import { UserIcon } from '@heroicons/vue/24/solid'
import { toasts } from '@/stores/toasts';

const route = useRoute();
const router = useRouter();

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Staff", link: { name: "staff" } },
] as BreadcrumbItem[]);

const staff = ref({} as Staff);
const form = ref(null);

const fetchStaff = async (first: boolean = false) => {
  staff.value = await apiRequest.get(`staff/${route.params.id}`);
  if (!first) {
    breadcrumbs.value.pop();
  }
  breadcrumbs.value.push({ title: `${staff.value.lastname} ${staff.value.firstname}`, current: true });
}
const update = async () => {
  const formData = Object.fromEntries(new FormData(form.value as never as HTMLFormElement).entries())
  const update = await apiRequest.put(`staff/${route.params.id}`, { ...formData, role: parseInt(formData.role as string) });
  if (update.message) {
    toasts.addToast({ message: update.message, type: 'success' });
    router.push({ name: 'staff' })
    await fetchStaff();
  }
}

onMounted(async () => { await fetchStaff(true) });

</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader :title="`${staff.lastname} ${staff.firstname}`" :icon-src="UserIcon"></PageHeader>
    <div style="padding: 0 15px; display: flex; justify-content: space-between; border-bottom:1px solid #333">
    </div>
    <div class="page-scroll-area">
      <form method="POST" v-on:submit.prevent="update" class="p-5" ref="form">
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div>
            <SelectField label="Role" name="role" v-model="staff.role" :options="[['Nurse', 3], ['Doctor', 2]]"
              required />
          </div>
          <div>
            <TextField label="First Name" placeholder="John" name="firstname" :value="staff.firstname"></TextField>
          </div>
          <div>
            <TextField label="Middle Name" placeholder="M." name="middlename" :value="staff.middlename"></TextField>
          </div>
          <div>
            <TextField label="Last Name" placeholder="Doe" name="lastname" :value="staff.lastname"></TextField>
          </div>
          <div>
            <SelectField label="Gender" name="sex" :options="[['Male'], ['Female']]" required v-model="staff.sex" />
          </div>
          <div>
            <TextField label="Phone Number" name="phone" type="tel" :value="staff.phone"></TextField>
          </div>
          <div>
            <TextField label="Email Address" name="email" :value="staff.email"></TextField>
          </div>
          <div>
            <TextField label="Password" name="password" placeholder="leave empty if you dont want to change it"></TextField>
          </div>
        </div>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
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

<style scoped></style>