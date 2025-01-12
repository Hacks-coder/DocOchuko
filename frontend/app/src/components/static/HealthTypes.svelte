<script>
	import Button from "../resuable/Button.svelte";
	import Text from "../resuable/Text.svelte";

</script>
<section class="grid grid-cols-3 gap-20 p-2 shadow-lg border-4 border-black shadow-black mt-10 bg-gray rounded-lg">
    <div class="w-32 h-32 border-4 border-black overflow-hidden">
        <img src="/img1.jpg" alt="" >
    </div>
    <div class="flex flex-col justify-between col-span-2">
        <div>
            <Text text = "Personal Health" className = "font-bold" />
            <Text text = "A happy family is a healthy family! From check-ups to wellness tips," className = "text-sm" />
        </div>
        <Button text = "Personal Health" className = "bg-purple w-48 font-bold text-white" />
    </div>
</section>