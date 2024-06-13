import vuetify from './vuetify'
import AvatarCropper from 'vue-avatar-cropper';
import router from '../routers/index';

export function registerPlugins (app) {
  app.use(vuetify)
  app.use(router)
  app.use(AvatarCropper)
}
