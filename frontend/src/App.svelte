<script lang="ts">
  import { greetClient } from "./lib/buf-client";
  import type { GreetResponse } from "@pb/greet_pb";

  let name = $state("");
  let loading = $state(false);
  let res = $state<GreetResponse | null>(null);

  async function greet() {
    loading = true;
    try {
      const r = await greetClient.greet({
        name,
      });

      res = r;
    } catch (e) {
      console.log("error greeting: ", e);
    } finally {
      loading = false;
    }
  }
</script>

<section id="center">
  <div>
    <h1>Protobuf, ConnectRPC, Go, Svelte</h1>
  </div>

  <div class="form">
    <input bind:value={name} name="name" class="input" placeholder="Name" />

    <button class="button" onclick={greet}>
      {#if loading}
        Loading...
      {:else}
        Greet
      {/if}
    </button>
  </div>

  {#if res}
    <div class="res">
      {res.greeting}
    </div>
  {/if}
</section>

<section id="spacer"></section>

<style>
  .form {
    display: flex;
    gap: 20px;
    align-items: center;
    justify-content: center;
  }

  .res {
    text-align: center;
    margin-top: 10px;
  }
</style>
