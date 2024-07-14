import { ref } from 'vue';

export const token = ref(localStorage.getItem('token'))

export const userInfo = ref({});

export const userId = ref(localStorage.getItem('userid'));

export const userName = ref(localStorage.getItem('username'));