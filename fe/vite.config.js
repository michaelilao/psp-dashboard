import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'
import dotenv from 'dotenv'
import path from 'path'

const defineEnv = {}
// console.log(process.env)
// dotenv.config({ path: path.resolve('..', '.env') })
// // eslint-disable-next-line no-undef
// for (const key in process.env) {
//   if (key.startsWith('VITE_')) {
//     // eslint-disable-next-line no-undef
//     defineEnv[`import.meta.env.${key}`] = JSON.stringify(process.env[key])
//   }
// }

// https://vite.dev/config/
export default defineConfig({
  // define: defineEnv,
  plugins: [
    react(),
    tailwindcss(),
  ],
})
