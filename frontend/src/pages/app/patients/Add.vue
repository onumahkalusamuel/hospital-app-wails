<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import PageHeader from '@/components/PageHeader.vue';
import apiRequest from '@/services/http/api-requests';
import TextField from '@/components/form/TextField.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';
import { ref } from 'vue';
import { router } from '@/router';
import { UserPlusIcon } from '@heroicons/vue/24/outline';
import { toasts } from '@/stores/toasts';
import CheckboxField from '@/components/form/CheckboxField.vue';

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Patients", link: { name: "patients" } },
  { title: "Add Patient", current: true },
] as BreadcrumbItem[]);

const createForm = ref(null);

const form = ref({
  sex: '',
  anc: '',
  marital_status: ''
});
const create = async () => {
  const formData = Object.fromEntries(new FormData(createForm.value as never as HTMLFormElement).entries())
  const create = await apiRequest.post('patients', formData);
  if (create.id) {
    toasts.addToast({ message: "account created successfully.", type: 'success' });
    router.push({ name: 'view-patient', params: { id: create.id } })
  }
}
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <PageHeader title="New Patient" :icon-src="UserPlusIcon"></PageHeader>
    <div style="padding: 0 15px; display: flex; justify-content: space-between; border-top:1px solid #333">
    </div>
    <hr class="ml-4 mr-4" />
    <div class="page-scroll-area">
      <form method="POST" v-on:submit.prevent="create" class="p-5" ref="createForm">
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <CheckboxField :options="['Mr.', 'Mrs.', 'Prof.', 'Dr.']" label="Title" name="title" model-value="Mr." />
          <CheckboxField :options="['Male', 'Female']" label="Gender" name="sex" />
          <CheckboxField :options="['Single', 'Married', 'Widowed']" label="Marital Status" name="marital_status" />
          <CheckboxField :options="['Yes', 'No']" label="ANC Patient?" required name="anc" model-value="No"/>
          <div>
            <TextField label="Last Name" placeholder="Doe" name="lastname" required></TextField>
          </div>
          <div>
            <TextField label="Middle Name" placeholder="M." name="middlename"></TextField>
          </div>
          <div>
            <TextField label="First Name" placeholder="John" name="firstname" required></TextField>
          </div>
          <div>
            <TextField label="Age" placeholder="42" name="age"></TextField>
          </div>
          <div>
            <TextField label="Card No" placeholder="HC0090" name="card_no"></TextField>
          </div>
          <div>
            <TextField label="Phone Number" placeholder="08123456789" name="phone" type="tel"></TextField>
          </div>
          <div>
            <TextField label="Email Address" placeholder="john.doe@example.com" name="email"></TextField>
          </div>
          <div>
            <TextField label="Home Address" placeholder="20 John Doe Avenue..." name="address" required></TextField>
          </div>
          <div>
            <TextField label="Occupation" placeholder="Engineer" name="occupation"></TextField>
          </div>
        </div>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div class="flex gap-2 mt-5">
            <PrimaryButton type="submit" class="w-full">Submit</PrimaryButton>
            <SecondaryButton type="button" class="w-full" v-on:click.prevent="() => $router.go(-1)">Cancel
            </SecondaryButton>
          </div>

        </div>
      </form>
    </div>
  </div>
</template>