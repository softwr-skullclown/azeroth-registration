<script lang="ts">
  import Realms from './lib/Realms.svelte'
  interface NavigationItem {
    label: string;
    url: string;
  }

  interface UIConfig {
    siteTitle: string;
    gameVersion: string;
    realmList: string;
  }

  // Initial state with loading values
  let pageConfig: UIConfig = $state({
    siteTitle: "Azeroth Registration",
    gameVersion: "3.3.5a",
    realmList: "logon.yourserver.com",
  });

  let navigation: NavigationItem[] = [
    {
      label: "Register",
      url: "/register",
    },
  ];

  function updatePageTitle(newTitle: string): void {
    document.title = newTitle;
  }

  // Error state
  let error: string | null = $state(null);
  let isLoading: boolean = $state(true);

  // Fetch configuration when component mounts
  async function fetchConfig(): Promise<void> {
    try {
      const response = await fetch('/api/config');
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data: UIConfig = await response.json();
      pageConfig = data;
      updatePageTitle(data.siteTitle);
      error = null;
    } catch (err) {
      error = err instanceof Error ? err.message : 'An error occurred while fetching configuration';
      console.error('Error fetching configuration:', err);
    } finally {
      isLoading = false;
    }
  }

  // Call the fetch function when component mounts
  fetchConfig();
</script>

<div class="min-h-screen bg-gray-50 dark:bg-gray-900">
  <!-- Navigation Bar -->
  <nav class="bg-white dark:bg-gray-800 shadow-md">
    <div class="container mx-auto px-4">
      <div class="relative flex items-center justify-center h-16">
        <!-- Navigation Items - Centered -->
        <div class="w-full flex justify-center items-center">
          <div class="flex space-x-8">
            {#each navigation as navItem (navItem.url)}
              <a 
                href={navItem.url} 
                class="text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white transition-colors"
              >
                {navItem.label}
              </a>
            {/each}
          </div>
        </div>
      </div>
    </div>
  </nav>

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
    <!-- Main Content Sections -->
    <main class="container mx-auto px-4 py-8">
      <div class="flex gap-6">
        <!-- Main Content Section (3/4 width) -->
        <section class="w-3/4 bg-white dark:bg-gray-800 rounded-lg shadow-sm p-6">
          <Realms />
        </section>

        <!-- Sidebar Section (1/4 width) -->
        <section class="w-1/4 bg-white dark:bg-gray-800 rounded-lg shadow-sm p-6">
          <h2 class="text-2xl font-bold text-gray-800 dark:text-white mb-4">
            Server Information
          </h2>
          <p class="text-gray-600 dark:text-gray-300">
            Realmlist: <span class="text-blue-500 dark:text-cyan-500">{pageConfig.realmList}</span><br />
            Game Version: <span class="text-blue-500 dark:text-cyan-500">{pageConfig.gameVersion}</span>
          </p>
        </section>
      </div>
    </main>
  {/if}
</div>

<style>
</style>
