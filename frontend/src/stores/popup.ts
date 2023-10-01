import { reactive } from 'vue';

export const popupStore = reactive({
  id: '' as string,
  show: false as boolean,
  child_show: false as boolean,
  child_id: '' as string,
});