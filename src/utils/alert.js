import { ref } from 'vue';

export const dialog = ref(false);
export const dialogMessage = ref('');

export const showAlert = (message,path) => {
  dialogMessage.value = message;
  dialog.value = true;
  if(path === "reload"){
    setTimeout(() => {
      location.reload();
    },500);
  }else if(path === ""){
    return
  }else {
    setTimeout(() => {
      window.location = path;
    },1000);
  }
};
