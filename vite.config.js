import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import legacy from '@vitejs/plugin-legacy'
import Components from 'unplugin-vue-components/vite'
import { VuetifyResolver } from 'unplugin-vue-components/resolvers'
import vuetify from 'vite-plugin-vuetify'

export default defineConfig({
  base: './',
  plugins: [
    vue(),
    legacy(),
    vuetify({}),
    Components({
      resolvers: [VuetifyResolver()],
    }),
  ],
  define: {
    'process.env': { ...process.env }
  },
})
