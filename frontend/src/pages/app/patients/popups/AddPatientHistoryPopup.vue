<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import apiRequest from '@/services/http/api-requests';
import PopUp from '@/components/PopUp.vue';
import { toasts, popupStore } from '@/stores';
import TextField from '@/components/form/TextField.vue';
import TextArea from '@/components/form/TextArea.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';
import { Patient, PatientHistoryTypes } from '@/interfaces';
import CheckboxField from '@/components/form/CheckboxField.vue';

defineProps<{ popupId: string, patient: Patient }>()
const addHistoryRef = ref(null);
const route = useRoute();
const emit = defineEmits(["update:closed"]);
const historyTypes = ref(['General', 'Appointment', 'Admission', 'Discharge', 'Diagnosis', 'Examination', 'TestResult', 'Treatment'] as PatientHistoryTypes[]);

const addHistory = async () => {
    const formData = new FormData(addHistoryRef.value as never as HTMLFormElement)
    const admit = await apiRequest.postMulti(`patients/${route.params.id}/patient-history`, formData);

    if (admit.id) {
        toasts.addToast({ message: "history created successfully.", type: 'success' });
        emit('update:closed');
        popupStore.show = false;
    }
}

const activeTab = ref('General' as PatientHistoryTypes);

</script>
<template>
    <pop-up :width="920" :title="`${patient.lastname} ${patient.firstname} - Add History`" :id="popupId"
        pop-class="min-h-[95vh]">
        <form method="POST" v-on:submit.prevent="addHistory" ref="addHistoryRef" enctype="multipart/form-data">
            <div class="pb-3">
                <TextField label="Subject" placeholder="Subject of history" name="subject" required></TextField>
            </div>
            <div>
                <CheckboxField :options="historyTypes" label="History Type" v-model="activeTab" />
            </div>

            <div class="flex flex-col py-3 gap-y-3" v-for="history, index in historyTypes" :key="index"
                :class="activeTab == history ? '' : 'hidden'">
                <div class="flex gap-x-3 w-full" v-if="history == 'Admission'">
                    <div class="w-full">
                        <TextField label="Ward Name" placeholder="Maternity Ward" name="details[admission][ward_number]">
                        </TextField>
                    </div>
                    <div class="w-full">
                        <TextField label="Room Number" placeholder="R-024" name="details[admission][room_number]">
                        </TextField>
                    </div>
                </div>
                <div class="flex gap-x-3 w-full" v-if="history == 'Appointment'">
                    <div class="w-full">
                        <CheckboxField :options="['Emergency', 'Regular']" label="Appointment Type"
                            name="details[appointment][appointment_type]" />
                    </div>
                </div>
                <div class="flex gap-x-3 w-full">
                    <div class="w-full"><TextArea :label="`${history} Note`" :placeholder="`${history} Note`"
                            :name="`details[${history.toLowerCase()}][note]`" rows="2"></TextArea></div>
                    <div class="w-full">
                        <TextField class="flex h-[65%]" :label="`Document (.png, .jpg, .jpeg)`" type="file"
                            :name="`details[${history.toLowerCase()}][document]`" accept=".png, .jpg, .jpeg"></TextField>
                    </div>
                </div>
            </div>

            <div class="flex gap-3 mt-5">
                <div>
                    <PrimaryButton type="submit">Submit</PrimaryButton>
                </div>
                <div>
                    <SecondaryButton type="button" @click.prevent="() => popupStore.show = false">Cancel</SecondaryButton>
                </div>
            </div>
        </form>
    </pop-up>
</template>