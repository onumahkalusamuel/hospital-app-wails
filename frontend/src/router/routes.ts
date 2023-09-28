import { RouteRecordRaw } from 'vue-router';
import { appGuard } from '@/guards/app-guard';

import MainLayout from '@/layouts/MainLayout.vue';
import AppLayout from '@/layouts/AppLayout.vue';

import Activate from '@/pages/Activate.vue';
import Login from '@/pages/Login.vue';
import CreateAdmin from '@/pages/CreateAdmin.vue';

import HospitalSetup from '@/pages/HospitalSetup.vue';
import Dashboard from '@/pages/app/Dashboard.vue';
import Settings from '@/pages/app/Settings.vue';

import Patients from '@/pages/app/patients/Index.vue';
import AddPatient from '@/pages/app/patients/Add.vue';
import ViewPatient from '@/pages/app/patients/View.vue';
import EditPatient from '@/pages/app/patients/Edit.vue';
import PatientHistoryPrintView from '@/pages/app/patients/HistoryPrintView.vue';
import DeliveriesIndex from '@/pages/app/patients/DeliveriesIndex.vue';
import BillingsIndex from '@/pages/app/patients/BillingsIndex.vue';

import Deliveries from '@/pages/app/deliveries/Index.vue';
import AddDelivery from '@/pages/app/deliveries/Add.vue';
import ViewDelivery from '@/pages/app/deliveries/View.vue';

import Staff from '@/pages/app/staff/Index.vue';
import AddStaff from '@/pages/app/staff/Add.vue';
import ViewStaff from '@/pages/app/staff/View.vue';

import Invoices from '@/pages/app/billings/Index.vue';
import AddInvoice from '@/pages/app/billings/Add.vue';
import ViewInvoice from '@/pages/app/billings/View.vue';
import PrintReceipt from '@/pages/app/billings/PrintReceipt.vue';
import PrintInvoice from '@/pages/app/billings/PrintInvoice.vue';

import Reports from '@/pages/app/reports/Index.vue';
import AncPatientPrintView from '@/pages/app/reports/AncPatientPrintView.vue';
import DeliveriesPrintPreview from '@/pages/app/reports/DeliveriesPrintView.vue';

export const routes: RouteRecordRaw[] = [
  {
    path: '/', component: MainLayout, children: [
      { path: '', redirect: { name: 'login' } },
      { path: 'activate/:code?', component: Activate, name: 'activate' },
      { path: 'create-admin', component: CreateAdmin, name: 'create-admin' },
      { path: 'hospital-setup', component: HospitalSetup, name: 'hospital-setup' },
    ]
  },
  { path: '/login', component: Login, name: 'login' },
  { path: '/app/patients/:id/patient-history/print', component: PatientHistoryPrintView, name: 'patient-history-print-view' },
  // billings
  { path: '/app/billings/:id/print-receipt', component: PrintReceipt, name: 'billings-print-receipt' },
  { path: '/app/billings/:id/print-invoice', component: PrintInvoice, name: 'billings-print-invoice' },
  // reports 
  { path: '/app/reports/anc-patient', component: AncPatientPrintView, name: 'anc-patient-print-view' },
  { path: '/app/reports/deliveries', component: DeliveriesPrintPreview, name: 'deliveries-print-view' },
  {
    path: '/app/', component: AppLayout, beforeEnter: appGuard, children: [
      { path: 'dashboard', component: Dashboard, name: 'dashboard' },
      { path: 'settings', component: Settings, name: 'settings' },
      { path: 'patients', component: Patients, name: 'patients' },
      { path: 'patients/add', component: AddPatient, name: 'add-patient' },
      { path: 'patients/:id', component: ViewPatient, name: 'view-patient' },
      { path: 'patients/:id/edit', component: EditPatient, name: 'edit-patient' },
      { path: 'patients/:id/patient-billings', component: BillingsIndex, name: 'patient-billings' },
      { path: 'patients/:id/patient-deliveries', component: DeliveriesIndex, name: 'patient-deliveries' },

      { path: 'deliveries', component: Deliveries, name: 'deliveries' },
      { path: 'deliveries/add/:patient_id?', component: AddDelivery, name: 'add-delivery' },
      { path: 'deliveries/:id', component: ViewDelivery, name: 'view-delivery' },

      { path: 'staff', component: Staff, name: 'staff' },
      { path: 'staff/add', component: AddStaff, name: 'add-staff' },
      { path: 'staff/:id', component: ViewStaff, name: 'view-staff' },

      { path: 'billings', component: Invoices, name: 'billings' },
      { path: 'billings/add/:patient_id?', component: AddInvoice, name: 'add-invoice' },
      { path: 'billings/:id', component: ViewInvoice, name: 'view-invoice' },

      { path: 'reports', component: Reports, name: 'reports' },

      { path: ':pathMatch(.*)', redirect: { name: 'dashboard' } },
    ]
  },
  { path: '/:pathMatch(.*)*', redirect: { name: 'login' } },
];