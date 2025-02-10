<script lang="ts">
  interface Realm {
    id: number,
    name: string,
    flag: number,
    icon: number,
    population: number,
  }

  let realms: Realm[] = [];

  // Error state
  let error: string | null = null;
  let isLoading: boolean = true;

  async function fetchRealms(): Promise<void> {
    try {
      const response = await fetch('/api/realms');
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data: Realm[] = await response.json();
      realms = data;
      
      error = null;
    } catch (err) {
      error = err instanceof Error ? err.message : 'An error occurred while fetching realms';
      console.error('Error fetching realms:', err);
    } finally {
      isLoading = false;
    }
  }

  fetchRealms();
</script>

<!-- Error Message -->
{#if error}
<div class="container mx-auto px-4 py-4">
  <div class="bg-red-100 dark:bg-red-900 border border-red-400 text-red-700 dark:text-red-200 px-4 py-3 rounded relative" role="alert">
    <strong class="font-bold">Error:</strong>
    <span class="block sm:inline"> {error}</span>
  </div>
</div>
{/if}

<!-- Loading State -->
{#if isLoading}
<div class="container mx-auto px-4 py-8">
  <div class="flex justify-center items-center">
    <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900 dark:border-white"></div>
  </div>
</div>
{:else}
  <h2 class="text-2xl font-bold text-gray-800 dark:text-white mb-4">
    Realms
  </h2>
  <div class="text-gray-600 dark:text-gray-300">
    {#each realms as realm}
        <div>{realm.name}</div>
    {/each}
  </div>
{/if}