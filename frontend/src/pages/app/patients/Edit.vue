<script lang="ts" setup>
import Breadcrumbs, { BreadcrumbItem } from '@/components/Breadcrumbs.vue';
import apiRequest from '@/services/http/api-requests';
import { Patient } from '@/interfaces'
import TextField from '@/components/form/TextField.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';
import { useRoute, useRouter } from 'vue-router';
import { onMounted, ref } from 'vue';
import { toasts } from '@/stores';
import CheckboxField from '@/components/form/CheckboxField.vue';

const route = useRoute();

const breadcrumbs = ref([
  { title: "Dashboard", link: { name: "dashboard" } },
  { title: "Patients", link: { name: "patients" } },
] as BreadcrumbItem[]);

const patient = ref({} as Patient);
const form = ref(null);
const router = useRouter();

const fetchPatient = async () => {
  patient.value = await apiRequest.get(`patients/${route.params.id}`);
  if (!breadcrumbs.value[breadcrumbs.value.length - 1]?.current) {
    breadcrumbs.value.push(
      { title: `${patient.value.lastname} ${patient.value.firstname}`, link: { name: 'view-patient' } },
      { title: 'Edit', current: true },
    );
  }
}
const update = async () => {
  const formData = Object.fromEntries(new FormData(form.value as never as HTMLFormElement).entries())
  const update = await apiRequest.put(`patients/${route.params.id}`, formData);
  if (update.message) {
    toasts.addToast({ message: "account updated successfully.", type: 'success' });
    router.go(-1)
  }
}

onMounted(async () => { await fetchPatient() });
</script>;

<template>
  <div class="page-wrapper">
    <Breadcrumbs :items="breadcrumbs"></Breadcrumbs>
    <div class="page-scroll-area">
      <form method="POST" v-on:submit.prevent="update" class="p-5" ref="form">
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <CheckboxField :options="['Mr.', 'Mrs.', 'Prof.', 'Dr.']" label="Title" name="title" v-model="patient.title" />
          <CheckboxField :options="['Male', 'Female']" label="Gender" name="sex" v-model="patient.sex" />
          <CheckboxField :options="['Single', 'Married', 'Widowed']" label="Marital Status" name="marital_status" v-model="patient.marital_status" />
          <CheckboxField :options="['Yes', 'No']" label="ANC Patient?" required name="anc" v-model="patient.anc" />
          <div>
            <TextField label="Last Name" placeholder="Doe" name="lastname" :value="patient.lastname"></TextField>
          </div>
          <div>
            <TextField label="Middle Name" placeholder="M." name="middlename" :value="patient.middlename"></TextField>
          </div>
          <div>
            <TextField label="First Name" placeholder="John" name="firstname" :value="patient.firstname"></TextField>
          </div>
          <div>
            <TextField label="Card No" :value="(patient.card_no as string)" readonly></TextField>
          </div>
          <div>
            <TextField label="Phone Number" name="phone" type="tel" :value="patient.phone"></TextField>
          </div>
          <div>
            <TextField label="Email Address" name="email" :value="patient.email"></TextField>
          </div>
          <div>
            <TextField label="Age" name="age" required :value="patient.age"></TextField>
          </div>
          <div>
            <TextField label="Home Address" placeholder="20 John Doe Avenue..." name="address" :value="patient.address"
              required></TextField>
          </div>
          <div>
            <TextField label="Occupation" placeholder="Engineer" name="occupation" :value="patient.occupation">
            </TextField>
          </div>
        </div>
        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <div class="flex gap-2 mt-5">
            <PrimaryButton type="submit" class="w-full">Update</PrimaryButton>
            <SecondaryButton type="button" class="w-full" v-on:click.prevent="()=>$router.go(-1)">Cancel</SecondaryButton>
          </div>
      </div>
    </form>
  </div>
</div></template>