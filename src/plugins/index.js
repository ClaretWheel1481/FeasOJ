import vuetify from './vuetify'
import router from '../routers/index';
import vue3AceEditor from './vue3-ace-editor';
import AvatarCropper from 'vue-avatar-cropper';

export function registerPlugins(app) {
  app.use(vuetify)
  app.use(router)
  app.use(AvatarCropper)
}
