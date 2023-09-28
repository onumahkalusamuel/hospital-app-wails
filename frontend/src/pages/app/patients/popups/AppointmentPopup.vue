<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import apiRequest from '@/services/http/api-requests';
import PopUp from '@/components/PopUp.vue';
import { toasts, popupStore } from '@/stores';
import TextArea from '@/components/form/TextArea.vue';
import SelectField from '@/components/form/SelectField.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';

defineProps<{ popupId: string }>()
const appointmentRef = ref(null);
const route = useRoute();
const emit = defineEmits(["update:closed"]);

const appointment = async () => {
  const formData = Object.fromEntries(new FormData(appointmentRef.value as never as HTMLFormElement).entries())
  const admit = await apiRequest.post(`patients/${route.params.id}/initiate-appointment`, formData);
  if (admit.message) {
    toasts.addToast({ message: "appointment set successfully.", type: 'success' });
    emit('update:closed');
    popupStore.show = false;
  }
}

</script>
<template>
  <pop-up title="Initiate Appointment" :id="popupId">
    <form method="POST" v-on:submit.prevent="appointment" ref="appointmentRef">
      <div class="flex gap-x-3">
        <div class="w-full">
          <SelectField label="Appointment Type" name="appointment_type" :options="[[''], ['Emergency'], ['Regular']]"
            required></SelectField>
        </div>
      </div>
      <div class="flex">
        <div class="w-full"><TextArea label="Note" placeholder="Note" name="note"></TextArea></div>
      </div>
      <div class="flex gap-3 mt-5">
        <div>
          <PrimaryButton type="submit">Submit</PrimaryButton>
        </div>
        <div>
        <SecondaryButton type="button" v-on:click.prevent="()=>popupStore.show=false">Cancel</SecondaryButton>
      </div>
    </div>
  </form>
</pop-up></template>