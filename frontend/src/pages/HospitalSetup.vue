<script lang="ts" setup>
import apiRequest from '@/services/http/api-requests';
import TextField from '@/components/form/TextField.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import { toasts } from '@/stores/toasts';

const router = useRouter();
const hospitalsetup = ref(null);

onMounted(async() => {
  try {
      const req = await apiRequest.get('hospital-details');
      if(req) { router.push({name: 'login'})}

    } catch(e) {
      console.log(e)
    }
})

const hospitalSetup = async () => {
  const formData = new FormData(hospitalsetup.value as never as HTMLFormElement)
  const hospitalSetup = await apiRequest.postMulti('hospital-setup', formData);
  if(hospitalSetup?.message) {
    toasts.addToast({message: hospitalSetup.message, type: 'success'});
    router.push({name: 'login'})
  }
}
</script>;

<template>
  <p class="text-xl">Hospital Setup</p>
  <form method="post" class="py-5" v-on:submit.prevent="hospitalSetup" ref="hospitalsetup" enctype="multipart/form-data">
    <div class="space-y-2">
      <TextField name="hospital_name" placeholder="Hospital Name" class="w-full"/>
      <TextField name="hospital_address" placeholder="Physical Address" class="w-full"/>
      <TextField name="hospital_email" placeholder="Email Address" class="w-full"/>
      <TextField name="hospital_phone" placeholder="Phone Number" class="w-full"/>
      <TextField label="Hospital Logo" name="hospital_logo" class="w-full font-xl" type="file" accept=".png, .jpg, .jpeg"/>
      <PrimaryButton class="w-full" type="submit">Submit</PrimaryButton>
    </div>
  </form>
</template>
