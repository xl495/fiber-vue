import { UserConfig, defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite';
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers';

// https://vitejs.dev/config/
export default defineConfig({
  // 代理
  server: {
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:3000/api/',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    }
  },
  resolve: {
    alias: {
      '@': '/src'
    }
  },
  plugins: [
    vue(),
    // Components({
    //   resolvers: [
    //     AntDesignVueResolver({
    //       importStyle: false, // css in js
    //     }),
    //   ],
    // })
  ],
} as UserConfig)
