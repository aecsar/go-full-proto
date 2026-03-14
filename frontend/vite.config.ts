import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import path from "node:path";

// https://vite.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
      "@pb": path.resolve(__dirname, "./src/gen/pb/aecsar/go_proto/v1"),
    },
  },
  plugins: [svelte()],
});
