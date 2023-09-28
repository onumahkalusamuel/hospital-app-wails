<script lang="ts" setup>
import { useRouter } from 'vue-router';
import apiRequest from '@/services/http/api-requests';
import { onMounted, ref } from 'vue';
import TextField from '@/components/form/TextField.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import { toasts } from '@/stores/toasts';

const router = useRouter();
const createadmin = ref(null);

onMounted(async () => {
  try {
    const req = await apiRequest.get('hospital-details');
    if (req) { router.push({ name: 'login' }) }

  } catch (e) {
    console.log(e)
  }
});

const createAdmin = async () => {
  const formData = new FormData(createadmin.value as never as HTMLFormElement)
  const createAdmin = await apiRequest.post('create-admin', Object.fromEntries(formData.entries()));
  if (createAdmin.message) {
    toasts.addToast({ message: createAdmin.message, type: 'success' });
    router.push({ name: 'hospital-setup' })
  }
}
</script>;

<template>
  <p class="text-xl">Create Super Admin</p>
  <form method="post" class="py-5" v-on:submit.prevent="createAdmin" ref="createadmin">
    <div class="space-y-2">
      <TextField name="firstname" placeholder="First Name" class="w-full" />
      <TextField name="lastname" placeholder="Last Name" class="w-full" />
      <TextField name="username" placeholder="Username" class="w-full" />
      <TextField name="password" placeholder="Password" class="w-full" />
      <PrimaryButton class="w-full" type="submit">Create Super Admin</PrimaryButton>
    </div>
  </form>
  <small>The super admin is the first user on the app. The account will have access to every section of the app.</small>
</template>
