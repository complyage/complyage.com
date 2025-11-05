//||------------------------------------------------------------------------------------------------||
//|| Imports
//||------------------------------------------------------------------------------------------------||

import { defineConfig } from 'vite'
import react          from '@vitejs/plugin-react'
import path           from 'path'

//||------------------------------------------------------------------------------------------------||
//|| Vite Configuration
//||------------------------------------------------------------------------------------------------||

import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'

export default defineConfig({
      plugins: [react()],
      envDir: path.resolve(__dirname, '..'),

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
