<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import apiRequest from '@/services/http/api-requests';
import PopUp from '@/components/PopUp.vue';
import { toasts, popupStore, hospital } from '@/stores';
import TextField from '@/components/form/TextField.vue';
import TextArea from '@/components/form/TextArea.vue';
import PrimaryButton from '@/components/form/PrimaryButton.vue';
import SecondaryButton from '@/components/form/SecondaryButton.vue';
import { Patient, PatientHistory, PatientHistoryTypes } from '@/interfaces';
import CheckboxField from '@/components/form/CheckboxField.vue';

const props = defineProps<{ popupId: string, patient: Patient, history: PatientHistory }>()
const updateHistoryRef = ref(null);
const route = useRoute();
const emit = defineEmits(["update:closed"]);

const updateHistory = async () => {
    const formData = new FormData(updateHistoryRef.value as never as HTMLFormElement)
    const update = await apiRequest.putMulti(`patients/${route.params.id}/patient-history/${props.history.id}`, formData);
    console.log(update)
    if (update.message) {
        toasts.addToast({ message: "history updated successfully", type: 'success' });
        emit('update:closed');
        popupStore.show = false;
    }
}

const historyTypes = ref(['General', 'Appointment', 'Admission', 'Discharge', 'Diagnosis', 'Examination', 'TestResult', 'Treatment'] as PatientHistoryTypes[]);
const activeTab = ref('General' as PatientHistoryTypes);

</script>
<template>
    <pop-up :width="600" :title="`${patient.lastname} ${patient.firstname} - Update History`" :id="popupId"
        pop-class="min-h-[95vh]">
        <form method="POST" v-on:submit.prevent="updateHistory" ref="updateHistoryRef" enctype="multipart/form-data">
            <div class="pb-3">
                <TextField label="Subject" placeholder="Subject of history" name="subject" :value="history.subject"
                    required></TextField>
            </div>
            <div>
                <CheckboxField :options="historyTypes" label="History Type" v-model="activeTab" />
            </div>

            <div class="flex flex-col py-3 gap-y-3" v-for="his, index in historyTypes" :key="index"
                :class="activeTab == his ? '' : 'hidden'">
                <div class="flex gap-x-3 w-full" v-if="his == 'Admission'">
                    <div class="w-full">
                        <TextField label="Ward Name" placeholder="Maternity Ward" name="details[admission][ward_number]"
                            :value="history.details.admission.ward_number">
                        </TextField>
                    </div>
                    <div class="w-full">
                        <TextField label="Room Number" placeholder="R-024" name="details[admission][room_number]"
                            :value="history.details.admission.room_number">
                        </TextField>
                    </div>
                </div>
                <div class="flex gap-x-3 w-full" v-if="his == 'Appointment'">
                    <div class="w-full">
                        <CheckboxField :options="['Emergency', 'Regular']" label="Appointment Type"
                            name="details[appointment][appointment_type]"
                            :value="history.details.appointment.appointment_type" />
                    </div>
                </div>
                <div class="flex gap-x-3 w-full">
                    <div class="w-full"><TextArea :label="`${his} Note`" :placeholder="`${his} Note`"
                            :name="`details[${his.toLowerCase()}][note]`" rows="2"
                            :value="history.details[`${his.toLowerCase()}`].note"></TextArea></div>
                    <div class="w-full">
                        <TextField class="flex h-[65%]" :label="`Document (.png, .jpg, .jpeg)`" type="file"
                            :name="`details[${his.toLowerCase()}][document]`" accept=".png, .jpg, .jpeg"></TextField>
                    </div>
                </div>

                <div class="mb-2 border-[1px] bg-blue-100 p-2 rounded pointer-cusrsor"
                    v-if="history.details[`${his.toLowerCase()}`].document.length">
                    <div class="title">Previous Document:</div>
                    <a :href="`${hospital.asset_base_url}/files/images_${patient.id}/${history.details[`${his.toLowerCase()}`].document}`"
                        target="_blank">
                        <img class="max-h-[150px]"
                            :src="`${hospital.asset_base_url}/files/images_${patient.id}/${history.details[`${his.toLowerCase()}`].document}`"
                            alt="Supporting Document" />
                    </a>
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

<style scoped>
.title {
    color: #0078d4;
    border-radius: 5px;
    font-weight: bold;
    margin-right: 5px;
}</style>