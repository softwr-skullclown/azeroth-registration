<script lang="ts">
    interface Props {
        raceId: string;
        genderId: string;
    }

    let { raceId, genderId }: Props = $props();

    let raceName: string = $state("");
    let genderName: string = $state("");

    switch(raceId) {
        case "1":
            raceName = "Human";
            break;
        case "2":
            raceName = "Orc";
            break;
        case "3":
            raceName = "Dwarf";
            break;
        case "4":
            raceName = "Nightelf";
            break;
        case "5":
            raceName = "Undead";
            break;
        case "6":
            raceName = "Tauren";
            break;
        case "7":
            raceName = "Gnome";
            break;
        case "8":
            raceName = "Troll";
            break;
        case "10":
            raceName = "Bloodelf";
            break;
        case "11":
            raceName = "Draenei";
            break;
    }

    switch(genderId) {
        case "0":
            genderName = "Male";
            break;
        case "1":
            genderName = "Female";
            break;
    }
</script>

<img src="/icons/race/{raceId}-{genderId}.gif" alt="{genderName} {raceName}" />