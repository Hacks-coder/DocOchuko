<script>
	import Button from "../resuable/Button.svelte";
	import Text from "../resuable/Text.svelte";
	import Channels from "./Channels.svelte";

  let backgroundUrl = "/img3.png";
</script>

<section class="bg-cover bg-center p-5" style={`background-image: url("${backgroundUrl}")`}>
    <div class="flex flex-col items-center text-center p-10">
        <Text text = "Don’t Wait—Your Health Matters! Book Now!" className = "text-lg font-bold" />
        <Button text = "Book" className = "bg-white shadow-sm shadow-black font-bold w-32 h-10 rounded-md"/>
    </div>
        <Channels />
        <Text text = "Trending Books" className = "" />
</section>