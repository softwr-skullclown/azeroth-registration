<script lang="ts">
  import CharacterClass from "./CharacterClass.svelte";
  import CharacterRace from "./CharacterRace.svelte";

  interface Props {
		realmId: number;
	}

  interface Character {
      id: number;
      name: string;
      level: number;
      class: number;
      race: number;
      gender: number;
  };

  let { realmId }: Props = $props();

  let characters: Character[] = $state([]);

  // Error state
  let error: string | null = $state(null);
  let isLoading: boolean = $state(true);

  async function fetchCharacters(): Promise<void> {
    try {
      const response = await fetch(`api/realms/${realmId}/online-characters`);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data: Character[] = await response.json();
      characters = data;
      
      error = null;
    } catch (err) {
      error = err instanceof Error ? err.message : 'An error occurred while fetching characters';
      console.error('Error fetching characters:', err);
    } finally {
      isLoading = false;
    }
  }

  fetchCharacters();
</script>

{#if characters.length > 0}
  <div class="w-full overflow-x-auto bg-white dark:bg-gray-800 rounded-lg shadow">
    <table class="w-full text-left">
      <thead>
        <tr class="bg-gray-100 dark:bg-gray-700">
          <th class="px-6 py-3 text-gray-800 dark:text-gray-200 font-semibold">Name</th>
          <th class="px-6 py-3 text-gray-800 dark:text-gray-200 font-semibold">Race</th>
          <th class="px-6 py-3 text-gray-800 dark:text-gray-200 font-semibold">Class</th>
          <th class="px-6 py-3 text-gray-800 dark:text-gray-200 font-semibold">Level</th>
        </tr>
      </thead>
      <tbody class="divide-y divide-gray-200 dark:divide-gray-600">
        {#each characters as character}
          <tr class="hover:bg-gray-50 dark:hover:bg-gray-700">
            <td class="px-6 py-4 text-gray-700 dark:text-gray-300">{character.name}</td>
            <td class="px-6 py-4 text-gray-700 dark:text-gray-300"><CharacterRace raceId={`${character.race}`} genderId={`${character.gender}`} /></td>
            <td class="px-6 py-4 text-gray-700 dark:text-gray-300"><CharacterClass classId={`${character.class}`} /></td>
            <td class="px-6 py-4 text-gray-700 dark:text-gray-300">{character.level}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
{:else}
    <div>No characters online.</div>
{/if}