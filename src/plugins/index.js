import vuetify from './vuetify'
import router from '../routers/index';
import vue3AceEditor from './vue3-ace-editor';

export function registerPlugins (app) {
  app.use(vuetify)
  app.use(router)
}
