<script>
	import Button from "../resuable/Button.svelte";


	import InputField from "../resuable/InputField.svelte";

</script>
<form action="" class="flex flex-col items-center mt-10">
    <InputField type = "text" placeholder = "Message Title" className = "w-72 h-10 bg-purple-1 opacity-45 border-none placeholder:text-white" />
    <textarea class="w-72 h-48 bg-purple-1 opacity-45 border-none mt-2"></textarea>
    <Button text = "Book Appointment" className = "bg-purple w-72 h-10 mt-2" />
</form>