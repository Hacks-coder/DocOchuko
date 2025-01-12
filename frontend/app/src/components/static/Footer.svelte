<script lang="ts">
	import Text from "../resuable/Text.svelte";
    import IconInstagram from '~icons/cib/instagram';
    import IconYouTube from '~icons/cib/youtube';
    import IconFacebook from '~icons/cib/facebook';
    import IconLinkedIn from '~icons/cib/linkedin';
    const footerPath = [
        {
            title: "Resources",
            text: [
                {title: "Consultation", href: "/"},
                {title: "Services", href: "/"},
                {title: "Pharmarcy", href: "/"},
                {title: "Article", href: "/"},
                {title: "Video", href: "/"}              
            ]
        },
        {
            title: "Resources",
            text: [
                {title: "Consultation", href: "/"},
                {title: "Services", href: "/"},
                {title: "Pharmarcy", href: "/"},
                {title: "Article", href: "/"},
                {title: "Video", href: "/"}              
            ]
        },
        {
            title: "Resources",
            text: [
                {title: "Consultation", href: "/"},
                {title: "Services", href: "/"},
                {title: "Pharmarcy", href: "/"},
                {title: "Article", href: "/"},
                {title: "Video", href: "/"}              
            ]
        },
    ]
</script>

<footer class="grid bg-purple p-10 grid-cols-1 md:grid-cols-2 md:place-items-center lg:grid-cols-4 mt-10">
    {#each footerPath as item}
        <section class="flex flex-col gap-4">
           <Text text = {item.title} className = "font-bold" />
            {#each item.text as sub_title}
                <a href={sub_title.href}>{sub_title.title}</a> 
            {/each}
        </section>
    {/each}
    <div class="mt-5">
        <Text text = "Contact" className = "font-bold" />
        <div class="flex gap-2 mt-2">
            <IconYouTube class = "text-2xl"/>
            <IconFacebook class = "text-xl"/>
            <IconLinkedIn class = "text-xl"/>
            <IconInstagram class = "text-xl"/>
        </div>
    </div>
</footer>

<!-- TODO -->