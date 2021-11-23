import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { VitePWA } from 'vite-plugin-pwa'

const vitePwaOptions = {
  mode: 'development',
  srcDir: 'src',
  filename: 'sw.ts',
  base: '/',
  strategies: 'injectManifest',
  includeAssets: ['/favicon.svg'], // <== don't remove slash, for testing purposes
  manifest: {
    name: 'PWA Inject Manifest',
    short_name: 'PWA Inject',
    theme_color: '#ffffff',
    icons: [
      {
        src: '/pwa-192x192.png', // <== don't remove slash, for testing purposes
        sizes: '192x192',
        type: 'image/png',
      },
      {
        src: '/pwa-512x512.png', // <== don't remove slash, for testing purposes
        sizes: '512x512',
        type: 'image/png',
      },
      {
        src: '/pwa-512x512.png', // <== don't remove slash, for testing purposes
        sizes: '512x512',
        type: 'image/png',
        purpose: 'any maskable',
      },
    ],
  },
}

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), VitePWA(vitePwaOptions)],
  define: {
    'process.env': {}
  },
  server: {
    port: 3000,
  },
})
