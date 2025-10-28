//||------------------------------------------------------------------------------------------------||
//|| Imports
//||------------------------------------------------------------------------------------------------||

import { defineConfig } from 'vite'
import react          from '@vitejs/plugin-react'
import path           from 'path'

//||------------------------------------------------------------------------------------------------||
//|| Vite Configuration
//||------------------------------------------------------------------------------------------------||

export default defineConfig({
      plugins: [react()],
      envDir: path.resolve(__dirname, '..'),

      //||------------------------------------------------------------------------------------------------||
      //|| Dev Server
      //||------------------------------------------------------------------------------------------------||
      server: {
            port: 5173,
            proxy: {
                  '/auth': {
                        target: 'http://localhost:8081',
                        changeOrigin: true,
                        secure: false,
                  },
                  '/api': {
                        target: 'http://localhost:8081',
                        changeOrigin: true,
                        secure: false,
                  },
                  '/user': {
                        target: 'http://localhost:8081',
                        changeOrigin: true,
                        secure: false,
                  },
                  '/public': {
                        target: 'http://localhost:8081',
                        changeOrigin: true,
                        secure: false,
                  },
                  '/verify': {
                        target: 'http://localhost:8081',
                        changeOrigin: true,
                        secure: false,
                  },
                  '/v1/api': {
                        target: 'http://localhost:8081',
                        changeOrigin: true,
                        secure: false,
                  },
            },
      },

      //||------------------------------------------------------------------------------------------------||
      //|| Build Optimization
      //||------------------------------------------------------------------------------------------------||
      build: {
            minify: 'esbuild',         // faster + smaller
            sourcemap: false,          // remove .map files from final build
            cssCodeSplit: true,        // split large CSS bundles
            rollupOptions: {
                  output: {
                        manualChunks: {
                              react: ['react', 'react-dom'],
                              stripe: ['@stripe/react-stripe-js', '@stripe/stripe-js'],
                              tf: ['@tensorflow/tfjs', '@tensorflow-models/face-landmarks-detection'],
                        },
                  },
            },
      },
})
