import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import legacy from '@vitejs/plugin-legacy'
import vuetify from 'vite-plugin-vuetify'
import compression from 'vite-plugin-compression';

export default defineConfig({
  base: './',
  plugins: [
    compression(),
    vue(),
    legacy(),
    vuetify({}),
  ],
  define: {
    'process.env': { ...process.env }
  },
})
