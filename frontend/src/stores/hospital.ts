import { reactive } from "vue";

type Keys =
  | "hospital_name"
  | "hospital_address"
  | "hospital_logo"
  | "hospital_email"
  | "hospital_phone"
  | "asset_base_url";
type Data = { [key: string]: any };

export const hospital = reactive({
  hospital_name: "",
  hospital_address: "",
  hospital_logo: "",
  hospital_email: "",
  hospital_phone: "",
  asset_base_url: "",
  setAll(data: Data) {
    (Object.keys(data) as Keys[]).forEach((key) => {
      this[key] = data[key];
    });
  },
  set(key: Keys, value: any) {
    this[key] = value;
    sessionStorage.setItem(`hospital_${key}`, value);
  },
  get(key: Keys) {
    if (!this[key] && sessionStorage.getItem(`hospital_${key}`)) {
      this[key] = sessionStorage.getItem(`hospital_${key}`) as any;
    }
    return this[key];
  },
});
