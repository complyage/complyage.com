//||------------------------------------------------------------------------------------------------||
//|| Imports
//||------------------------------------------------------------------------------------------------||

import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { fileURLToPath } from 'node:url'
import { dirname, resolve } from 'node:path'

//||------------------------------------------------------------------------------------------------||
//|| Vite Configuration
//||------------------------------------------------------------------------------------------------||

const __filename = fileURLToPath(import.meta.url)
const __dirname = dirname(__filename)

export default defineConfig({
      plugins: [react()],
      envDir: resolve(__dirname, '..'),

      server: {
            port: 3040, // local docs dev port
      },

      build: {
            outDir: 'dist',
            minify: 'esbuild',
            cssCodeSplit: true,
            rollupOptions: {
                  output: {
                        manualChunks: {
                              react: ['react', 'react-dom'],
                        },
                  },
            },
      },
})
