<script>
    import IconInstagram from '~icons/cib/instagram';
    import IconYouTube from '~icons/cib/youtube';
    import IconFacebook from '~icons/cib/facebook';
    import IconLinkedIn from '~icons/cib/linkedin';
	import Text from '../resuable/Text.svelte';
    const channelsIcon = [
        {icon: IconYouTube, text: "Youtube"},
        {icon: IconYouTube, text: "Youtube"},
        {icon: IconYouTube, text: "Youtube"},
        {icon: IconYouTube, text: "Youtube"}
    ]
</script>

<section class="w-fit bg-white shadow-md shadow-black p-3 rounded-lg    ">
    <Text text = "Channels" className = "font-bold" />
    {#each channelsIcon as item}
        <div class="flex gap-3">
            <item.icon />
            <Text text = {item.text} className = "text-sm" />
        </div>
    {/each}
</section>