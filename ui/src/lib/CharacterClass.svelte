<script lang="ts">
    interface Props {
        classId: string;
    }

    let { classId }: Props = $props();

    let className: string = $state("");

    switch(classId) {
        case "1":
            className = "Warrior";
            break;
        case "2":
            className = "Paladin";
            break;
        case "3":
            className = "Hunter";
            break;
        case "4":
            className = "Rogue";
            break;
        case "5":
            className = "Priest";
            break;
        case "6":
            className = "Deathknight";
            break;
        case "7":
            className = "Shaman";
            break;
        case "8":
            className = "Mage";
            break;
        case "9":
            className = "Warlock";
            break;
        case "11":
            className = "Druid";
            break;
    }
</script>

<img src="/icons/class/{classId}.gif" alt={className} />