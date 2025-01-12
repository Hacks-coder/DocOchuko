<script>
    import IconPatient from '~icons/mdi/patient';
    import IconLibrary from '~icons/hugeicons/library';
    import IconWork from '~icons/material-symbols/work';
	import Text from '../resuable/Text.svelte';
    const ratings = [
        {icon: IconWork, text: "Active", percentage: "100%", bg: "#047064"},
        {icon: IconLibrary, text: "Books", percentage: "100%", bg: "#160094"},
        {icon: IconPatient, text: "Patients", percentage: "100%", bg: "#960000"}
    ]
</script>


<section class="flex justify-center gap-5 mt-10">
    {#each ratings as item}
        <div class="text-center">
            <div class={`w-20 h-20 shadow-md rounded-full flex items-center justify-center ${item.bg}`} >
                <item.icon class="text-5xl"/>           
            </div>
            <Text text = {item.text} className = "text-aqua font-bold text-xl"/>
            <Text text = {item.percentage} className = "text-gold"/>
        </div>
    {/each}
</section>