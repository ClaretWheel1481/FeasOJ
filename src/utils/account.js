import { ref } from 'vue';

export const token = ref(localStorage.getItem('token'));

export const userName = ref(localStorage.getItem('username'));