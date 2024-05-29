import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import ViteMonacoPlugin from 'vite-plugin-monaco-editor'

export default defineConfig({
  plugins: [
    vue(),
    ViteMonacoPlugin,
  ],
  define: {
    'process.env': { ...process.env }
  },
})
